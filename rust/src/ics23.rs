extern crate failure;


use crate::proofs;
use crate::verify::{CommitmentRoot, verify_existence};

pub fn verify_membership(proof: &proofs::CommitmentProof, spec: &proofs::ProofSpec, root: CommitmentRoot, key: &[u8], value: &[u8]) -> bool {
    match &proof.proof {
        Some(proofs::CommitmentProof_oneof_proof::exist(ex)) => {
            let valid = verify_existence(&ex, spec, root, key, value);
            valid.is_ok() && valid.unwrap()
        }
        _ => false,
    }
}

pub fn iavl_spec() -> proofs::ProofSpec {
    let mut leaf = proofs::LeafOp::new();
    leaf.set_prefix(vec![0u8]);
    leaf.set_hash(proofs::HashOp::SHA256);
    leaf.set_prehash_value(proofs::HashOp::SHA256);
    leaf.set_length(proofs::LengthOp::VAR_PROTO);

    let mut spec = proofs::ProofSpec::new();
    spec.set_leaf_spec(leaf);
    spec
}

pub fn tendermint_spec() -> proofs::ProofSpec {
    let mut leaf = proofs::LeafOp::new();
    leaf.set_prefix(vec![0u8]);
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

    use serde::{Deserialize};
    use protobuf::Message;
    use failure::ensure;
    use std::fs::File;
    use std::io::prelude::*;

    use crate::helpers::{Result};

    #[derive(Deserialize, Debug)]
    struct TestVector {
        pub root: String,
        pub existence: String,
    }

    fn verify_test_vector(filename: &str, spec: &proofs::ProofSpec) -> Result<()> {
        let mut file = File::open(filename)?;
        let mut contents = String::new();
        file.read_to_string(&mut contents)?;

        let data: TestVector = serde_json::from_str(&contents)?;
        let proto_bin = hex::decode(&data.existence)?;
        let root = hex::decode(data.root)?;

        let mut parsed = proofs::ExistenceProof::new();
        parsed.merge_from_bytes(&proto_bin)?;
        let valid = verify_existence(&parsed, spec, root, &parsed.key, &parsed.value)?;
        ensure!(valid, "invalid test vector");
        Ok(())
    }

    #[test]
    fn test_vector_iavl1() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/existence1.json", &spec)
    }

    #[test]
    fn test_vector_iavl2() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/existence2.json", &spec)
    }

    #[test]
    fn test_vector_iavl3() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/existence3.json", &spec)
    }

    #[test]
    fn test_vector_iavl4() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/existence4.json", &spec)
    }

    #[test]
    fn test_vector_tendermint1() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/existence1.json", &spec)
    }

    #[test]
    fn test_vector_tendermint2() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/existence2.json", &spec)
    }

    #[test]
    fn test_vector_tendermint3() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/existence3.json", &spec)
    }

    #[test]
    fn test_vector_tendermint4() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/existence4.json", &spec)
    }

}