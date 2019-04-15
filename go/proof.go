package proofs

import fmt "fmt"

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
func (p *ExistenceProof) Calculate() ([]byte, error) {
	if len(p.Steps) == 0 {
		return nil, fmt.Errorf("Existence Proof needs at least one step")
	}

	first, rem := p.Steps[0], p.Steps[1:]
	// first step takes the key and value as input
	res, err := first.Apply(p.Key, p.Value)
	if err != nil {
		return nil, err
	}

	// the rest just take the output of the last step (reducing it)
	for _, step := range rem {
		res, err = step.Apply(res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
