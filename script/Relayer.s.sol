// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../contracts/StateBridge.sol";
import "../contracts/ValidatorRegistry.sol";
import "../contracts/libraries/SimpleSerialize.sol";

contract RelayerScript is Script {
    StateBridge public stateBridge;
    ValidatorRegistry public validatorRegistry;
    uint256 public deployerPrivateKey;

    // Storage variables to avoid stack too deep
    bytes32 public parentHash;
    bytes32 public stateRoot;
    bytes32 public transactionsRoot;
    bytes32 public receiptsRoot;
    uint256 public timestamp;
    uint256 public number;
    BeaconBlockHeader public header;
    bytes32 public signingRoot;
    uint8 public v;
    bytes32 public r;
    bytes32 public s;

    // Constants for sync committee
    uint256 constant NEXT_SYNC_COMMITTEE_INDEX = 55; // From BeaconLightClient.sol
    uint256 constant SLOTS_PER_PERIOD = 8192; // From BeaconLightClient.sol

    function setUp() public {
        deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        stateBridge = StateBridge(0x74Cee9B5EFF0349766e4Dd727f784a42De37B2b3);
        validatorRegistry = ValidatorRegistry(
            0x5FC26b06c865B9233069d42a09feB08EEAB20888
        );
    }

    function createBlockData() public {
        parentHash = keccak256("parent");
        // Create a mock sync committee root
        bytes32 syncCommitteeRoot = keccak256("sync_committee");
        // Create Merkle proof and get state root
        bytes[] memory proof = createMerkleProof(
            syncCommitteeRoot,
            NEXT_SYNC_COMMITTEE_INDEX,
            bytes32(0) // root parameter is unused
        );

        bytes32[] memory proof32 = convertBranch(proof);

        bytes32 _stateRoot = SimpleSerialize.restoreMerkleRoot(
            syncCommitteeRoot,
            NEXT_SYNC_COMMITTEE_INDEX,
            proof32
        );

        stateRoot = _stateRoot;
        transactionsRoot = keccak256("transactions");
        receiptsRoot = keccak256("receipts");
        timestamp = block.timestamp;
        number = 1;
    }

    function createBeaconHeader() public {
        // Create block hash first
        bytes32 blockHash = keccak256(
            abi.encodePacked(
                parentHash,
                stateRoot,
                transactionsRoot,
                receiptsRoot,
                timestamp,
                number
            )
        );

        header = BeaconBlockHeader({
            slot: uint64(number),
            proposerIndex: 0,
            parentRoot: parentHash,
            stateRoot: stateRoot,
            bodyRoot: blockHash
        });
    }

    function signMessage() public {
        // Compute signing root using SimpleSerialize
        signingRoot = SimpleSerialize.computeSigningRoot(
            header,
            bytes4(0x00000000), // defaultForkVersion
            0x1367160dcd01806c3ce79c58695cdf73ef86888130237370eda1408827c55e6d // genesisValidatorsRoot
        );

        // Create message for signing
        bytes32 message = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", signingRoot)
        );

        // Sign the message
        (v, r, s) = vm.sign(deployerPrivateKey, message);
    }

    function createMerkleProof(
        bytes32 leaf,
        uint256 index,
        bytes32 root
    ) public pure returns (bytes[] memory) {
        // Create a proof that will verify against the given root
        bytes[] memory proof = new bytes[](32);

        // Start with the leaf
        bytes32 value = leaf;

        // Build the proof by working backwards from the root
        for (uint i = 0; i < 32; i++) {
            if ((index / (2 ** i)) % 2 == 1) {
                // If index bit is 1, we need the left sibling
                bytes32 sibling = sha256(bytes.concat(bytes32(0), value));
                proof[i] = abi.encodePacked(sibling);
                value = sha256(bytes.concat(sibling, value));
            } else {
                // If index bit is 0, we need the right sibling
                bytes32 sibling = sha256(bytes.concat(value, bytes32(0)));
                proof[i] = abi.encodePacked(sibling);
                value = sha256(bytes.concat(value, sibling));
            }
        }

        return proof;
    }

    function convertBranch(
        bytes[] memory branch
    ) internal pure returns (bytes32[] memory) {
        bytes32[] memory branch32 = new bytes32[](branch.length);
        for (uint256 i = 0; i < branch.length; i++) {
            require(branch[i].length == 32, "Invalid branch element length");
            branch32[i] = bytes32(branch[i]);
        }
        return branch32;
    }

    function getSyncCommitteePeriod(uint64 slot) public pure returns (uint64) {
        return uint64(slot / SLOTS_PER_PERIOD);
    }

    function setupSyncCommittee() public {
        // Create a mock sync committee root
        bytes32 syncCommitteeRoot = keccak256("sync_committee");

        // Create a Merkle proof that will verify against our state root
        bytes[] memory proof = createMerkleProof(
            syncCommitteeRoot,
            NEXT_SYNC_COMMITTEE_INDEX,
            stateRoot
        );

        bytes32[] memory proof32 = convertBranch(proof);

        bytes32 _stateRoot = SimpleSerialize.restoreMerkleRoot(
            syncCommitteeRoot,
            NEXT_SYNC_COMMITTEE_INDEX,
            proof32
        );

        // First submit a block to set the state root
        stateBridge.submitBlock(
            header.bodyRoot,
            stateRoot,
            true,
            StateBridge.BlockData({
                parentHash: parentHash,
                stateRoot: _stateRoot,
                transactionsRoot: transactionsRoot,
                receiptsRoot: receiptsRoot,
                timestamp: timestamp,
                number: number,
                proposerV: v,
                proposerR: r,
                proposerS: s
            })
        );

        // Then set the sync committee for the correct period
        uint64 period = getSyncCommitteePeriod(uint64(number));
        stateBridge.setSyncCommittee(period, syncCommitteeRoot, proof);
    }

    function submitBlock() public {
        stateBridge.submitBlock(
            header.bodyRoot,
            stateRoot,
            false,
            StateBridge.BlockData({
                parentHash: parentHash,
                stateRoot: stateRoot,
                transactionsRoot: transactionsRoot,
                receiptsRoot: receiptsRoot,
                timestamp: timestamp,
                number: number,
                proposerV: v,
                proposerR: r,
                proposerS: s
            })
        );
    }

    function challengeBlock() public {
        bytes32 invalidStateRoot = keccak256("invalid");
        StateBridge.FraudProof memory fraudProof = StateBridge.FraudProof({
            correctBlockHash: keccak256(
                abi.encodePacked(
                    parentHash,
                    invalidStateRoot,
                    transactionsRoot,
                    receiptsRoot,
                    timestamp,
                    number
                )
            ),
            executionStateRoot: invalidStateRoot,
            proof: new bytes[](0) // Empty proof for testing
        });

        stateBridge.challengeBlock(number, fraudProof);
    }

    function run() external {
        setUp();
        vm.startBroadcast(deployerPrivateKey);

        // Create block data
        createBlockData();

        // Create beacon header
        createBeaconHeader();

        // Sign message
        signMessage();

        // Setup sync committee
        setupSyncCommittee();

        // Submit challenge
        challengeBlock();

        vm.stopBroadcast();
    }
}
