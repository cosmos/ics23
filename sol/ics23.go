package ics23

import (
	"math/big"

	ics23 "github.com/confio/ics23/go"
)

func LeafOpToABI(op *ics23.LeafOp) ICS23LeafOp {
	if op == nil {
		return ICS23LeafOp{}
	}
	return ICS23LeafOp{
		Valid: true,
		Hash:         uint8(op.Hash),
		PrehashKey:   uint8(op.PrehashKey),
		PrehashValue: uint8(op.PrehashValue),
		Len:          uint8(op.Length),
		Prefix:       op.Prefix,
	}
}

func InnerOpToABI(op *ics23.InnerOp) ICS23InnerOp {
	if op == nil {
		return ICS23InnerOp{}
	}
	return ICS23InnerOp{
		Valid: true,
		Hash:   uint8(op.Hash),
		Prefix: op.Prefix,
		Suffix: op.Suffix,
	}
}

func ExistenceProofToABI(op *ics23.ExistenceProof) ICS23ExistenceProof {
	if op == nil {
		return ICS23ExistenceProof{}
	}
	path := make([]ICS23InnerOp, len(op.Path))
	for i, op := range op.Path {
		path[i] = InnerOpToABI(op)
	}
	return ICS23ExistenceProof{
		Valid: true,
		Key:   op.Key,
		Value: op.Value,
		Leaf:  LeafOpToABI(op.Leaf),
		Path:  path,
	}
}

func NonExistenceProofToABI(op *ics23.NonExistenceProof) ICS23NonExistenceProof {
	if op == nil {
		return ICS23NonExistenceProof{}
	}
	return ICS23NonExistenceProof{
		Valid: true,
		Key:   op.Key,
		Left:  ExistenceProofToABI(op.Left),
		Right: ExistenceProofToABI(op.Right),
	}
}

func ProofSpecToABI(spec *ics23.ProofSpec) ICS23ProofSpec {
	childOrder := make([]*big.Int, len(spec.InnerSpec.ChildOrder))
	for i, x := range spec.InnerSpec.ChildOrder {
		childOrder[i] = big.NewInt(int64(x))
	}

	return ICS23ProofSpec{
		LeafSpec: LeafOpToABI(spec.LeafSpec),
		InnerSpec: ICS23InnerSpec{
			ChildOrder:      childOrder,
			ChildSize:       big.NewInt(int64(spec.InnerSpec.ChildSize)),
			MinPrefixLength: big.NewInt(int64(spec.InnerSpec.MinPrefixLength)),
			MaxPrefixLength: big.NewInt(int64(spec.InnerSpec.MaxPrefixLength)),
			EmptyChild:      spec.InnerSpec.EmptyChild,
			Hash:            uint8(spec.InnerSpec.Hash),
		},
		MaxDepth: big.NewInt(int64(spec.MaxDepth)),
		MinDepth: big.NewInt(int64(spec.MinDepth)),
	}
}
