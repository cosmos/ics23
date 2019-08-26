use std::collections::HashMap;
use std::hash::BuildHasher;

use crate::compress::{decompress, is_compressed};
use crate::ics23;
use crate::verify::{verify_existence, verify_non_existence, CommitmentRoot};

// Use CommitmentRoot vs &[u8] to stick with ics naming
#[allow(clippy::ptr_arg)]
pub fn verify_membership(
    proof: &ics23::CommitmentProof,
    spec: &ics23::ProofSpec,
    root: &CommitmentRoot,
    key: &[u8],
    value: &[u8],
) -> bool {
    // ugly attempt to conditionally decompress...
    let mut proof = proof;
    let my_proof;
    if is_compressed(proof) {
        if let Ok(p) = decompress(proof) {
            my_proof = p;
            proof = &my_proof;
        } else {
            return false;
        }
    }

    //    if let Some(ics23::commitment_proof::Proof::Exist(ex)) = &proof.proof {
    if let Some(ex) = get_exist_proof(proof, key) {
        let valid = verify_existence(&ex, spec, root, key, value);
        valid.is_ok()
    } else {
        false
    }
}

// Use CommitmentRoot vs &[u8] to stick with ics naming
#[allow(clippy::ptr_arg)]
pub fn verify_non_membership(
    proof: &ics23::CommitmentProof,
    spec: &ics23::ProofSpec,
    root: &CommitmentRoot,
    key: &[u8],
) -> bool {
    // ugly attempt to conditionally decompress...
    let mut proof = proof;
    let my_proof;
    if is_compressed(proof) {
        if let Ok(p) = decompress(proof) {
            my_proof = p;
            proof = &my_proof;
        } else {
            return false;
        }
    }

    if let Some(non) = get_nonexist_proof(proof, key) {
        let valid = verify_non_existence(&non, spec, root, key);
        valid.is_ok()
    } else {
        false
    }
}

#[allow(clippy::ptr_arg)]
pub fn verify_batch_membership<G: BuildHasher>(
    proof: &ics23::CommitmentProof,
    spec: &ics23::ProofSpec,
    root: &CommitmentRoot,
    items: HashMap<&[u8], &[u8], G>,
) -> bool {
    // ugly attempt to conditionally decompress...
    let mut proof = proof;
    let my_proof;
    if is_compressed(proof) {
        if let Ok(p) = decompress(proof) {
            my_proof = p;
            proof = &my_proof;
        } else {
            return false;
        }
    }

    items
        .iter()
        .all(|(key, value)| verify_membership(proof, spec, root, key, value))
}

#[allow(clippy::ptr_arg)]
pub fn verify_batch_non_membership(
    proof: &ics23::CommitmentProof,
    spec: &ics23::ProofSpec,
    root: &CommitmentRoot,
    keys: &[&[u8]],
) -> bool {
    // ugly attempt to conditionally decompress...
    let mut proof = proof;
    let my_proof;
    if is_compressed(proof) {
        if let Ok(p) = decompress(proof) {
            my_proof = p;
            proof = &my_proof;
        } else {
            return false;
        }
    }

    keys.iter()
        .all(|key| verify_non_membership(proof, spec, root, key))
}

fn get_exist_proof<'a>(
    proof: &'a ics23::CommitmentProof,
    key: &[u8],
) -> Option<&'a ics23::ExistenceProof> {
    match &proof.proof {
        Some(ics23::commitment_proof::Proof::Exist(ex)) => Some(ex),
        Some(ics23::commitment_proof::Proof::Batch(batch)) => {
            for entry in &batch.entries {
                if let Some(ics23::batch_entry::Proof::Exist(ex)) = &entry.proof {
                    if ex.key == key {
                        return Some(ex);
                    }
                }
            }
            None
        }
        _ => None,
    }
}

fn get_nonexist_proof<'a>(
    proof: &'a ics23::CommitmentProof,
    key: &[u8],
) -> Option<&'a ics23::NonExistenceProof> {
    match &proof.proof {
        Some(ics23::commitment_proof::Proof::Nonexist(non)) => Some(non),
        Some(ics23::commitment_proof::Proof::Batch(batch)) => {
            for entry in &batch.entries {
                if let Some(ics23::batch_entry::Proof::Nonexist(non)) = &entry.proof {
                    // use iter/all - true if None, must check if Some
                    if non.left.iter().all(|x| x.key.as_slice() < key)
                        && non.right.iter().all(|x| x.key.as_slice() > key)
                    {
                        return Some(non);
                    }
                }
            }
            None
        }
        _ => None,
    }
}

#[warn(clippy::ptr_arg)]
pub fn iavl_spec() -> ics23::ProofSpec {
    let leaf = ics23::LeafOp {
        hash: ics23::HashOp::Sha256.into(),
        prehash_key: 0,
        prehash_value: ics23::HashOp::Sha256.into(),
        length: ics23::LengthOp::VarProto.into(),
        prefix: vec![0_u8],
    };
    let inner = ics23::InnerSpec {
        child_order: vec![0, 1],
        min_prefix_length: 4,
        max_prefix_length: 12,
        child_size: 33,
        empty_child: vec![],
    };
    ics23::ProofSpec {
        leaf_spec: Some(leaf),
        inner_spec: Some(inner),
    }
}

pub fn tendermint_spec() -> ics23::ProofSpec {
    let leaf = ics23::LeafOp {
        hash: ics23::HashOp::Sha256.into(),
        prehash_key: 0,
        prehash_value: ics23::HashOp::Sha256.into(),
        length: ics23::LengthOp::VarProto.into(),
        prefix: vec![0_u8],
    };
    let inner = ics23::InnerSpec {
        child_order: vec![0, 1],
        min_prefix_length: 1,
        max_prefix_length: 1,
        child_size: 32,
        empty_child: vec![],
    };
    ics23::ProofSpec {
        leaf_spec: Some(leaf),
        inner_spec: Some(inner),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use failure::{bail, ensure};
    use prost::Message;
    use serde::Deserialize;
    use std::fs::File;
    use std::io::prelude::*;
    use std::vec::Vec;

    use crate::compress::compress;
    use crate::helpers::Result;

    #[derive(Deserialize, Debug)]
    struct TestVector {
        pub root: String,
        pub proof: String,
        pub key: String,
        pub value: String,
    }

    struct RefData {
        pub root: Vec<u8>,
        pub key: Vec<u8>,
        pub value: Option<Vec<u8>>,
    }

    fn load_file(filename: &str) -> Result<(ics23::CommitmentProof, RefData)> {
        let mut file = File::open(filename)?;
        let mut contents = String::new();
        file.read_to_string(&mut contents)?;

        let data: TestVector = serde_json::from_str(&contents)?;
        let proto_bin = hex::decode(&data.proof)?;
        let mut parsed = ics23::CommitmentProof { proof: None };
        parsed.merge(&proto_bin)?;

        let root = hex::decode(data.root)?;
        let key = hex::decode(data.key)?;
        let value = if data.value.is_empty() {
            None
        } else {
            Some(hex::decode(data.value)?)
        };
        let data = RefData { root, key, value };

        Ok((parsed, data))
    }

    fn verify_test_vector(filename: &str, spec: &ics23::ProofSpec) -> Result<()> {
        let (proof, data) = load_file(filename)?;

        if let Some(value) = data.value {
            let valid = super::verify_membership(&proof, spec, &data.root, &data.key, &value);
            ensure!(valid, "invalid test vector");
            Ok(())
        } else {
            let valid = super::verify_non_membership(&proof, spec, &data.root, &data.key);
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

    fn load_batch(files: &[&str]) -> Result<(ics23::CommitmentProof, Vec<RefData>)> {
        let mut entries = Vec::new();
        let mut data = Vec::new();

        for &file in files {
            let (proof, datum) = load_file(file)?;
            data.push(datum);
            match proof.proof {
                Some(ics23::commitment_proof::Proof::Nonexist(non)) => {
                    entries.push(ics23::BatchEntry {
                        proof: Some(ics23::batch_entry::Proof::Nonexist(non)),
                    })
                }
                Some(ics23::commitment_proof::Proof::Exist(ex)) => {
                    entries.push(ics23::BatchEntry {
                        proof: Some(ics23::batch_entry::Proof::Exist(ex)),
                    })
                }
                _ => bail!("unknown proof type to batch"),
            }
        }

        let batch = ics23::CommitmentProof {
            proof: Some(ics23::commitment_proof::Proof::Batch(ics23::BatchProof {
                entries,
            })),
        };

        Ok((batch, data))
    }

    fn verify_batch(
        spec: &ics23::ProofSpec,
        proof: &ics23::CommitmentProof,
        data: &RefData,
    ) -> Result<()> {
        if let Some(value) = &data.value {
            let valid = super::verify_membership(&proof, spec, &data.root, &data.key, &value);
            ensure!(valid, "invalid test vector");
            let mut items = HashMap::new();
            items.insert(data.key.as_slice(), value.as_slice());
            let valid = super::verify_batch_membership(&proof, spec, &data.root, items);
            ensure!(valid, "invalid test vector");
            Ok(())
        } else {
            let valid = super::verify_non_membership(&proof, spec, &data.root, &data.key);
            ensure!(valid, "invalid test vector");
            let keys = &[data.key.as_slice()];
            let valid = super::verify_batch_non_membership(&proof, spec, &data.root, keys);
            ensure!(valid, "invalid test vector");
            Ok(())
        }
    }

    #[test]
    fn test_vector_iavl_batch_exist() -> Result<()> {
        let spec = iavl_spec();
        let (proof, data) = load_batch(&[
            "../testdata/iavl/exist_left.json",
            "../testdata/iavl/exist_right.json",
            "../testdata/iavl/exist_middle.json",
            "../testdata/iavl/nonexist_left.json",
            "../testdata/iavl/nonexist_right.json",
            "../testdata/iavl/nonexist_middle.json",
        ])?;
        verify_batch(&spec, &proof, &data[0])
    }

    #[test]
    fn compressed_iavl_batch_exist() -> Result<()> {
        let spec = iavl_spec();
        let (proof, data) = load_batch(&[
            "../testdata/iavl/exist_left.json",
            "../testdata/iavl/exist_right.json",
            "../testdata/iavl/exist_middle.json",
            "../testdata/iavl/nonexist_left.json",
            "../testdata/iavl/nonexist_right.json",
            "../testdata/iavl/nonexist_middle.json",
        ])?;
        let comp = compress(&proof)?;
        verify_batch(&spec, &comp, &data[0])
    }

    #[test]
    fn test_vector_iavl_batch_nonexist() -> Result<()> {
        let spec = iavl_spec();
        let (proof, data) = load_batch(&[
            "../testdata/iavl/exist_left.json",
            "../testdata/iavl/exist_right.json",
            "../testdata/iavl/exist_middle.json",
            "../testdata/iavl/nonexist_left.json",
            "../testdata/iavl/nonexist_right.json",
            "../testdata/iavl/nonexist_middle.json",
        ])?;
        verify_batch(&spec, &proof, &data[4])
    }

    #[test]
    fn compressed_iavl_batch_nonexist() -> Result<()> {
        let spec = iavl_spec();
        let (proof, data) = load_batch(&[
            "../testdata/iavl/exist_left.json",
            "../testdata/iavl/exist_right.json",
            "../testdata/iavl/exist_middle.json",
            "../testdata/iavl/nonexist_left.json",
            "../testdata/iavl/nonexist_right.json",
            "../testdata/iavl/nonexist_middle.json",
        ])?;
        let comp = compress(&proof)?;
        verify_batch(&spec, &comp, &data[4])
    }

    #[test]
    fn test_vector_tendermint_batch_exist() -> Result<()> {
        let spec = tendermint_spec();
        let (proof, data) = load_batch(&[
            "../testdata/tendermint/exist_left.json",
            "../testdata/tendermint/exist_right.json",
            "../testdata/tendermint/exist_middle.json",
            "../testdata/tendermint/nonexist_left.json",
            "../testdata/tendermint/nonexist_right.json",
            "../testdata/tendermint/nonexist_middle.json",
        ])?;
        verify_batch(&spec, &proof, &data[2])
    }

    #[test]
    fn test_vector_tendermint_batch_nonexist() -> Result<()> {
        let spec = tendermint_spec();
        let (proof, data) = load_batch(&[
            "../testdata/tendermint/exist_left.json",
            "../testdata/tendermint/exist_right.json",
            "../testdata/tendermint/exist_middle.json",
            "../testdata/tendermint/nonexist_left.json",
            "../testdata/tendermint/nonexist_right.json",
            "../testdata/tendermint/nonexist_middle.json",
        ])?;
        verify_batch(&spec, &proof, &data[5])
    }
}
