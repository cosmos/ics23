package ics23

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"
)

func TestLeafOp(t *testing.T) {
	cases := LeafOpTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Op.Apply(tc.Key, tc.Value)
			// short-circuit with error case
			if tc.IsErr && err == nil {
				t.Fatal("expected error, but got none")
			}
			if !tc.IsErr && err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestInnerOp(t *testing.T) {
	cases := InnerOpTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Op.Apply(tc.Child)
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got none")
			}
			if !tc.IsErr && err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestDoHash(t *testing.T) {
	cases := DoHashTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := doHash(tc.HashOp, []byte(tc.Preimage))
			if err != nil {
				t.Fatal(err)
			}
			hexRes := hex.EncodeToString(res)
			if hexRes != tc.ExpectedHash {
				t.Fatalf("Expected %s got %s", tc.ExpectedHash, hexRes)
			}
		})
	}
}

func TestInnerOpCheckAgainstSpec(t *testing.T) {
	var (
		spec    *ProofSpec
		innerOp *InnerOp
	)

	cases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {},
			nil,
		},
		{
			"failure: empty spec",
			func() {
				spec = nil
			},
			errors.New("op and spec must be both non-nil"),
		},
		{
			"failure: empty inner spec",
			func() {
				spec.InnerSpec = nil
			},
			errors.New("spec.InnerSpec must be non-nil"),
		},
		{
			"failure: empty leaf spec",
			func() {
				spec.LeafSpec = nil
			},
			errors.New("spec.LeafSpec must be non-nil"),
		},
		{
			"failure: incorrect hash function inner op",
			func() {
				innerOp.Hash = HashOp_BITCOIN
			},
			fmt.Errorf("unexpected HashOp: %d", HashOp_BITCOIN),
		},
		{
			"failure: spec fails validation",
			func() {
				innerOp.Prefix = []byte{0x01}
			},
			fmt.Errorf("wrong value in IAVL leaf op"),
		},
		{
			"failure: inner prefix starts with leaf prefix",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Prefix = append(spec.LeafSpec.Prefix, innerOp.Prefix...)
			},
			fmt.Errorf("inner Prefix starts with %X", []byte{0}),
		},
		{
			"failure: inner prefix too short",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Prefix = []byte{0x01}
			},
			errors.New("innerOp prefix too short (1)"),
		},
		{
			"failure: inner prefix too long",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Prefix = []byte("AgQIIGe3bHuC1g6+5/Qd0RoCU0waFu+nDCFzEDViMN/VrQwgIA==")
			},
			errors.New("innerOp prefix too long (52)"),
		},
		{
			"failure: child size must be greater than zero",
			func() {
				spec.InnerSpec.ChildSize = 0
			},
			errors.New("spec.InnerSpec.ChildSize must be >= 1"),
		},
		{
			"failure: inner op suffix malformed",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Suffix = []byte{0x01}
			},
			fmt.Errorf("InnerOp suffix malformed"),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			spec = &ProofSpec{ // IavlSpec
				LeafSpec: &LeafOp{
					Prefix:       []byte{0},
					PrehashKey:   HashOp_NO_HASH,
					Hash:         HashOp_SHA256,
					PrehashValue: HashOp_SHA256,
					Length:       LengthOp_VAR_PROTO,
				},
				InnerSpec: &InnerSpec{
					ChildOrder:      []int32{0, 1},
					MinPrefixLength: 4,
					MaxPrefixLength: 12,
					ChildSize:       33, // (with length byte)
					EmptyChild:      nil,
					Hash:            HashOp_SHA256,
				},
			}

			innerOp = &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: generateInnerOpPrefix(),
				Suffix: []byte(""),
			}
			tc.malleate()

			err := innerOp.CheckAgainstSpec(spec, 1) // use a layer number of 1
			if tc.expError == nil && err != nil {
				t.Fatal(err)
			}

			if tc.expError != nil && err.Error() != tc.expError.Error() {
				t.Fatalf("expected: %v, got: %v", tc.expError, err)
			}
		})

	}
}

// generatePrefix generates a valid iavl prefix for an inner op.
func generateInnerOpPrefix() []byte {
	var (
		varintBuf  [binary.MaxVarintLen64]byte
		lengthByte byte = 0x20
	)
	height, size, version := 5, 10, 20
	prefix := convertVarIntToBytes(int64(height), varintBuf)
	prefix = append(prefix, convertVarIntToBytes(int64(size), varintBuf)...)
	prefix = append(prefix, convertVarIntToBytes(int64(version), varintBuf)...)
	prefix = append(prefix, lengthByte)
	return prefix
}

func convertVarIntToBytes(orig int64, buf [binary.MaxVarintLen64]byte) []byte {
	n := binary.PutVarint(buf[:], orig)
	return buf[:n]
}
