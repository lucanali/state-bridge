// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "./libraries/SimpleSerialize.sol";
import "./Struct.sol";

library BeaconLightClient {
    // Time window (in seconds) for optimistic updates before they can be finalized
    uint256 constant OPTIMISTIC_UPDATE_TIMEOUT = 86400;

    // Number of slots in a single epoch of the beacon chain
    uint256 constant SLOTS_PER_EPOCH = 32;

    // Number of slots in a sync committee period (approximately 27.3 hours)
    uint256 constant SLOTS_PER_SYNC_COMMITTEE_PERIOD = 8192;

    // Minimum number of sync committee participants required for a valid update
    uint256 constant MIN_SYNC_COMMITTEE_PARTICIPANTS = 10;

    // Total size of the sync committee (number of validators in the committee)
    uint256 constant SYNC_COMMITTEE_SIZE = 512;

    // Index of the finalized root in the beacon state
    uint256 constant FINALIZED_ROOT_INDEX = 105;

    // Index of the next sync committee in the beacon state
    uint256 constant NEXT_SYNC_COMMITTEE_INDEX = 55;

    // Index of the execution state root in the beacon state
    uint256 constant EXECUTION_STATE_ROOT_INDEX = 402;

    // LightClientState stores the current state of the light client
    // It maintains the necessary data to verify beacon chain blocks and sync committee updates
    struct LightClientState {
        // Root of the genesis validators set, used for signature verification
        bytes32 genesisValidatorsRoot;
        // Timestamp of the genesis block, used to calculate current slot
        uint256 genesisTime;
        // Time duration of each slot in seconds (12 seconds for Ethereum)
        uint256 secondsPerSlot;
        // Current fork version of the beacon chain
        bytes4 defaultForkVersion;
        // Current head slot number of the light client
        uint64 head;
        // Mapping of slot numbers to beacon block headers
        // Stores block headers for verification and state updates
        mapping(uint64 => BeaconBlockHeader) headers;
        // Mapping of slot numbers to execution state roots
        // Stores the execution state roots for each block
        mapping(uint64 => bytes32) executionStateRoots;
        // Mapping of sync committee periods to committee roots
        // Used to verify sync committee updates and block signatures
        mapping(uint256 => bytes32) syncCommitteeRootByPeriod;
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

    function verifyValidatorSignature(
        LightClientState storage state,
        BeaconBlockHeader memory header,
        uint8 v,
        bytes32 r,
        bytes32 s,
        address validator
    ) internal view returns(bool) {
        bytes32 signingRoot = SimpleSerialize.computeSigningRoot(
            header,
            state.defaultForkVersion,
            state.genesisValidatorsRoot
        );

        // Verify the validator's signature using standard Ethereum signature
        // Recover the signer's address from the signature
        address recoveredSigner = ecrecover(
            keccak256(
                abi.encodePacked(
                    "\x19Ethereum Signed Message:\n32",
                    signingRoot
                )
            ),
            v,
            r,
            s
        );

        // Check if the recovered address matches the validator
        return(recoveredSigner == validator);
    }

    function verifyCorrectBlockHash(
        LightClientState storage state,
        uint256 blockNumber,
        bytes32 correctBlockHash,
        bytes32[] memory finalityBranch32
    ) internal view {
        // Try to get the next block header
        BeaconBlockHeader memory nextHeader = state.headers[
            uint64(blockNumber + 1)
        ];

        // If next block exists, verify using its parentRoot
        if (nextHeader.slot != 0) {
            require(
                SimpleSerialize.isValidMerkleBranch(
                    correctBlockHash,
                    FINALIZED_ROOT_INDEX,
                    finalityBranch32,
                    nextHeader.parentRoot
                ),
                "Invalid correct block hash"
            );
        }

        // If next block doesn't exist, verify using the current block's stateRoot
        BeaconBlockHeader memory currentHeader = state.headers[
            uint64(blockNumber)
        ];
        require(currentHeader.slot != 0, "Current block not found");

        require(
            SimpleSerialize.isValidMerkleBranch(
                correctBlockHash,
                FINALIZED_ROOT_INDEX,
                finalityBranch32,
                currentHeader.stateRoot
            ),
            "Invalid correct block hash"
        );
    }

    function verifyExecutionStateRoot(
        bytes32 executionStateRoot,
        bytes32[] memory executionStateRootBranch32,
        bytes32 correctBlockHash
    ) internal pure {
        require(
            SimpleSerialize.isValidMerkleBranch(
                executionStateRoot,
                EXECUTION_STATE_ROOT_INDEX,
                executionStateRootBranch32,
                correctBlockHash
            ),
            "Invalid execution state root proof"
        );
    }

    function getSyncCommitteePeriodFromSlot(
        uint64 slot
    ) internal pure returns (uint64) {
        return uint64(slot / SLOTS_PER_SYNC_COMMITTEE_PERIOD);
    }

    function updateSyncCommittee(
        LightClientState storage state,
        uint64 period,
        bytes32 syncCommitteeRoot,
        bytes[] memory proof
    ) internal {
        // Verify the sync committee update
        bytes32[] memory proof32 = convertBranch(proof);
        require(
            SimpleSerialize.isValidMerkleBranch(
                syncCommitteeRoot,
                NEXT_SYNC_COMMITTEE_INDEX,
                proof32,
                state.headers[state.head].stateRoot
            ),
            "Invalid sync committee proof"
        );

        // Update the sync committee root
        state.syncCommitteeRootByPeriod[period] = syncCommitteeRoot;
    }

    function verifyFraudProof(
        LightClientState storage state,
        uint256 blockNumber,
        bytes32 correctBlockHash,
        bytes32 executionStateRoot,
        bytes[] memory proof
    ) internal view returns (bool) {
        // Get the current sync committee period
        uint64 currentPeriod = getSyncCommitteePeriodFromSlot(
            uint64(blockNumber)
        );

        // Get the sync committee root for this period
        bytes32 syncCommitteeRoot = state.syncCommitteeRootByPeriod[
            currentPeriod
        ];
        require(
            syncCommitteeRoot != bytes32(0),
            "Sync committee root not found"
        );

        // Verify the proof using sync committee
        bytes32[] memory proof32 = convertBranch(proof);
        if (
            !SimpleSerialize.isValidMerkleBranch(
                correctBlockHash,
                FINALIZED_ROOT_INDEX,
                proof32,
                syncCommitteeRoot
            )
        ) {
            return false;
        }

        // Verify execution state root if provided
        if (executionStateRoot != bytes32(0)) {
            return (
                SimpleSerialize.isValidMerkleBranch(
                    executionStateRoot,
                    EXECUTION_STATE_ROOT_INDEX,
                    proof32,
                    correctBlockHash
                )
            );
        }

        return true;
    }

    function setHead(
        LightClientState storage state,
        BeaconBlockHeader memory header
    ) internal {
        state.head = header.slot;
        state.headers[state.head] = header;
    }

    function getCurrentSlot(
        LightClientState storage state
    ) internal view returns (uint64) {
        return
            uint64(
                (block.timestamp - state.genesisTime) / state.secondsPerSlot
            );
    }

    function setExecutionStateRoot(
        LightClientState storage state,
        uint64 slot,
        bytes32 _executionStateRoot
    ) internal {
        state.executionStateRoots[slot] = _executionStateRoot;
    }
}
