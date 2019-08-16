package proofs

import (
	"bytes"

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
	InnerSpec: &InnerSpec{
		ChildOrder: []int32{0, 1},
		MinPrefixLength: 4,
		MaxPrefixLength: 12,
		ChildSize: 33, // (with length byte)
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
		return nil, errors.New("Existence Proof needs defined LeafOp")
	}

	// leaf step takes the key and value as input
	res, err := p.Leaf.Apply(p.Key, p.Value)
	if err != nil {
		return nil, errors.WithMessage(err, "leaf")
	}

	// the rest just take the output of the last step (reducing it)
	for _, step := range p.Path {
		res, err = step.Apply(res)
		if err != nil {
			return nil, errors.WithMessage(err, "inner")
		}
	}
	return res, nil
}

// CheckAgainstSpec will verify the leaf and all path steps are in the format defined in spec
func (p *ExistenceProof) CheckAgainstSpec(spec *ProofSpec) error {
	if p.GetLeaf() == nil {
		return errors.New("Existence Proof needs defined LeafOp")
	}
	err := p.Leaf.CheckAgainstSpec(spec)
	if err != nil {
		return errors.WithMessage(err, "leaf")
	}
	for _, inner := range p.Path {
		if err := inner.CheckAgainstSpec(spec); err != nil {
			return errors.WithMessage(err, "inner")
		}
	}
	return nil
}

// Verify does all checks to ensure the proof has valid non-existence proofs,
// and they ensure the given key is not in the CommitmentState
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
		if !p.Right.IsLeftmost(spec.InnerSpec) {
			return errors.New("left proof missing, right proof must be left-most")
		}
		return nil
	}

	if rightKey == nil {
		if !p.Left.IsRightmost(spec.InnerSpec) {
			return errors.New("right proof missing, left proof must be right-most")
		}
		return nil
	}

	// TODO: enforce neighbors
	return nil
}

// IsLeftmost returns true if this proof is the left-most path in the tree
func (p *ExistenceProof)IsLeftmost(spec *InnerSpec) bool {
	for _, step := range p.Path {
		// we only want a prefix, no child....
		if len(step.Prefix) < int(spec.MinPrefixLength) {
			return false
		}
		if len(step.Prefix) > int(spec.MaxPrefixLength) {
			return false
		}

		// and one child as suffix
		if len(step.Suffix) != int(spec.ChildSize) {
			return false
		}
	}
	return true
}

// IsRightmost returns true if this proof is the left-most path in the tree
func (p *ExistenceProof)IsRightmost(spec *InnerSpec) bool {
	for _, step := range p.Path {
		// we only want a prefix plus child....
		if len(step.Prefix) < int(spec.MinPrefixLength + spec.ChildSize) {
			return false
		}
		if len(step.Prefix) > int(spec.MaxPrefixLength + spec.ChildSize) {
			return false
		}

		// and no suffix
		if len(step.Suffix) != 0 {
			return false
		}
	}
	return true
}
