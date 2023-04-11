# Changelog

# 0.10.0

## Go / Rust / JavaScript

This release introduces one single boolean new parameter to the top-level `ProofSpec`: `prehash_compare_key`.
When set to `true`, this flag causes keys to be consistently compared lexicographically according to their hashes
within nonexistence proof verification, using the same hash function as specified by the already-extant `prehash_key` field.

This is a backwards-compatible change, as it requires opt-in via setting the `prehash_compare_key` flag to `true` in the `ProofSpec`.
All existing ProofSpecs will continue to behave identically.

- feat: Add `prehash_compare_key` to allow proving nonexistence in sparse trees
  ([\#136](https://github.com/cosmos/ics23/pull/136))

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
