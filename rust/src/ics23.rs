use crate::proofs;
use crate::verify::{CommitmentRoot, verify_existence, verify_non_existence};

// Use CommitmentRoot vs &[u8] to stick with ics naming
#[allow(clippy::ptr_arg)]
pub fn verify_membership(
    proof: &proofs::CommitmentProof,
    spec: &proofs::ProofSpec,
    root: &CommitmentRoot,
    key: &[u8],
    value: &[u8],
) -> bool {
    if let Some(proofs::commitment_proof::Proof::Exist(ex)) = &proof.proof {
        let valid = verify_existence(&ex, spec, root, key, value);
        valid.is_ok()
    } else {
        false
    }
}
#[warn(clippy::ptr_arg)]

// Use CommitmentRoot vs &[u8] to stick with ics naming
#[allow(clippy::ptr_arg)]
pub fn verify_non_membership(
    proof: &proofs::CommitmentProof,
    spec: &proofs::ProofSpec,
    root: &CommitmentRoot,
    key: &[u8],
) -> bool {
    if let Some(proofs::commitment_proof::Proof::Nonexist(non)) = &proof.proof {
        let valid = verify_non_existence(&non, spec, root, key);
        valid.is_ok()
    } else {
        false
    }
}
#[warn(clippy::ptr_arg)]


pub fn iavl_spec() -> proofs::ProofSpec {
    let leaf = proofs::LeafOp{
        hash: proofs::HashOp::Sha256.into(),
        prehash_key: 0,
        prehash_value: proofs::HashOp::Sha256.into(),
        length: proofs::LengthOp::VarProto.into(),
        prefix: vec![0_u8]
    };
    let inner = proofs::InnerSpec{
        child_order: vec![0, 1],
        min_prefix_length: 4,
        max_prefix_length: 12,
        child_size: 33,
        empty_child: vec![]
    };
    proofs::ProofSpec{
        leaf_spec: Some(leaf),
        inner_spec: Some(inner),
    }
}

pub fn tendermint_spec() -> proofs::ProofSpec {
    let leaf = proofs::LeafOp{
        hash: proofs::HashOp::Sha256.into(),
        prehash_key: 0,
        prehash_value: proofs::HashOp::Sha256.into(),
        length: proofs::LengthOp::VarProto.into(),
        prefix: vec![0_u8]
    };
    let inner = proofs::InnerSpec{
        child_order: vec![0, 1],
        min_prefix_length: 1,
        max_prefix_length: 1,
        child_size: 32,
        empty_child: vec![]
    };
    proofs::ProofSpec{
        leaf_spec: Some(leaf),
        inner_spec: Some(inner)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use failure::{ensure};
    use prost::Message;
    use serde::Deserialize;
    use std::fs::File;
    use std::io::prelude::*;

    use crate::helpers::Result;

    #[derive(Deserialize, Debug)]
    struct TestVector {
        pub root: String,
        pub proof: String,
        pub key: String,
        pub value: String,
    }

    fn verify_test_vector(filename: &str, spec: &proofs::ProofSpec) -> Result<()> {
        let mut file = File::open(filename)?;
        let mut contents = String::new();
        file.read_to_string(&mut contents)?;

        let data: TestVector = serde_json::from_str(&contents)?;
        let proto_bin = hex::decode(&data.proof)?;
        let root = hex::decode(data.root)?;
        let key = hex::decode(data.key)?;

        let mut parsed = proofs::CommitmentProof{ proof: None };
        parsed.merge(&proto_bin)?;

        if data.value.is_empty() {
             let valid = super::verify_non_membership(&parsed, spec, &root, &key);
             ensure!(valid, "invalid test vector");
             Ok(())
        } else {
            let value = hex::decode(data.value)?;
            let valid = super::verify_membership(&parsed, spec, &root, &key, &value);
            ensure!(valid, "invalid test vector");
            Ok(())
        }
    }

    #[test]
    fn test_vector_iavl_left() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/exist_left.json", &spec)
    }

    #[test]
    fn test_vector_iavl_right() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/exist_right.json", &spec)
    }

    #[test]
    fn test_vector_iavl_middle() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/exist_middle.json", &spec)
    }

    #[test]
    fn test_vector_iavl_left_non() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/nonexist_left.json", &spec)
    }

    #[test]
    fn test_vector_iavl_right_non() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/nonexist_right.json", &spec)
    }

    #[test]
    fn test_vector_iavl_middle_non() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/nonexist_middle.json", &spec)
    }

    #[test]
    fn test_vector_tendermint_left() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/exist_left.json", &spec)
    }

    #[test]
    fn test_vector_tendermint_right() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/exist_right.json", &spec)
    }

    #[test]
    fn test_vector_tendermint_middle() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/exist_middle.json", &spec)
    }

    #[test]
    fn test_vector_tendermint_left_non() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/nonexist_left.json", &spec)
    }

    #[test]
    fn test_vector_tendermint_right_non() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/nonexist_right.json", &spec)
    }

    #[test]
    fn test_vector_tendermint_middle_non() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/nonexist_middle.json", &spec)
    }
}
