package ics23

import (

	proofs "github.com/confio/proofs/go"
)

type LeafOp = Struct0
type InnerOp = Struct1
type ExistenceProof = Struct2

func LeafOpToABI(op *proofs.LeafOp) LeafOp {
	return LeafOp{
		Hash:         uint8(op.Hash),
		PrehashKey:   uint8(op.PrehashKey),
		PrehashValue: uint8(op.PrehashValue),
		Len:          uint8(op.Length),
		Prefix:       op.Prefix,
	}
}

func InnerOpToABI(op *proofs.InnerOp) InnerOp {
	return InnerOp{
		Hash:   uint8(op.Hash),
		Prefix: op.Prefix,
		Suffix: op.Suffix,
	}
}

func ExistenceProofToABI(op *proofs.ExistenceProof) ExistenceProof {
	path := make([]InnerOp, len(op.Path))
	for i, op := range op.Path {
		path[i] = InnerOpToABI(op)
	}
	return Struct2{
		Key:   op.Key,
		Value: op.Value,
		Leaf:  LeafOpToABI(op.Leaf),
		Path:  path,
	}
}
