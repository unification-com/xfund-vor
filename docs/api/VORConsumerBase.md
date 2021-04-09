# 
PURPOSE

Reggie the Random Oracle (not his real job) wants to provide randomness
to Vera the verifier in such a way that Vera can be sure he's not
making his output up to suit himself. Reggie provides Vera a public key
to which he knows the secret key. Each time Vera provides a seed to
Reggie, he gives back a value which is computed completely
deterministically from the seed and the secret key.

Reggie provides a proof by which Vera can verify that the output was
correctly computed once Reggie tells it to her, but without that proof,
the output is indistinguishable to her from a uniform random sample
from the output space.

The purpose of this contract is to make it easy for unrelated contracts
to talk to Vera the verifier about the work Reggie is doing, to provide
simple access to a verifiable source of randomness.
*****************************************************************************
USAGE

Calling contracts must inherit from VORConsumerBase, and can
initialize VORConsumerBase's attributes in their constructor as
shown:

  contract VORConsumer {
    constuctor(<other arguments>, address _vorCoordinator, address _xfund)
      VORConsumerBase(_vorCoordinator, _xfund) public {
        <initialization with other arguments goes here>
      }
  }

The oracle will have given you an ID for the VOR keypair they have
committed to (let's call it keyHash), and have told you the minimum xFUND
price for VOR service. Make sure your contract has sufficient xFUND, and
call requestRandomness(keyHash, fee, seed), where seed is the input you
want to generate randomness from.

Once the VORCoordinator has received and validated the oracle's response
to your request, it will call your contract's fulfillRandomness method.

The randomness argument to fulfillRandomness is the actual random value
generated from your seed.

The requestId argument is generated from the keyHash and the seed by
makeRequestId(keyHash, seed). If your contract could have concurrent
requests open, you can use the requestId to track which seed is
associated with which randomness. See VORRequestIDBase.sol for more
details. (See "SECURITY CONSIDERATIONS" for principles to keep in mind,
if your contract could have multiple requests in flight simultaneously.)

Colliding `requestId`s are cryptographically impossible as long as seeds
differ. (Which is critical to making unpredictable randomness! See the
next section.)

*****************************************************************************
SECURITY CONSIDERATIONS

A method with the ability to call your fulfillRandomness method directly
could spoof a VOR response with any random value, so it's critical that
it cannot be directly called by anything other than this base contract
(specifically, by the VORConsumerBase.rawFulfillRandomness method).

For your users to trust that your contract's random behavior is free
from malicious interference, it's best if you can write it so that all
behaviors implied by a VOR response are executed *during* your
fulfillRandomness method. If your contract must store the response (or
anything derived from it) and use it later, you must ensure that any
user-significant behavior which depends on that stored value cannot be
manipulated by a subsequent VOR request.

Similarly, both miners and the VOR oracle itself have some influence
over the order in which VOR responses appear on the blockchain, so if
your contract could have multiple VOR requests in flight simultaneously,
you must ensure that the order in which the VOR responses arrive cannot
be used to manipulate your contract's user-significant behavior.

Since the ultimate input to the VOR is mixed with the block hash of the
block in which the request is made, user-provided seeds have no impact
on its economic security properties. They are only included for API
compatability with previous versions of this contract.

Since the block hash of the block which contains the requestRandomness
call is mixed into the input to the VOR *last*, a sufficiently powerful
miner could, in principle, fork the blockchain to evict the block
containing the request, forcing the request to be included in a
different block with a different hash, and therefore a different input
to the VOR. However, such an attack would incur a substantial economic
cost. This cost scales with the number of blocks the VOR oracle waits
until it calls responds to a request.

## Functions:
- [`constructor(address _vorCoordinator, address _xfund)`](#VORConsumerBase-constructor-address-address-)
- [`rawFulfillRandomness(bytes32 requestId, uint256 randomness)`](#VORConsumerBase-rawFulfillRandomness-bytes32-uint256-)
- [`receive()`](#VORConsumerBase-receive--)



<a name="VORConsumerBase-constructor-address-address-"></a>
### Function `constructor(address _vorCoordinator, address _xfund)`
No description
#### Parameters:
- `_vorCoordinator`: address of VORCoordinator contract

- `_xfund`: address of xFUND token contract
<a name="VORConsumerBase-rawFulfillRandomness-bytes32-uint256-"></a>
### Function `rawFulfillRandomness(bytes32 requestId, uint256 randomness)`
No description
<a name="VORConsumerBase-receive--"></a>
### Function `receive()`
No description


