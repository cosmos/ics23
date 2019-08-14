extern crate protobuf;

use std::vec::Vec;

use crate::proofs;
use crate::ops::{apply_inner, apply_leaf, Result};

pub type CommitmentRoot = Vec<u8>;

pub fn verify_existence(
  proof: proofs::ExistenceProof,
  _spec: proofs::ProofSpec,
  root: CommitmentRoot,
  key: &[u8],
  value: &[u8],
) -> Result<bool> {

// TODO:  ensureSpec(proof, spec);
  if !proof.key.eq(&key) {
      return Err("Provided key doesn't match proof");
  }
  if !proof.value.eq(&value) {
      return Err("Provided value doesn't match proof");
  }

  let calc = calculate_existence_root(&proof)?;
  if !calc.eq(&root) {
      return Err("Root hash doesn't match");
  }
  Ok(true)
}


// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
pub fn calculate_existence_root(proof: &proofs::ExistenceProof) -> Result<CommitmentRoot> {
    if proof.key.len() == 0 {
        return Err("Existence proof must have key set");
    }
    if proof.value.len() == 0 {
        return Err("Existence proof must have value set");
    }
    if proof.leaf.is_none() {
        return Err("No leaf operation set");
    }
    let mut hash = apply_leaf(proof.leaf.get_ref(), &proof.key, &proof.value)?;

    for step in proof.path.iter() {
        hash = apply_inner(step, hash)?;
    }
    Ok(hash)
}
