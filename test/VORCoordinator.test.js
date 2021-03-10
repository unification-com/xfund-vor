const { BN, expectEvent, expectRevert, constants } = require('@openzeppelin/test-helpers');

const chai = require('chai');
chai.use(require('chai-as-promised'));
const { expect } = chai;

const MockERC20 = artifacts.require('MockERC20');
const BlockhashStore = artifacts.require('BlockhashStore');
const VORCoordinator = artifacts.require('VORCoordinator');
const VORD20 = artifacts.require('VORD20');

contract('VORCoordinator', ([owner, oracle, alice]) => {
    beforeEach(async () => {
        this.keyHash = web3.utils.fromAscii('keyHash');
        this.fee = web3.utils.toWei('0.1', 'ether');

        this.deposit = web3.utils.toWei('1', 'ether');

        this.xFund = await MockERC20.new('xFUND', 'xFUND', web3.utils.toWei('1000000000', 'ether'), { from: owner });
        this.blockhashStore = await BlockhashStore.new({ from: owner });
        this.vorCoordinator = await VORCoordinator.new(this.xFund.address, this.blockhashStore.address, { from: owner });

        this.vorD20 = await VORD20.new(this.vorCoordinator.address, this.xFund.address, this.keyHash, this.fee, { from: owner });
        await this.xFund.transfer(this.vorD20.address, this.deposit, { from: owner });
    });

    it('returns the correct serviceAgreements', async () => {
        const publicProvingKey = [new BN('0'), new BN('0')];
        const keyHash = await this.vorCoordinator.hashOfKey(publicProvingKey);

        const newServiceAgreement = await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey, true);
        expectEvent(newServiceAgreement, 'NewServiceAgreement', { keyHash, fee: this.fee });

        const serviceAgreements = await this.vorCoordinator.serviceAgreements.call(keyHash);
        expect(serviceAgreements.vOROracle).to.be.equal(oracle);
        expect(serviceAgreements.fee).to.be.bignumber.equal(new BN(this.fee));
        expect(serviceAgreements.providerPaysGas).to.be.equal(true);
    });

    it('registerProvingKey rejects', async () => {
        const newFee = web3.utils.toWei('10000000000', 'ether');
        const publicProvingKey = [new BN('0'), new BN('0')];

        await expectRevert(
            this.vorCoordinator.registerProvingKey(this.fee, constants.ZERO_ADDRESS, publicProvingKey, true),
            '_oracle must not be 0x0'
        );

        await expectRevert(
            this.vorCoordinator.registerProvingKey(newFee, oracle, publicProvingKey, true),
            `you can't charge more than all the xFUND in the world, greedy`
        );

        await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey, true);

        await expectRevert(
            this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey, true),
            `please register a new key`
        );
    });

    it('returns the correct fee', async () => {
        const newFee = web3.utils.toWei('0.2', 'ether');
        const publicProvingKey = [new BN('0'), new BN('0')];

        const keyHash = await this.vorCoordinator.hashOfKey(publicProvingKey);

        const newServiceAgreement = await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey, true);
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
        const seed = 12345;
        const rollResult = await this.vorD20.rollDice(seed, alice);

        const requestId = rollResult.logs[0].args.requestId;
        
        const callbacks = await this.vorCoordinator.callbacks.call(requestId);
        expect(callbacks.callbackContract).to.be.equal(this.vorD20.address);
        expect(callbacks.randomnessFee).to.be.bignumber.equal(new BN(this.fee));
    });

    it('randomnessRequest reject', async () => {
        const seed = 12345;

        await expectRevert(
            this.vorCoordinator.randomnessRequest(this.keyHash, seed, this.fee),
            `request can only be made by a contract`
        );
    });

    it('returns the gas top up', async () => {
        const oldGasTopUp = await this.vorCoordinator.getGasTopUpLimit();
        expect(oldGasTopUp).to.be.bignumber.equal(new BN(web3.utils.toWei('1', 'ether')));

        await this.vorCoordinator.setGasTopUpLimit(web3.utils.toWei('2', 'ether'), { from: owner });
        const gasTopUp = await this.vorCoordinator.getGasTopUpLimit();
        expect(gasTopUp).to.be.bignumber.equal(new BN(web3.utils.toWei('2', 'ether')));
    });

    it('setGasTopUpLimit reject', async () => {
        await expectRevert(
            this.vorCoordinator.setGasTopUpLimit(web3.utils.toWei('2', 'ether'), { from: oracle }),
            `Ownable: caller is not the owner`
        );

        await expectRevert(
            this.vorCoordinator.setGasTopUpLimit(0, { from: owner }),
            `_gasTopUpLimit must be > 0`
        );
    });

    it('returns the correct provider pays gas', async () => {
        const publicProvingKey = [new BN('0'), new BN('0')];
        const keyHash = await this.vorCoordinator.hashOfKey(publicProvingKey);

        await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey, true);

        let serviceAgreements = await this.vorCoordinator.serviceAgreements.call(keyHash);
        expect(serviceAgreements.providerPaysGas).to.be.equal(true);
        
        await this.vorCoordinator.setProviderPaysGas(publicProvingKey, false, { from: oracle });
        serviceAgreements = await this.vorCoordinator.serviceAgreements.call(keyHash);
        expect(serviceAgreements.providerPaysGas).to.be.equal(false);
    });

    it('setProviderPaysGas reject', async () => {
        const publicProvingKey = [new BN('0'), new BN('0')];
        await this.vorCoordinator.registerProvingKey(this.fee, oracle, publicProvingKey, true);

        await expectRevert(
            this.vorCoordinator.setProviderPaysGas(publicProvingKey, false, { from: owner }),
            `only oracle can change who will pay gas`
        );
    });

    it('returns the correct total gas deposits', async () => {
        const actualBalance = await web3.eth.getBalance(alice);
        expect(actualBalance).to.be.bignumber.equal(new BN(web3.utils.toWei('100', 'ether')));

        await web3.eth.sendTransaction({ from: alice, to: this.vorD20.address, value: web3.utils.toWei('1', 'ether') });

        await this.vorD20.topUpGas(web3.utils.toWei('1', 'ether'));
        const totalGasDeposits = await this.vorCoordinator.getTotalGasDeposits();
        expect(totalGasDeposits).to.be.bignumber.equal(new BN(web3.utils.toWei('1', 'ether')));
    });

    it('topUpGas reject', async () => {
        await expectRevert(
            this.vorCoordinator.topUpGas(oracle, { from: owner, value: web3.utils.toWei('1', 'ether') }),
            `only a contract can top up gas`
        );
    });
});