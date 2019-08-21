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
    if let Some(proofs::CommitmentProof_oneof_proof::exist(ex)) = &proof.proof {
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
    if let Some(proofs::CommitmentProof_oneof_proof::nonexist(non)) = &proof.proof {
        let valid = verify_non_existence(&non, spec, root, key);
        valid.is_ok()
    } else {
        false
    }
}
#[warn(clippy::ptr_arg)]


pub fn iavl_spec() -> proofs::ProofSpec {
    let mut leaf = proofs::LeafOp::new();
    leaf.set_prefix(vec![0_u8]);
    leaf.set_hash(proofs::HashOp::SHA256);
    leaf.set_prehash_value(proofs::HashOp::SHA256);
    leaf.set_length(proofs::LengthOp::VAR_PROTO);

    let mut spec = proofs::ProofSpec::new();
    spec.set_leaf_spec(leaf);
    spec
}

pub fn tendermint_spec() -> proofs::ProofSpec {
    let mut leaf = proofs::LeafOp::new();
    leaf.set_prefix(vec![0_u8]);
    leaf.set_hash(proofs::HashOp::SHA256);
    leaf.set_prehash_value(proofs::HashOp::SHA256);
    leaf.set_length(proofs::LengthOp::VAR_PROTO);

    let mut spec = proofs::ProofSpec::new();
    spec.set_leaf_spec(leaf);
    spec
}

#[cfg(test)]
mod tests {
    use super::*;

    extern crate protobuf;
    extern crate serde;
    extern crate serde_json;

    use failure::{bail, ensure};
    use protobuf::Message;
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

        let mut parsed = proofs::CommitmentProof::new();
        parsed.merge_from_bytes(&proto_bin)?;

        if data.value.is_empty() {
            // non existence
            bail!("non membership not yet implemented");
            // let valid = super::verify_non_membership(spec, &root, &parsed, key);
            // ensure!(valid, "invalid test vector");
            // Ok(())
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
}
