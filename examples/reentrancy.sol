pragma solidity =0.4.24;

contract Vulnerable {
	mapping (address => uint) userBalances;

	function transfer(address to, uint amount) {
        if (userBalances[msg.sender] >= amount) {
            userBalances[to] += amount;
            userBalances[msg.sender] -= amount;
        }
	}

	function withdraw() {
        uint amount = userBalances[msg.sender];
        require(msg.sender.call.value(amount)());
        userBalances[msg.sender] = 0;
	}
}
