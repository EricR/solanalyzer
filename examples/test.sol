pragma solidity 0.4.25;

contract Vulnerable {
	mapping (address => uint) userBalances;

	function transfer(address to, uint amount) onlyUser() public returns(uint) {
        if (userBalances[msg.sender] >= amount) {
            userBalances[to] += amount;
            userBalances[msg.sender] -= amount;
        }
	}

	function withdraw() public {
        var (testa, testb) = (1, 2);
        var a = 2 ** (4 + 4);
        a += 10;
        var asdf = 6 + 4 / 2 + 10;
        var aaaa = 256 >> 333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333;
	}
}
