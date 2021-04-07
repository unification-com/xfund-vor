# BlockhashStore



## Functions:
- [`storeEarliest()`](#BlockhashStore-storeEarliest--)
- [`getBlockhash(uint256 n)`](#BlockhashStore-getBlockhash-uint256-)
- [`store(uint256 n)`](#BlockhashStore-store-uint256-)
- [`storeVerifyHeader(uint256 n, bytes header)`](#BlockhashStore-storeVerifyHeader-uint256-bytes-)



<a name="BlockhashStore-storeEarliest--"></a>
### Function `storeEarliest()`
No description
<a name="BlockhashStore-getBlockhash-uint256-"></a>
### Function `getBlockhash(uint256 n) -> bytes32`
No description
#### Parameters:
- `n`: the number of the block whose blockhash should be returned
<a name="BlockhashStore-store-uint256-"></a>
### Function `store(uint256 n)`
No description
#### Parameters:
- `n`: the number of the block whose blockhash should be stored
<a name="BlockhashStore-storeVerifyHeader-uint256-bytes-"></a>
### Function `storeVerifyHeader(uint256 n, bytes header)`
No description
#### Parameters:
- `n`: the number of the block whose blockhash should be stored

- `header`: the rlp-encoded blockheader of block n+1. We verify its correctness by checking
  that it hashes to a stored blockhash, and then extract parentHash to get the n-th blockhash.


