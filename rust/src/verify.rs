// we want to name functions verify_* to match ics23
#![allow(clippy::module_name_repetitions)]

use alloc::vec::Vec;

use anyhow::{bail, ensure};

use crate::api::{ensure_inner_prefix, ensure_leaf_prefix};
use crate::helpers::Result;
use crate::host_functions::HostFunctionsProvider;
use crate::ics23;
use crate::ops::{apply_inner, apply_leaf};

pub type CommitmentRoot = Vec<u8>;

pub fn verify_existence<H: HostFunctionsProvider>(
    proof: &ics23::ExistenceProof,
    spec: &ics23::ProofSpec,
    root: &[u8],
    key: &[u8],
    value: &[u8],
) -> Result<()> {
    check_existence_spec(proof, spec)?;
    ensure!(proof.key == key, "Provided key doesn't match proof");
    ensure!(proof.value == value, "Provided value doesn't match proof");

    let calc = calculate_existence_root_for_spec::<H>(proof, Some(spec))?;
    ensure!(calc == root, "Root hash doesn't match");
    Ok(())
}

pub fn verify_non_existence<H: HostFunctionsProvider>(
    proof: &ics23::NonExistenceProof,
    spec: &ics23::ProofSpec,
    root: &[u8],
    key: &[u8],
) -> Result<()> {
    if let Some(left) = &proof.left {
        verify_existence::<H>(left, spec, root, &left.key, &left.value)?;
        ensure!(key > left.key.as_slice(), "left key isn't before key");
    }
    if let Some(right) = &proof.right {
        verify_existence::<H>(right, spec, root, &right.key, &right.value)?;
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
pub fn calculate_existence_root<H: HostFunctionsProvider>(
    proof: &ics23::ExistenceProof,
) -> Result<CommitmentRoot> {
    calculate_existence_root_for_spec::<H>(proof, None)
}

fn calculate_existence_root_for_spec<H: HostFunctionsProvider>(
    proof: &ics23::ExistenceProof,
    spec: Option<&ics23::ProofSpec>,
) -> Result<CommitmentRoot> {
    ensure!(!proof.key.is_empty(), "Existence proof must have key set");
    ensure!(
        !proof.value.is_empty(),
        "Existence proof must have value set"
    );

    if let Some(leaf_node) = &proof.leaf {
        let mut hash = apply_leaf::<H>(leaf_node, &proof.key, &proof.value)?;
        for step in &proof.path {
            hash = apply_inner::<H>(step, &hash)?;

            if let Some(inner_spec) = spec.and_then(|spec| spec.inner_spec.as_ref()) {
                if hash.len() > inner_spec.child_size as usize && inner_spec.child_size >= 32 {
                    bail!("Invalid inner operation (child_size)")
                }
            }
        }
        Ok(hash)
    } else {
        bail!("No leaf operation set")
    }
}

fn check_existence_spec(proof: &ics23::ExistenceProof, spec: &ics23::ProofSpec) -> Result<()> {
    if let (Some(leaf), Some(leaf_spec)) = (&proof.leaf, &spec.leaf_spec) {
        ensure_leaf_prefix(&leaf.prefix, spec)?;
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
        for (idx, step) in proof.path.iter().enumerate() {
            ensure_inner_prefix(&step.prefix, spec, (idx as i64) + 1, step.hash)?;
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
                "Inner prefix too long: {}",
                inner.prefix.len(),
            );
            ensure!(
                inner.suffix.len() % (inner_spec.child_size as usize) == 0,
                "InnerOp suffix malformed"
            );
            Ok(())
        }
        (_, _) => bail!("Spec requires both leaf_spec and inner_spec"),
    }
}

// ensure_left_most fails unless this is the left-most path in the tree, excluding placeholder (empty child) nodes
fn ensure_left_most(spec: &ics23::InnerSpec, path: &[ics23::InnerOp]) -> Result<()> {
    let pad = get_padding(spec, 0)?;
    // ensure every step has a prefix and suffix defined to be leftmost, unless it is a placeholder node
    for step in path {
        if !has_padding(step, &pad) && !left_branches_are_empty(spec, step)? {
            bail!("step not leftmost")
        }
    }
    Ok(())
}

// ensure_right_most returns true if this is the right-most path in the tree, excluding placeholder (empty child) nodes
fn ensure_right_most(spec: &ics23::InnerSpec, path: &[ics23::InnerOp]) -> Result<()> {
    let idx = spec.child_order.len() - 1;
    let pad = get_padding(spec, idx as i32)?;
    // ensure every step has a prefix and suffix defined to be rightmost, unless it is a placeholder node
    for step in path {
        if !has_padding(step, &pad) && !right_branches_are_empty(spec, step)? {
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

// left_branches_are_empty returns true if the padding bytes correspond to all empty children
// on the left side of this branch, ie. it's a valid placeholder on a leftmost path
fn left_branches_are_empty(spec: &ics23::InnerSpec, op: &ics23::InnerOp) -> Result<bool> {
    let idx = order_from_padding(spec, op)?;
    // count branches to left of this
    let left_branches = idx as usize;
    if left_branches == 0 {
        return Ok(false);
    }
    let child_size = spec.child_size as usize;
    // compare prefix with the expected number of empty branches
    let actual_prefix = match op.prefix.len().checked_sub(left_branches * child_size) {
        Some(n) => n,
        _ => return Ok(false),
    };
    for i in 0..left_branches {
        let idx = spec.child_order.iter().find(|&&x| x == i as i32).unwrap();
        let idx = *idx as usize;
        let from = actual_prefix + idx * child_size;
        if spec.empty_child != op.prefix[from..from + child_size] {
            return Ok(false);
        }
    }
    Ok(true)
}

// right_branches_are_empty returns true if the padding bytes correspond to all empty children
// on the right side of this branch, ie. it's a valid placeholder on a rightmost path
fn right_branches_are_empty(spec: &ics23::InnerSpec, op: &ics23::InnerOp) -> Result<bool> {
    let idx = order_from_padding(spec, op)?;
    // count branches to right of this one
    let right_branches = spec.child_order.len() - 1 - idx as usize;
    // compare suffix with the expected number of empty branches
    if right_branches == 0 {
        return Ok(false);
    }
    if op.suffix.len() != spec.child_size as usize {
        return Ok(false);
    }
    for i in 0..right_branches {
        let idx = spec.child_order.iter().find(|&&x| x == i as i32).unwrap();
        let idx = *idx as usize;
        let from = idx * spec.child_size as usize;
        if spec.empty_child != op.suffix[from..from + spec.child_size as usize] {
            return Ok(false);
        }
    }
    Ok(true)
}

#[cfg(test)]
mod tests {
    use super::*;

    use crate::api;
    use crate::host_functions::host_functions_impl::HostFunctionsManager;
    use crate::ics23::{ExistenceProof, HashOp, InnerOp, InnerSpec, LeafOp, LengthOp, ProofSpec};

    use alloc::collections::btree_map::BTreeMap as HashMap;
    use alloc::vec;

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
            calculate_existence_root::<HostFunctionsManager>(&proof).unwrap(),
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
            calculate_existence_root::<HostFunctionsManager>(&proof).unwrap(),
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
        impl InnerOp {
            fn with_height(mut self, height: u8) -> InnerOp {
                self.prefix[0] = height;
                self
            }
        }

        let leaf = LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: HashOp::Sha256.into(),
            length: LengthOp::VarProto.into(),
            prefix: vec![0u8, 2, 2],
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
            prefix: vec![2u8, 2, 2, 0],
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
                            valid_inner.clone().with_height(4),
                            valid_inner.clone().with_height(6),
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
                check.expect(name);
            } else {
                assert!(check.is_err(), "{} should be an error", name);
            }
        }
    }

    fn spec_with_empty_child() -> ProofSpec {
        let leaf = LeafOp {
            hash: ics23::HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: ics23::HashOp::Sha256.into(),
            length: 0,
            prefix: vec![0_u8],
        };
        let inner = InnerSpec {
            child_order: vec![0, 1],
            child_size: 32,
            min_prefix_length: 1,
            max_prefix_length: 1,
            empty_child: b"32_empty_child_placeholder_bytes".to_vec(),
            hash: ics23::HashOp::Sha256.into(),
        };
        ProofSpec {
            leaf_spec: Some(leaf),
            inner_spec: Some(inner),
            min_depth: 0,
            max_depth: 0,
        }
    }

    struct EmptyBranchCase<'a> {
        op: InnerOp,
        spec: &'a ProofSpec,
        is_left: bool,
        is_right: bool,
    }

    #[test]
    fn check_empty_branch() -> Result<()> {
        let spec = &spec_with_empty_child();
        let inner_spec = spec.inner_spec.as_ref().unwrap();
        let empty_child = inner_spec.empty_child.clone();

        let non_empty_spec = &api::tendermint_spec();
        let non_empty_inner = non_empty_spec.inner_spec.as_ref().unwrap();

        let cases = vec![
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: [&[1u8], &empty_child[..]].concat().to_vec(),
                    suffix: vec![],
                    hash: inner_spec.hash,
                },
                spec,
                is_left: true,
                is_right: false,
            },
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: vec![1u8],
                    suffix: empty_child.clone(),
                    hash: inner_spec.hash,
                },
                spec,
                is_left: false,
                is_right: true,
            },
            // non-empty cases
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: [&[1u8], &[0u8; 32] as &[u8]].concat().to_vec(),
                    suffix: vec![],
                    hash: inner_spec.hash,
                },
                spec,
                is_left: false,
                is_right: false,
            },
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: vec![1u8],
                    suffix: vec![0u8; 32],
                    hash: inner_spec.hash,
                },
                spec,
                is_left: false,
                is_right: false,
            },
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: [&[1u8], &empty_child[..28], b"xxxx"].concat().to_vec(),
                    suffix: vec![],
                    hash: inner_spec.hash,
                },
                spec,
                is_left: false,
                is_right: false,
            },
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: vec![1u8],
                    suffix: [&empty_child[..28], b"xxxx"].concat().to_vec(),
                    hash: inner_spec.hash,
                },
                spec,
                is_left: false,
                is_right: false,
            },
            // some cases using a spec with no empty child
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: [&[1u8], &[0u8; 32] as &[u8]].concat().to_vec(),
                    suffix: vec![],
                    hash: non_empty_inner.hash,
                },
                spec: non_empty_spec,
                is_left: false,
                is_right: false,
            },
            EmptyBranchCase {
                op: ics23::InnerOp {
                    prefix: vec![1u8],
                    suffix: vec![0u8; 32],
                    hash: non_empty_inner.hash,
                },
                spec: non_empty_spec,
                is_left: false,
                is_right: false,
            },
        ];

        for (i, case) in cases.iter().enumerate() {
            ensure_inner(&case.op, case.spec)?;
            let inner = &case.spec.inner_spec.as_ref().unwrap();
            assert_eq!(
                case.is_left,
                left_branches_are_empty(inner, &case.op)?,
                "case {}",
                i
            );
            assert_eq!(
                case.is_right,
                right_branches_are_empty(inner, &case.op)?,
                "case {}",
                i
            );
        }
        Ok(())
    }
}
