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

The minimum supported Rust version (MSRV) is 1.60.


## Fuzzing

Requires nightly and `cargo install cargo-fuzz`. Targets live in `fuzz/fuzz_targets`; `fuzz/seeds` holds the shared test vectors as protobuf (`verify_arbitrary_spec` pairs each with its built-in spec as a `Case` message).

```
cargo fuzz run verify_builtin_specs fuzz/corpus/verify_builtin_specs fuzz/seeds/verify_builtin_specs
cargo fuzz run verify_arbitrary_spec fuzz/corpus/verify_arbitrary_spec fuzz/seeds/verify_arbitrary_spec
```
