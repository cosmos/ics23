#![no_main]
// Input: protobuf-encoded CommitmentProof, verified against each shipped spec.
use libfuzzer_sys::fuzz_target;
use prost::Message;

fuzz_target!(|data: &[u8]| {
    let Ok(proof) = ics23::CommitmentProof::decode(data) else {
        return;
    };
    for spec in [
        ics23::iavl_spec(),
        ics23::tendermint_spec(),
        ics23::smt_spec(),
    ] {
        ics23_fuzz::verify_all(&proof, &spec);
    }
});
