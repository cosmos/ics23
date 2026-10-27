use ics23::{
    batch_entry, commitment_proof, compress, BatchEntry, BatchProof, CommitmentProof,
    ExistenceProof, HashOp, InnerOp, InnerSpec, LeafOp, LengthOp, NonExistenceProof, ProofSpec,
};
use proptest::prelude::*;
use prost::Message;

const HASH_OPS: &[(HashOp, i32, &str)] = &[
    (HashOp::NoHash, 0, "NO_HASH"),
    (HashOp::Sha256, 1, "SHA256"),
    (HashOp::Sha512, 2, "SHA512"),
    (HashOp::Keccak256, 3, "KECCAK256"),
    (HashOp::Ripemd160, 4, "RIPEMD160"),
    (HashOp::Bitcoin, 5, "BITCOIN"),
    (HashOp::Sha512256, 6, "SHA512_256"),
    (HashOp::Blake2b512, 7, "BLAKE2B_512"),
    (HashOp::Blake2s256, 8, "BLAKE2S_256"),
    (HashOp::Blake3, 9, "BLAKE3"),
];

const LENGTH_OPS: &[(LengthOp, i32, &str)] = &[
    (LengthOp::NoPrefix, 0, "NO_PREFIX"),
    (LengthOp::VarProto, 1, "VAR_PROTO"),
    (LengthOp::VarRlp, 2, "VAR_RLP"),
    (LengthOp::Fixed32Big, 3, "FIXED32_BIG"),
    (LengthOp::Fixed32Little, 4, "FIXED32_LITTLE"),
    (LengthOp::Fixed64Big, 5, "FIXED64_BIG"),
    (LengthOp::Fixed64Little, 6, "FIXED64_LITTLE"),
    (LengthOp::Require32Bytes, 7, "REQUIRE_32_BYTES"),
    (LengthOp::Require64Bytes, 8, "REQUIRE_64_BYTES"),
];

#[test]
fn enum_names_and_numbers() {
    for &(op, number, name) in HASH_OPS {
        assert_eq!(op as i32, number);
        assert_eq!(HashOp::try_from(number), Ok(op));
        assert_eq!(op.as_str_name(), name);
        assert_eq!(HashOp::from_str_name(name), Some(op));
    }
    for &(op, number, name) in LENGTH_OPS {
        assert_eq!(op as i32, number);
        assert_eq!(LengthOp::try_from(number), Ok(op));
        assert_eq!(op.as_str_name(), name);
        assert_eq!(LengthOp::from_str_name(name), Some(op));
    }
    for name in ["", "sha256", "var_proto", "UNKNOWN"] {
        assert_eq!(HashOp::from_str_name(name), None);
        assert_eq!(LengthOp::from_str_name(name), None);
    }
    for number in [-1, 10, i32::MAX] {
        assert!(HashOp::try_from(number).is_err());
        assert!(LengthOp::try_from(number).is_err());
    }
    assert!(LengthOp::try_from(9).is_err());
}

fn bytes() -> impl Strategy<Value = Vec<u8>> {
    prop::collection::vec(any::<u8>(), 0..=64)
}

fn hash_op() -> impl Strategy<Value = i32> {
    prop::sample::select(HASH_OPS).prop_map(|(op, _, _)| op as i32)
}

fn leaf_op() -> impl Strategy<Value = LeafOp> {
    (
        hash_op(),
        hash_op(),
        hash_op(),
        prop::sample::select(LENGTH_OPS),
        bytes(),
    )
        .prop_map(
            |(hash, prehash_key, prehash_value, (length, _, _), prefix)| LeafOp {
                hash,
                prehash_key,
                prehash_value,
                length: length as i32,
                prefix,
            },
        )
}

fn inner_op() -> impl Strategy<Value = InnerOp> {
    (hash_op(), bytes(), bytes()).prop_map(|(hash, prefix, suffix)| InnerOp {
        hash,
        prefix,
        suffix,
    })
}

fn existence_proof() -> impl Strategy<Value = ExistenceProof> {
    (
        bytes(),
        bytes(),
        prop::option::of(leaf_op()),
        prop::collection::vec(inner_op(), 0..=8),
    )
        .prop_map(|(key, value, leaf, path)| ExistenceProof {
            key,
            value,
            leaf,
            path,
        })
}

fn nonexistence_proof() -> impl Strategy<Value = NonExistenceProof> {
    (
        bytes(),
        prop::option::of(existence_proof()),
        prop::option::of(existence_proof()),
    )
        .prop_map(|(key, left, right)| NonExistenceProof { key, left, right })
}

fn batch_proof() -> impl Strategy<Value = BatchProof> {
    let entry = prop::option::of(prop_oneof![
        existence_proof().prop_map(batch_entry::Proof::Exist),
        nonexistence_proof().prop_map(batch_entry::Proof::Nonexist),
    ])
    .prop_map(|proof| BatchEntry { proof });
    prop::collection::vec(entry, 0..=4).prop_map(|entries| BatchProof { entries })
}

fn proof_spec() -> impl Strategy<Value = ProofSpec> {
    let inner_spec = (
        prop::collection::vec(any::<i32>(), 0..=4),
        any::<i32>(),
        any::<i32>(),
        any::<i32>(),
        bytes(),
        hash_op(),
    )
        .prop_map(
            |(child_order, child_size, min_prefix_length, max_prefix_length, empty_child, hash)| {
                InnerSpec {
                    child_order,
                    child_size,
                    min_prefix_length,
                    max_prefix_length,
                    empty_child,
                    hash,
                }
            },
        );
    (
        prop::option::of(leaf_op()),
        prop::option::of(inner_spec),
        any::<i32>(),
        any::<i32>(),
        any::<bool>(),
    )
        .prop_map(
            |(leaf_spec, inner_spec, max_depth, min_depth, prehash_key_before_comparison)| {
                ProofSpec {
                    leaf_spec,
                    inner_spec,
                    max_depth,
                    min_depth,
                    prehash_key_before_comparison,
                }
            },
        )
}

fn assert_protobuf_roundtrip<T: Message + Default + PartialEq>(value: &T) {
    assert_eq!(&T::decode(value.encode_to_vec().as_slice()).unwrap(), value);
}

#[cfg(feature = "serde")]
fn assert_json_roundtrip<T>(value: &T)
where
    T: serde::Serialize + serde::de::DeserializeOwned + PartialEq + std::fmt::Debug,
{
    let json = serde_json::to_string(value).unwrap();
    assert_eq!(&serde_json::from_str::<T>(&json).unwrap(), value);
}

proptest! {
    #![proptest_config(ProptestConfig::with_cases(128))]

    #[test]
    fn commitment_proof_roundtrip(
        exist in existence_proof(),
        nonexist in nonexistence_proof(),
        batch in batch_proof(),
    ) {
        let batch = CommitmentProof { proof: Some(commitment_proof::Proof::Batch(batch)) };
        let compressed = compress(&batch).unwrap();
        for proof in [
            CommitmentProof::default(),
            CommitmentProof { proof: Some(commitment_proof::Proof::Exist(exist)) },
            CommitmentProof { proof: Some(commitment_proof::Proof::Nonexist(nonexist)) },
            batch,
            compressed,
        ] {
            assert_protobuf_roundtrip(&proof);
            #[cfg(feature = "serde")]
            assert_json_roundtrip(&proof);
        }
    }

    #[test]
    fn proof_spec_roundtrip(spec in proof_spec()) {
        assert_protobuf_roundtrip(&spec);
        #[cfg(feature = "serde")]
        assert_json_roundtrip(&spec);
    }
}

#[cfg(feature = "serde")]
mod json {
    use super::*;
    use serde_json::{from_str, from_value, json, to_value};

    #[test]
    fn enum_names_and_numbers() {
        for &(op, number, name) in HASH_OPS {
            assert_eq!(to_value(op).unwrap(), json!(name));
            assert_eq!(from_value::<HashOp>(json!(name)).unwrap(), op);
            assert_eq!(from_value::<HashOp>(json!(number)).unwrap(), op);
        }
        for &(op, number, name) in LENGTH_OPS {
            assert_eq!(to_value(op).unwrap(), json!(name));
            assert_eq!(from_value::<LengthOp>(json!(name)).unwrap(), op);
            assert_eq!(from_value::<LengthOp>(json!(number)).unwrap(), op);
        }
    }

    #[test]
    fn field_names_bytes_and_defaults() {
        let leaf = LeafOp {
            hash: HashOp::Sha256 as i32,
            prehash_key: HashOp::Sha512 as i32,
            prehash_value: HashOp::Ripemd160 as i32,
            length: LengthOp::VarProto as i32,
            prefix: vec![0, 255, 128],
        };
        assert_eq!(
            to_value(&leaf).unwrap(),
            json!({
                "hash": "SHA256", "prehashKey": "SHA512", "prehashValue": "RIPEMD160",
                "length": "VAR_PROTO", "prefix": "AP+A"
            })
        );
        assert_eq!(
            from_value::<LeafOp>(json!({
                "hash": 1, "prehash_key": 2, "prehash_value": 4, "length": 1, "prefix": "AP+A"
            }))
            .unwrap(),
            leaf
        );

        let spec = ProofSpec {
            leaf_spec: Some(LeafOp::default()),
            max_depth: 8,
            prehash_key_before_comparison: true,
            ..Default::default()
        };
        assert_eq!(
            to_value(&spec).unwrap(),
            json!({
                "leafSpec": {}, "maxDepth": 8, "prehashKeyBeforeComparison": true
            })
        );
        assert_eq!(
            from_value::<ProofSpec>(json!({
                "leaf_spec": {}, "max_depth": 8, "prehash_key_before_comparison": true
            }))
            .unwrap(),
            spec
        );

        assert_eq!(to_value(LeafOp::default()).unwrap(), json!({}));
        assert_eq!(from_str::<LeafOp>("{}").unwrap(), LeafOp::default());
        assert_eq!(to_value(ProofSpec::default()).unwrap(), json!({}));
        assert_eq!(from_str::<ProofSpec>("{}").unwrap(), ProofSpec::default());
        assert_eq!(to_value(CommitmentProof::default()).unwrap(), json!({}));
        assert_eq!(
            from_str::<CommitmentProof>("{}").unwrap(),
            CommitmentProof::default()
        );
    }

    #[test]
    fn rejects_invalid_fields_and_enums() {
        for json in [
            r#"{"prefix":"!"}"#,
            r#"{"unknown":0}"#,
            r#"{"hash":"UNKNOWN"}"#,
            r#"{"hash":-1}"#,
            r#"{"hash":10}"#,
            r#"{"hash":2147483648}"#,
            r#"{"length":"UNKNOWN"}"#,
            r#"{"length":9}"#,
            r#"{"hash":1,"hash":2}"#,
            r#"{"prehashKey":1,"prehash_key":2}"#,
        ] {
            assert!(from_str::<LeafOp>(json).is_err(), "{json}");
        }
        for leaf in [
            LeafOp {
                hash: -1,
                ..Default::default()
            },
            LeafOp {
                prehash_key: 10,
                ..Default::default()
            },
            LeafOp {
                prehash_value: 10,
                ..Default::default()
            },
            LeafOp {
                length: 9,
                ..Default::default()
            },
        ] {
            assert!(to_value(leaf).is_err());
        }
    }

    #[test]
    fn rejects_conflicting_proof_variants() {
        for json in [
            r#"{"exist":{},"exist":{}}"#,
            r#"{"exist":{},"nonexist":{}}"#,
            r#"{"nonexist":{},"exist":{}}"#,
            r#"{"batch":{},"compressed":{}}"#,
            r#"{"compressed":{},"batch":{}}"#,
        ] {
            assert!(from_str::<CommitmentProof>(json).is_err(), "{json}");
        }
    }
}
