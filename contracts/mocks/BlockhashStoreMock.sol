// SPDX-License-Identifier: MIT
pragma solidity 0.6.6;

import "../dev/BlockhashStore.sol";

contract BlockhashStoreMock is BlockhashStore {
    function godmodeSetHash(uint256 n, bytes32 h) public {
        _sBlockhashes[n] = h;
    }
}