/* global artifacts, contract, it, beforeEach, describe, before, web3 */
/* eslint no-unused-expressions: 0 */
/* eslint consistent-return: 0 */
/* eslint-disable max-len */

process.env.NODE_ENV = 'test';

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(web3.BigNumber))
  .should();

const ValidatorMgr = artifacts.require('ValidatorMgr.sol');
const PublicResolver = artifacts.require('PublicResolver.sol');
const KNS = artifacts.require('KNSRegistry.sol');
const FIFSRegistrar = artifacts.require('FIFSRegistrar.sol');
const MiningTokenMock = artifacts.require('TokenMock.sol');
const namehash = require('eth-ens-namehash');
const { EVMError } = require('../helpers/testUtils.js');

contract('Validator Manager', ([_, owner, newOwner, notOwner]) => {
  beforeEach(async () => {
    this.kns = await KNS.new({ from: owner });
    this.registrar = await FIFSRegistrar.new(this.kns.address, namehash('kowala'));
    this.resolver = await PublicResolver.new(this.kns.address);

    await this.kns.setSubnodeOwner(0, web3.sha3('kowala'), this.registrar.address, { from: owner });
    await this.registrar.register(web3.sha3('miningtoken'), owner, { from: owner });
    await this.kns.setResolver(namehash('miningtoken.kowala'), this.resolver.address, { from: owner });
    this.miningToken = await MiningTokenMock.new();
    await this.resolver.setAddr(namehash('miningtoken.kowala'), this.miningToken.address, { from: owner });
    this.validator = await ValidatorMgr.new(1, 2, 3, 1, this.kns.address, { from: owner });
  });

  it('should set MiningToken Address from KNS during creation', async () => {
    // given
    const knsResolverAddr = await this.validator.knsResolver();
    const resolver = await PublicResolver.at(knsResolverAddr);

    // when
    const miningTokenAddr = await resolver.addr(namehash('miningtoken.kowala'));
    
    // then
    await miningTokenAddr.should.be.equal(this.miningToken.address);
  });
});
