// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../VORRequestIDBase.sol";

contract VORRandomnessRequestMock is Ownable, VORRequestIDBase {
    event RandomnessRequest(
        bytes32 keyHash,
        uint256 seed,
        address sender,
        uint256 fee,
        bytes32 requestID
    );

    function randomnessRequest(
        bytes32 _keyHash,
        uint256 _consumerSeed,
        uint256 _feePaid
    ) external {
        uint256 preSeed =
        makeVORInputSeed(_keyHash, _consumerSeed, _msgSender(), 0);
        bytes32 requestId = makeRequestId(_keyHash, preSeed);
        emit RandomnessRequest(
            _keyHash,
            preSeed,
            _msgSender(),
            _feePaid,
            requestId
        );
    }
}
