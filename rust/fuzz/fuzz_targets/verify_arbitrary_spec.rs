#![no_main]
// Input: protobuf-encoded `ics23_fuzz::Case`, i.e. ProofSpec and CommitmentProof both attacker-chosen.
use libfuzzer_sys::fuzz_target;
use prost::Message;

fuzz_target!(|data: &[u8]| {
    let Ok(case) = ics23_fuzz::Case::decode(data) else {
        return;
    };
    let (Some(spec), Some(proof)) = (case.spec, case.proof) else {
        return;
    };
    ics23_fuzz::verify_all(&proof, &spec);
});
