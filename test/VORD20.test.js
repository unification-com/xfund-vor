const { BN, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');

const chai = require('chai');
chai.use(require('chai-as-promised'));
const { expect } = chai;

const MockERC20 = artifacts.require('MockERC20');
const VORD20 = artifacts.require('VORD20');
const VORCoordinator = artifacts.require('VORCoordinatorMock');
const BlockhashStore = artifacts.require('BlockhashStore');

contract('VORD20', ([owner, alice, dummy]) => {
    beforeEach(async () => {
        this.fee = web3.utils.toWei('0.1', 'ether');
        this.keyHash = web3.utils.fromAscii('keyHash');

        this.deposit = web3.utils.toWei('1', 'ether');
        this.seed = 12345;

        this.xFund = await MockERC20.new('xFUND', 'xFUND', web3.utils.toWei('1000000000', 'ether'), { from: owner });
        this.blockhashStore = await BlockhashStore.new({ from: owner });
        this.vorCoordinator = await VORCoordinator.new({ from: owner });
        this.vorD20 = await VORD20.new(this.vorCoordinator.address, this.xFund.address, this.keyHash, this.fee, { from: owner });

        await this.xFund.transfer(this.vorD20.address, this.deposit, { from: owner });
    });

    describe('#withdrawToken', () => {
        describe('failure', () => {
            it('reverts when called by a non-owner', async () => {
                await expectRevert(this.vorD20.withdrawToken(alice, this.deposit, { from: alice }), 'Ownable: caller is not the owner');
            });

            it('reverts when not enough xFUND in the contract', async () => {
                const withdrawAmount = (new BN(this.deposit)).mul(new BN('2'));
                await expectRevert(this.vorD20.withdrawToken(owner, withdrawAmount, { from: owner }), 'ERC20: transfer amount exceeds balance');
            });
        });

        describe('success', () => {
            it('withdraws xFUND', async () => {
                const startingAmount = await this.xFund.balanceOf(owner);
                await this.vorD20.withdrawToken(owner, this.deposit, { from: owner });
                const actualAmount = await this.xFund.balanceOf(owner);
                expect(actualAmount).to.be.bignumber.equal((new BN(startingAmount)).add(new BN(this.deposit)));
            });
        });
    });

    describe('#setKeyHash', () => {
        const newHash = web3.utils.fromAscii('newhash');

        describe('failure', () => {
            it('reverts when called by a non-owner', async () => {
                await expectRevert(this.vorD20.setKeyHash(newHash, { from: alice }), 'Ownable: caller is not the owner');
            });
        });

        describe('success', () => {
            it('sets the key hash', async () => {
                await this.vorD20.setKeyHash(newHash, { from: owner });
                const actualHash = await this.vorD20.keyHash();
                expect(actualHash).to.be.equal(newHash.padEnd(66, '0'));
            });
        });
    });

    describe('#setFee', () => {
        const newFee = 1234;
        
        describe('failure', () => {
            it('reverts when called by a non-owner', async () => {
                await expectRevert(this.vorD20.setFee(newFee, { from: alice }), 'Ownable: caller is not the owner');
            });
        });

        describe('success', () => {
            it('sets the fee', async () => {
                await this.vorD20.setFee(newFee, { from: owner });
                const actualFee = await this.vorD20.fee();
                expect(actualFee).to.be.bignumber.equal(new BN(newFee));
            });
        });
    });

    describe('#setVORCoordinator', () => {
        const newFee = 1234;

        describe('failure', () => {
            it('reverts when called by a non-owner', async () => {
                await expectRevert(this.vorD20.setVORCoordinator(dummy, { from: alice }), 'Ownable: caller is not the owner');
            });
        });

        describe('success', () => {
            it('sets the VORCoordinator', async () => {
                await this.vorD20.setVORCoordinator(dummy, { from: owner });
                const actualVORCoordinator = await this.vorD20.getVORCoordinator();
                expect(actualVORCoordinator).to.be.equal(dummy);
            });
        });
    });

    describe('#house', () => {
        describe('failure', () => {
            it('reverts when dice not rolled', async () => {
                await expectRevert(this.vorD20.house(alice), 'Dice not rolled');
            });

            it('reverts when dice roll is in progress', async () => {
                await this.vorD20.rollDice(this.seed, alice);
                await expectRevert(this.vorD20.house(alice), 'Roll in progress');
            });
        });

        describe('success', () => {
            it('returns the correct house', async () => {
                const randomness = 98765;
                const expectedHouse = 'Martell';

                const rollResult = await this.vorD20.rollDice(this.seed, alice);
                expectEvent(rollResult, 'DiceRolled', { roller: alice });
                
                // Getting requestId from an event 
                const requestId = rollResult.logs[0].args.requestId.toString();
                await this.vorCoordinator.callBackWithRandomness(requestId, randomness, this.vorD20.address);

                const response = await this.vorD20.house(alice);
                expect(response).to.be.equal(expectedHouse);
            });
        });
    });

    describe('#rollDice', () => {
        describe('failure', () => {
            it('reverts when xFUND balance is zero', async () => {
                const vorD20 = await VORD20.new(this.vorCoordinator.address, this.xFund.address, this.keyHash, this.fee, { from: owner });
                await expectRevert(vorD20.rollDice(this.seed, alice), 'Not enough xFUND to pay fee');
            });

            it('reverts when called by a non-owner', async () => {
                await expectRevert(this.vorD20.rollDice(this.seed, alice, { from: alice }), 'Ownable: caller is not the owner');
            });

            it('reverts when the roller rolls more than once', async () => {
                await this.vorD20.rollDice(this.seed, alice);
                await expectRevert(this.vorD20.rollDice(this.seed, alice), 'Already rolled');
            });
        });

        describe('success', () => {
            it('emits a RandomnessRequest event from the VORCoordinator', async () => {
                const result = await this.vorD20.rollDice(this.seed, alice);
                expectEvent.inTransaction(result, 'RandomnessRequest', { sender: this.vorD20.address, seed: this.seed });
            });
        });
    });
});