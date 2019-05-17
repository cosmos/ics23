package proofs

import (
	"bytes"
	"fmt"
)

var IavlSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *ExistenceProof) Calculate() ([]byte, error) {
	if p.GetLeaf() == nil {
		return nil, fmt.Errorf("Existence Proof needs defined LeafOp")
	}

	// leaf step takes the key and value as input
	res, err := p.Leaf.Apply(p.Key, p.Value)
	if err != nil {
		return nil, err
	}

	// the rest just take the output of the last step (reducing it)
	for _, step := range p.Path {
		res, err = step.Apply(res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (p *ExistenceProof) CheckAgainstSpec(spec *ProofSpec) error {
	if p.GetLeaf() == nil {
		return fmt.Errorf("Existence Proof needs defined LeafOp")
	}
	err := checkLeaf(p.Leaf, spec.LeafSpec)
	if err != nil {
		return err
	}
	for _, inner := range p.Path {
		if err := checkInner(inner, spec.LeafSpec.Prefix); err != nil {
			return err
		}
	}
	return nil
}

func checkLeaf(leaf *LeafOp, spec *LeafOp) error {
	if leaf.Hash != spec.Hash {
		return fmt.Errorf("Unexpected HashOp: %d", leaf.Hash)
	}
	if leaf.PrehashKey != spec.PrehashKey {
		return fmt.Errorf("Unexpected PrehashKey: %d", leaf.PrehashKey)
	}
	if leaf.PrehashValue != spec.PrehashValue {
		return fmt.Errorf("Unexpected PrehashValue: %d", leaf.PrehashValue)
	}
	if leaf.Length != spec.Length {
		return fmt.Errorf("Unexpected LengthOp: %d", leaf.Length)
	}
	if !bytes.HasPrefix(leaf.Prefix, spec.Prefix) {
		return fmt.Errorf("Leaf Prefix doesn't start with %X", spec.Prefix)
	}
	return nil
}

func checkInner(inner *InnerOp, leafPrefix []byte) error {
	if bytes.HasPrefix(inner.Prefix, leafPrefix) {
		return fmt.Errorf("Inner Prefix starts with %X", leafPrefix)
	}
	return nil
}
