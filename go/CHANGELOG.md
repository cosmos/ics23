# Changelog

# Unreleased

- deps: bump golang to v1.22 ([#363](https://github.com/cosmos/ics23/pull/363)).
- fix: guarantee that `spec.InnerSpec.MaxPrefixLength` < `spec.InnerSpec.MinPrefixLength` + `spec.InnerSpec.ChildSize` ([#369](https://github.com/cosmos/ics23/pull/369))
- refactor: support for `BatchProof` and `CompressedBatchProof` is being dropped. 
    * The API's `BatchVerifyMembership`, `BatchVerifyNonMembership`, and `CombineProofs` have been removed. ([#390](https://github.com/cosmos/ics23/pull/390))
    * The API's `IsCompressed`, `Compress`, and `Decompress` have been removed. ([#392](https://github.com/cosmos/ics23/pull/392))

# v0.11.0

- fix(go): interpret max_depth in proof specs as 128 if left to 0 ([#352](https://github.com/cosmos/ics23/pull/352)).
- deps: bump cosmos/gogoproto to v1.7.0.
- deps: bump x/crypto v0.26.0.
- feat: add support for Blake2b/2s/3 hash functions ([#212](https://github.com/cosmos/ics23/pull/212)).
- fix(go): bullet-proof against nil dereferences, add more fuzzers ([\#244](https://github.com/cosmos/ics23/pull/244)).

# v0.10.0

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

# v0.9.1

This release is a backport into the `release/v0.9.x` line of the feature that added the `prehash_compare_key` boolean parameter to the top-level `ProofSpec`.
When set to `true`, this flag causes keys to be consistently compared lexicographically according to their hashes
within nonexistence proof verification, using the same hash function as specified by the already-extant `prehash_key` field.

This is a backwards-compatible change, as it requires opt-in via setting the `prehash_compare_key` flag to `true` in the `ProofSpec`.
All existing `ProofSpec`s will continue to behave identically.

## Full changes

- feat(go): Add `prehash_compare_key` to allow proving nonexistence in sparse trees ([#136](https://github.com/cosmos/ics23/pull/136))

# v0.9.0

Release of ics23/go including changes made in the fork of ics23/go housed in the [Cosmos SDK](http://github.com/cosmos/cosmos-sdk).

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
