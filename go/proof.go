package proofs

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
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
	if err := p.CheckAgainstSpec(spec); err != nil {
		return err
	}

	if !bytes.Equal(key, p.Key) {
		return errors.Errorf("Provided key doesn't match proof")
	}
	if !bytes.Equal(value, p.Value) {
		return errors.Errorf("Provided value doesn't match proof")
	}

	calc, err := p.Calculate()
	if err != nil {
		return errors.Wrap(err, "Error calculating root")
	}
	if !bytes.Equal(root, calc) {
		return errors.Errorf("Calculcated root doesn't match provided root")
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

// Verify does all checks to ensure this proof proves this key, value -> root
// and matches the spec.
func (p *NonExistenceProof) Verify(spec *ProofSpec, root CommitmentRoot, key []byte) error {
	// ensure the existence proofs are valid
	var leftKey, rightKey []byte 
	if p.Left != nil {
		if err := p.Left.Verify(spec, root, p.Left.Key, p.Left.Value); err != nil {
			return errors.Wrap(err, "left proof")
		}
		leftKey = p.Left.Key
	}
	if p.Right != nil {
		if err := p.Right.Verify(spec, root, p.Right.Key, p.Right.Value); err != nil {
			return errors.Wrap(err, "right proof")
		}
		rightKey = p.Right.Key
	}

	// If both proofs are missing, this is not a valid proof
	if leftKey == nil && rightKey == nil {
		return errors.New("both left and right proofs missing")
	}

	// Ensure in valid range
	if rightKey != nil {
		if bytes.Compare(key, rightKey) >= 0 {
			return errors.New("key is not left of right proof")
		}
	}
	if leftKey != nil {
		if bytes.Compare(key, leftKey) <= 0 {
			return errors.New("key is not right of left proof")
		}
	}

	if leftKey == nil {
		// TODO: enforce left-most
		return nil
	}

	if rightKey == nil {
		// TODO: enforce right-most
		return nil
	}

	// TODO: enforce neighbors
	return nil
}
