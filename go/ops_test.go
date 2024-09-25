package ics23

import (
	"bytes"
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

func TestForgeNonExistenceProofWithIncorrectMaxPrefixLength(t *testing.T) {
	spec := &ProofSpec{ // TendermintSpec
		LeafSpec: &LeafOp{
			Prefix:       []byte{0},
			PrehashKey:   HashOp_NO_HASH,
			Hash:         HashOp_SHA256,
			PrehashValue: HashOp_SHA256,
			Length:       LengthOp_VAR_PROTO,
		},
		InnerSpec: &InnerSpec{
			ChildOrder:      []int32{0, 1},
			MinPrefixLength: 1,
			MaxPrefixLength: 1,
			ChildSize:       32, // (no length byte)
			Hash:            HashOp_SHA256,
		},
	}

	spec.InnerSpec.MaxPrefixLength = 33
	leafOp := spec.LeafSpec
	aLeaf, _ := leafOp.Apply([]byte("a"), []byte("a"))
	bLeaf, _ := leafOp.Apply([]byte("b"), []byte("b"))
	b2Leaf, _ := leafOp.Apply([]byte("b2"), []byte("b2"))

	cLeaf, _ := leafOp.Apply([]byte("c"), []byte("c"))
	aExist := ExistenceProof{
		Key:   []byte("a"),
		Value: []byte("a"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: append(bLeaf, b2Leaf...),
			},
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: cLeaf,
			},
		},
	}
	bExist := ExistenceProof{
		Key:   []byte("b"),
		Value: []byte("b"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: append([]byte{1}, aLeaf...),
				Suffix: b2Leaf,
			},
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: cLeaf,
			},
		},
	}
	b2Exist := ExistenceProof{
		Key:   []byte("b2"),
		Value: []byte("b2"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: append(append([]byte{1}, aLeaf...), bLeaf...),
				Suffix: []byte{},
			},
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: cLeaf,
			},
		},
	}
	yHash, _ := aExist.Path[0].Apply(aLeaf)
	cExist := ExistenceProof{
		Key:   []byte("c"),
		Value: []byte("c"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			&InnerOp{
				Hash:   spec.InnerSpec.Hash,
				Prefix: append([]byte{1}, yHash...),
				Suffix: []byte{},
			},
		},
	}
	aNotExist := NonExistenceProof{
		Key:   []byte("a"),
		Left:  nil,
		Right: &bExist,
	}
	root, err := aExist.Calculate()
	if err != nil {
		t.Fatal("failed to calculate existence proof of leaf a")
	}

	expError := fmt.Errorf("inner, %w", errors.New("spec.InnerSpec.MaxPrefixLength must be < spec.InnerSpec.MinPrefixLength + spec.InnerSpec.ChildSize"))
	err = aExist.Verify(spec, root, []byte("a"), []byte("a"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of leaf a returned incorrect error")
	}

	err = bExist.Verify(spec, root, []byte("b"), []byte("b"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of leaf b returned incorrect error")
	}

	err = b2Exist.Verify(spec, root, []byte("b2"), []byte("b2"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of third leaf returned incorrect error")
	}

	err = cExist.Verify(spec, root, []byte("c"), []byte("c"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of leaf c returned incorrect error")
	}

	err = aNotExist.Verify(spec, root, []byte("a"))
	expError = fmt.Errorf("right proof, %w", expError)
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove non-existence of leaf a returned incorrect error")
	}
}
