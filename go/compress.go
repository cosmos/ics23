package ics23

import "fmt"

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
func Decompress(proof *CommitmentProof) (*CommitmentProof, error) {
	comp := proof.GetCompressed()
	if comp != nil {
		batch, err := decompress(comp)
		if err != nil {
			return nil, err
		}
		return &CommitmentProof{
			Proof: &CommitmentProof_Batch{
				Batch: batch,
			},
		}, nil
	}
	return proof, nil
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
		Entries:      centries,
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
	}

	non := entry.GetNonexist()
	return &CompressedBatchEntry{
		Proof: &CompressedBatchEntry_Nonexist{
			Nonexist: &CompressedNonExistenceProof{
				Key:   non.Key,
				Left:  compressExist(non.Left, lookup, registry),
				Right: compressExist(non.Right, lookup, registry),
			},
		},
	}
}

func compressExist(exist *ExistenceProof, lookup *[]*InnerOp, registry map[string]int32) *CompressedExistenceProof {
	if exist == nil {
		return nil
	}
	res := &CompressedExistenceProof{
		Key:   exist.Key,
		Value: exist.Value,
		Leaf:  exist.Leaf,
		Path:  make([]int32, len(exist.Path)),
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

	// load from cache if there
	if num, ok := registry[string(bz)]; ok {
		return num
	}

	// create new step if not there
	num := int32(len(*lookup))
	*lookup = append(*lookup, step)
	registry[string(bz)] = num
	return num
}

func decompress(comp *CompressedBatchProof) (*BatchProof, error) {
	lookup := comp.LookupInners

	var entries []*BatchEntry

	for _, centry := range comp.Entries {
		entry, err := decompressEntry(centry, lookup)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return &BatchProof{
		Entries: entries,
	}, nil
}

func decompressEntry(entry *CompressedBatchEntry, lookup []*InnerOp) (*BatchEntry, error) {
	if exist := entry.GetExist(); exist != nil {
		decompressedExist, err := decompressExist(exist, lookup)
		if err != nil {
			return nil, err
		}
		return &BatchEntry{
			Proof: &BatchEntry_Exist{
				Exist: decompressedExist,
			},
		}, nil
	}

	non := entry.GetNonexist()
	decompressedLeft, err := decompressExist(non.Left, lookup)
	if err != nil {
		return nil, err
	}

	decompressedRight, err := decompressExist(non.Right, lookup)
	if err != nil {
		return nil, err
	}

	return &BatchEntry{
		Proof: &BatchEntry_Nonexist{
			Nonexist: &NonExistenceProof{
				Key:   non.Key,
				Left:  decompressedLeft,
				Right: decompressedRight,
			},
		},
	}, nil
}

func decompressExist(exist *CompressedExistenceProof, lookup []*InnerOp) (*ExistenceProof, error) {
	if exist == nil {
		return nil, nil
	}
	res := &ExistenceProof{
		Key:   exist.Key,
		Value: exist.Value,
		Leaf:  exist.Leaf,
		Path:  make([]*InnerOp, len(exist.Path)),
	}
	for i, step := range exist.Path {
		if int(step) >= len(lookup) {
			return nil, fmt.Errorf("compressed existence proof at index %d has lookup index out of bounds", i)
		}
		res.Path[i] = lookup[step]
	}
	return res, nil
}
