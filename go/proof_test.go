package ics23

import (
	"bytes"
	"testing"
)

func TestExistenceProof(t *testing.T) {
	cases := ExistenceProofTestData()

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Proof.Calculate()
			// short-circuit with error case
			if tc.IsErr {
				if err == nil {
					t.Fatal("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestCheckLeaf(t *testing.T) {
	cases := CheckLeafTestData()
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.Leaf.CheckAgainstSpec(&ProofSpec{LeafSpec: tc.Spec})
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.IsErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCheckAgainstSpec(t *testing.T) {
	validInner := &InnerOp{
		Hash:   HashOp_SHA256,
		Prefix: fromHex("aabbccdd"),
	}
	invalidInner := &InnerOp{
		Hash:   HashOp_SHA256,
		Prefix: fromHex("00aabbccdd"),
		Suffix: fromHex("bb"),
	}
	invalidInner2 := &InnerOp{
		Hash:   HashOp_SHA512,
		Prefix: fromHex("aabbccdd"),
	}

	// this is a copy of IavlSpec with a min and max depth parameters set
	depthLimitedSpec := &ProofSpec{
		LeafSpec: &LeafOp{
			Prefix:       []byte{0},
			Hash:         HashOp_SHA256,
			PrehashValue: HashOp_SHA256,
			Length:       LengthOp_VAR_PROTO,
		},
		InnerSpec: &InnerSpec{
			ChildOrder:      []int32{0, 1},
			MinPrefixLength: 4,
			MaxPrefixLength: 12,
			ChildSize:       33, // (with length byte)
			Hash:            HashOp_SHA256,
		},
		MaxDepth: 4,
		MinDepth: 2,
	}

	cases := map[string]struct {
		proof *ExistenceProof
		spec  *ProofSpec
		isErr bool
	}{
		"empty proof fails": {
			proof: &ExistenceProof{
				Key:   []byte("foo"),
				Value: []byte("bar"),
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"accepts one proper leaf": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
			},
			spec: IavlSpec,
		},
		"rejects invalid leaf": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf: &LeafOp{
					Prefix: []byte{0},
					Hash:   HashOp_SHA256,
					Length: LengthOp_VAR_PROTO,
				},
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"rejects only inner proof": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Path: []*InnerOp{
					validInner,
				},
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"accepts leaf with valid inner proofs": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					validInner,
				},
			},
			spec: IavlSpec,
		},
		"rejects leaf with invalid inner proofs": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					invalidInner,
					validInner,
				},
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"rejects invalid inner proof (hash mismatch)": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					invalidInner2,
					validInner,
					validInner,
				},
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"allows depth limited in proper range": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					validInner,
					validInner,
				},
			},
			spec:  depthLimitedSpec,
			isErr: false,
		},
		"reject depth limited with too few inner nodes": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
				},
			},
			spec:  depthLimitedSpec,
			isErr: true,
		},
		"reject depth limited with too many inner nodes": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					validInner,
					validInner,
					validInner,
					validInner,
				},
			},
			spec:  depthLimitedSpec,
			isErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.proof.CheckAgainstSpec(tc.spec)
			if tc.isErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.isErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}
