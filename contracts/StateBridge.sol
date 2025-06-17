// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "./ValidatorRegistry.sol";
import "./BeaconLightClient.sol";
import "./Struct.sol";

contract StateBridge {
    using BeaconLightClient for BeaconLightClient.LightClientState;

    ValidatorRegistry public registry;
    BeaconLightClient.LightClientState public lightClientState;

    struct FraudProof {
        bytes32 correctBlockHash;
        bytes32 executionStateRoot;
        bytes[] proof;
    }

    struct BlockData {
        bytes32 parentHash;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        uint256 timestamp;
        uint256 number;
        uint8 proposerV;
        bytes32 proposerR;
        bytes32 proposerS;
    }

    struct BlockUpdate {
        bytes32 blockHash;
        uint256 challengeTimestamp;
        address proposer;
        bool challenged;
        bool isCritical;
        bytes32 executionStateRoot;
        BlockData blockData;
    }

    uint256 public updateDelay = 24 hours;
    uint256 public blockNumber;
    uint256 public lastBlockNumber;
    mapping(uint256 => BlockUpdate) public updates;

    // Add slashing parameters
    uint256 public constant SLASH_AMOUNT = 0.5 ether; // 0.5 ETH

    // Add slashing event
    event ValidatorSlashed(
        address indexed validator,
        address indexed challenger,
        uint256 amount
    );

    event BlockSubmitted(
        address indexed proposer,
        uint256 blockNumber,
        bool isCritical
    );
    event Challenged(uint256 indexed blockNumber, address indexed challenger);
    event Slashed(address indexed validator, uint256 amount);
    event TreasuryWithdrawn(address indexed to, uint256 amount);

    // Add owner state variable
    address public owner;
    uint256 public treasuryBalance;

    // Add owner modifier
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner");
        _;
    }

    constructor(
        address _registry,
        bytes32 genesisValidatorsRoot,
        uint256 genesisTime,
        uint256 secondsPerSlot,
        bytes4 forkVersion
    ) {
        registry = ValidatorRegistry(_registry);
        owner = msg.sender;

        // Initialize light client state
        lightClientState.genesisValidatorsRoot = genesisValidatorsRoot;
        lightClientState.genesisTime = genesisTime;
        lightClientState.secondsPerSlot = secondsPerSlot;
        lightClientState.defaultForkVersion = forkVersion;
    }

    function submitBlock(
        bytes32 blockHash,
        bytes32 executionStateRoot,
        bool isCritical,
        BlockData memory blockData
    ) external {
        require(registry.isValidator(msg.sender), "Only validator");
        uint256 _blockNumber = blockData.number;
        require(_blockNumber > 0, "Invalid block number");
        {
            BlockUpdate memory _lastUpdate = updates[_blockNumber];
            require(
                _lastUpdate.blockHash == bytes32(0) || _lastUpdate.challenged,
                "block number already submitted"
            );
        }

        {
            BlockUpdate memory lastBlockNumberUpdate = updates[blockNumber];
            if (
                lastBlockNumberUpdate.isCritical &&
                !lastBlockNumberUpdate.challenged
            ) {
                require(
                    block.timestamp >= lastBlockNumberUpdate.challengeTimestamp,
                    "Time-lock active"
                );
            }
        }

        require(
            executionStateRoot != bytes32(0),
            "Invalid execution state root"
        );

        {
            // Create beacon header for light client
            BeaconBlockHeader memory header = BeaconBlockHeader({
                slot: uint64(_blockNumber),
                proposerIndex: 0, // This will be set by the light client
                parentRoot: blockData.parentHash,
                stateRoot: blockData.stateRoot,
                bodyRoot: blockHash
            });

            // Verify proposer signature
            lightClientState.verifyValidatorSignature(
                header,
                blockData.proposerV,
                blockData.proposerR,
                blockData.proposerS,
                msg.sender
            );

            // Update light client state with the new header
            lightClientState.setHead(header);
            if (executionStateRoot != bytes32(0)) {
                lightClientState.setExecutionStateRoot(
                    uint64(_blockNumber),
                    executionStateRoot
                );
            }
        }

        updates[_blockNumber] = BlockUpdate({
            blockHash: blockHash,
            challengeTimestamp: block.timestamp + updateDelay,
            proposer: msg.sender,
            challenged: false,
            isCritical: isCritical,
            executionStateRoot: executionStateRoot,
            blockData: blockData
        });

        lastBlockNumber = blockNumber;
        blockNumber = _blockNumber;

        emit BlockSubmitted(msg.sender, _blockNumber, isCritical);
    }

    function challengeBlock(
        uint256 _blockNumber,
        FraudProof memory proof
    ) external {
        require(registry.isValidator(msg.sender), "Only validator");
        BlockUpdate memory update = updates[_blockNumber];
        verifyBlockChallenged(update.challenged);
        verifyCriticalTimestamp(update.isCritical, update.challengeTimestamp);

        require(
            proof.correctBlockHash != update.blockHash,
            "No fraud detected"
        );

        // Verify block data matches the claimed block hash
        bytes32 computedBlockHash = keccak256(
            abi.encodePacked(
                update.blockData.parentHash,
                update.blockData.stateRoot,
                update.blockData.transactionsRoot,
                update.blockData.receiptsRoot,
                update.blockData.timestamp,
                update.blockData.number
            )
        );
        if (computedBlockHash == update.blockHash) {
            // Verify the fraud proof using light client library
            bool isValid = lightClientState.verifyFraudProof(
                _blockNumber,
                proof.correctBlockHash,
                proof.executionStateRoot,
                proof.proof
            );

            require(!isValid, "Invalid fraud proof");
        }

        // Revert the header update in light client state
        if (lightClientState.head == uint64(_blockNumber)) {
            // Clear the invalid header
            lightClientState.head = uint64(lastBlockNumber);
            delete lightClientState.headers[uint64(_blockNumber)];
            delete lightClientState.executionStateRoots[uint64(_blockNumber)];
        }

        updates[_blockNumber].challenged = true;

        _slashValidator(update.proposer);
        emit Challenged(_blockNumber, msg.sender);
    }

    function _slashValidator(address validator) internal {
        // Slash the validator
        registry.slash(validator, SLASH_AMOUNT);

        treasuryBalance += SLASH_AMOUNT;

        // Emit event
        emit ValidatorSlashed(validator, msg.sender, SLASH_AMOUNT);
    }

    function getUpdate(
        uint256 _blockNumber
    ) external view returns (BlockUpdate memory) {
        return updates[_blockNumber];
    }

    function withdrawTreasuryBalance(address to) external onlyOwner {
        require(to != address(0), "Invalid recipient address");
        require(treasuryBalance > 0, "No balance to withdraw");

        uint256 amount = treasuryBalance;
        treasuryBalance = 0;

        (bool success, ) = to.call{value: amount}("");
        require(success, "Transfer failed");

        emit TreasuryWithdrawn(to, amount);
    }

    function setSyncCommittee(
        uint64 period,
        bytes32 syncCommitteeRoot,
        bytes[] memory proof
    ) external onlyOwner {
        // Update the sync committee using the light client library
        lightClientState.updateSyncCommittee(period, syncCommitteeRoot, proof);
    }

    function verifyCriticalTimestamp(
        bool isCritical,
        uint256 challengeTimestamp
    ) internal view {
        if (isCritical) {
            require(block.timestamp < challengeTimestamp, "Time-lock active");
        } else {
            revert("Block state update is not critical");
        }
    }

    function verifyBlockChallenged(bool challenged) internal pure {
        require(!challenged, "Challenged");
    }

    // Add function to get sync committee root for a period
    function syncCommitteeRootByPeriod(
        uint256 period
    ) external view returns (bytes32) {
        return lightClientState.syncCommitteeRootByPeriod[period];
    }

    // Add function to get header for a slot
    function headers(
        uint64 slot
    ) external view returns (BeaconBlockHeader memory) {
        return lightClientState.headers[slot];
    }
}
