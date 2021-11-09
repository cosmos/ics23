package ics23_sol

import (
	"math"

	ics23 "github.com/confio/ics23/go"
)

func translateHashOp(op ics23.HashOp) uint8 {
	if op > math.MaxUint8 || op < 0 {
		panic("translateHashOp overflow")
	}
	return uint8(op)
}

func translateLengthOp(op ics23.LengthOp) uint8 {
	if op > math.MaxUint8 || op < 0 {
		panic("translateLengthOp overflow")
	}
	return uint8(op)
}

func translateInnerOp(op *ics23.InnerOp) InnerOpData {
	return InnerOpData{
		Hash:   translateHashOp(op.Hash),
		Prefix: op.Prefix,
		Suffix: op.Suffix,
	}
}

func translateInnerOps(ops ...*ics23.InnerOp) []InnerOpData {
	if len(ops) == 0 {
		return nil
	}
	icss := make([]InnerOpData, len(ops))
	for idx, op := range ops {
		icss[idx] = translateInnerOp(op)
	}
	return icss
}

func translateLeafOp(op *ics23.LeafOp) LeafOpData {
	if op == nil {
		return LeafOpData{}
	}
	return LeafOpData{
		Hash:         translateHashOp(op.Hash),
		PrehashKey:   translateHashOp(op.PrehashKey),
		PrehashValue: translateHashOp(op.PrehashValue),
		Length:       translateLengthOp(op.Length),
		Prefix:       op.Prefix,
	}
}

func translateProofSpec(spec *ics23.ProofSpec) ProofSpecData {
	return ProofSpecData{
		LeafSpec:  translateLeafOp(spec.LeafSpec),
		InnerSpec: translateInnerSpec(spec.InnerSpec),
		MinDepth:  spec.MinDepth,
		MaxDepth:  spec.MaxDepth,
	}
}

func translateInnerSpec(spec *ics23.InnerSpec) InnerSpecData {
	return InnerSpecData{
		ChildOrder:      spec.ChildOrder,
		ChildSize:       spec.ChildSize,
		MinPrefixLength: spec.MinPrefixLength,
		MaxPrefixLength: spec.MaxPrefixLength,
		EmptyChild:      spec.EmptyChild,
		Hash:            translateHashOp(spec.Hash),
	}
}

func translateExistenceProof(proof *ics23.ExistenceProof) ExistenceProofData {
	if proof == nil {
		return ExistenceProofData{}
	}
	return ExistenceProofData{
		Key:   proof.Key,
		Value: proof.Value,
		Leaf:  translateLeafOp(proof.Leaf),
		Path:  translateInnerOps(proof.Path...),
	}
}

func translateNonExistenceProof(proof *ics23.NonExistenceProof) NonExistenceProofData {
	if proof == nil {
		return NonExistenceProofData{}
	}
	return NonExistenceProofData{
		Key:   proof.Key,
		Left:  translateExistenceProof(proof.Left),
		Right: translateExistenceProof(proof.Right),
	}
}

func translateBatchEntry(entry *ics23.BatchEntry) BatchEntryData {
	return BatchEntryData{
		Exist:    translateExistenceProof(entry.GetExist()),
		Nonexist: translateNonExistenceProof(entry.GetNonexist()),
	}
}
func translateBatchEntries(entries ...*ics23.BatchEntry) []BatchEntryData {
	if len(entries) == 0 {
		return nil
	}
	entriesData := make([]BatchEntryData, len(entries))
	for idx, entry := range entries {
		entriesData[idx] = translateBatchEntry(entry)
	}
	return entriesData
}

func translateBatchProof(proof *ics23.BatchProof) BatchProofData {
	if proof == nil {
		return BatchProofData{}
	}
	return BatchProofData{
		Entries: translateBatchEntries(proof.Entries...),
	}
}

func translateCompressedExistenceProof(proof *ics23.CompressedExistenceProof) CompressedExistenceProofData {
	if proof == nil {
		return CompressedExistenceProofData{}
	}
	return CompressedExistenceProofData{
		Key:   proof.Key,
		Value: proof.Value,
		Leaf:  translateLeafOp(proof.Leaf),
		Path:  proof.Path,
	}
}

func translateCompressedNonExistenceProof(proof *ics23.CompressedNonExistenceProof) CompressedNonExistenceProofData {
	if proof == nil {
		return CompressedNonExistenceProofData{}
	}
	return CompressedNonExistenceProofData{
		Key:   proof.Key,
		Left:  translateCompressedExistenceProof(proof.Left),
		Right: translateCompressedExistenceProof(proof.Right),
	}
}

func translateCompressedBatchEntry(proof *ics23.CompressedBatchEntry) CompressedBatchEntryData {
	return CompressedBatchEntryData{
		Exist:    translateCompressedExistenceProof(proof.GetExist()),
		Nonexist: translateCompressedNonExistenceProof(proof.GetNonexist()),
	}
}

func translateCompressedBatchEntries(entries ...*ics23.CompressedBatchEntry) []CompressedBatchEntryData {
	if len(entries) == 0 {
		return nil
	}
	entriesData := make([]CompressedBatchEntryData, len(entries))
	for idx, entry := range entries {
		entriesData[idx] = translateCompressedBatchEntry(entry)
	}
	return entriesData
}

func translateCompressedBatchProof(proof *ics23.CompressedBatchProof) CompressedBatchProofData {
	if proof == nil {
		return CompressedBatchProofData{}
	}
	return CompressedBatchProofData{
		Entries:      translateCompressedBatchEntries(proof.GetEntries()...),
		LookupInners: translateInnerOps(proof.GetLookupInners()...),
	}
}

// WARNING: does not work
func translateCommitmentProof(proof *ics23.CommitmentProof) CommitmentProofData {
	return CommitmentProofData{
		Exist:      translateExistenceProof(proof.GetExist()),
		Nonexist:   translateNonExistenceProof(proof.GetNonexist()),
		Batch:      translateBatchProof(proof.GetBatch()),
		Compressed: translateCompressedBatchProof(proof.GetCompressed()),
	}
}
