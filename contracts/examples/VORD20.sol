// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../VORConsumerBase.sol";

/**
 * @title VORD20
 * @notice A VOR consumer which uses randomness to mimic the rolling
 * of a 20 sided die
 * This is only an example implementation and not necessarily suitable for mainnet.
 */
contract VORD20 is Ownable, VORConsumerBase {
    using SafeMath for uint256;

    uint256 private constant _ROLL_IN_PROGRESS = 42;

    bytes32 private _sKeyHash;
    uint256 private _sFee;
    mapping(bytes32 => address) private _sRollers;
    mapping(address => uint256) private _sResults;

    event DiceRolled(bytes32 indexed requestId, address indexed roller);
    event DiceLanded(bytes32 indexed requestId, uint256 indexed result);

    /**
     * @notice Constructor inherits VORConsumerBase
     *
     *
     * @param vorCoordinator address of the VOR Coordinator
     * @param xfund address of the xFUND token
     * @param keyHash bytes32 representing the hash of the VOR provider
     * @param fee uint256 fee to pay the VOR oracle
     */
    constructor(address vorCoordinator, address xfund, bytes32 keyHash, uint256 fee)
        public
        VORConsumerBase(vorCoordinator, xfund)
    {
        _sKeyHash = keyHash;
        _sFee = fee;
    }

    /**
     * @notice Requests randomness from a user-provided seed
     * Warning: if the VOR response is delayed, avoid calling requestRandomness repeatedly
     * as that would give miners/VOR operators latitude about which VOR response arrives first.
     * You must review your implementation details with extreme care.
     *
     * @param userProvidedSeed uint256 unpredictable seed
     * @param roller address of the roller
     */
    function rollDice(uint256 userProvidedSeed, address roller) public onlyOwner returns (bytes32 requestId) {
        require(xFUND.balanceOf(address(this)) >= _sFee, "Not enough xFUND to pay fee");
        require(_sResults[roller] == 0, "Already rolled");
        requestId = requestRandomness(_sKeyHash, _sFee, userProvidedSeed);
        _sRollers[requestId] = roller;
        _sResults[roller] = _ROLL_IN_PROGRESS;
        emit DiceRolled(requestId, roller);
    }

    /**
     * @notice Callback function used by VOR Coordinator to return the random number
     * to this contract.
     * Some action on the contract state should be taken here, like storing the result.
     * WARNING: take care to avoid having multiple VOR requests in flight if their order of arrival would result
     * in contract states with different outcomes. Otherwise miners or the VOR operator would could take advantage
     * by controlling the order.
     * The VOR Coordinator will only send this function verified responses, and the parent VORConsumerBase
     * contract ensures that this method only receives randomness from the designated VORCoordinator.
     *
     * @param requestId bytes32
     * @param randomness The random result returned by the oracle
     */
    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal override {
        uint256 d20Value = randomness.mod(20).add(1);
        _sResults[_sRollers[requestId]] = d20Value;
        emit DiceLanded(requestId, d20Value);
    }

    /**
     * @notice Example wrapper function for the VORConsumerBase increaseVorCoordinatorAllowance function.
     * Wrapped around an Ownable modifier to ensure only the contract owner can call it.
     * Allows contract owner to increase the xFUND allowance for the VORCoordinator contract
     * enabling it to pay request fees on behalf of this contract's owner.
     * NOTE: This contract must have an xFUND balance in order to request randomness
     *
     * @param _amount uint256 amount to increase allowance by
     */
    function increaseVorAllowance(uint256 _amount) external onlyOwner {
        _increaseVorCoordinatorAllowance(_amount);
    }

    /**
     * @notice Example wrapper function for the VORConsumerBase withdrawXFUND function.
     * Wrapped around an Ownable modifier to ensure only the contract owner can call it.
     * Allows contract owner to withdraw any xFUND currently held by this contract
     */
    function withdrawToken(address to, uint256 value) external onlyOwner {
        require(xFUND.transfer(to, value), "Not enough xFUND");
    }

    /**
     * @notice Get the house assigned to the player once the address has rolled
     * @param player address
     * @return house as a string
     */
    function house(address player) public view returns (string memory) {
        require(_sResults[player] != 0, "Dice not rolled");
        require(_sResults[player] != _ROLL_IN_PROGRESS, "Roll in progress");
        return getHouseName(_sResults[player]);
    }

    /**
     * @notice Set the key hash for the oracle
     *
     * @param keyHash bytes32
     */
    function setKeyHash(bytes32 keyHash) public onlyOwner {
        _sKeyHash = keyHash;
    }

    /**
     * @notice Get the current key hash
     *
     * @return bytes32
     */
    function keyHash() public view returns (bytes32) {
        return _sKeyHash;
    }

    /**
     * @notice Set the oracle fee for requesting randomness
     *
     * @param fee uint256
     */
    function setFee(uint256 fee) public onlyOwner {
        _sFee = fee;
    }

    /**
     * @notice Get the current fee
     *
     * @return uint256
     */
    function fee() public view returns (uint256) {
        return _sFee;
    }

    /**
     * @notice Get the house namne from the id
     * @param id uint256
     * @return house name string
     */
    function getHouseName(uint256 id) private pure returns (string memory) {
        string[20] memory houseNames = [
            "Targaryen",
            "Lannister",
            "Stark",
            "Tyrell",
            "Baratheon",
            "Martell",
            "Tully",
            "Bolton",
            "Greyjoy",
            "Arryn",
            "Frey",
            "Mormont",
            "Tarley",
            "Dayne",
            "Umber",
            "Valeryon",
            "Manderly",
            "Clegane",
            "Glover",
            "Karstark"
        ];
        return houseNames[id.sub(1)];
    }
}