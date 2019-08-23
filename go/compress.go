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
	var centries []*CompressedBatchEntry
	var lookup []*InnerOp
	registry := make(map[string]int32)

	for _, entry := range batch.Entries {
		centry := compressEntry(entry, &lookup, registry)
		centries = append(centries, centry)
	}

	return &CompressedBatchProof{
		Entries: centries,
		LookupInners: lookup,
	}
}

func compressEntry(entry *BatchEntry, lookup *[]*InnerOp, registry map[string]int32) *CompressedBatchEntry {
	if exist := entry.GetExist(); exist != nil {
		return &CompressedBatchEntry{
			Proof: &CompressedBatchEntry_Exist{
				Exist: compressExist(exist, lookup, registry),
			},
		}
	} else {
		non := entry.GetNonexist()
		return &CompressedBatchEntry{
			Proof: &CompressedBatchEntry_Nonexist{
				Nonexist: &CompressedNonExistenceProof{
					Left: compressExist(non.Left, lookup, registry),
					Right: compressExist(non.Right, lookup, registry),
				},
			},
		}
	}
}

func compressExist(exist *ExistenceProof, lookup *[]*InnerOp, registry map[string]int32) *CompressedExistenceProof {
	res := &CompressedExistenceProof{
		Key: exist.Key,
		Value: exist.Value,
		Leaf: exist.Leaf,
		Path: make([]int32, len(exist.Path)),
	}
	for i, step := range exist.Path {
		res.Path[i] = compressStep(step, lookup, registry)
	}
	return res
}

func compressStep(step *InnerOp, lookup *[]*InnerOp, registry map[string]int32) int32 {
	bz, err := step.Marshal()
	if err != nil {
		panic(err)
	}
	sig := string(bz)

	// load from cache if there
	if num, ok := registry[sig]; ok {
		return num
	}

	// create new step if not there
	num := int32(len(*lookup))
	*lookup = append(*lookup, step)
	registry[sig] = num
	return num
}

func decompress(comp *CompressedBatchProof) *BatchProof {
	return nil
}