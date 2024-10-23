/*
*
This implements the client side functions as specified in
https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments

In particular:

	// Assumes ExistenceProof
	type verifyMembership = (root: CommitmentRoot, proof: CommitmentProof, key: Key, value: Value) => boolean

	// Assumes NonExistenceProof
	type verifyNonMembership = (root: CommitmentRoot, proof: CommitmentProof, key: Key) => boolean

	// Assumes BatchProof - required ExistenceProofs may be a subset of all items proven
	type batchVerifyMembership = (root: CommitmentRoot, proof: CommitmentProof, items: Map<Key, Value>) => boolean

	// Assumes BatchProof - required NonExistenceProofs may be a subset of all items proven
	type batchVerifyNonMembership = (root: CommitmentRoot, proof: CommitmentProof, keys: Set<Key>) => boolean

We make an adjustment to accept a Spec to ensure the provided proof is in the format of the expected merkle store.
This can avoid an range of attacks on fake preimages, as we need to be careful on how to map key, value -> leaf
and determine neighbors
*/
package ics23

import (
	"bytes"
)

// CommitmentRoot is a byte slice that represents the merkle root of a tree that can be used to validate proofs
type CommitmentRoot []byte

// VerifyMembership returns true iff
// proof is (contains) an ExistenceProof for the given key and value AND
// calculating the root for the ExistenceProof matches the provided CommitmentRoot
func VerifyMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, key []byte, value []byte) bool {
	// decompress it before running code (no-op if not compressed)
	proof = Decompress(proof)
	ep := getExistProofForKey(proof, key)
	if ep == nil {
		return false
	}
	err := ep.Verify(spec, root, key, value)
	return err == nil
}

// VerifyNonMembership returns true iff
// proof is (contains) a NonExistenceProof
// both left and right sub-proofs are valid existence proofs (see above) or nil
// left and right proofs are neighbors (or left/right most if one is nil)
// provided key is between the keys of the two proofs
func VerifyNonMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, key []byte) bool {
	// decompress it before running code (no-op if not compressed)
	proof = Decompress(proof)
	np := getNonExistProofForKey(spec, proof, key)
	if np == nil {
		return false
	}
	err := np.Verify(spec, root, key)
	return err == nil
}

func getExistProofForKey(proof *CommitmentProof, key []byte) *ExistenceProof {
	if proof == nil {
		return nil
	}

	switch p := proof.Proof.(type) {
	case *CommitmentProof_Exist:
		ep := p.Exist
		if bytes.Equal(ep.Key, key) {
			return ep
		}
	case *CommitmentProof_Batch:
		for _, sub := range p.Batch.Entries {
			if ep := sub.GetExist(); ep != nil && bytes.Equal(ep.Key, key) {
				return ep
			}
		}
	}
	return nil
}

func getNonExistProofForKey(spec *ProofSpec, proof *CommitmentProof, key []byte) *NonExistenceProof {
	switch p := proof.Proof.(type) {
	case *CommitmentProof_Nonexist:
		np := p.Nonexist
		if isLeft(spec, np.Left, key) && isRight(spec, np.Right, key) {
			return np
		}
	case *CommitmentProof_Batch:
		for _, sub := range p.Batch.Entries {
			if np := sub.GetNonexist(); np != nil && isLeft(spec, np.Left, key) && isRight(spec, np.Right, key) {
				return np
			}
		}
	}
	return nil
}

func isLeft(spec *ProofSpec, left *ExistenceProof, key []byte) bool {
	return left == nil || bytes.Compare(keyForComparison(spec, left.Key), keyForComparison(spec, key)) < 0
}

func isRight(spec *ProofSpec, right *ExistenceProof, key []byte) bool {
	return right == nil || bytes.Compare(keyForComparison(spec, right.Key), keyForComparison(spec, key)) > 0
}
