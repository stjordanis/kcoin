pragma solidity 0.4.24;

import "zos-lib/contracts/migrations/Initializable.sol";

contract MyContractV1 is Initializable {
    uint256 public value;
    
    function initialize(uint256 _value) isInitializer public {
        value = _value;
    }

    function add(uint256 _value) public {
        value = value + _value;
    }
}