use crate::proofs;
use crate::verify::{CommitmentRoot, verify_existence};

pub fn verify_membership(proof: proofs::CommitmentProof, _spec: proofs::ProofSpec, root: CommitmentRoot, key: &[u8], value: &[u8]) -> bool {
    match proof.proof {
        Some(proofs::CommitmentProof_oneof_proof::exist(ex)) => verify_existence(ex, _spec, root, key, value).is_ok(),
        _ => false,
    }
}