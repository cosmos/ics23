package proofs

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
				Key:   []byte("foo"),
				Value: []byte("bar"),
			},
			isErr: true,
		},
		// copied from ops_test / TestLeafOp
		"executes one leaf step": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Steps: []*ProofOp{
					WrapLeaf(&LeafOp{
						Hash:   HashOp_SHA256,
						Length: LengthOp_VAR_PROTO,
					}),
				},
			},
			expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		// iavl leaf: start with 0, length 3
		// inner prefix: !start with 0, length >= 4
		"demonstrate maliability of leaf if we change leaf algorithm": {
			proof: &ExistenceProof{
				Key:   append([]byte{4}, []byte("food")...),
				Value: append([]byte{16}, []byte("some longer text")...),
				Steps: []*ProofOp{
					WrapLeaf(&LeafOp{
						Hash: HashOp_SHA256,
					}),
				},
			},
			expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		"demonstrate maliability of leaf if we change leaf prefix": {
			proof: &ExistenceProof{
				Key:   append([]byte("od"), byte(16)),
				Value: []byte("some longer text"),
				Steps: []*ProofOp{
					WrapLeaf(&LeafOp{
						Prefix: []byte{4, 'f', 'o'},
						Hash:   HashOp_SHA256,
					}),
				},
			},
			expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		"cannot execute two leafs": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Steps: []*ProofOp{
					WrapLeaf(&LeafOp{
						Hash:   HashOp_SHA256,
						Length: LengthOp_VAR_PROTO,
					}),
					WrapLeaf(&LeafOp{
						Hash:   HashOp_SHA256,
						Length: LengthOp_VAR_PROTO,
					}),
				},
			},
			isErr: true,
		},
		"cannot execute inner first": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Steps: []*ProofOp{
					WrapInner(&InnerOp{
						Hash:   HashOp_SHA256,
						Prefix: fromHex("deadbeef00cafe00"),
					}),
				},
			},
			isErr: true,
		},
		"executes lead then inner op": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Steps: []*ProofOp{
					WrapLeaf(&LeafOp{
						Hash:   HashOp_SHA256,
						Length: LengthOp_VAR_PROTO,
					}),
					// output: b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265
					WrapInner(&InnerOp{
						Hash:   HashOp_SHA256,
						Prefix: fromHex("deadbeef00cafe00"),
					}),
					// echo -n deadbeef00cafe00b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265 | xxd -r -p | sha256sum
				},
			},
			expected: fromHex("836ea236a6902a665c2a004c920364f24cad52ded20b1e4f22c3179bfe25b2a9"),
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
			err := checkLeaf(tc.leaf, tc.spec)
			if tc.isErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.isErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCheckAgainstSpec(t *testing.T) {
	validInner := WrapInner(&InnerOp{
		Prefix: fromHex("aa"),
	})
	invalidInner := WrapInner(&InnerOp{
		Prefix: fromHex("00aa"),
		Suffix: fromHex("bb"),
	})

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
				Steps: []*ProofOp{
					WrapLeaf(IavlSpec.LeafSpec),
				},
			},
			spec: IavlSpec,
		},
		"rejects invalid leaf": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Steps: []*ProofOp{
					WrapLeaf(&LeafOp{
						Prefix: []byte{0},
						Hash:   HashOp_SHA256,
						Length: LengthOp_VAR_PROTO,
					}),
				},
			},
			spec:  IavlSpec,
			isErr: true,
		},
		"rejects only inner proof": {
			proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Steps: []*ProofOp{
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
				Steps: []*ProofOp{
					WrapLeaf(IavlSpec.LeafSpec),
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
				Steps: []*ProofOp{
					WrapLeaf(IavlSpec.LeafSpec),
					validInner,
					invalidInner,
					validInner,
				},
			},
			spec:  IavlSpec,
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
