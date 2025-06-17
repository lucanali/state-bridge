// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract ValidatorRegistry {
    mapping(address => uint256) public stakes;
    mapping(address => bool) public isValidator;

    function register() external payable {
        require(msg.value >= 1 ether, "Minimum stake: 1 ETH");
        require(!isValidator[msg.sender], "Already registered");
        isValidator[msg.sender] = true;
        stakes[msg.sender] = msg.value;
    }

    function slash(address validator, uint256 slashAmount) external {
        require(isValidator[validator], "Not validator");
        uint256 _amount =  slashAmount <= stakes[validator] ? slashAmount : stakes[validator];   
        stakes[validator] -= _amount;
        // Send to treasury or burn
    }

    function getStake(address validator) external view returns (uint256) {
        return stakes[validator];
    }
}
