package ics23

import (
	"bytes"
	"testing"
)

func TestExistenceProof(t *testing.T) {
	cases := map[string]struct {
		proof    *ExistenceProof
		isErr    bool
		expected []byte
	}{
		"must have at least one step": {
			proof: &ExistenceProof{
				Key:       []byte("foo"),
				ValueHash: []byte("bar"),
			},
			isErr: true,
		},
		// copied from ops_test / TestLeafOp
		"executes one leaf step": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("some longer text"),
				Leaf: &LeafOp{
					Hash:   HashOp_SHA256,
					Length: LengthOp_VAR_PROTO,
				},
			},
			expected: fromHex("7496fdaa49e3764d635f9f21e60cec0a75b8d7595a9a3bb013692bb45d14e326"),
		},
		// iavl leaf: start with 0, length 3
		// inner prefix: !start with 0, length >= 4
		"demonstrate maliability of leaf if we change leaf algorithm": {
			proof: &ExistenceProof{
				Key:       append([]byte{4}, []byte("food")...),
				ValueHash: append([]byte{16}, []byte("some longer text")...),
				Leaf: &LeafOp{
					Hash: HashOp_SHA256,
				},
			},
			expected: fromHex("7496fdaa49e3764d635f9f21e60cec0a75b8d7595a9a3bb013692bb45d14e326"),
		},
		"demonstrate maliability of leaf if we change leaf prefix": {
			proof: &ExistenceProof{
				Key:       append([]byte("od"), byte(16)),
				ValueHash: []byte("some longer text"),
				Leaf: &LeafOp{
					Prefix: []byte{4, 'f', 'o'},
					Hash:   HashOp_SHA256,
				},
			},
			expected: fromHex("7496fdaa49e3764d635f9f21e60cec0a75b8d7595a9a3bb013692bb45d14e326"),
		},
		"cannot execute inner first": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("some longer text"),
				Path: []*InnerOp{
					&InnerOp{
						Hash:   HashOp_SHA256,
						Prefix: fromHex("deadbeef00cafe00"),
					},
				},
			},
			isErr: true,
		},
		"executes leaf then inner op": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("some longer text"),
				Leaf: &LeafOp{
					Hash:   HashOp_SHA256,
					Length: LengthOp_VAR_PROTO,
				},
				// output: b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265
				Path: []*InnerOp{
					&InnerOp{
						Hash:   HashOp_SHA256,
						Prefix: fromHex("deadbeef00cafe00"),
					},
					// echo -n deadbeef00cafe00b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265 | xxd -r -p | sha256sum
				},
			},
			expected: fromHex("0a9acd00a6a8b95a65a46263dfddb3c4731e376bb12e25c7fb96cac5c7885ffc"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.proof.Calculate()
			// short-circuit with error case
			if tc.isErr {
				if err == nil {
					t.Fatal("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.expected))
			}
		})
	}
}

func TestCheckLeaf(t *testing.T) {
	cases := map[string]struct {
		leaf  *LeafOp
		spec  *LeafOp
		isErr bool
	}{
		"empty spec, empty leaf": {
			leaf: &LeafOp{},
			spec: &LeafOp{},
		},
		"empty spec allows prefix": {
			leaf: &LeafOp{Prefix: fromHex("aabb")},
			spec: &LeafOp{},
		},
		"empty spec doesn't allow hashop": {
			leaf:  &LeafOp{Hash: HashOp_SHA256},
			spec:  &LeafOp{},
			isErr: true,
		},
		"spec with different prefixes": {
			leaf:  &LeafOp{Prefix: fromHex("aabb")},
			spec:  &LeafOp{Prefix: fromHex("bb")},
			isErr: true,
		},
		"leaf with empty prefix (but spec has one)": {
			leaf:  &LeafOp{},
			spec:  &LeafOp{Prefix: fromHex("bb")},
			isErr: true,
		},
		"leaf and spec match, all fields full": {
			leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
		},
		"leaf and spec differ on hash": {
			leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA256,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			isErr: true,
		},
		"leaf and spec differ on length": {
			leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_NO_PREFIX,
			},
			spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			isErr: true,
		},
		"leaf and spec differ on prehash key": {
			leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_SHA256,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			isErr: true,
		},
		"leaf and spec differ on prehash value": {
			leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_NO_HASH,
				Length:       LengthOp_VAR_PROTO,
			},
			spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			isErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.leaf.CheckAgainstSpec(&ProofSpec{LeafSpec: tc.spec})
			if tc.isErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.isErr && err != nil {
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
				Key:       []byte("foo"),
				ValueHash: []byte("bar"),
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"accepts one proper leaf": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
			},
			spec: IavlSpec,
		},
		"rejects invalid leaf": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
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
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Path: []*InnerOp{
					validInner,
				},
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"accepts leaf with valid inner proofs": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					validInner,
				},
			},
			spec: IavlSpec,
		},
		"rejects leaf with invalid inner proofs": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
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
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
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
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
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
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
				},
			},
			spec:  depthLimitedSpec,
			isErr: true,
		},
		"reject depth limited with too many inner nodes": {
			proof: &ExistenceProof{
				Key:       []byte("food"),
				ValueHash: []byte("bar"),
				Leaf:      IavlSpec.LeafSpec,
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
