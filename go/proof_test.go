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
