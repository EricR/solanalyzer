pragma solidity =0.4.24;

import "test2.sol";

contract Vulnerable {
  function a(int z) {
    a();
  }
  
  function b(int z) {
    a(100);
    b(1);
  }
  
  function a() returns(bool) {
    return false; 
  }
}
