# Rust Proof Validation

## Codegen

To avoid direct dependencies on `protoc` in the build system, I have separated
`codegen` into a sub-crate. This will generate the rust `proofs.rs` file from
the `proofs.proto` file. The rest of the main build/test cycle is now independent
of the `protoc` binary.

To rebuild protobuf, simply: `cargo protoc` (on a dev machine with `protoc` in path).
Unless you modify the protobuf file, you can ignore this step.

## Formatting

`cargo fmt`

## Testing

`cargo test`

## Linting

`cargo clippy -- --test -W clippy::pedantic`

## Code Coverage

`cargo llvm-cov`

## MSRV

The minimum supported Rust version (MSRV) is 1.56.1.

