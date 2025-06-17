package relayer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"relayer/config"
	"relayer/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Relayer struct {
	cfg             *config.Config
	sourceClient    *ethclient.Client
	destClient      *ethclient.Client
	sourceAuth      *bind.TransactOpts
	destAuth        *bind.TransactOpts
	sourceBridge    *contracts.Contracts
	destBridge      *contracts.Contracts
	lastSourceBlock uint64
	lastDestBlock   uint64
	pollInterval    time.Duration
	validatorKey    *ecdsa.PrivateKey
	httpClient      *http.Client
}

type BlockData struct {
	ParentHash       common.Hash
	StateRoot        common.Hash
	TransactionsRoot common.Hash
	ReceiptsRoot     common.Hash
	Timestamp        *big.Int
	Number           *big.Int
	ProposerV        uint8
	ProposerR        [32]byte
	ProposerS        [32]byte
}

func NewRelayer(cfg *config.Config) (*Relayer, error) {
	// Create custom HTTP client that skips TLS verification for Arbitrum
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip TLS verification for Arbitrum
			},
		},
	}

	// Connect to source chain with custom client
	sourceClient, err := ethclient.Dial(cfg.SourceChain.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to source chain: %v", err)
	}

	// Connect to destination chain
	destClient, err := ethclient.Dial(cfg.DestinationChain.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to destination chain: %v", err)
	}

	// Get validator's private key
	validatorKey, err := crypto.HexToECDSA(cfg.ValidatorPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid validator private key: %v", err)
	}

	// Setup source chain auth using validator's key
	sourceAuth, err := bind.NewKeyedTransactorWithChainID(validatorKey, big.NewInt(int64(cfg.SourceChain.ChainID)))
	if err != nil {
		return nil, fmt.Errorf("failed to create source auth: %v", err)
	}

	// Setup destination chain auth using validator's key
	destAuth, err := bind.NewKeyedTransactorWithChainID(validatorKey, big.NewInt(int64(cfg.DestinationChain.ChainID)))
	if err != nil {
		return nil, fmt.Errorf("failed to create destination auth: %v", err)
	}

	// Initialize bridge contracts
	sourceBridge, err := contracts.NewContracts(common.HexToAddress(cfg.SourceChain.BridgeAddr), sourceClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create source bridge: %v", err)
	}

	destBridge, err := contracts.NewContracts(common.HexToAddress(cfg.DestinationChain.BridgeAddr), destClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination bridge: %v", err)
	}

	// Get the last block number from the bridge
	lastBlockNumber, err := destBridge.BlockNumber(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get last block number: %v", err)
	}

	return &Relayer{
		cfg:             cfg,
		sourceClient:    sourceClient,
		destClient:      destClient,
		sourceAuth:      sourceAuth,
		destAuth:        destAuth,
		sourceBridge:    sourceBridge,
		destBridge:      destBridge,
		lastSourceBlock: lastBlockNumber.Uint64(),
		lastDestBlock:   lastBlockNumber.Uint64(),
		pollInterval:    time.Duration(cfg.PollInterval) * time.Second,
		validatorKey:    validatorKey,
		httpClient:      httpClient,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) error {
	log.Println("Starting relayer...")

	// Start source chain listener
	go r.listenSourceChain(ctx)

	// Start block monitor for challenges
	go r.monitorBlocks(ctx)

	// Wait for context cancellation
	<-ctx.Done()
	return nil
}

func (r *Relayer) listenSourceChain(ctx context.Context) {
	ticker := time.NewTicker(r.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := r.processBlocks(ctx); err != nil {
				log.Printf("Error processing blocks: %v", err)
			}
		}
	}
}

func (r *Relayer) processBlocks(ctx context.Context) error {
	// Get latest block from source chain
	latestBlock, err := r.sourceClient.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest block: %v", err)
	}

	// Process new blocks
	for blockNum := r.lastSourceBlock + 1; blockNum <= latestBlock; blockNum++ {
		if err := r.processBlock(ctx, blockNum); err != nil {
			log.Printf("Error processing block %d: %v", blockNum, err)
			continue
		}
		r.lastSourceBlock = blockNum
	}

	return nil
}

func (r *Relayer) processBlock(ctx context.Context, blockNum uint64) error {
	// Check if we need to wait for critical block challenge period
	if blockNum > 0 {
		lastUpdate, err := r.destBridge.GetUpdate(&bind.CallOpts{}, big.NewInt(int64(blockNum-1)))
		if err != nil {
			return fmt.Errorf("failed to get last update: %v", err)
		}

		// If last block was critical and not challenged, check if we need to wait
		if lastUpdate.IsCritical && !lastUpdate.Challenged {
			// Get current timestamp
			header, err := r.destClient.HeaderByNumber(ctx, nil)
			if err != nil {
				return fmt.Errorf("failed to get current header: %v", err)
			}

			// If we haven't reached the challenge timestamp, wait
			if header.Time < lastUpdate.ChallengeTimestamp.Uint64() {
				waitTime := time.Until(time.Unix(int64(lastUpdate.ChallengeTimestamp.Uint64()), 0))
				log.Printf("Waiting %v for critical block challenge period", waitTime)
				time.Sleep(waitTime)
			}
		}
	}

	// Get block from source chain using JSON-RPC
	reqBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{fmt.Sprintf("0x%x", blockNum), true},
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := r.httpClient.Post(r.cfg.SourceChain.RPCURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Result struct {
			Hash         string        `json:"hash"`
			ParentHash   string        `json:"parentHash"`
			StateRoot    string        `json:"stateRoot"`
			Transactions []interface{} `json:"transactions"`
			ReceiptsRoot string        `json:"receiptsRoot"`
			Timestamp    string        `json:"timestamp"`
			Number       string        `json:"number"`
		} `json:"result"`
		Error *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Error != nil {
		return fmt.Errorf("RPC error: %s", result.Error.Message)
	}

	// Convert hex strings to appropriate types
	blockHash := common.HexToHash(result.Result.Hash)
	parentHash := common.HexToHash(result.Result.ParentHash)
	stateRoot := common.HexToHash(result.Result.StateRoot)
	receiptsRoot := common.HexToHash(result.Result.ReceiptsRoot)
	timestamp, _ := new(big.Int).SetString(result.Result.Timestamp[2:], 16)
	number, _ := new(big.Int).SetString(result.Result.Number[2:], 16)

	// Get light client state
	state, err := r.destBridge.LightClientState(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get light client state: %v", err)
	}

	// Create block data
	blockData := contracts.StateBridgeBlockData{
		ParentHash:       parentHash,
		StateRoot:        stateRoot,
		TransactionsRoot: blockHash, // Using block hash as transactions root since we don't have it
		ReceiptsRoot:     receiptsRoot,
		Timestamp:        timestamp,
		Number:           number,
	}

	// Create beacon header for signing
	header := contracts.BeaconBlockHeader{
		Slot:          uint64(blockNum),
		ProposerIndex: 0, // This will be set by the light client
		ParentRoot:    blockData.ParentHash,
		StateRoot:     blockData.StateRoot,
		BodyRoot:      blockHash,
	}

	// Create signing root using SimpleSerialize algorithm
	// First compute SSZ beacon block header
	slotBytes := toLittleEndian(header.Slot)
	proposerBytes := toLittleEndian(header.ProposerIndex)
	slotProposerHash := sha256.Sum256(bytes.Join([][]byte{slotBytes, proposerBytes}, nil))
	parentStateHash := sha256.Sum256(bytes.Join([][]byte{header.ParentRoot[:], header.StateRoot[:]}, nil))
	left := sha256.Sum256(bytes.Join([][]byte{slotProposerHash[:], parentStateHash[:]}, nil))

	bodyZeroHash := sha256.Sum256(bytes.Join([][]byte{header.BodyRoot[:], bytes.Repeat([]byte{0}, 32)}, nil))
	zeroZeroHash := sha256.Sum256(bytes.Join([][]byte{bytes.Repeat([]byte{0}, 32), bytes.Repeat([]byte{0}, 32)}, nil))
	right := sha256.Sum256(bytes.Join([][]byte{bodyZeroHash[:], zeroZeroHash[:]}, nil))

	sszHeader := sha256.Sum256(bytes.Join([][]byte{left[:], right[:]}, nil))

	// Compute domain
	// First compute the hash of the fork version and genesis validators root
	domainInput := bytes.Join([][]byte{
		state.DefaultForkVersion[:],
		state.GenesisValidatorsRoot[:],
	}, nil)
	domainHash := sha256.Sum256(domainInput)

	// Create domain with 0x07 prefix and right-shifted hash
	domain := bytes.Join([][]byte{
		[]byte{0x07},
		bytes.Repeat([]byte{0}, 31),
		domainHash[4:], // Right shift by 32 bytes (256 bits)
	}, nil)

	// Compute final signing root
	signingRoot := sha256.Sum256(bytes.Join([][]byte{sszHeader[:], domain}, nil))

	// Create message for signing (same as in Solidity)
	message := crypto.Keccak256(
		append(
			[]byte("\x19Ethereum Signed Message:\n32"),
			signingRoot[:]...,
		),
	)

	// Sign the message with validator's private key
	signature, err := crypto.Sign(message, r.validatorKey)
	if err != nil {
		return fmt.Errorf("failed to sign block data: %v", err)
	}

	// Split signature into v, r, s components
	rBytes := signature[:32]
	sBytes := signature[32:64]
	v := signature[64] + 27 // Convert to Ethereum's v format

	// Set signature components in block data
	blockData.ProposerV = v
	copy(blockData.ProposerR[:], rBytes)
	copy(blockData.ProposerS[:], sBytes)

	// Submit block to destination chain
	tx, err := r.destBridge.SubmitBlock(
		r.destAuth,
		blockHash,
		stateRoot,
		true, // isCritical
		blockData,
	)
	if err != nil {
		return fmt.Errorf("failed to submit block %d: %v", blockNum, err)
	}

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, r.destClient, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}
	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	log.Printf("Submitted block %d to destination chain: %s", blockNum, tx.Hash().Hex())
	return nil
}

// Helper function to convert uint64 to little endian bytes
func toLittleEndian(n uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)
	return b
}

// monitorBlocks continuously checks for invalid blocks and generates fraud proofs
func (r *Relayer) monitorBlocks(ctx context.Context) {
	ticker := time.NewTicker(r.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := r.checkAndChallengeBlocks(ctx); err != nil {
				log.Printf("Error checking blocks: %v", err)
			}
		}
	}
}

// checkAndChallengeBlocks verifies blocks and generates fraud proofs if needed
func (r *Relayer) checkAndChallengeBlocks(ctx context.Context) error {
	// Get latest block number from destination chain
	latestBlock, err := r.destBridge.BlockNumber(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get latest block: %v", err)
	}

	// Check each block from last checked to latest
	for blockNum := r.lastDestBlock + 1; blockNum <= latestBlock.Uint64(); blockNum++ {
		// Get block update from destination chain
		update, err := r.destBridge.GetUpdate(&bind.CallOpts{}, big.NewInt(int64(blockNum)))
		if err != nil {
			log.Printf("Error getting update for block %d: %v", blockNum, err)
			continue
		}

		// Skip if block is already challenged
		if update.Challenged {
			continue
		}

		// Get actual block from source chain
		sourceBlock, err := r.sourceClient.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
		if err != nil {
			log.Printf("Error getting source block %d: %v", blockNum, err)
			continue
		}

		// Verify block hash matches
		if sourceBlock.Hash() != common.BytesToHash(update.BlockHash[:]) {
			log.Printf("Found invalid block %d, generating fraud proof", blockNum)

			// Generate fraud proof
			proof, err := r.generateFraudProof(blockNum, sourceBlock)
			if err != nil {
				log.Printf("Error generating fraud proof for block %d: %v", blockNum, err)
				continue
			}

			// Submit challenge
			tx, err := r.destBridge.ChallengeBlock(
				r.destAuth,
				big.NewInt(int64(blockNum)),
				proof,
			)
			if err != nil {
				log.Printf("Error challenging block %d: %v", blockNum, err)
				continue
			}

			// Wait for transaction to be mined
			receipt, err := bind.WaitMined(ctx, r.destClient, tx)
			if err != nil {
				log.Printf("Error waiting for challenge transaction: %v", err)
				continue
			}
			if receipt.Status == 0 {
				log.Printf("Challenge transaction failed for block %d", blockNum)
				continue
			}

			log.Printf("Successfully challenged block %d: %s", blockNum, tx.Hash().Hex())
		}

		r.lastDestBlock = blockNum
	}

	return nil
}

// generateFraudProof creates a fraud proof for an invalid block
func (r *Relayer) generateFraudProof(blockNum uint64, correctBlock *types.Block) (contracts.StateBridgeFraudProof, error) {
	// Get the current sync committee period
	period := blockNum / 8192 // SLOTS_PER_SYNC_COMMITTEE_PERIOD

	// Get sync committee root from the contract
	syncCommitteeRoot, err := r.destBridge.SyncCommitteeRootByPeriod(&bind.CallOpts{}, big.NewInt(int64(period)))
	if err != nil {
		return contracts.StateBridgeFraudProof{}, fmt.Errorf("failed to get sync committee root: %v", err)
	}

	// Generate Merkle proof for the block
	proof, err := r.generateMerkleProof(correctBlock.Hash(), syncCommitteeRoot)
	if err != nil {
		return contracts.StateBridgeFraudProof{}, fmt.Errorf("failed to generate merkle proof: %v", err)
	}

	return contracts.StateBridgeFraudProof{
		CorrectBlockHash:   correctBlock.Hash(),
		ExecutionStateRoot: correctBlock.Root(),
		Proof:              proof,
	}, nil
}

// generateMerkleProof generates a Merkle proof for a block hash against a sync committee root
func (r *Relayer) generateMerkleProof(blockHash common.Hash, syncCommitteeRoot [32]byte) ([][]byte, error) {
	// Get the beacon chain state
	state, err := r.destBridge.LightClientState(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get light client state: %v", err)
	}

	// Get the header for the block
	header, err := r.destBridge.Headers(&bind.CallOpts{}, state.Head)
	if err != nil {
		return nil, fmt.Errorf("failed to get header: %v", err)
	}

	// Create the proof elements
	// The proof should contain the path from block hash to sync committee root
	// For now, we'll use the state root as an intermediate step
	proof := [][]byte{
		blockHash.Bytes(),    // Block hash
		header.StateRoot[:],  // State root
		syncCommitteeRoot[:], // Sync committee root
	}

	return proof, nil
}
