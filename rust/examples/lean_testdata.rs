//! Phase 2a differential oracle: generate `lean/Ics23/TestVectors.lean` from
//! the shared `testdata/` vectors the Rust and Go suites verify.
//!
//! Each vector's hex protobuf `CommitmentProof` is decoded and re-emitted as a
//! literal of the Lean model's proof types, together with a `native_decide`
//! acceptance check mirroring the Rust call (`verify_membership` /
//! `verify_non_membership`). `lake build` then re-verifies every shared vector
//! against the Lean model with real SHA-256 — any accept/reject divergence
//! between the model and the implementations fails the proof build.
//!
//! Run from `rust/`: `cargo run --example lean_testdata`
//! CI regenerates the file and fails if it drifts from the committed copy.

use anyhow::{bail, Context, Result};
use prost::Message;
use serde::Deserialize;
use std::fmt::Write as _;
use std::fs;

#[derive(Deserialize)]
struct TestVector {
    root: String,
    proof: String,
    key: String,
    value: String,
}

fn hash_op(i: i32) -> Result<&'static str> {
    Ok(match i {
        0 => ".noHash",
        1 => ".sha256",
        2 => ".sha512",
        3 => ".keccak256",
        4 => ".ripemd160",
        5 => ".bitcoin",
        6 => ".sha512256",
        7 => ".blake2b512",
        8 => ".blake2s256",
        9 => ".blake3",
        _ => bail!("unknown HashOp {i}"),
    })
}

fn length_op(i: i32) -> Result<&'static str> {
    Ok(match i {
        0 => ".noPrefix",
        1 => ".varProto",
        3 => ".fixed32Big",
        4 => ".fixed32Little",
        5 => ".fixed64Big",
        6 => ".fixed64Little",
        7 => ".require32Bytes",
        8 => ".require64Bytes",
        _ => bail!("unsupported LengthOp {i}"),
    })
}

/// A Lean `Bytes` literal, wrapped at 16 bytes per line at the given indent.
fn bytes_lit(b: &[u8], indent: &str) -> String {
    if b.is_empty() {
        return "[]".to_string();
    }
    let mut out = String::from("[");
    for (i, byte) in b.iter().enumerate() {
        if i > 0 {
            if i % 16 == 0 {
                out.push_str(",\n");
                out.push_str(indent);
            } else {
                out.push_str(", ");
            }
        }
        write!(out, "0x{byte:02x}").unwrap();
    }
    out.push(']');
    out
}

/// Emit `def <name> : ExistenceProof where ...` (where-style keeps every
/// structure brace at a fixed shallow column, as Lean's indentation-sensitive
/// structure-instance syntax requires).
fn emit_exist_def(out: &mut String, name: &str, p: &ics23::ExistenceProof) -> Result<()> {
    let leaf = p.leaf.as_ref().context("existence proof without leaf")?;
    writeln!(out, "def {name} : ExistenceProof where").unwrap();
    writeln!(out, "  key := {}", bytes_lit(&p.key, "    ")).unwrap();
    writeln!(out, "  value := {}", bytes_lit(&p.value, "    ")).unwrap();
    writeln!(out, "  leaf :=").unwrap();
    writeln!(
        out,
        "    {{ hash := {}, prehashKey := {}, prehashValue := {},",
        hash_op(leaf.hash)?,
        hash_op(leaf.prehash_key)?,
        hash_op(leaf.prehash_value)?,
    )
    .unwrap();
    writeln!(
        out,
        "      length := {}, prefixBytes := {} }}",
        length_op(leaf.length)?,
        bytes_lit(&leaf.prefix, "        "),
    )
    .unwrap();
    if p.path.is_empty() {
        writeln!(out, "  path := []").unwrap();
    } else {
        writeln!(out, "  path := [").unwrap();
        for (i, op) in p.path.iter().enumerate() {
            writeln!(out, "    {{ hash := {},", hash_op(op.hash)?).unwrap();
            writeln!(
                out,
                "      prefixBytes := {},",
                bytes_lit(&op.prefix, "        ")
            )
            .unwrap();
            writeln!(
                out,
                "      suffix := {} }}{}",
                bytes_lit(&op.suffix, "        "),
                if i + 1 == p.path.len() { "" } else { "," }
            )
            .unwrap();
        }
        writeln!(out, "  ]").unwrap();
    }
    out.push('\n');
    Ok(())
}

fn emit_vector(out: &mut String, dir: &str, stem: &str, spec_name: &str) -> Result<()> {
    let path = format!("../testdata/{dir}/{stem}.json");
    let contents = fs::read_to_string(&path).with_context(|| path.clone())?;
    let data: TestVector = serde_json::from_str(&contents)?;
    let proto_bin = hex::decode(&data.proof)?;
    let parsed = ics23::CommitmentProof::decode(proto_bin.as_slice())?;
    let root = hex::decode(&data.root)?;
    let key = hex::decode(&data.key)?;

    // camelCase vector name: iavl/exist_left -> iavlExistLeft
    let mut name = String::from(dir);
    for part in stem.split('_') {
        let mut cs = part.chars();
        if let Some(c) = cs.next() {
            name.push(c.to_ascii_uppercase());
            name.push_str(cs.as_str());
        }
    }

    writeln!(out, "/-! ### `testdata/{dir}/{stem}.json` -/\n").unwrap();
    match parsed.proof {
        Some(ics23::commitment_proof::Proof::Exist(ep)) => {
            if data.value.is_empty() {
                bail!("{path}: existence vector without value");
            }
            let value = hex::decode(&data.value)?;
            emit_exist_def(out, &name, &ep)?;
            writeln!(out, "def {name}Root : Bytes :=").unwrap();
            writeln!(out, "  {}\n", bytes_lit(&root, "   ")).unwrap();
            writeln!(out, "def {name}Key : Bytes :=").unwrap();
            writeln!(out, "  {}\n", bytes_lit(&key, "   ")).unwrap();
            writeln!(out, "def {name}Value : Bytes :=").unwrap();
            writeln!(out, "  {}\n", bytes_lit(&value, "   ")).unwrap();
            writeln!(
                out,
                "example : verifyExistence concreteHash {name} {spec_name} {name}Root\n    \
                 {name}Key {name}Value = true := by native_decide\n"
            )
            .unwrap();
        }
        Some(ics23::commitment_proof::Proof::Nonexist(np)) => {
            if !data.value.is_empty() {
                bail!("{path}: non-existence vector with value");
            }
            let left = match &np.left {
                None => "none".to_string(),
                Some(ep) => {
                    emit_exist_def(out, &format!("{name}L"), ep)?;
                    format!("some {name}L")
                }
            };
            let right = match &np.right {
                None => "none".to_string(),
                Some(ep) => {
                    emit_exist_def(out, &format!("{name}R"), ep)?;
                    format!("some {name}R")
                }
            };
            writeln!(out, "def {name} : NonExistenceProof where").unwrap();
            writeln!(out, "  key := {}", bytes_lit(&np.key, "    ")).unwrap();
            writeln!(out, "  left := {left}").unwrap();
            writeln!(out, "  right := {right}\n").unwrap();
            writeln!(out, "def {name}Root : Bytes :=").unwrap();
            writeln!(out, "  {}\n", bytes_lit(&root, "   ")).unwrap();
            writeln!(out, "def {name}Key : Bytes :=").unwrap();
            writeln!(out, "  {}\n", bytes_lit(&key, "   ")).unwrap();
            writeln!(
                out,
                "example : verifyNonExistence concreteHash {name} {spec_name} {name}Root\n    \
                 {name}Key = true := by native_decide\n"
            )
            .unwrap();
        }
        _ => bail!("{path}: unsupported proof variant"),
    }
    Ok(())
}

fn main() -> Result<()> {
    let mut out = String::new();
    out.push_str(
        "/-\nShared test vectors (Phase 2a differential oracle). GENERATED FILE — do not\n\
         edit by hand; regenerate with `cargo run --example lean_testdata` in `rust/`.\n\n\
         Each vector under `testdata/{iavl,tendermint,smt}/` is the decoded protobuf\n\
         `CommitmentProof` the Rust and Go suites verify, re-checked here against the\n\
         Lean model with real SHA-256 (`native_decide`). An accept/reject divergence\n\
         between the model and the implementations fails this build.\n-/\n\
         import Ics23.Executable\n\nnamespace Ics23.TestVectors\n\nopen Ics23\n\n",
    );

    let stems = [
        "exist_left",
        "exist_right",
        "exist_middle",
        "nonexist_left",
        "nonexist_right",
        "nonexist_middle",
    ];
    for (dir, spec_name) in [
        ("iavl", "iavlSpec"),
        ("tendermint", "tendermintSpec"),
        ("smt", "smtSpec"),
    ] {
        for stem in stems {
            emit_vector(&mut out, dir, stem, spec_name)?;
        }
    }

    out.push_str("end Ics23.TestVectors\n");
    fs::write("../lean/Ics23/TestVectors.lean", &out)?;
    println!("wrote lean/Ics23/TestVectors.lean");
    Ok(())
}
