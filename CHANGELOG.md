# Changelog

# Unreleased

- fix(go): interpret max_depth in proof specs as 128 if left to 0 [#352](https://github.com/cosmos/ics23/pull/352).
- deps(go): bump golang to v1.22 ([#363](https://github.com/cosmos/ics23/pull/363).

# 0.12.0

## Rust

- chore(rust): Update `prost` to v0.13 ([#335](https://github.com/cosmos/ics23/pull/335), [#336](https://github.com/cosmos/ics23/pull/336))

# 0.11.3

## Rust

- fix(rust): Enforce that `spec.InnerSpec.ChildSize` is >= 1  ([#339](https://github.com/cosmos/ics23/pull/339))

# 0.11.2

## Rust

This release was yanked, please use v0.11.3 instead.

# 0.11.1

## Rust

- chore(rust): Update `informalsystems-pbjson` to v0.7.0 ([#274](https://github.com/cosmos/ics23/pull/274))

# 0.11.0

## Rust

- chore(rust): update `prost` to v0.12 ([#202](https://github.com/cosmos/ics23/pull/202))

# 0.10.2

## Rust

This release re-enables `no_std` support for ProtoJSON `Serialize` and `Deserialize` instances,
by swapping out `pbjson` for the `no_std`-compatible fork `informalsystems-pbjson`.

## Full changes

- feat(rust): enable no_std support for pbjson ([#158](https://github.com/cosmos/ics23/pull/146))

# 0.10.1

## Rust

The only change in this release of the `ics23` crate is the addition of a `serde` feature
which adds ProtoJSON-compatible `Serialize` and `Deserialize` instances on all Protobuf definitions via
the [`pbjson-build`](https://docs.rs/pbjson-build/latest/pbjson_build/) crate.

## Full changes

- feat(rust): Add ProtoJSON-compatible `Serialize` and `Deserialize` instances on all Protobuf definitions via `pbjson` ([#146](https://github.com/cosmos/ics23/pull/146))

# 0.10.0

## Go / Rust / TypeScript

This release introduces one single boolean new parameter to the top-level `ProofSpec`: `prehash_compare_key`.
When set to `true`, this flag causes keys to be consistently compared lexicographically according to their hashes
within nonexistence proof verification, using the same hash function as specified by the already-extant `prehash_key` field.

This is a backwards-compatible change, as it requires opt-in via setting the `prehash_compare_key` flag to `true` in the `ProofSpec`.
All existing `ProofSpec`s will continue to behave identically.

Please note that the version of the TypeScript library has been bump from 0.6.8 to 0.10.0 to align it with 
the Go and Rust implementations.

## Full changes

- feat: Add `prehash_compare_key` to allow proving nonexistence in sparse trees ([#136](https://github.com/cosmos/ics23/pull/136))
- chore: retract previous versions ([#131](https://github.com/cosmos/ics23/pull/131))
- fix: protobuf formatting using clang-format ([#129](https://github.com/cosmos/ics23/pull/129))
- fix: use /go suffix in go package option ([#127](https://github.com/cosmos/ics23/pull/127))
- add buf support to repo ([#126](https://github.com/cosmos/ics23/pull/126))
- chore: Add Cosmos, license and coverage badges to the README ([#122](https://github.com/cosmos/ics23/pull/122))
- ci: Enable code coverage for TypeScript version ([#123](https://github.com/cosmos/ics23/pull/123))
- ci: Add tags to codecov reports ([#121](https://github.com/cosmos/ics23/pull/121)) (4 months ago)
- ci: Refactor GitHub workflows and add code coverage job ([#120](https://github.com/cosmos/ics23/pull/120))

# 0.9.1

## Go

This release is a backport into the `release/v0.9.x` line of the feature that added the `prehash_compare_key` boolean parameter to the top-level `ProofSpec`.
When set to `true`, this flag causes keys to be consistently compared lexicographically according to their hashes
within nonexistence proof verification, using the same hash function as specified by the already-extant `prehash_key` field.

This is a backwards-compatible change, as it requires opt-in via setting the `prehash_compare_key` flag to `true` in the `ProofSpec`.
All existing `ProofSpec`s will continue to behave identically.

## Full changes

- feat(go): Add `prehash_compare_key` to allow proving nonexistence in sparse trees ([#136](https://github.com/cosmos/ics23/pull/136))

# 0.9.0

## Go

Release of ics23/go including changes made in the fork of ics23/go housed in the [Cosmos SDK](http://github.com/cosmos/cosmos-sdk).

## Rust

This release includes the same changes as its Go counterpart.

# 0.8.1

## Rust ([`ics23`](https://crates.io/crates/ics23))

- Fix no\_std compatibility and add check for this on CI ([#104](https://github.com/confio/ics23/pull/104))

# 0.8.0

## Rust ([`ics23`](https://crates.io/crates/ics23))

The following functions have been made generic over a new trait `HostFunctionProvider`:

- [x] `calculate_existence_root`
- [x] `verify_batch_membership`
- [x] `verify_batch_non_membership`
- [x] `verify_membership`
- [x] `verify_non_membership`

For `wasm32-unknown-unknown` environments this trait allows you to delegate hashing functions to a native implementation  through host functions.

With the `host-functions` feature (enabled by default), you can simply use `ics23::HostFunctionManager` as this provides a default implementation of this trait.

# v0.7.0

This handles non-existence tests for empty branches properly. This
is needed for properly handling proofs on Tries, like the SMT being
integrated with the Cosmos SDK.

This is used in ibc-go v3

# 0.6.x

This handles proofs for normal merkle trees, where every branch is full.
This works for tendermint merkle hashes and iavl hashes, and should work
for merk (nomic's db) proofs.

This was used in the original ibc release (cosmos sdk v0.40) and up until
ibc-go v2.
