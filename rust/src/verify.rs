// we want to name functions verify_* to match ics23
#![allow(clippy::module_name_repetitions)]

use anyhow::{bail, ensure};

use crate::helpers::Result;
use crate::ics23;
use crate::ops::{apply_inner, apply_leaf};

pub type CommitmentRoot = ::std::vec::Vec<u8>;

pub fn verify_existence(
    proof: &ics23::ExistenceProof,
    spec: &ics23::ProofSpec,
    root: &[u8],
    key: &[u8],
    value: &[u8],
) -> Result<()> {
    check_existence_spec(proof, spec)?;
    ensure!(proof.key == key, "Provided key doesn't match proof");
    ensure!(proof.value == value, "Provided value doesn't match proof");

    let calc = calculate_existence_root(&proof)?;
    ensure!(calc == root, "Root hash doesn't match");
    Ok(())
}

pub fn verify_non_existence(
    proof: &ics23::NonExistenceProof,
    spec: &ics23::ProofSpec,
    root: &[u8],
    key: &[u8],
) -> Result<()> {
    if let Some(left) = &proof.left {
        verify_existence(&left, spec, root, &left.key, &left.value)?;
        ensure!(key > left.key.as_slice(), "left key isn't before key");
    }
    if let Some(right) = &proof.right {
        verify_existence(&right, spec, root, &right.key, &right.value)?;
        ensure!(key < right.key.as_slice(), "right key isn't after key");
    }

    if let Some(inner) = &spec.inner_spec {
        match (&proof.left, &proof.right) {
            (Some(left), None) => ensure_right_most(inner, &left.path),
            (None, Some(right)) => ensure_left_most(inner, &right.path),
            (Some(left), Some(right)) => ensure_left_neighbor(inner, &left.path, &right.path),
            (None, None) => bail!("neither left nor right proof defined"),
        }
    } else {
        bail!("Inner Spec missing")
    }
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
pub fn calculate_existence_root(proof: &ics23::ExistenceProof) -> Result<CommitmentRoot> {
    ensure!(!proof.key.is_empty(), "Existence proof must have key set");
    ensure!(
        !proof.value.is_empty(),
        "Existence proof must have value set"
    );

    if let Some(leaf_node) = &proof.leaf {
        let mut hash = apply_leaf(leaf_node, &proof.key, &proof.value)?;
        for step in &proof.path {
            hash = apply_inner(step, &hash)?;
        }
        Ok(hash)
    } else {
        bail!("No leaf operation set")
    }
}

fn check_existence_spec(proof: &ics23::ExistenceProof, spec: &ics23::ProofSpec) -> Result<()> {
    if let (Some(leaf), Some(leaf_spec)) = (&proof.leaf, &spec.leaf_spec) {
        ensure_leaf(leaf, leaf_spec)?;
        // ensure min/max depths
        if spec.min_depth != 0 {
            ensure!(
                proof.path.len() >= spec.min_depth as usize,
                "Too few InnerOps: {}",
                proof.path.len(),
            );
            ensure!(
                proof.path.len() <= spec.max_depth as usize,
                "Too many InnerOps: {}",
                proof.path.len(),
            );
        }
        for step in &proof.path {
            ensure_inner(step, spec)?;
        }
        Ok(())
    } else {
        bail!("Leaf and Leaf Spec must be set")
    }
}

fn ensure_leaf(leaf: &ics23::LeafOp, leaf_spec: &ics23::LeafOp) -> Result<()> {
    ensure!(
        leaf_spec.hash == leaf.hash,
        "Unexpected hashOp: {:?}",
        leaf.hash
    );
    ensure!(
        leaf_spec.prehash_key == leaf.prehash_key,
        "Unexpected prehash_key: {:?}",
        leaf.prehash_key
    );
    ensure!(
        leaf_spec.prehash_value == leaf.prehash_value,
        "Unexpected prehash_value: {:?}",
        leaf.prehash_value
    );
    ensure!(
        leaf_spec.length == leaf.length,
        "Unexpected lengthOp: {:?}",
        leaf.length
    );
    ensure!(
        has_prefix(&leaf_spec.prefix, &leaf.prefix),
        "Incorrect prefix on leaf"
    );
    Ok(())
}

fn has_prefix(prefix: &[u8], data: &[u8]) -> bool {
    if prefix.len() > data.len() {
        return false;
    }
    prefix == &data[..prefix.len()]
}

fn ensure_inner(inner: &ics23::InnerOp, spec: &ics23::ProofSpec) -> Result<()> {
    match (&spec.leaf_spec, &spec.inner_spec) {
        (Some(leaf_spec), Some(inner_spec)) => {
            ensure!(
                inner.hash == inner_spec.hash,
                "Unexpected hashOp: {:?}",
                inner.hash,
            );
            ensure!(
                !has_prefix(&leaf_spec.prefix, &inner.prefix),
                "Inner node with leaf prefix",
            );
            ensure!(
                inner.prefix.len() >= (inner_spec.min_prefix_length as usize),
                "Inner prefix too short: {}",
                inner.prefix.len(),
            );
            let max_left_child_bytes =
                (inner_spec.child_order.len() - 1) as i32 * inner_spec.child_size;
            ensure!(
                inner.prefix.len()
                    <= (inner_spec.max_prefix_length + max_left_child_bytes) as usize,
                "Inner prefix too short: {}",
                inner.prefix.len(),
            );
            Ok(())
        }
        (_, _) => bail!("Spec requires both leaf_spec and inner_spec"),
    }
}

fn ensure_left_most(spec: &ics23::InnerSpec, path: &[ics23::InnerOp]) -> Result<()> {
    let pad = get_padding(spec, 0)?;
    for step in path {
        if !has_padding(step, &pad) {
            bail!("step not leftmost")
        }
    }
    Ok(())
}

fn ensure_right_most(spec: &ics23::InnerSpec, path: &[ics23::InnerOp]) -> Result<()> {
    let idx = spec.child_order.len() - 1;
    let pad = get_padding(spec, idx as i32)?;
    for step in path {
        if !has_padding(step, &pad) {
            bail!("step not leftmost")
        }
    }
    Ok(())
}

fn ensure_left_neighbor(
    spec: &ics23::InnerSpec,
    left: &[ics23::InnerOp],
    right: &[ics23::InnerOp],
) -> Result<()> {
    let mut mut_left = Vec::from(left);
    let mut mut_right = Vec::from(right);

    let mut top_left = mut_left.pop().unwrap();
    let mut top_right = mut_right.pop().unwrap();

    while top_left.prefix == top_right.prefix && top_left.suffix == top_right.suffix {
        top_left = mut_left.pop().unwrap();
        top_right = mut_right.pop().unwrap();
    }

    if !is_left_step(spec, &top_left, &top_right)? {
        bail!("Not left neighbor at first divergent step");
    }

    ensure_right_most(spec, &mut_left)?;
    ensure_left_most(spec, &mut_right)
}

fn is_left_step(
    spec: &ics23::InnerSpec,
    left: &ics23::InnerOp,
    right: &ics23::InnerOp,
) -> Result<bool> {
    let left_idx = order_from_padding(spec, left)?;
    let right_idx = order_from_padding(spec, right)?;
    Ok(left_idx + 1 == right_idx)
}

fn order_from_padding(spec: &ics23::InnerSpec, op: &ics23::InnerOp) -> Result<i32> {
    let len = spec.child_order.len() as i32;
    for branch in 0..len {
        let padding = get_padding(spec, branch)?;
        if has_padding(op, &padding) {
            return Ok(branch);
        }
    }
    bail!("padding doesn't match any branch");
}

struct Padding {
    min_prefix: usize,
    max_prefix: usize,
    suffix: usize,
}

fn has_padding(op: &ics23::InnerOp, pad: &Padding) -> bool {
    (op.prefix.len() >= pad.min_prefix)
        && (op.prefix.len() <= pad.max_prefix)
        && (op.suffix.len() == pad.suffix)
}

fn get_padding(spec: &ics23::InnerSpec, branch: i32) -> Result<Padding> {
    if let Some(&idx) = spec.child_order.iter().find(|&&x| x == branch) {
        let prefix = idx * spec.child_size;
        let suffix = spec.child_size as usize * (spec.child_order.len() - 1 - idx as usize);
        Ok(Padding {
            min_prefix: (prefix + spec.min_prefix_length) as usize,
            max_prefix: (prefix + spec.max_prefix_length) as usize,
            suffix,
        })
    } else {
        bail!("Branch {} not found", branch);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::api;
    use crate::ics23::{ExistenceProof, HashOp, InnerOp, LeafOp, LengthOp, ProofSpec};
    use std::collections::HashMap;

    #[test]
    fn calculate_root_from_leaf() {
        let leaf = ics23::LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: LengthOp::VarProto.into(),
            prefix: vec![],
        };

        let proof = ics23::ExistenceProof {
            key: b"food".to_vec(),
            value: b"some longer text".to_vec(),
            leaf: Some(leaf),
            path: vec![],
        };

        let expected =
            hex::decode("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265")
                .unwrap();
        assert_eq!(
            expected,
            calculate_existence_root(&proof).unwrap(),
            "invalid root hash"
        );
    }

    #[test]
    fn calculate_root_from_leaf_and_inner() {
        let leaf = ics23::LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: LengthOp::VarProto.into(),
            prefix: vec![],
        };

        let inner = ics23::InnerOp {
            hash: HashOp::Sha256.into(),
            prefix: hex::decode("deadbeef00cafe00").unwrap(),
            suffix: vec![],
        };

        let proof = ics23::ExistenceProof {
            key: b"food".to_vec(),
            value: b"some longer text".to_vec(),
            leaf: Some(leaf),
            path: vec![inner],
        };

        let expected =
            hex::decode("836ea236a6902a665c2a004c920364f24cad52ded20b1e4f22c3179bfe25b2a9")
                .unwrap();
        assert_eq!(
            expected,
            calculate_existence_root(&proof).unwrap(),
            "invalid root hash"
        );
    }

    #[derive(Debug, Clone)]
    struct ExistenceCase {
        proof: ExistenceProof,
        spec: ProofSpec,
        valid: bool,
    }

    #[test]
    fn enforce_existence_spec() {
        let leaf = LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: HashOp::Sha256.into(),
            length: LengthOp::VarProto.into(),
            prefix: vec![0_u8],
        };
        let invalid_leaf = LeafOp {
            hash: HashOp::Sha512.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: LengthOp::VarProto.into(),
            prefix: vec![0_u8],
        };

        let valid_inner = InnerOp {
            hash: HashOp::Sha256.into(),
            prefix: hex::decode("deadbeef00cafe00").unwrap(),
            suffix: vec![],
        };
        let invalid_inner = InnerOp {
            hash: HashOp::Sha256.into(),
            prefix: hex::decode("aa").unwrap(),
            suffix: vec![],
        };
        let invalid_inner_hash = InnerOp {
            hash: HashOp::Sha512.into(),
            prefix: hex::decode("deadbeef00cafe00").unwrap(),
            suffix: vec![],
        };

        let mut depth_limited_spec = api::iavl_spec();
        depth_limited_spec.min_depth = 2;
        depth_limited_spec.max_depth = 4;

        let cases: HashMap<&'static str, ExistenceCase> = [
            (
                "empty proof fails",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: None,
                        path: vec![],
                    },
                    spec: api::iavl_spec(),
                    valid: false,
                },
            ),
            (
                "accepts one valid leaf",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf.clone()),
                        path: vec![],
                    },
                    spec: api::iavl_spec(),
                    valid: true,
                },
            ),
            (
                "rejects invalid leaf",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(invalid_leaf),
                        path: vec![],
                    },
                    spec: api::iavl_spec(),
                    valid: false,
                },
            ),
            (
                "rejects only inner (no leaf)",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: None,
                        path: vec![valid_inner.clone()],
                    },
                    spec: api::iavl_spec(),
                    valid: false,
                },
            ),
            (
                "accepts leaf and valid inner",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf.clone()),
                        path: vec![valid_inner.clone()],
                    },
                    spec: api::iavl_spec(),
                    valid: true,
                },
            ),
            (
                "rejects invalid inner (prefix)",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf.clone()),
                        path: vec![invalid_inner],
                    },
                    spec: api::iavl_spec(),
                    valid: false,
                },
            ),
            (
                "rejects invalid inner (hash)",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf.clone()),
                        path: vec![invalid_inner_hash],
                    },
                    spec: api::iavl_spec(),
                    valid: false,
                },
            ),
            (
                "accepts depth limited with proper number of inner nodes",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf.clone()),
                        path: vec![
                            valid_inner.clone(),
                            valid_inner.clone(),
                            valid_inner.clone(),
                        ],
                    },
                    spec: depth_limited_spec.clone(),
                    valid: true,
                },
            ),
            (
                "reject depth limited with too few inner nodes",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf.clone()),
                        path: vec![valid_inner.clone()],
                    },
                    spec: depth_limited_spec.clone(),
                    valid: false,
                },
            ),
            (
                "reject depth limited with too many inner nodes",
                ExistenceCase {
                    proof: ExistenceProof {
                        key: b"foo".to_vec(),
                        value: b"bar".to_vec(),
                        leaf: Some(leaf),
                        path: vec![
                            valid_inner.clone(),
                            valid_inner.clone(),
                            valid_inner.clone(),
                            valid_inner.clone(),
                            valid_inner,
                        ],
                    },
                    spec: depth_limited_spec,
                    valid: false,
                },
            ),
        ]
        .iter()
        .cloned()
        .collect();

        for (name, tc) in cases {
            let check = check_existence_spec(&tc.proof, &tc.spec);
            if tc.valid {
                assert!(
                    check.is_ok(),
                    "{} should be ok, got err {}",
                    name,
                    check.unwrap_err()
                );
            } else {
                assert!(check.is_err(), "{} should be an error", name);
            }
        }
    }
}
