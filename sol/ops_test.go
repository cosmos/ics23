package ics23

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"math/rand"
	"testing"

	"golang.org/x/crypto/ripemd160"

	ics23 "github.com/confio/ics23/go"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func testDoHash(t *testing.T, cont ICS23Session) {
	for i := 0; i < 1000; i++ {
		image := make([]byte, 64)
		rand.Read(image)

		// NO_HASH
		hashed, err := cont.DoHashOrNoop(0, image)
		require.NoError(t, err)
		require.Equal(t, image, hashed)

		// SHA256
		hashed, err = cont.DoHash(1, image)
		require.NoError(t, err)
		exp32 := sha256.Sum256(image)
		require.Equal(t, exp32[:], hashed)

		// Not supported in solidity
		// SHA512
		hashed, err = cont.DoHash(2, image)
		//require.NoError(t, err)
		//exp64 := sha512.Sum512(image)
		//require.Equal(t, exp64[:], hashed)
		require.Error(t, err)

		// KECCAK
		hashed, err = cont.DoHash(3, image)
		require.NoError(t, err)
		exp32 = crypto.Keccak256Hash(image)
		require.Equal(t, exp32[:], hashed)

		// RIPEMD160
		hashed, err = cont.DoHash(4, image)
		require.NoError(t, err)
		hash := ripemd160.New()
		hash.Write(image)
		require.Equal(t, hash.Sum(nil)[:], hashed)

		// BITCOIN
		hashed, err = cont.DoHash(5, image)
		require.NoError(t, err)
		sha := sha256.Sum256(image)
		hash = ripemd160.New()
		hash.Write(sha[:])
		require.Equal(t, hash.Sum(nil)[:], hashed)

		// INVALID
		_, err = cont.DoHash(6, image)
		require.Error(t, err)
	}
}

func Initialize(t *testing.T) ICS23Session {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	address := auth.From
	gAlloc := map[common.Address]core.GenesisAccount{
		address: {Balance: big.NewInt(100000000000)},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 100000000000)

	_, _, cont, err := DeployICS23(auth, sim)

	require.NoError(t, err)

	session := ICS23Session{cont, bind.CallOpts{}, bind.TransactOpts{}}

	sim.Commit()

	return session
}

func TestHelperFunctions(t *testing.T) {
	session := Initialize(t)

	testDoHash(t, session)
}

func TestLeafOp(t *testing.T) {
	session := Initialize(t)

	cases := ics23.LeafOpTestData()

	// Set IsErr = true for all SHA512
	for name, tc := range cases {
		if tc.Op.Hash == ics23.HashOp_SHA512 {
			tc.IsErr = true
			cases[name] = tc
		}
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := session.ApplyLeaf(LeafOpToABI(tc.Op), tc.Key, tc.Value)
			if tc.IsErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, res, tc.Expected)
			}
		})
	}
}

func TestInnerOp(t *testing.T) {
	session := Initialize(t)

	cases := ics23.InnerOpTestData()

	for name, tc := range cases {
		if tc.IsErr {
			delete(cases, name)
		}
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := session.ApplyInner(InnerOpToABI(tc.Op), tc.Child)
			if tc.IsErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, res, tc.Expected)
			}
		})
	}
}

func TestDoHash(t *testing.T) {
	session := Initialize(t)

	cases := ics23.DoHashTestData()

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := session.DoHash(uint8(tc.HashOp), []byte(tc.Preimage))
			if tc.HashOp == ics23.HashOp_SHA512 {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedHash, hex.EncodeToString(res))
			}
		})
	}
}
