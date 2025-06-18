# Dione Bridge

A trustless bridge implementation between Ethereum and other EVM-compatible chains.

## Architecture

The bridge consists of three main components:

1. **StateBridge Contract**: The core contract that handles state verification and fraud proofs
2. **Validator Registry**: Manages validators and their stakes
3. **Beacon Light Client**: Verifies Ethereum beacon chain state

### StateBridge Contract

The StateBridge contract is the main component that:

- Accepts block submissions from validators
- Verifies block data using the light client
- Handles fraud proofs and challenges
- Manages sync committee updates
- Processes slashing for malicious validators

Key features:

- Validators submit blocks with state roots and execution data
- Relayers can submit blocks on behalf of validators
- Anyone can challenge invalid blocks with fraud proofs
- Owner can update sync committee for light client verification
- Invalid headers are automatically reverted when challenged

#### Block Submission

Blocks can be submitted by:

- By validators directly using `submitBlock`

Both methods require:

- Block hash
- Execution state root
- Block data (parent hash, state root, etc.)
- Critical flag for time-lock

#### Fraud Proofs

The contract supports fraud proofs through the `challengeBlock` function:

- Anyone can submit a fraud proof
- Proof must include:
  - Correct block hash
  - Execution state root
  - Merkle proof for verification
- The Merkle proof is used to verify:
  1. The correct block hash against the sync committee root
  2. The execution state root against the block hash
- If challenge is successful:
  - Invalid header is reverted
  - Validator is slashed
  - Treasury receives slashed funds

#### Sync Committee Updates

The owner can update the sync committee using `setSyncCommittee`:

- Requires period number
- Sync committee root
- Only owner can call this function

**Trust Assumptions:**

- The sync committee is assumed to be trusted
- The committee members are responsible for verifying and signing blocks
- This trust model can be made trustless in the future by:
  1. Using zero-knowledge proofs to verify committee membership
  2. Implementing a more decentralized committee update mechanism
  3. Adding additional verification layers

### Validator Registry

The ValidatorRegistry contract:

- Manages validator registration
- Handles staking and slashing
- Tracks validator stakes and status

### Beacon Light Client

The BeaconLightClient library:

- Verifies beacon chain state
- Handles sync committee updates
- Manages header verification
- Supports fraud proof verification

### Signing Process

**Current Issue**: The signing process in the relayer is not currently working correctly. The relayer attempts to implement a signing process that matches the contract's verification logic, but there are issues with the signature generation and verification.

#### Current Implementation

The relayer currently implements a signing process that:

1. Creates a beacon block header with block data
2. Computes SSZ hash using SimpleSerialize algorithm
3. Computes domain using `defaultForkVersion` and `genesisValidatorsRoot`
4. Creates signing root by concatenating SSZ header and domain
5. Signs using Ethereum signed message format
6. Splits signature into v, r, s components

#### Known Issues

- The signature verification in the contract's `verifyValidatorSignature()` function may not match the relayer's signing process
- There may be discrepancies in how the signing root is computed between the relayer and contract
- The domain computation or SSZ header creation may have implementation differences
- The Ethereum signed message format may not be correctly implemented

#### Required Fixes

To make the signing process work correctly, the following need to be addressed:

1. **Align signing root computation** between relayer and contract
2. **Verify domain computation** matches `SimpleSerialize.computeDomain()`
3. **Ensure SSZ header creation** matches `SimpleSerialize.sszBeaconBlockHeader()`
4. **Fix Ethereum signed message format** implementation
5. **Test signature verification** end-to-end

The current implementation serves as a foundation but requires debugging and fixes to function properly with the contract's verification logic.

## Security

The bridge implements several security measures:

- Time-lock for critical updates
- Slashing for malicious validators
- Fraud proof system
- Sync committee verification
- Header reversion on invalid blocks

## Development

### Prerequisites

- Node.js 16+
- Go 1.19+
- Solidity 0.8.28

### Setup

1. Install dependencies:

```bash
yarn install
```

2. Compile contracts:

```bash
npx hardhat compile
```

### Deployment

1. Set up environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

2. Deploy contracts:

```bash
npx hardhat run scripts/deploy.js --network <network>
```

### Relayer Configuration

The relayer requires a `config.json` file in the `relayer` directory with the following structure:

```json
{
  "sourceChain": {
    "rpcUrl": "http://localhost:8545",
    "chainId": 1337,
    "bridgeAddr": "0x...",
    "startBlock": 0
  },
  "destinationChain": {
    "rpcUrl": "http://localhost:8546",
    "chainId": 1338,
    "bridgeAddr": "0x..."
  },
  "pollInterval": 12,
  "validatorPrivateKey": "your_validator_private_key"
}
```

Configuration fields:

- `sourceChain`: Configuration for the source chain (Ethereum)
  - `rpcUrl`: RPC endpoint URL
  - `chainId`: Chain ID (must be a number, not a string)
  - `bridgeAddr`: Address of the StateBridge contract
  - `startBlock`: Block number to start listening from
- `destinationChain`: Configuration for the destination chain
  - Same fields as sourceChain (except startBlock)
- `pollInterval`: Time in seconds between polling for new blocks
- `validatorPrivateKey`: Validator's private key used for signing blocks and sending transactions

To run the relayer:

```bash
cd relayer
go run main.go
```

# State Bridge

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat node
npx hardhat run scripts/deploy.js
```

## Deployment

To deploy the contracts:

```bash
# Deploy using Hardhat
npx hardhat run scripts/deploy.js --network <network>
```

## Running the Relayer

To run the relayer script:

```bash
# Using Foundry
forge script script/Relayer.s.sol --rpc-url <rpc_url> --broadcast
```

Make sure to set your private key in the environment:

```bash
export PRIVATE_KEY=your_private_key_here
```
