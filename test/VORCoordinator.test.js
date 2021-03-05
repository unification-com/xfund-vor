const { BN, expectEvent, expectRevert, constants } = require('@openzeppelin/test-helpers');

const chai = require('chai');
chai.use(require('chai-as-promised'));
const { expect } = chai;

const MockERC20 = artifacts.require('MockERC20');
const BlockhashStore = artifacts.require('BlockhashStore');
const VORCoordinator = artifacts.require('VORCoordinator');

contract('VORCoordinator', ([owner, oracle, contractSender]) => {
    beforeEach(async () => {
        this.fee = web3.utils.toWei('0.1', 'ether');

        this.xFund = await MockERC20.new('xFUND', 'xFUND', web3.utils.toWei('1000000000', 'ether'), { from: owner });
        this.blockhashStore = await BlockhashStore.new({ from: owner });
        this.vorCoordinator = await VORCoordinator.new(this.xFund.address, this.blockhashStore.address, { from: owner });
    });

    it('returns the correct serviceAgreements', async () => {
        const publicProvingKey = [new BN('0'), new BN('0')];
        const keyHash = await this.vorCoordinator.hashOfKey(publicProvingKey);

        const newServiceAgreement = await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey);
        expectEvent(newServiceAgreement, 'NewServiceAgreement', { keyHash, fee: this.fee });

        const serviceAgreements = await this.vorCoordinator.serviceAgreements.call(keyHash);
        expect(serviceAgreements.vOROracle).to.be.equal(oracle);
        expect(serviceAgreements.fee).to.be.bignumber.equal(new BN(this.fee));
    });

    it('registerProvingKey rejects', async () => {
        const newFee = web3.utils.toWei('10000000000', 'ether');
        const publicProvingKey = [new BN('0'), new BN('0')];

        await expectRevert(
            this.vorCoordinator.registerProvingKey(this.fee, constants.ZERO_ADDRESS, publicProvingKey),
            '_oracle must not be 0x0'
        );

        await expectRevert(
            this.vorCoordinator.registerProvingKey(newFee, oracle, publicProvingKey),
            `you can't charge more than all the xFUND in the world, greedy`
        );
    });

    it('returns the correct fee', async () => {
        const newFee = web3.utils.toWei('0.2', 'ether');
        const publicProvingKey = [new BN('0'), new BN('0')];

        const keyHash = await this.vorCoordinator.hashOfKey(publicProvingKey);

        const newServiceAgreement = await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey);
        expectEvent(newServiceAgreement, 'NewServiceAgreement', { keyHash, fee: this.fee });

        const changeFee = await this.vorCoordinator.changeFee(publicProvingKey, newFee, { from: oracle });
        expectEvent(changeFee, 'ChangeFee', { keyHash, fee: newFee });

        await expectRevert(
            this.vorCoordinator.changeFee(publicProvingKey, web3.utils.toWei('10000000000', 'ether')),
            `only oracle can change the commission`
        );

        await expectRevert(
            this.vorCoordinator.changeFee(publicProvingKey, web3.utils.toWei('10000000000', 'ether'), { from: oracle }),
            `you can't charge more than all the xFUND in the world, greedy`
        );
    });

    it('returns the correct callbacks', async () => {
        await this.xFund.transfer(contractSender, this.fee, { from: owner });
        await this.xFund.approve(this.vorCoordinator.address, this.fee, { from: contractSender });

        const keyHash = web3.utils.fromAscii('keyHash');
        const seed = 12345;

        const randomnessRequest = await this.vorCoordinator.randomnessRequest(keyHash, seed, this.fee, { from: contractSender });
        expectEvent(
            randomnessRequest,
            'RandomnessRequest',
            { keyHash: keyHash.padEnd(66, '0'), fee: this.fee, sender: contractSender }
        );

        const requestId = randomnessRequest.logs[0].args.requestID;
        
        const callbacks = await this.vorCoordinator.callbacks.call(requestId);
        expect(callbacks.callbackContract).to.be.equal(contractSender);
        expect(callbacks.randomnessFee).to.be.bignumber.equal(new BN(this.fee));
    });
});