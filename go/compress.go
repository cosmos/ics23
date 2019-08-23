package ics23

// IsCompressed returns true if the proof was compressed
func IsCompressed(proof *CommitmentProof) bool {
	return proof.GetCompressed() != nil
}

// Compress will return a CompressedBatchProof if the input is BatchProof
// Otherwise it will return the input.
// This is safe to call multiple times (idempotent)
func Compress(proof *CommitmentProof) *CommitmentProof {
	batch := proof.GetBatch()
	if batch == nil {
		return proof
	}
	return &CommitmentProof{
		Proof: &CommitmentProof_Compressed{
			Compressed: compress(batch),
		},
	}
}

// Decompress will return a BatchProof if the input is CompressedBatchProof
// Otherwise it will return the input.
// This is safe to call multiple times (idempotent)
func Decompress(proof *CommitmentProof) *CommitmentProof {
	comp := proof.GetCompressed()
	if comp == nil {
		return proof
	}
	return &CommitmentProof{
		Proof: &CommitmentProof_Batch{
			Batch: decompress(comp),
		},
	}
}

func compress(batch *BatchProof) *CompressedBatchProof {
	return nil
}

func decompress(comp *CompressedBatchProof) *BatchProof {
	return nil
}