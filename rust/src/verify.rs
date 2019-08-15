extern crate protobuf;
extern crate failure;

use failure::{ensure};
use crate::helpers::{Result};
use crate::proofs;
use crate::ops::{apply_inner, apply_leaf};

pub type CommitmentRoot = ::std::vec::Vec<u8>;

pub fn verify_existence(
  proof: &proofs::ExistenceProof,
  spec: &proofs::ProofSpec,
  root: CommitmentRoot,
  key: &[u8],
  value: &[u8],
) -> Result<bool> {
  check_existence_spec(proof, spec)?;
  ensure!(proof.key.eq(&key), "Provided key doesn't match proof");
  ensure!(proof.value.eq(&value), "Provided value doesn't match proof");

  let calc = calculate_existence_root(&proof)?;
  ensure!(calc.eq(&root), "Root hash doesn't match");
  Ok(true)
}


// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
pub fn calculate_existence_root(proof: &proofs::ExistenceProof) -> Result<CommitmentRoot> {
    ensure!(proof.key.len() > 0, "Existence proof must have key set");
    ensure!(proof.value.len() > 0, "Existence proof must have value set");
    ensure!(proof.leaf.is_some(), "No leaf operation set");

    let mut hash = apply_leaf(proof.leaf.get_ref(), &proof.key, &proof.value)?;
    for step in proof.path.iter() {
        hash = apply_inner(step, &hash)?;
    }
    Ok(hash)
}

fn check_existence_spec(proof: &proofs::ExistenceProof, spec: &proofs::ProofSpec) -> Result<()> {
  ensure!(proof.leaf.is_some(), "Existence proof must start with a leaf operation");
  ensure!(spec.leaf_spec.is_some(), "Spec must include leafSpec");
//   ensureLeaf(proof.leaf, spec.leafSpec);
//   const path = proof.path || [];
//   for (const inner of path) {
//     ensureInner(inner, spec.leafSpec.prefix);
//   }
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    use crate::proofs::{HashOp, LengthOp};

    #[test]
    fn calculate_root_from_leaf() {
        let mut leaf = proofs::LeafOp::new();
        leaf.set_hash(HashOp::SHA256);
        leaf.set_length(LengthOp::VAR_PROTO);
        let mut proof = proofs::ExistenceProof::new();
        proof.set_leaf(leaf);
        proof.set_key(b"food".to_vec());
        proof.set_value(b"some longer text".to_vec());

        let expected = hex::decode("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265").unwrap();
        assert_eq!(expected, calculate_existence_root(&proof).unwrap());
    }

    #[test]
    fn calculate_root_from_leaf_and_inner() {
        let mut leaf = proofs::LeafOp::new();
        leaf.set_hash(HashOp::SHA256);
        leaf.set_length(LengthOp::VAR_PROTO);

        let mut inner = proofs::InnerOp::new();
        inner.set_hash(HashOp::SHA256);
        inner.set_prefix(hex::decode("deadbeef00cafe00").unwrap());

        let mut proof = proofs::ExistenceProof::new();
        proof.set_key(b"food".to_vec());
        proof.set_value(b"some longer text".to_vec());
        proof.set_leaf(leaf);
        proof.set_path(protobuf::RepeatedField::from_slice(&[inner]));

        let expected = hex::decode("836ea236a6902a665c2a004c920364f24cad52ded20b1e4f22c3179bfe25b2a9").unwrap();
        assert_eq!(expected, calculate_existence_root(&proof).unwrap());
    }

}
