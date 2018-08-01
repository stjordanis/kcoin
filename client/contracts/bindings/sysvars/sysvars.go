package sysvars

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/sysvars/SystemVars.sol
//go:generate ../../../build/bin/abigen -abi build/SystemVars.abi -bin build/SystemVars.bin -pkg sysvars -type SystemVars -out ./gen_system.go

var mapSystemToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x17C56D5aC0cddFd63aC860237197827cB4639CDA"),
}

type System interface {
	CurrencySupply() (*big.Int, error)
	PrevCurrencyPrice() (*big.Int, error)
	CurrencyPrice() (*big.Int, error)
	MintedAmount() (*big.Int, error)
	MintedReward() (*big.Int, error)
	OracleDeduction(*big.Int) (*big.Int, error)
	OracleReward() (*big.Int, error)
	Address() common.Address
}

type system struct {
	*SystemVarsSession
	addr common.Address
}

// Bind returns a binding to the current consensus engine
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, ok := mapSystemToAddr[chainID.Uint64()]
	if !ok {
		return nil, bindings.ErrNoAddress
	}

	sys, err := NewSystemVars(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &system{
		SystemVarsSession: &SystemVarsSession{
			Contract: sys,
			CallOpts: bind.CallOpts{},
		},
		addr: addr,
	}, nil
}

func (sys *system) Address() common.Address { return sys.addr }