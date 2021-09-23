// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;

interface BlockHashStoreInterface {
    function getBlockhash(uint256 number) external view returns (bytes32);
}
