package ics23

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecompressExist(t *testing.T) {
	leafOp := &LeafOp{
		Hash:         HashOp_SHA256,
		PrehashKey:   HashOp_NO_HASH,
		PrehashValue: HashOp_NO_HASH,
		Length:       LengthOp_NO_PREFIX,
		Prefix:       []byte{},
	}
	innerOps := []*InnerOp{
		{
			Hash:   HashOp_SHA256,
			Prefix: generateInnerOpPrefix(),
		},
		{
			Hash:   HashOp_SHA256,
			Prefix: generateInnerOpPrefix(),
			Suffix: []byte{1},
		},
		{
			Hash:   HashOp_SHA256,
			Prefix: generateInnerOpPrefix(),
			Suffix: []byte{2},
		},
	}

	var (
		compressedExistProof *CompressedExistenceProof
		lookup               []*InnerOp
	)

	cases := []struct {
		name     string
		malleate func()
		expProof *ExistenceProof
		expError error
	}{
		{
			"success: single lookup",
			func() {
				compressedExistProof.Path = []int32{0}
			},
			&ExistenceProof{
				Key:   []byte{0},
				Value: []byte{0},
				Leaf:  leafOp,
				Path:  []*InnerOp{innerOps[0]},
			},
			nil,
		},
		{
			"success: multiple lookups",
			func() {
				compressedExistProof.Path = []int32{0, 1, 0, 2}
			},
			&ExistenceProof{
				Key:   []byte{0},
				Value: []byte{0},
				Leaf:  leafOp,
				Path:  []*InnerOp{innerOps[0], innerOps[1], innerOps[0], innerOps[2]},
			},
			nil,
		},
		{
			"success: empty exist proof",
			func() {
				compressedExistProof = nil
			},
			nil,
			nil,
		},
		{
			"failure: path index out of bounds",
			func() {
				compressedExistProof.Path = []int32{0}
				lookup = nil
			},
			nil,
			fmt.Errorf("compressed existence proof at index %d has lookup index out of bounds", 0),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			// reset default values for compressedExistProof and lookup
			compressedExistProof = &CompressedExistenceProof{
				Key:   []byte{0},
				Value: []byte{0},
				Leaf:  leafOp,
			}

			lookup = innerOps

			tc.malleate()

			proof, err := decompressExist(compressedExistProof, lookup)
			if !reflect.DeepEqual(tc.expProof, proof) {
				t.Fatalf("expexted proof: %v, got: %v", tc.expProof, proof)
			}

			if tc.expError == nil && err != nil {
				t.Fatal(err)
			}

			if tc.expError != nil && err.Error() != tc.expError.Error() {
				t.Fatalf("expected: %v, got: %v", tc.expError, err)
			}

		})
	}
}
