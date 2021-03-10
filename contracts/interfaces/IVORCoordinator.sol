// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

interface IVORCoordinator {
    function getProviderAddress(bytes32 _keyHash) external view returns (address);
    function randomnessRequest(bytes32 keyHash, uint256 consumerSeed, uint256 feePaid) external;
    function topUpGas(address _provider) external payable returns (bool success);
}