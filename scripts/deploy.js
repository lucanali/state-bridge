const hre = require("hardhat");

async function main() {
  const [deployer] = await hre.ethers.getSigners();

  // Deploy ValidatorRegistry
  const ValidatorRegistry = await hre.ethers.getContractFactory(
    "ValidatorRegistry"
  );
  const validatorRegistry = await ValidatorRegistry.deploy();
  console.log("ValidatorRegistry deployed to:", validatorRegistry.target);

  // Register deployer as validator with 1 ETH stake
  const stakeAmount = hre.ethers.parseEther("1.0"); // 1 ETH
  await validatorRegistry.register({ value: stakeAmount });

  // For testing, we'll use a simple hash of the validator address as the genesis root
  const genesisValidatorsRoot = hre.ethers.keccak256(deployer.address);
  console.log("Generated genesis validators root:", genesisValidatorsRoot);

  // Deploy StateBridge
  const StateBridge = await hre.ethers.getContractFactory("StateBridge");
  const stateBridge = await StateBridge.deploy(
    validatorRegistry.target,
    genesisValidatorsRoot, // Use our generated root
    Math.floor(Date.now() / 1000), // genesisTime
    12, // secondsPerSlot
    "0x00000000" // forkVersion
  );
  console.log("StateBridge deployed to:", stateBridge.target);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
