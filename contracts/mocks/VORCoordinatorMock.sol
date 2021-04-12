// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;

import "@openzeppelin/contracts/GSN/Context.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "../VORConsumerBase.sol";

contract VORCoordinatorMock is Context {
    using Address for address;

    event RandomnessRequest(
        address indexed sender,
        bytes32 indexed keyHash,
        uint256 indexed seed,
        uint256 fee
    );

    function randomnessRequest(
        bytes32 keyHash,
        uint256 consumerSeed,
        uint256 feePaid
    ) public {
        require(address(_msgSender()).isContract(), "request can only be made by a contract");
        emit RandomnessRequest(_msgSender(), keyHash, consumerSeed, feePaid);
    }

    function callBackWithRandomness(
        bytes32 requestId,
        uint256 randomness,
        address consumerContract
    ) public {
        VORConsumerBase v;
        bytes memory resp = abi.encodeWithSelector(v.rawFulfillRandomness.selector, requestId, randomness);
        uint256 b = 206000;
        require(gasleft() >= b, "not enough gas for consumer");
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = consumerContract.call(resp);
        (success);
    }
}
