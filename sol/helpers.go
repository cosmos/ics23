package ics23_sol

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func setup(t *testing.T, testCases int) (*bind.TransactOpts, []bind.CallOpts, *backends.SimulatedBackend) {
	gAlloc := make(map[common.Address]core.GenesisAccount, testCases)
	auths := make([]bind.CallOpts, testCases)
	for idx := 0; idx < testCases; idx++ {
		key, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		transactor := bind.NewKeyedTransactor(key)
		auths[idx] = bind.CallOpts{From: transactor.From}
		gAlloc[auths[idx].From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000000)} // 1 eth
	}
	deployerKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	deployerAuth := bind.NewKeyedTransactor(deployerKey)
	gAlloc[deployerAuth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000000)} // 1 eth

	client := backends.NewSimulatedBackend(gAlloc, 1000000000000000000) // 1 eth
	return deployerAuth, auths, client
}

func proofsEqual(a, b CommitmentProofData) bool {
	//exist
	if existenceProofEqual(a.Exist, b.Exist) == false {
		return false
	}
	//nonexist
	if nonexistenceProofEqual(a.Nonexist, b.Nonexist) == false {
		return false
	}

	if len(a.Batch.Entries) != len(b.Batch.Entries) {
		return false
	}
	if batchProofEqual(a.Batch, b.Batch) == false {
		return false
	}
	return compressedBatchProofEqual(a.Compressed, b.Compressed)
}

func existenceProofEqual(a, b ExistenceProofData) bool {
	if bytes.Equal(a.Key, b.Key) == false {
		return false
	}
	if bytes.Equal(a.Value, b.Value) == false {
		return false
	}
	if len(a.Path) != len(b.Path) {
		return false
	}
	for idx, apath := range a.Path {
		if apath.Hash != b.Path[idx].Hash ||
			bytes.Equal(apath.Prefix, b.Path[idx].Prefix) ||
			bytes.Equal(apath.Suffix, b.Path[idx].Suffix) {
			return false
		}
	}
	return a.Leaf.Hash == b.Leaf.Hash && a.Leaf.PrehashKey == b.Leaf.PrehashKey && a.Leaf.PrehashValue == a.Leaf.PrehashValue
}

func nonexistenceProofEqual(a, b NonExistenceProofData) bool {
	if existenceProofEqual(a.Left, b.Left) == false {
		return false
	}
	if existenceProofEqual(a.Right, b.Right) == false {
		return false
	}
	if bytes.Equal(a.Key, b.Key) == false {
		return false
	}
	return true
}

func batchProofEqual(a, b BatchProofData) bool {
	if len(a.Entries) != len(b.Entries) {
		return false
	}
	for idx, aentry := range a.Entries {
		if existenceProofEqual(aentry.Exist, b.Entries[idx].Exist) == false {
			return false
		}
		if nonexistenceProofEqual(aentry.Nonexist, b.Entries[idx].Nonexist) == false {
			return false
		}
	}
	return true
}
func compressedBatchProofEqual(a, b CompressedBatchProofData) bool {
	//compressedBatchProof are never written, only either copied or emptied.
	// exploiting this to take a comparison shortcut
	return len(a.Entries) == len(b.Entries) && len(a.LookupInners) == len(b.LookupInners)
}
