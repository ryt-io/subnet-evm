// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Defines the interface for the configuration and execution of a precompile contract
package contract

import (
	"math/big"

	"github.com/ryt-io/ryt-v2/snow"
	"github.com/ryt-io/ryt-v2/utils/set"
	"github.com/ryt-io/ryt-v2/vms/evm/predicate"
	"github.com/ryt-io/libevm/common"
	"github.com/ryt-io/libevm/libevm/stateconf"
	"github.com/holiman/uint256"

	"github.com/ryt-io/subnet-evm/precompile/precompileconfig"

	ethtypes "github.com/ryt-io/libevm/core/types"
)

// StatefulPrecompiledContract is the interface for executing a precompiled contract
type StatefulPrecompiledContract interface {
	// Run executes the precompiled contract.
	Run(accessibleState AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error)
}

type StateReader interface {
	GetState(common.Address, common.Hash, ...stateconf.StateDBStateOption) common.Hash
}

// StateDB is the interface for accessing EVM state
type StateDB interface {
	GetState(common.Address, common.Hash, ...stateconf.StateDBStateOption) common.Hash
	SetState(common.Address, common.Hash, common.Hash, ...stateconf.StateDBStateOption)

	SetNonce(common.Address, uint64)
	GetNonce(common.Address) uint64

	GetBalance(common.Address) *uint256.Int
	AddBalance(common.Address, *uint256.Int)

	CreateAccount(common.Address)
	Exist(common.Address) bool

	AddLog(*ethtypes.Log)
	GetPredicate(address common.Address, index int) (predicate.Predicate, bool)

	TxHash() common.Hash

	Snapshot() int
	RevertToSnapshot(int)
}

// AccessibleState defines the interface exposed to stateful precompile contracts
type AccessibleState interface {
	GetStateDB() StateDB
	GetBlockContext() BlockContext
	GetSnowContext() *snow.Context
	GetRules() precompileconfig.Rules
}

// ConfigurationBlockContext defines the interface required to configure a precompile.
type ConfigurationBlockContext interface {
	Number() *big.Int
	Timestamp() uint64
}

type BlockContext interface {
	ConfigurationBlockContext
	// GetPredicateResults returns the result of verifying the predicates of the
	// given transaction, precompile address pair.
	GetPredicateResults(txHash common.Hash, precompileAddress common.Address) set.Bits
}

type Configurator interface {
	MakeConfig() precompileconfig.Config
	Configure(
		chainConfig precompileconfig.ChainConfig,
		precompileconfig precompileconfig.Config,
		state StateDB,
		blockContext ConfigurationBlockContext,
	) error
}
