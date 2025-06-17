const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("StateBridgeModule", (m) => {
  // Constants for StateBridge deployment
  const GENESIS_TIME = process.env.GENESIS_TIME || 1606824023;
  const SECONDS_PER_SLOT = process.env.SECONDS_PER_SLOT || 12;
  const FORK_VERSION = process.env.FORK_VERSION || "0x00000000";
  const GENESIS_VALIDATORS_ROOT =
    process.env.GENESIS_VALIDATORS_ROOT ||
    "0x0000000000000000000000000000000000000000000000000000000000000000";
    
    // Deploy ValidatorRegistry first
    const validatorRegistry = m.contract("ValidatorRegistry", []);

  // Deploy StateBridge with ValidatorRegistry address
  const stateBridge = m.contract("StateBridge", [
    validatorRegistry, // _registry
    GENESIS_VALIDATORS_ROOT, // genesisValidatorsRoot
    GENESIS_TIME, // genesisTime
    SECONDS_PER_SLOT, // secondsPerSlot
    FORK_VERSION, // forkVersion
  ]);

  return { validatorRegistry, stateBridge };
});
