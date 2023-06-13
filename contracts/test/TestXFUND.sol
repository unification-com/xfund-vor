// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract TestXFUND is ERC20 {
    using SafeMath for uint256;

    uint8 private decs;

    /**
     * See {ERC20-constructor}.
     */
    constructor(string memory name, string memory symbol, uint256 initSupply, uint8 _decs) public ERC20(name, symbol) {
        decs = _decs;
        if(initSupply > 0) {
            _mint(msg.sender, initSupply);
        }
    }

    function decimals() public view override returns (uint8) {
        return decs;
    }

    function gimme() external {
        uint256 amount = uint256(1).mul(uint256(10) ** decimals());
        _mint(msg.sender, amount);
    }
}
