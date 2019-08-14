extern crate protobuf;

use crate::proofs;
use crate::verify::{CommitmentRoot, verify_existence};

pub fn verify_membership(proof: &proofs::CommitmentProof, _spec: &proofs::ProofSpec, root: CommitmentRoot, key: &[u8], value: &[u8]) -> bool {
    match &proof.proof {
        Some(proofs::CommitmentProof_oneof_proof::exist(ex)) => verify_existence(&ex, _spec, root, key, value).unwrap(),
        _ => false,
    }
}

// pub static IavlSpec: proofs::ProofSpec = make_iavl_spec();

pub fn make_iavl_spec() -> proofs::ProofSpec {
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
    use protobuf;
    use protobuf::Message;

    use std::fs::File;
    use std::io::prelude::*;

    struct TestVector {
        pub root: String,
        pub existence: String,
    }

    fn verify_test_vector(filename: &str, spec: &proofs::ProofSpec) -> Result<bool, &'static str> {
        let file = File::open(filename)?;
        let mut contents = String::new();
        file.read_to_string(&mut contents)?;

        let data: TestVector;

    // const { root, existence } = JSON.parse(content);
    // expect(existence).toBeDefined();
    // expect(root).toBeDefined();

        let mut parsed = proofs::ExistenceProof::new();
        parsed.merge_from_bytes(data.root.as_bytes())?;
        let root = hex::decode(data.root)?;

        verify_existence(&parsed, spec, root, &parsed.key, &parsed.value)
    }

    #[test]
    fn test_vector() {
        let iavl_spec = make_iavl_spec();
        let valid = verify_test_vector("../testdata/iavl/existence1.json", &iavl_spec);
        assert_eq!(true, valid.is_ok());
        assert_eq!(true, valid.unwrap());
    }
}