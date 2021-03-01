// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "./interfaces/XFundTokenInterface.sol";
import "./interfaces/IVORCoordinator.sol";
import "./VORRequestIDBase.sol";

/** ****************************************************************************
 * @notice Interface for contracts using VOR randomness
 * *****************************************************************************
 * @dev PURPOSE
 *
 * @dev Reggie the Random Oracle (not his real job) wants to provide randomness
 * @dev to Vera the verifier in such a way that Vera can be sure he's not
 * @dev making his output up to suit himself. Reggie provides Vera a public key
 * @dev to which he knows the secret key. Each time Vera provides a seed to
 * @dev Reggie, he gives back a value which is computed completely
 * @dev deterministically from the seed and the secret key.
 *
 * @dev Reggie provides a proof by which Vera can verify that the output was
 * @dev correctly computed once Reggie tells it to her, but without that proof,
 * @dev the output is indistinguishable to her from a uniform random sample
 * @dev from the output space.
 *
 * @dev The purpose of this contract is to make it easy for unrelated contracts
 * @dev to talk to Vera the verifier about the work Reggie is doing, to provide
 * @dev simple access to a verifiable source of randomness.
 * *****************************************************************************
 * @dev USAGE
 *
 * @dev Calling contracts must inherit from VORConsumerBase, and can
 * @dev initialize VORConsumerBase's attributes in their constructor as
 * @dev shown:
 *
 * @dev   contract VORConsumer {
 * @dev     constuctor(<other arguments>, address _vorCoordinator, address _xfund)
 * @dev       VORConsumerBase(_vorCoordinator, _xfund) public {
 * @dev         <initialization with other arguments goes here>
 * @dev       }
 * @dev   }
 *
 * @dev The oracle will have given you an ID for the VOR keypair they have
 * @dev committed to (let's call it keyHash), and have told you the minimum xFUND
 * @dev price for VOR service. Make sure your contract has sufficient xFUND, and
 * @dev call requestRandomness(keyHash, fee, seed), where seed is the input you
 * @dev want to generate randomness from.
 *
 * @dev Once the VORCoordinator has received and validated the oracle's response
 * @dev to your request, it will call your contract's fulfillRandomness method.
 *
 * @dev The randomness argument to fulfillRandomness is the actual random value
 * @dev generated from your seed.
 *
 * @dev The requestId argument is generated from the keyHash and the seed by
 * @dev makeRequestId(keyHash, seed). If your contract could have concurrent
 * @dev requests open, you can use the requestId to track which seed is
 * @dev associated with which randomness. See VORRequestIDBase.sol for more
 * @dev details. (See "SECURITY CONSIDERATIONS" for principles to keep in mind,
 * @dev if your contract could have multiple requests in flight simultaneously.)
 *
 * @dev Colliding `requestId`s are cryptographically impossible as long as seeds
 * @dev differ. (Which is critical to making unpredictable randomness! See the
 * @dev next section.)
 *
 * *****************************************************************************
 * @dev SECURITY CONSIDERATIONS
 *
 * @dev A method with the ability to call your fulfillRandomness method directly
 * @dev could spoof a VOR response with any random value, so it's critical that
 * @dev it cannot be directly called by anything other than this base contract
 * @dev (specifically, by the VORConsumerBase.rawFulfillRandomness method).
 *
 * @dev For your users to trust that your contract's random behavior is free
 * @dev from malicious interference, it's best if you can write it so that all
 * @dev behaviors implied by a VOR response are executed *during* your
 * @dev fulfillRandomness method. If your contract must store the response (or
 * @dev anything derived from it) and use it later, you must ensure that any
 * @dev user-significant behavior which depends on that stored value cannot be
 * @dev manipulated by a subsequent VOR request.
 *
 * @dev Similarly, both miners and the VOR oracle itself have some influence
 * @dev over the order in which VOR responses appear on the blockchain, so if
 * @dev your contract could have multiple VOR requests in flight simultaneously,
 * @dev you must ensure that the order in which the VOR responses arrive cannot
 * @dev be used to manipulate your contract's user-significant behavior.
 *
 * @dev Since the ultimate input to the VOR is mixed with the block hash of the
 * @dev block in which the request is made, user-provided seeds have no impact
 * @dev on its economic security properties. They are only included for API
 * @dev compatability with previous versions of this contract.
 *
 * @dev Since the block hash of the block which contains the requestRandomness
 * @dev call is mixed into the input to the VOR *last*, a sufficiently powerful
 * @dev miner could, in principle, fork the blockchain to evict the block
 * @dev containing the request, forcing the request to be included in a
 * @dev different block with a different hash, and therefore a different input
 * @dev to the VOR. However, such an attack would incur a substantial economic
 * @dev cost. This cost scales with the number of blocks the VOR oracle waits
 * @dev until it calls responds to a request.
 */
abstract contract VORConsumerBase is VORRequestIDBase {
    using SafeMath for uint256;

    /**
     * @notice fulfillRandomness handles the VOR response. Your contract must
     * @notice implement it. See "SECURITY CONSIDERATIONS" above for important
     * @notice principles to keep in mind when implementing your fulfillRandomness
     * @notice method.
     *
     * @dev VORConsumerBase expects its subcontracts to have a method with this
     * @dev signature, and will call it once it has verified the proof
     * @dev associated with the randomness. (It is triggered via a call to
     * @dev rawFulfillRandomness, below.)
     *
     * @param requestId The Id initially returned by requestRandomness
     * @param randomness the VOR output
     */
    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal virtual;

    /**
     * @notice requestRandomness initiates a request for VOR output given _seed
     *
     * @dev The fulfillRandomness method receives the output, once it's provided
     * @dev by the Oracle, and verified by the vorCoordinator.
     *
     * @dev The _keyHash must already be registered with the VORCoordinator, and
     * @dev the _fee must exceed the fee specified during registration of the
     * @dev _keyHash.
     *
     * @dev The _seed parameter is vestigial, and is kept only for API
     * @dev compatibility with older versions. It can't *hurt* to mix in some of
     * @dev your own randomness, here, but it's not necessary because the VOR
     * @dev oracle will mix the hash of the block containing your request into the
     * @dev VOR seed it ultimately uses.
     *
     * @param _keyHash ID of public key against which randomness is generated
     * @param _fee The amount of xFUND to send with the request
     * @param _seed seed mixed into the input of the VOR.
     *
     * @return requestId unique ID for this request
     *
     * @dev The returned requestId can be used to distinguish responses to
     * @dev concurrent requests. It is passed as the first argument to
     * @dev fulfillRandomness.
     */
    function requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) internal returns (bytes32 requestId) {
        xFUND.approve(vorCoordinator, _fee);
        IVORCoordinator(vorCoordinator).randomnessRequest(_keyHash, _seed, _fee, address(this));
        // This is the seed passed to VORCoordinator. The oracle will mix this with
        // the hash of the block containing this request to obtain the seed/input
        // which is finally passed to the VOR cryptographic machinery.
        uint256 vORSeed = makeVORInputSeed(_keyHash, _seed, address(this), nonces[_keyHash]);
        // nonces[_keyHash] must stay in sync with
        // VORCoordinator.nonces[_keyHash][this], which was incremented by the above
        // successful xFUND.transferAndCall (in VORCoordinator.randomnessRequest).
        // This provides protection against the user repeating their input seed,
        // which would result in a predictable/duplicate output, if multiple such
        // requests appeared in the same block.
        nonces[_keyHash] = nonces[_keyHash].add(1);
        return makeRequestId(_keyHash, vORSeed);
    }

    XFundTokenInterface internal immutable xFUND;
    address private immutable vorCoordinator;

    // Nonces for each VOR key from which randomness has been requested.
    //
    // Must stay in sync with VORCoordinator[_keyHash][this]
    /* keyHash */
    /* nonce */
    mapping(bytes32 => uint256) private nonces;

    /**
     * @param _vorCoordinator address of VORCoordinator contract
     * @param _xfund address of xFUND token contract
     */
    constructor(address _vorCoordinator, address _xfund) public {
        vorCoordinator = _vorCoordinator;
        xFUND = XFundTokenInterface(_xfund);
    }

    // rawFulfillRandomness is called by VORCoordinator when it receives a valid VOR
    // proof. rawFulfillRandomness then calls fulfillRandomness, after validating
    // the origin of the call
    function rawFulfillRandomness(bytes32 requestId, uint256 randomness) external {
        require(msg.sender == vorCoordinator, "Only VORCoordinator can fulfill");
        fulfillRandomness(requestId, randomness);
    }
}