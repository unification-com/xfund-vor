// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

interface IVORCoordinator {
    function randomnessRequest(bytes32 keyHash, uint256 consumerSeed, uint256 feePaid, address sender) external;
}