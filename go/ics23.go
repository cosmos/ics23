/**
This implements the client side functions as specified in
https://github.com/cosmos/ics/tree/master/spec/ics-023-vector-commitments

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
package proofs

import (
	"bytes"
	"fmt"
)

// CommitmentRoot is a byte slice that represents the merkle root of a tree that can be used to validate proofs
type CommitmentRoot []byte

// VerifyMembership returns true iff
// proof is (contains) an ExistenceProof for the given key and value AND
// calculating the root for the ExistenceProof matches the provided CommitmentRoot
func VerifyMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, key []byte, value []byte) bool {
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
	// TODO: handle batch
	np := getNonExistProofForKey(proof, key)
	if np == nil {
		return false
	}
	err := np.Verify(spec, root, key)
	return err == nil
}

// BatchVerifyMembership will ensure all items are also proven by the CommitmentProof (which should be a BatchProof,
// unless there is one item, when a ExistenceProof may work)
func BatchVerifyMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, items map[string][]byte) bool {
	for k, v := range items {
		valid := VerifyMembership(spec, root, proof, []byte(k), v)
		fmt.Printf("Validate %t\n", valid)
		if !valid {
			return false
		}
	}
	return true
}

// BatchVerifyNonMembership will ensure all items are also proven to not be in the Commitment by the CommitmentProof
// (which should be a BatchProof, unless there is one item, when a NonExistenceProof may work)
func BatchVerifyNonMembership(spec *ProofSpec, root CommitmentRoot, proof *CommitmentProof, keys [][]byte) bool {
	for _, k := range keys {
		valid := VerifyNonMembership(spec, root, proof, k)
		fmt.Printf("Validate non %t\n", valid)
		if !valid {
			return false
		}
	}
	return true
}


func getExistProofForKey(proof *CommitmentProof, key []byte) *ExistenceProof{
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

func getNonExistProofForKey(proof *CommitmentProof, key []byte) *NonExistenceProof{
	switch p := proof.Proof.(type) {
	case *CommitmentProof_Nonexist:
		np := p.Nonexist
		if isLeft(np.Left, key) && isRight(np.Right, key) {
			return np
		}
	case *CommitmentProof_Batch:
		for _, sub := range p.Batch.Entries {
			if np := sub.GetNonexist(); np != nil && isLeft(np.Left, key) && isRight(np.Right, key) {
				return np
			}
		}
	}
	return nil
}

func isLeft(left *ExistenceProof, key []byte) bool {
	return left == nil || bytes.Compare(left.Key, key) < 0
}

func isRight(right *ExistenceProof, key []byte) bool {
	return right == nil || bytes.Compare(right.Key, key) > 0
}