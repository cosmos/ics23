use ics23::{
    batch_entry, commitment_proof, CommitmentProof, ExistenceProof, HostFunctionsManager as H,
    ProofSpec,
};

/// Every existence proof inside `proof`, plus the key of the first non-existence proof.
fn parts(proof: &CommitmentProof) -> (Vec<&ExistenceProof>, Option<&[u8]>) {
    use commitment_proof::Proof::*;
    match &proof.proof {
        Some(Exist(ex)) => (vec![ex], None),
        Some(Nonexist(non)) => (non.left.iter().chain(&non.right).collect(), Some(&non.key)),
        Some(Batch(batch)) => {
            let mut exist = Vec::new();
            let mut key = None;
            for entry in &batch.entries {
                match &entry.proof {
                    Some(batch_entry::Proof::Exist(ex)) => exist.push(ex),
                    Some(batch_entry::Proof::Nonexist(non)) => {
                        exist.extend(non.left.iter().chain(&non.right));
                        key.get_or_insert(non.key.as_slice());
                    }
                    None => {}
                }
            }
            (exist, key)
        }
        _ => (Vec::new(), None),
    }
}

/// Runs the whole public API over `proof`. The root is derived from the proof itself: a malicious
/// counterparty controls its own state tree, so any proof that hashes consistently is a realistic input.
pub fn verify_all(proof: &CommitmentProof, spec: &ProofSpec) {
    let _ = ics23::compress(proof);
    let decompressed;
    let proof = if ics23::is_compressed(proof) {
        match ics23::decompress(proof) {
            Ok(p) => {
                decompressed = p;
                &decompressed
            }
            Err(_) => return,
        }
    } else {
        proof
    };
    let (exist, nonexist_key) = parts(proof);
    let mut roots = vec![Vec::new()];
    roots.extend(
        exist
            .iter()
            .filter_map(|ex| ics23::calculate_existence_root::<H>(ex).ok()),
    );
    for root in &roots {
        for ex in &exist {
            ics23::verify_membership::<H>(proof, spec, root, &ex.key, &ex.value);
        }
        if let Some(key) = nonexist_key {
            ics23::verify_non_membership::<H>(proof, spec, root, key);
        }
    }
}

/// Input of the `verify_arbitrary_spec` target: plain protobuf so seeds can be built from the vectors.
#[derive(Clone, PartialEq, prost::Message)]
pub struct Case {
    #[prost(message, optional, tag = "1")]
    pub spec: Option<ProofSpec>,
    #[prost(message, optional, tag = "2")]
    pub proof: Option<CommitmentProof>,
}
