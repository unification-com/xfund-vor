// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;

interface IVORCoordinator {
    function getProviderAddress(bytes32 _keyHash) external view returns (address);
    function randomnessRequest(bytes32 keyHash, uint256 consumerSeed, uint256 feePaid) external;
}
