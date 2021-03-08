package ics23

import (
	ics23 "github.com/confio/ics23/go"
)

func LeafOpToABI(op *ics23.LeafOp) ICS23LeafOp {
	return ICS23LeafOp{
		Hash:         uint8(op.Hash),
		PrehashKey:   uint8(op.PrehashKey),
		PrehashValue: uint8(op.PrehashValue),
		Len:          uint8(op.Length),
		Prefix:       op.Prefix,
	}
}

func InnerOpToABI(op *ics23.InnerOp) ICS23InnerOp {
	return ICS23InnerOp{
		Hash:   uint8(op.Hash),
		Prefix: op.Prefix,
		Suffix: op.Suffix,
	}
}

func ExistenceProofToABI(op *ics23.ExistenceProof) ICS23ExistenceProof {
	path := make([]ICS23InnerOp, len(op.Path))
	for i, op := range op.Path {
		path[i] = InnerOpToABI(op)
	}
	return ICS23ExistenceProof{
		Key:   op.Key,
		Value: op.Value,
		Leaf:  LeafOpToABI(op.Leaf),
		Path:  path,
	}
}
