// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract ValidatorRegistry {

    uint256 public treasuryBalance;
    // Add owner state variable
    address public owner;

    mapping(address => uint256) public stakes;
    mapping(address => bool) public isValidator;

    event TreasuryWithdrawn(address indexed to, uint256 amount);

    // Add owner modifier
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

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
        treasuryBalance += _amount;
    }

    function unRegister() external {
        require(isValidator[msg.sender], "Already registered");
        isValidator[msg.sender] = false;
        uint256 amount = stakes[msg.sender];
        stakes[msg.sender] = 0;

        (bool success, ) = payable(msg.sender).call{value: amount}("");
        require(success, "Transfer failed");
    }


    function withdrawTreasuryBalance(address to) external onlyOwner {
        require(to != address(0), "Invalid recipient address");
        require(treasuryBalance > 0, "No balance to withdraw");

        uint256 amount = treasuryBalance;
        treasuryBalance = 0;

        (bool success, ) = payable(to).call{value: amount}("");
        require(success, "Transfer failed");

        emit TreasuryWithdrawn(to, amount);
    }

    function getStake(address validator) external view returns (uint256) {
        return stakes[validator];
    }
}
