# Changelog

# Unreleased

- fix: guarantee that `spec.InnerSpec.MaxPrefixLength` < `spec.InnerSpec.MinPrefixLength` + `spec.InnerSpec.ChildSize` ([#369](https://github.com/cosmos/ics23/pull/369))

# v0.12.0

- chore(rust): Update `prost` to v0.13 ([#335](https://github.com/cosmos/ics23/pull/335), [#336](https://github.com/cosmos/ics23/pull/336))

# v0.11.3

- fix(rust): Enforce that `spec.InnerSpec.ChildSize` is >= 1  ([#339](https://github.com/cosmos/ics23/pull/339))

# v0.11.2

This release was yanked, please use v0.11.3 instead.

# v0.11.1

- chore(rust): Update `informalsystems-pbjson` to v0.7.0 ([#274](https://github.com/cosmos/ics23/pull/274))

# v0.11.0

- chore(rust): update `prost` to v0.12 ([#202](https://github.com/cosmos/ics23/pull/202))

# v0.10.2

This release re-enables `no_std` support for ProtoJSON `Serialize` and `Deserialize` instances,
by swapping out `pbjson` for the `no_std`-compatible fork `informalsystems-pbjson`.

## Full changes

- feat(rust): enable no_std support for pbjson ([#158](https://github.com/cosmos/ics23/pull/146))

# v0.10.1

The only change in this release of the `ics23` crate is the addition of a `serde` feature
which adds ProtoJSON-compatible `Serialize` and `Deserialize` instances on all Protobuf definitions via
the [`pbjson-build`](https://docs.rs/pbjson-build/latest/pbjson_build/) crate.

## Full changes

- feat(rust): Add ProtoJSON-compatible `Serialize` and `Deserialize` instances on all Protobuf definitions via `pbjson` ([#146](https://github.com/cosmos/ics23/pull/146))

# v0.10.0

This release introduces one single boolean new parameter to the top-level `ProofSpec`: `prehash_compare_key`.
When set to `true`, this flag causes keys to be consistently compared lexicographically according to their hashes
within nonexistence proof verification, using the same hash function as specified by the already-extant `prehash_key` field.

This is a backwards-compatible change, as it requires opt-in via setting the `prehash_compare_key` flag to `true` in the `ProofSpec`.
All existing `ProofSpec`s will continue to behave identically.

## Full changes

- feat: Add `prehash_compare_key` to allow proving nonexistence in sparse trees ([#136](https://github.com/cosmos/ics23/pull/136))
- fix: protobuf formatting using clang-format ([#129](https://github.com/cosmos/ics23/pull/129))
- add buf support to repo ([#126](https://github.com/cosmos/ics23/pull/126))
- chore: Add Cosmos, license and coverage badges to the README ([#122](https://github.com/cosmos/ics23/pull/122))
- ci: Add tags to codecov reports ([#121](https://github.com/cosmos/ics23/pull/121)) (4 months ago)
- ci: Refactor GitHub workflows and add code coverage job ([#120](https://github.com/cosmos/ics23/pull/120))

# v0.9.0

Release of the `ics23` create, including changes that reflect changes made in the Go implementation in the fork of `ics23` housed in the [Cosmos SDK](http://github.com/cosmos/cosmos-sdk).

# v0.8.1

- Fix no\_std compatibility and add check for this on CI ([#104](https://github.com/confio/ics23/pull/104))

# v0.8.0

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
