package proofs

import (
	"bytes"
	"fmt"
)

// IavlSpec constrains the format from proofs-iavl (iavl merkle proofs)
var IavlSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
}

// TendermintSpec constrains the format from proofs-tendermint (crypto/merkle SimpleProof)
var TendermintSpec = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
		Length:       LengthOp_VAR_PROTO,
	},
}

// Verify does all checks to ensure this proof proves this key, value -> root
// and matches the spec.
func (p *ExistenceProof) Verify(spec *ProofSpec, root CommitmentRoot, key []byte, value []byte) error {
	if err := exist.CheckAgainstSpec(spec); err != nil {
		return err
	}

	if !bytes.Equal(key, exist.Key) {
		return fmt.Errorf("Provided key doesn't match proof")
	}
	if !bytes.Equal(value, exist.Value) {
		return fmt.Errorf("Provided value doesn't match proof")
	}

	calc, err := exist.Calculate()
	if err != nil {
		return fmt.Errorf("Error calculating root: %s", err)
	}
	if !bytes.Equal(root, calc) {
		return fmt.Errorf("Calculcated root doesn't match provided root")
	}

	return nil

}


// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *ExistenceProof) Calculate() (CommitmentRoot, error) {
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
