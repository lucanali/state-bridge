#!/bin/bash

# Compile the contract
solc --abi contracts/StateBridge.sol -o build
solc --bin contracts/StateBridge.sol -o build

# Generate Go bindings
abigen --abi=build/StateBridge.abi --bin=build/StateBridge.bin --pkg=contracts --out=relayer/contracts/StateBridge.go