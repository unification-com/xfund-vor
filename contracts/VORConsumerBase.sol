// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;

import "./vendor/VORSafeMath.sol";
import "./interfaces/IERC20_Ex.sol";
import "./interfaces/IVORCoordinator.sol";
import "./VORRequestIDBase.sol";

/**
 * @title VORConsumerBase
 * @notice Interface for contracts using VOR randomness
 * @dev PURPOSE
 *
 * @dev Reggie the Random Oracle (not his real job) wants to provide randomness
 * to Vera the verifier in such a way that Vera can be sure he's not
 * making his output up to suit himself. Reggie provides Vera a public key
 * to which he knows the secret key. Each time Vera provides a seed to
 * Reggie, he gives back a value which is computed completely
 * deterministically from the seed and the secret key.
 *
 * @dev Reggie provides a proof by which Vera can verify that the output was
 * correctly computed once Reggie tells it to her, but without that proof,
 * the output is indistinguishable to her from a uniform random sample
 * from the output space.
 *
 * @dev The purpose of this contract is to make it easy for unrelated contracts
 * to talk to Vera the verifier about the work Reggie is doing, to provide
 * simple access to a verifiable source of randomness.
 *
 * @dev USAGE
 *
 * @dev Calling contracts must inherit from VORConsumerBase, and can
 * initialize VORConsumerBase's attributes in their constructor as
 * shown:
 *
 * ```
 *   contract VORConsumer {
 *     constuctor(<other arguments>, address _vorCoordinator, address _xfund)
 *       VORConsumerBase(_vorCoordinator, _xfund) public {
 *         <initialization with other arguments goes here>
 *       }
 *   }
 * ```
 * @dev The oracle will have given you an ID for the VOR keypair they have
 * committed to (let's call it keyHash), and have told you the minimum xFUND
 * price for VOR service. Make sure your contract has sufficient xFUND, and
 * call requestRandomness(keyHash, fee, seed), where seed is the input you
 * want to generate randomness from.
 *
 * @dev Once the VORCoordinator has received and validated the oracle's response
 * to your request, it will call your contract's fulfillRandomness method.
 *
 * @dev The randomness argument to fulfillRandomness is the actual random value
 * generated from your seed.
 *
 * @dev The requestId argument is generated from the keyHash and the seed by
 * makeRequestId(keyHash, seed). If your contract could have concurrent
 * requests open, you can use the requestId to track which seed is
 * associated with which randomness. See VORRequestIDBase.sol for more
 * details. (See "SECURITY CONSIDERATIONS" for principles to keep in mind,
 * if your contract could have multiple requests in flight simultaneously.)
 *
 * @dev Colliding `requestId`s are cryptographically impossible as long as seeds
 * differ. (Which is critical to making unpredictable randomness! See the
 * next section.)
 *
 * @dev SECURITY CONSIDERATIONS
 *
 * @dev A method with the ability to call your fulfillRandomness method directly
 * could spoof a VOR response with any random value, so it's critical that
 * it cannot be directly called by anything other than this base contract
 * (specifically, by the VORConsumerBase.rawFulfillRandomness method).
 *
 * @dev For your users to trust that your contract's random behavior is free
 * from malicious interference, it's best if you can write it so that all
 * behaviors implied by a VOR response are executed *during* your
 * fulfillRandomness method. If your contract must store the response (or
 * anything derived from it) and use it later, you must ensure that any
 * user-significant behavior which depends on that stored value cannot be
 * manipulated by a subsequent VOR request.
 *
 * @dev Similarly, both miners and the VOR oracle itself have some influence
 * over the order in which VOR responses appear on the blockchain, so if
 * your contract could have multiple VOR requests in flight simultaneously,
 * you must ensure that the order in which the VOR responses arrive cannot
 * be used to manipulate your contract's user-significant behavior.
 *
 * @dev Since the ultimate input to the VOR is mixed with the block hash of the
 * block in which the request is made, user-provided seeds have no impact
 * on its economic security properties. They are only included for API
 * compatability with previous versions of this contract.
 *
 * @dev Since the block hash of the block which contains the requestRandomness
 * call is mixed into the input to the VOR *last*, a sufficiently powerful
 * miner could, in principle, fork the blockchain to evict the block
 * containing the request, forcing the request to be included in a
 * different block with a different hash, and therefore a different input
 * to the VOR. However, such an attack would incur a substantial economic
 * cost. This cost scales with the number of blocks the VOR oracle waits
 * until it calls responds to a request.
 */
abstract contract VORConsumerBase is VORRequestIDBase {
    using VORSafeMath for uint256;

    /**
     * @notice fulfillRandomness handles the VOR response. Your contract must
     * @notice implement it. See "SECURITY CONSIDERATIONS" above for important
     * @notice principles to keep in mind when implementing your fulfillRandomness
     * @notice method.
     *
     * @dev VORConsumerBase expects its subcontracts to have a method with this
     * signature, and will call it once it has verified the proof
     * associated with the randomness. (It is triggered via a call to
     * rawFulfillRandomness, below.)
     *
     * @param requestId The Id initially returned by requestRandomness
     * @param randomness the VOR output
     */
    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal virtual;

    /**
     * @notice requestRandomness initiates a request for VOR output given _seed
     *
     * @dev The fulfillRandomness method receives the output, once it's provided
     * by the Oracle, and verified by the vorCoordinator.
     *
     * @dev The _keyHash must already be registered with the VORCoordinator, and
     * the _fee must exceed the fee specified during registration of the
     * _keyHash.
     *
     * @dev The _seed parameter is vestigial, and is kept only for API
     * compatibility with older versions. It can't *hurt* to mix in some of
     * your own randomness, here, but it's not necessary because the VOR
     * oracle will mix the hash of the block containing your request into the
     * VOR seed it ultimately uses.
     *
     * @param _keyHash ID of public key against which randomness is generated
     * @param _fee The amount of xFUND to send with the request
     * @param _seed seed mixed into the input of the VOR.
     *
     * @return requestId unique ID for this request
     *
     * The returned requestId can be used to distinguish responses to
     * concurrent requests. It is passed as the first argument to
     * fulfillRandomness.
     */
    function requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) internal returns (bytes32 requestId) {
        IVORCoordinator(vorCoordinator).randomnessRequest(_keyHash, _seed, _fee);
        // This is the seed passed to VORCoordinator. The oracle will mix this with
        // the hash of the block containing this request to obtain the seed/input
        // which is finally passed to the VOR cryptographic machinery.
        uint256 vORSeed = makeVORInputSeed(_keyHash, _seed, address(this), nonces[_keyHash]);
        // nonces[_keyHash] must stay in sync with
        // VORCoordinator.nonces[_keyHash][this], which was incremented by the above
        // successful VORCoordinator.randomnessRequest.
        // This provides protection against the user repeating their input seed,
        // which would result in a predictable/duplicate output, if multiple such
        // requests appeared in the same block.
        nonces[_keyHash] = nonces[_keyHash].safeAdd(1);
        return makeRequestId(_keyHash, vORSeed);
    }

    /**
     * @notice _increaseVorCoordinatorAllowance is a helper function to increase token allowance for
     * the VORCoordinator
     * Allows this contract to increase the xFUND allowance for the VORCoordinator contract
     * enabling it to pay request fees on behalf of this contract.
     * NOTE: it is hightly recommended to wrap this around a function that uses,
     * for example, OpenZeppelin's onlyOwner modifier
     *
     * @param _amount uint256 amount to increase allowance by
     */
    function _increaseVorCoordinatorAllowance(uint256 _amount) internal returns (bool) {
        require(xFUND.increaseAllowance(vorCoordinator, _amount), "failed to increase allowance");
        return true;
    }

    /**
     * @notice _setVORCoordinator is a helper function to enable setting the VORCoordinator address
     * NOTE: it is hightly recommended to wrap this around a function that uses,
     * for example, OpenZeppelin's onlyOwner modifier
     *
     * @param _vorCoordinator address new VORCoordinator address
     */
    function _setVORCoordinator(address _vorCoordinator) internal {
        vorCoordinator = _vorCoordinator;
    }

    address internal immutable xFUNDAddress;
    IERC20_Ex internal immutable xFUND;
    address internal vorCoordinator;

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
    constructor(address _vorCoordinator, address _xfund) internal {
        vorCoordinator = _vorCoordinator;
        xFUNDAddress = _xfund;
        xFUND = IERC20_Ex(_xfund);
    }

    /**
     * @notice rawFulfillRandomness is called by VORCoordinator when it receives a valid VOR
     * proof. rawFulfillRandomness then calls fulfillRandomness, after validating
     * the origin of the call
     */
    function rawFulfillRandomness(bytes32 requestId, uint256 randomness) external {
        require(msg.sender == vorCoordinator, "Only VORCoordinator can fulfill");
        fulfillRandomness(requestId, randomness);
    }
}
