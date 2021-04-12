# VOR


## Functions:
- [`bigModExp(uint256 base, uint256 exponent) internal`](#VOR-bigModExp-uint256-uint256-)
- [`squareRoot(uint256 x) internal`](#VOR-squareRoot-uint256-)
- [`ySquared(uint256 x) internal`](#VOR-ySquared-uint256-)
- [`isOnCurve(uint256[2] p) internal`](#VOR-isOnCurve-uint256-2--)
- [`fieldHash(bytes b) internal`](#VOR-fieldHash-bytes-)
- [`newCandidateSecp256k1Point(bytes b) internal`](#VOR-newCandidateSecp256k1Point-bytes-)
- [`hashToCurve(uint256[2] pk, uint256 input) internal`](#VOR-hashToCurve-uint256-2--uint256-)
- [`ecmulVerify(uint256[2] multiplicand, uint256 scalar, uint256[2] product) internal`](#VOR-ecmulVerify-uint256-2--uint256-uint256-2--)
- [`projectiveSub(uint256 x1, uint256 z1, uint256 x2, uint256 z2) internal`](#VOR-projectiveSub-uint256-uint256-uint256-uint256-)
- [`projectiveMul(uint256 x1, uint256 z1, uint256 x2, uint256 z2) internal`](#VOR-projectiveMul-uint256-uint256-uint256-uint256-)
- [`projectiveECAdd(uint256 px, uint256 py, uint256 qx, uint256 qy) internal`](#VOR-projectiveECAdd-uint256-uint256-uint256-uint256-)
- [`affineECAdd(uint256[2] p1, uint256[2] p2, uint256 invZ) internal`](#VOR-affineECAdd-uint256-2--uint256-2--uint256-)
- [`verifyLinearCombinationWithGenerator(uint256 c, uint256[2] p, uint256 s, address lcWitness) internal`](#VOR-verifyLinearCombinationWithGenerator-uint256-uint256-2--uint256-address-)
- [`linearCombination(uint256 c, uint256[2] p1, uint256[2] cp1Witness, uint256 s, uint256[2] p2, uint256[2] sp2Witness, uint256 zInv) internal`](#VOR-linearCombination-uint256-uint256-2--uint256-2--uint256-uint256-2--uint256-2--uint256-)
- [`scalarFromCurvePoints(uint256[2] hash, uint256[2] pk, uint256[2] gamma, address uWitness, uint256[2] v) internal`](#VOR-scalarFromCurvePoints-uint256-2--uint256-2--uint256-2--address-uint256-2--)
- [`verifyVORProof(uint256[2] pk, uint256[2] gamma, uint256 c, uint256 s, uint256 seed, address uWitness, uint256[2] cGammaWitness, uint256[2] sHashWitness, uint256 zInv) internal`](#VOR-verifyVORProof-uint256-2--uint256-2--uint256-uint256-uint256-address-uint256-2--uint256-2--uint256-)
- [`randomValueFromVORProof(bytes proof) internal`](#VOR-randomValueFromVORProof-bytes-)



<a name="VOR-bigModExp-uint256-uint256-"></a>
### Function `bigModExp(uint256 base, uint256 exponent) internal  -> uint256 exponentiation`
No description
<a name="VOR-squareRoot-uint256-"></a>
### Function `squareRoot(uint256 x) internal  -> uint256`
No description
<a name="VOR-ySquared-uint256-"></a>
### Function `ySquared(uint256 x) internal  -> uint256`
No description
<a name="VOR-isOnCurve-uint256-2--"></a>
### Function `isOnCurve(uint256[2] p) internal  -> bool`
No description
<a name="VOR-fieldHash-bytes-"></a>
### Function `fieldHash(bytes b) internal  -> uint256 x_`
No description
<a name="VOR-newCandidateSecp256k1Point-bytes-"></a>
### Function `newCandidateSecp256k1Point(bytes b) internal  -> uint256[2] p`
No description
<a name="VOR-hashToCurve-uint256-2--uint256-"></a>
### Function `hashToCurve(uint256[2] pk, uint256 input) internal  -> uint256[2] rv`
No description
<a name="VOR-ecmulVerify-uint256-2--uint256-uint256-2--"></a>
### Function `ecmulVerify(uint256[2] multiplicand, uint256 scalar, uint256[2] product) internal  -> bool verifies`
Based on Vitalik Buterin's idea in ethresear.ch post cited below.


#### Parameters:
- `secp256k1`: point

- `zero`: GF(GROUP_ORDER) scalar

- `secp256k1`: expected to be multiplier * multiplicand

#### Return Values:
- verifies true iff product==scalar*multiplicand, with cryptographically high probability
/
    func
<a name="VOR-projectiveSub-uint256-uint256-uint256-uint256-"></a>
### Function `projectiveSub(uint256 x1, uint256 z1, uint256 x2, uint256 z2) internal  -> uint256 x3, uint256 z3`
No description
<a name="VOR-projectiveMul-uint256-uint256-uint256-uint256-"></a>
### Function `projectiveMul(uint256 x1, uint256 z1, uint256 x2, uint256 z2) internal  -> uint256 x3, uint256 z3`
No description
<a name="VOR-projectiveECAdd-uint256-uint256-uint256-uint256-"></a>
### Function `projectiveECAdd(uint256 px, uint256 py, uint256 qx, uint256 qy) internal  -> uint256 sx, uint256 sy, uint256 sz`
No description
<a name="VOR-affineECAdd-uint256-2--uint256-2--uint256-"></a>
### Function `affineECAdd(uint256[2] p1, uint256[2] p2, uint256 invZ) internal  -> uint256[2]`
No description
<a name="VOR-verifyLinearCombinationWithGenerator-uint256-uint256-2--uint256-address-"></a>
### Function `verifyLinearCombinationWithGenerator(uint256 c, uint256[2] p, uint256 s, address lcWitness) internal  -> bool`
No description
<a name="VOR-linearCombination-uint256-uint256-2--uint256-2--uint256-uint256-2--uint256-2--uint256-"></a>
### Function `linearCombination(uint256 c, uint256[2] p1, uint256[2] cp1Witness, uint256 s, uint256[2] p2, uint256[2] sp2Witness, uint256 zInv) internal  -> uint256[2]`
No description
<a name="VOR-scalarFromCurvePoints-uint256-2--uint256-2--uint256-2--address-uint256-2--"></a>
### Function `scalarFromCurvePoints(uint256[2] hash, uint256[2] pk, uint256[2] gamma, address uWitness, uint256[2] v) internal  -> uint256 s`
No description
<a name="VOR-verifyVORProof-uint256-2--uint256-2--uint256-uint256-uint256-address-uint256-2--uint256-2--uint256-"></a>
### Function `verifyVORProof(uint256[2] pk, uint256[2] gamma, uint256 c, uint256 s, uint256 seed, address uWitness, uint256[2] cGammaWitness, uint256[2] sHashWitness, uint256 zInv) internal `
No description
<a name="VOR-randomValueFromVORProof-bytes-"></a>
### Function `randomValueFromVORProof(bytes proof) internal  -> uint256 output`
No description


