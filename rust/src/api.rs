use alloc::collections::btree_map::BTreeMap;
use alloc::vec;

use anyhow::{anyhow, ensure};
use bytes::{Buf, Bytes};

use crate::compress::{decompress, is_compressed};
use crate::helpers::Result;
use crate::host_functions::HostFunctionsProvider;
use crate::ics23;
use crate::verify::{verify_existence, verify_non_existence, CommitmentRoot};

// Use CommitmentRoot vs &[u8] to stick with ics naming
#[allow(clippy::ptr_arg)]
pub fn verify_membership<H: HostFunctionsProvider>(
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
        let valid = verify_existence::<H>(ex, spec, root, key, value);
        valid.is_ok()
    } else {
        false
    }
}

// Use CommitmentRoot vs &[u8] to stick with ics naming
#[allow(clippy::ptr_arg)]
pub fn verify_non_membership<H: HostFunctionsProvider>(
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
        let valid = verify_non_existence::<H>(non, spec, root, key);
        valid.is_ok()
    } else {
        false
    }
}

#[allow(clippy::ptr_arg)]
pub fn verify_batch_membership<H: HostFunctionsProvider>(
    proof: &ics23::CommitmentProof,
    spec: &ics23::ProofSpec,
    root: &CommitmentRoot,
    items: BTreeMap<&[u8], &[u8]>,
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
        .all(|(key, value)| verify_membership::<H>(proof, spec, root, key, value))
}

#[allow(clippy::ptr_arg)]
pub fn verify_batch_non_membership<H: HostFunctionsProvider>(
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
        .all(|key| verify_non_membership::<H>(proof, spec, root, key))
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
        hash: ics23::HashOp::Sha256.into(),
    };
    ics23::ProofSpec {
        leaf_spec: Some(leaf),
        inner_spec: Some(inner),
        min_depth: 0,
        max_depth: 0,
    }
}

fn read_varint<B: Buf>(buf: &mut B) -> Option<i64> {
    let ux = prost::encoding::decode_varint(buf).ok()?;
    let mut x: i64 = (ux >> 1).try_into().ok()?;
    if ux & 1 != 0 {
        x = !x;
    }
    Some(x)
}

fn ensure_iavl_prefix(prefix: &[u8], min_height: i64) -> Result<Option<usize>> {
    let mut prefix = Bytes::copy_from_slice(prefix);

    let height = read_varint(&mut prefix)
        .ok_or_else(|| anyhow!("error decoding height in IAVL prefix encoding"))?;
    ensure!(height >= min_height, "wrong height in IAVL prefix encoding");

    let size = read_varint(&mut prefix)
        .ok_or_else(|| anyhow!("error decoding size in IAVL prefix encoding"))?;
    ensure!(size >= 0, "wrong size in IAVL prefix encoding");

    let version = read_varint(&mut prefix)
        .ok_or_else(|| anyhow!("error decoding version in IAVL prefix encoding"))?;
    ensure!(version >= 0, "wrong version in IAVL prefix encoding");

    Ok(Some(prefix.remaining()))
}

fn is_iavl_spec(spec: &ics23::ProofSpec) -> bool {
    // we use a relaxed impl that over-declares equality
    spec_equals(spec, &iavl_spec())
}

fn spec_equals(left: &ics23::ProofSpec, right: &ics23::ProofSpec) -> bool {
    match (
        &left.leaf_spec,
        &left.inner_spec,
        &right.leaf_spec,
        &right.inner_spec,
    ) {
        (
            Some(left_leaf_spec),
            Some(left_inner_spec),
            Some(right_leaf_spec),
            Some(right_inner_spec),
        ) => {
            leaf_spec_equals(left_leaf_spec, right_leaf_spec)
                && inner_spec_equals(left_inner_spec, right_inner_spec)
        }
        _ => false,
    }
}

fn leaf_spec_equals(left: &ics23::LeafOp, right: &ics23::LeafOp) -> bool {
    left.hash == right.hash
        && left.prehash_key == right.prehash_key
        && left.prehash_value == right.prehash_value
        && left.length == right.length
}

fn inner_spec_equals(left: &ics23::InnerSpec, right: &ics23::InnerSpec) -> bool {
    left.hash == right.hash
        && left.min_prefix_length == right.min_prefix_length
        && left.max_prefix_length == right.max_prefix_length
        && left.child_size == right.child_size
        && left.child_order.len() == right.child_order.len()
}

pub(crate) fn ensure_leaf_prefix(prefix: &[u8], spec: &ics23::ProofSpec) -> Result<()> {
    if !is_iavl_spec(spec) {
        return Ok(());
    }

    let remaining = ensure_iavl_prefix(prefix, 0)?;
    match remaining {
        None => Ok(()),
        Some(remaining_bytes) => {
            ensure!(remaining_bytes == 0, "bad prefix in leaf");
            Ok(())
        }
    }
}

pub(crate) fn ensure_inner_prefix(
    prefix: &[u8],
    spec: &ics23::ProofSpec,
    min_height: i64,
    hash_op: i32,
) -> Result<()> {
    if !is_iavl_spec(spec) {
        return Ok(());
    }

    let remaining = ensure_iavl_prefix(prefix, min_height)?;
    match remaining {
        None => Ok(()),
        Some(remaining_bytes) => {
            // 1 byte due to containing length prefix for left hash.
            // 33 bytes due to IAVL length prefix + left hash + next IAVL legnth prefix
            ensure!(
                remaining_bytes == 1 || remaining_bytes == 34,
                "bad prefix in layer {}",
                min_height
            );
            ensure!(hash_op == ics23::HashOp::Sha256 as i32, "bad hash op");
            Ok(())
        }
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
        hash: ics23::HashOp::Sha256.into(),
    };
    ics23::ProofSpec {
        leaf_spec: Some(leaf),
        inner_spec: Some(inner),
        min_depth: 0,
        max_depth: 0,
    }
}

pub fn smt_spec() -> ics23::ProofSpec {
    let leaf = ics23::LeafOp {
        hash: ics23::HashOp::Sha256.into(),
        prehash_key: 0,
        prehash_value: ics23::HashOp::Sha256.into(),
        length: 0,
        prefix: vec![0_u8],
    };
    let inner = ics23::InnerSpec {
        child_order: vec![0, 1],
        min_prefix_length: 1,
        max_prefix_length: 1,
        child_size: 32,
        empty_child: vec![0; 32],
        hash: ics23::HashOp::Sha256.into(),
    };
    ics23::ProofSpec {
        leaf_spec: Some(leaf),
        inner_spec: Some(inner),
        min_depth: 0,
        max_depth: 0,
    }
}

#[cfg(feature = "std")]
#[cfg(test)]
mod tests {
    use super::*;

    use alloc::string::String;
    use alloc::vec::Vec;
    use anyhow::{bail, ensure};
    use bytes::{BufMut, BytesMut};
    use prost::Message;
    use serde::Deserialize;
    #[cfg(feature = "std")]
    use std::fs::File;
    #[cfg(feature = "std")]
    use std::io::prelude::*;

    use crate::compress::compress;
    use crate::helpers::Result;
    use crate::host_functions::host_functions_impl::HostFunctionsManager;

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

    #[cfg(feature = "std")]
    fn load_file(filename: &str) -> Result<(ics23::CommitmentProof, RefData)> {
        let mut file = File::open(filename)?;
        let mut contents = String::new();
        file.read_to_string(&mut contents)?;

        let data: TestVector = serde_json::from_str(&contents)?;
        let proto_bin = hex::decode(&data.proof)?;
        let mut parsed = ics23::CommitmentProof { proof: None };
        parsed.merge(proto_bin.as_slice())?;

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

    #[cfg(feature = "std")]
    fn verify_test_vector(filename: &str, spec: &ics23::ProofSpec) -> Result<()> {
        let (proof, data) = load_file(filename)?;

        if let Some(value) = data.value {
            let valid = verify_membership::<HostFunctionsManager>(
                &proof, spec, &data.root, &data.key, &value,
            );
            ensure!(valid, "invalid test vector");
            Ok(())
        } else {
            let valid =
                verify_non_membership::<HostFunctionsManager>(&proof, spec, &data.root, &data.key);
            ensure!(valid, "invalid test vector");
            Ok(())
        }
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_iavl_left() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/exist_left.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_iavl_right() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/exist_right.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_iavl_middle() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/exist_middle.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_iavl_left_non() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/nonexist_left.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_iavl_right_non() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/nonexist_right.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_iavl_middle_non() -> Result<()> {
        let spec = iavl_spec();
        verify_test_vector("../testdata/iavl/nonexist_middle.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_tendermint_left() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/exist_left.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_tendermint_right() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/exist_right.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_tendermint_middle() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/exist_middle.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_tendermint_left_non() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/nonexist_left.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_tendermint_right_non() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/nonexist_right.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_tendermint_middle_non() -> Result<()> {
        let spec = tendermint_spec();
        verify_test_vector("../testdata/tendermint/nonexist_middle.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_left() -> Result<()> {
        let spec = smt_spec();
        verify_test_vector("../testdata/smt/exist_left.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_right() -> Result<()> {
        let spec = smt_spec();
        verify_test_vector("../testdata/smt/exist_right.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_middle() -> Result<()> {
        let spec = smt_spec();
        verify_test_vector("../testdata/smt/exist_middle.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_left_non() -> Result<()> {
        let spec = smt_spec();
        verify_test_vector("../testdata/smt/nonexist_left.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_right_non() -> Result<()> {
        let spec = smt_spec();
        verify_test_vector("../testdata/smt/nonexist_right.json", &spec)
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_middle_non() -> Result<()> {
        let spec = smt_spec();
        verify_test_vector("../testdata/smt/nonexist_middle.json", &spec)
    }

    #[cfg(feature = "std")]
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
            let valid = verify_membership::<HostFunctionsManager>(
                proof, spec, &data.root, &data.key, value,
            );
            ensure!(valid, "invalid test vector");
            let mut items = BTreeMap::new();
            items.insert(data.key.as_slice(), value.as_slice());
            let valid =
                verify_batch_membership::<HostFunctionsManager>(proof, spec, &data.root, items);
            ensure!(valid, "invalid test vector");
            Ok(())
        } else {
            let valid =
                verify_non_membership::<HostFunctionsManager>(proof, spec, &data.root, &data.key);
            ensure!(valid, "invalid test vector");
            let keys = &[data.key.as_slice()];
            let valid =
                verify_batch_non_membership::<HostFunctionsManager>(proof, spec, &data.root, keys);
            ensure!(valid, "invalid test vector");
            Ok(())
        }
    }

    #[test]
    #[cfg(feature = "std")]
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
    #[cfg(feature = "std")]
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
    #[cfg(feature = "std")]
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
    #[cfg(feature = "std")]
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
    #[cfg(feature = "std")]
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
    #[cfg(feature = "std")]
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

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_batch_exist() -> Result<()> {
        let spec = smt_spec();
        let (proof, data) = load_batch(&[
            "../testdata/smt/exist_left.json",
            "../testdata/smt/exist_right.json",
            "../testdata/smt/exist_middle.json",
            "../testdata/smt/nonexist_left.json",
            "../testdata/smt/nonexist_right.json",
            "../testdata/smt/nonexist_middle.json",
        ])?;
        verify_batch(&spec, &proof, &data[0])
    }

    #[test]
    #[cfg(feature = "std")]
    fn compressed_smt_batch_exist() -> Result<()> {
        let spec = smt_spec();
        let (proof, data) = load_batch(&[
            "../testdata/smt/exist_left.json",
            "../testdata/smt/exist_right.json",
            "../testdata/smt/exist_middle.json",
            "../testdata/smt/nonexist_left.json",
            "../testdata/smt/nonexist_right.json",
            "../testdata/smt/nonexist_middle.json",
        ])?;
        let comp = compress(&proof)?;
        verify_batch(&spec, &comp, &data[0])
    }

    #[test]
    #[cfg(feature = "std")]
    fn test_vector_smt_batch_nonexist() -> Result<()> {
        let spec = smt_spec();
        let (proof, data) = load_batch(&[
            "../testdata/smt/exist_left.json",
            "../testdata/smt/exist_right.json",
            "../testdata/smt/exist_middle.json",
            "../testdata/smt/nonexist_left.json",
            "../testdata/smt/nonexist_right.json",
            "../testdata/smt/nonexist_middle.json",
        ])?;
        verify_batch(&spec, &proof, &data[4])
    }

    #[test]
    #[cfg(feature = "std")]
    fn compressed_smt_batch_nonexist() -> Result<()> {
        let spec = smt_spec();
        let (proof, data) = load_batch(&[
            "../testdata/smt/exist_left.json",
            "../testdata/smt/exist_right.json",
            "../testdata/smt/exist_middle.json",
            "../testdata/smt/nonexist_left.json",
            "../testdata/smt/nonexist_right.json",
            "../testdata/smt/nonexist_middle.json",
        ])?;
        let comp = compress(&proof)?;
        verify_batch(&spec, &comp, &data[4])
    }

    #[test]
    fn test_read_varint() {
        fn put_varint<B: BufMut>(x: i64, buf: &mut B) {
            let mut ux = (x as u64) << 1;
            if x < 0 {
                ux = !ux;
            }
            prost::encoding::encode_varint(ux, buf);
        }

        fn test_varint(x: i64) {
            let mut buf = BytesMut::new();
            put_varint(x, &mut buf);
            let y = read_varint(&mut buf).unwrap();
            assert_eq!(x, y)
        }

        const TESTS: [i64; 18] = [
            -1i64 << 63,
            (-1 << 63) + 1,
            -1,
            0,
            1,
            2,
            10,
            20,
            63,
            64,
            65,
            127,
            128,
            129,
            255,
            256,
            257,
            1 << (63 - 1),
        ];

        for x in TESTS {
            test_varint(x);
            test_varint(0i64.wrapping_sub(x));
        }

        let mut x = 0x7i64;
        while x != 0 {
            test_varint(x);
            test_varint(0i64.wrapping_sub(x));
            x <<= 1;
        }
    }
}
