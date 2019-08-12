fn main() {
    prost_build::compile_protos(&["../proofs.proto"],
                                &["../"]).unwrap();
}