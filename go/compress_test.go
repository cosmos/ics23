package ics23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecompressIndexPanic(t *testing.T) {
	testCases := []struct {
		name string
		path []int32
	}{
		{"out of bounds positive", []int32{0x7FFFFFFF}},
		{"out of bounds negative", []int32{-1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			proof := &CommitmentProof{
				Proof: &CommitmentProof_Compressed{
					Compressed: &CompressedBatchProof{
						LookupInners: []*InnerOp{},
						Entries: []*CompressedBatchEntry{{
							Proof: &CompressedBatchEntry_Exist{
								Exist: &CompressedExistenceProof{
									Key:   []byte("k"),
									Value: []byte("v"),
									Leaf:  &LeafOp{},
									Path:  tc.path,
								},
							},
						}},
					},
				},
			}

			var decompressed *CommitmentProof
			require.NotPanics(t, func() {
				decompressed = Decompress(proof)
			})
			if decompressed != nil {
				batch := decompressed.GetBatch()
				if batch != nil && len(batch.Entries) > 0 {
					require.Nil(t, batch.Entries[0].GetExist())
				}
			}
		})
	}
}
