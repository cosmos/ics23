use alloc::borrow::ToOwned;
use alloc::collections::btree_map::BTreeMap as HashMap;
use alloc::vec::Vec;
use prost::Message;

use crate::helpers::Result;
use crate::ics23;

pub fn is_compressed(proof: &ics23::CommitmentProof) -> bool {
    matches!(
        &proof.proof,
        Some(ics23::commitment_proof::Proof::Compressed(_))
    )
}

pub fn compress(proof: &ics23::CommitmentProof) -> Result<ics23::CommitmentProof> {
    if let Some(ics23::commitment_proof::Proof::Batch(batch)) = &proof.proof {
        compress_batch(batch)
    } else {
        Ok(proof.to_owned())
    }
}

pub fn decompress(proof: &ics23::CommitmentProof) -> Result<ics23::CommitmentProof> {
    if let Some(ics23::commitment_proof::Proof::Compressed(compressed)) = &proof.proof {
        decompress_batch(compressed)
    } else {
        Ok(proof.to_owned())
    }
}

pub fn compress_batch(proof: &ics23::BatchProof) -> Result<ics23::CommitmentProof> {
    let mut lookup = Vec::new();
    let mut registry = HashMap::new();

    let centries = proof
        .entries
        .iter()
        .map(|entry| match &entry.proof {
            Some(ics23::batch_entry::Proof::Exist(ex)) => {
                let exist = compress_exist(ex, &mut lookup, &mut registry)?;
                Ok(ics23::CompressedBatchEntry {
                    proof: Some(ics23::compressed_batch_entry::Proof::Exist(exist)),
                })
            }
            Some(ics23::batch_entry::Proof::Nonexist(non)) => {
                let left = non
                    .left
                    .as_ref()
                    .map(|l| compress_exist(l, &mut lookup, &mut registry))
                    .transpose()?;
                let right = non
                    .right
                    .as_ref()
                    .map(|r| compress_exist(r, &mut lookup, &mut registry))
                    .transpose()?;
                let nonexist = ics23::CompressedNonExistenceProof {
                    key: non.key.clone(),
                    left,
                    right,
                };
                Ok(ics23::CompressedBatchEntry {
                    proof: Some(ics23::compressed_batch_entry::Proof::Nonexist(nonexist)),
                })
            }
            None => Ok(ics23::CompressedBatchEntry { proof: None }),
        })
        .collect::<Result<_>>()?;

    Ok(ics23::CommitmentProof {
        proof: Some(ics23::commitment_proof::Proof::Compressed(
            ics23::CompressedBatchProof {
                entries: centries,
                lookup_inners: lookup,
            },
        )),
    })
}

pub fn compress_exist(
    exist: &ics23::ExistenceProof,
    lookup: &mut Vec<ics23::InnerOp>,
    registry: &mut HashMap<Vec<u8>, i32>,
) -> Result<ics23::CompressedExistenceProof> {
    let path = exist
        .path
        .iter()
        .map(|x| {
            let buf = x.encode_to_vec();
            let idx = *registry.entry(buf).or_insert_with(|| {
                let idx = lookup.len() as i32;
                lookup.push(x.to_owned());
                idx
            });
            Ok(idx)
        })
        .collect::<Result<_>>()?;

    Ok(ics23::CompressedExistenceProof {
        key: exist.key.clone(),
        value: exist.value.clone(),
        leaf: exist.leaf.clone(),
        path,
    })
}

pub fn decompress_batch(proof: &ics23::CompressedBatchProof) -> Result<ics23::CommitmentProof> {
    let lookup = &proof.lookup_inners;
    let entries = proof
        .entries
        .iter()
        .map(|cent| match &cent.proof {
            Some(ics23::compressed_batch_entry::Proof::Exist(ex)) => {
                let exist = decompress_exist(ex, lookup);
                Ok(ics23::BatchEntry {
                    proof: Some(ics23::batch_entry::Proof::Exist(exist)),
                })
            }
            Some(ics23::compressed_batch_entry::Proof::Nonexist(non)) => {
                let left = non.left.as_ref().map(|l| decompress_exist(l, lookup));
                let right = non.right.as_ref().map(|r| decompress_exist(r, lookup));
                let nonexist = ics23::NonExistenceProof {
                    key: non.key.clone(),
                    left,
                    right,
                };
                Ok(ics23::BatchEntry {
                    proof: Some(ics23::batch_entry::Proof::Nonexist(nonexist)),
                })
            }
            None => Ok(ics23::BatchEntry { proof: None }),
        })
        .collect::<Result<_>>()?;

    Ok(ics23::CommitmentProof {
        proof: Some(ics23::commitment_proof::Proof::Batch(ics23::BatchProof {
            entries,
        })),
    })
}

fn decompress_exist(
    exist: &ics23::CompressedExistenceProof,
    lookup: &[ics23::InnerOp],
) -> ics23::ExistenceProof {
    let path = exist
        .path
        .iter()
        .map(|&x| lookup.get(x as usize).cloned())
        .collect::<Option<Vec<_>>>()
        .unwrap_or_default();
    ics23::ExistenceProof {
        key: exist.key.clone(),
        value: exist.value.clone(),
        leaf: exist.leaf.clone(),
        path,
    }
}
