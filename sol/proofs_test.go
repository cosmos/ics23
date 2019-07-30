package proofs

import (
	"crypto/sha256"
	"math/big"
	"math/rand"
	"testing"
	"fmt"

	"golang.org/x/crypto/ripemd160"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func testDoHash(t *testing.T, cont ProofsSession) {
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

func TestHelperFunctions(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	address := auth.From
	gAlloc := map[common.Address]core.GenesisAccount{
		address: {Balance: big.NewInt(100000000000)},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 100000000000)

	_, _, cont, err := DeployProofs(auth, sim)

	require.NoError(t, err)

	session := ProofsSession{cont, bind.CallOpts{}, bind.TransactOpts{}}

	sim.Commit()

	testDoHash(t, session)
}

func TestABIEncoding(t *testing.T) {
	ty := new(abi.Argument)
	err := ty.UnmarshalJSON([]byte(`
{
	"Name": "LeafOp",
	"Type": "tuple",
	"Components": [
		{
			"Name": "hash",
			"Type": "uint8"
		},
		{
			"Name": "prehash_key",
			"Type": "uint8"
		},
		{
			"Name": "prehash_value",
			"Type": "uint8"
		},
		{
			"Name": "len",
			"Type": "uint8"
		},
		{
			"Name": "prefix",
			"Type": "bytes"
		}
	]
}`))
	require.NoError(t, err)

	tys := abi.Arguments{*ty}
	bz, err := tys.Pack(struct {
		Hash uint8
		PrehashKey uint8
		PrehashValue uint8
		Len uint8
		Prefix []byte
	}{
		1,
		1,
		1,
		0,
		[]byte{0x00},
	})
	require.NoError(t, err)
	fmt.Println(bz)
}
