[![Cosmos ecosystem][cosmos-shield]][cosmos-link]

# ICS 23 [![Apache 2.0 Licensed][license-badge]][license-link]

| Language           | Test Suite                                        | Code Coverage                                  |
| ------------------ | ------------------------------------------------- | ---------------------------------------------- |
| [Go](./go)         | [![Go Test][go-test-badge]][go-test-link]         | [![Go Cov][go-cov-badge]][go-cov-link]         |
| [Rust](./rust)     | [![Rust Test][rust-test-badge]][rust-test-link]   | [![Rust Cov][rust-cov-badge]][rust-cov-link]   |

[cosmos-link]: https://cosmos.network
[go-test-link]: https://github.com/cosmos/ics23/actions/workflows/go.yml
[go-test-badge]: https://github.com/cosmos/ics23/actions/workflows/go.yml/badge.svg?branch=master
[go-cov-link]: https://sonarcloud.io/project/configuration?id=ics23-go
[go-cov-badge]: https://sonarcloud.io/api/project_badges/measure?project=ics23-go&metric=coverage
[rust-test-link]: https://github.com/cosmos/ics23/actions/workflows/rust.yml
[rust-test-badge]: https://github.com/cosmos/ics23/actions/workflows/rust.yml/badge.svg?branch=master
[rust-cov-link]: https://codecov.io/gh/cosmos/ics23/tree/master/rust
[rust-cov-badge]: https://codecov.io/github/cosmos/ics23/branch/master/graph/badge.svg?token=xlGriS907o&flag=rust
[license-link]: https://github.com/cosmos/ics23/blob/master/LICENSE
[license-badge]: https://img.shields.io/badge/license-Apache2.0-blue.svg

## Proofs

This is an attempt to define a generic, cross-language, binary representation
of merkle proofs, which can be generated by many underlying merkle tree storage
implementations and validated by many client libraries over various languages.

The end goal is to provide a standard representation not only for light-client
proofs of blockchains, but more specifically for proofs that accompany IBC
(inter-blockchain communication) packets as defined in the cosmos specification.

## Feature set

The features and naming follow the [ICS23: Vector Commitments](https://github.com/cosmos/ibc/tree/master/spec/core/ics-023-vector-commitments) Specification

* Proof of Existence (key-value pair linked to root hash)
* Hold Existence Proof to db-specific spec (avoid reinterpretation of bytes to different key-value pair)
* Proof of Non-Existence (key proven not to be inside tree with given root hash)

### Future features

* Batch Proof or Range Proof (prove many keys at once, more efficiently than separate proofs)

## Organization

The top level package will provide language-agnostic documentation and protobuf specifications.
Every language should have a sub-folder, containing both protobuf generated code,
as well as client-side code for validating such proofs. The server-side code should
live alongside the various merkle tree representations (eg. not in this repository).

### Supported client languages

* [Go](./go)
* [Rust](./rust)

The repository used to have a TypeScript implementation, but due to lack of maintenance and usage, it was removed in [#353](https://github.com/cosmos/ics23/pull/353).

Anyone interested in adding support for Solidity could pick up where [#58](https://github.com/cosmos/ics23/pull/58) left off.

### Compatibility table

| ICS 023 Spec  | Go                | Rust                |
|---------------|-------------------|---------------------|
| 2019-08-25    | 0.9.x             | 0.9.x               |
| 2019-08-25    | 0.10.x            | 0.10.x, 0.11.x      |

## Supported merkle stores

* [cosmos/iavl](https://github.com/cosmos/iavl)
* [cometbft/crypto/merkle](https://github.com/cometbft/cometbft/tree/main/crypto/merkle)
* [penumbra-zone/jmt](https://github.com/penumbra-zone/jmt)

Supported merkle stores must be lexographically ordered to mainatin soundness. 

### Unsupported

* [turbofish-org/merk](https://github.com/turbofish-org/merk)
* [go-ethereum Merkle Patricia Trie](https://github.com/ethereum/go-ethereum/blob/master/trie/trie.go)

Ethan Frey [spent quite some time](https://github.com/confio/proofs-ethereum) to wrestle out a well-defined serialization and a validation logic that didn't involve importing too much code from go-ethereum (copying parts and stripping it down to the minimum). At the end, he realized the key is only present from parsing the entire path and is quite a painstaking process, even with go code that imports rlp and has the custom patricia key encodings. After a long time reflecting, he cannot see any way to implement this that doesn't either: (A) allow the provider to forge a different key that cannot be detected by the verifier or (B) include a huge amount of custom code in the client side.

If anyone has a solution to this, please open an issue in this repository.

## Known limitations

This format is designed to support any merklized data store that encodes key-value pairs into
a node, and computes a root hash by repeatedly concatenating hashes and re-hashing them.

Notably, this requires the *key* to be present in the leaf node in order to enforce the structure properly
and prove the *key* provided matches the proof without extensive db-dependent code. Since some
tries (such as Ethereum's Patricia Trie) do not store the key in the leaf, but require precise analysis of
every step along the path in order to reconstruct the path, these are not supported. Doing so would
require a much more complex format and most likely custom code for each such database. The design goal
was to be able to add new data sources with only configuration (spec object), rather than custom code.

[cosmos-shield]: https://img.shields.io/static/v1?label=&labelColor=1B1E36&color=1B1E36&message=cosmos%20ecosystem&style=for-the-badge&logo=data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPCEtLSBHZW5lcmF0b3I6IEFkb2JlIElsbHVzdHJhdG9yIDI0LjMuMCwgU1ZHIEV4cG9ydCBQbHVnLUluIC4gU1ZHIFZlcnNpb246IDYuMDAgQnVpbGQgMCkgIC0tPgo8c3ZnIHZlcnNpb249IjEuMSIgaWQ9IkxheWVyXzEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4IgoJIHZpZXdCb3g9IjAgMCAyNTAwIDI1MDAiIHN0eWxlPSJlbmFibGUtYmFja2dyb3VuZDpuZXcgMCAwIDI1MDAgMjUwMDsiIHhtbDpzcGFjZT0icHJlc2VydmUiPgo8c3R5bGUgdHlwZT0idGV4dC9jc3MiPgoJLnN0MHtmaWxsOiM2RjczOTA7fQoJLnN0MXtmaWxsOiNCN0I5Qzg7fQo8L3N0eWxlPgo8cGF0aCBjbGFzcz0ic3QwIiBkPSJNMTI1Mi42LDE1OS41Yy0xMzQuOSwwLTI0NC4zLDQ4OS40LTI0NC4zLDEwOTMuMXMxMDkuNCwxMDkzLjEsMjQ0LjMsMTA5My4xczI0NC4zLTQ4OS40LDI0NC4zLTEwOTMuMQoJUzEzODcuNSwxNTkuNSwxMjUyLjYsMTU5LjV6IE0xMjY5LjQsMjI4NGMtMTUuNCwyMC42LTMwLjksNS4xLTMwLjksNS4xYy02Mi4xLTcyLTkzLjItMjA1LjgtOTMuMi0yMDUuOAoJYy0xMDguNy0zNDkuOC04Mi44LTExMDAuOC04Mi44LTExMDAuOGM1MS4xLTU5Ni4yLDE0NC03MzcuMSwxNzUuNi03NjguNGM2LjctNi42LDE3LjEtNy40LDI0LjctMmM0NS45LDMyLjUsODQuNCwxNjguNSw4NC40LDE2OC41CgljMTEzLjYsNDIxLjgsMTAzLjMsODE3LjksMTAzLjMsODE3LjljMTAuMywzNDQuNy01Ni45LDczMC41LTU2LjksNzMwLjVDMTM0MS45LDIyMjIuMiwxMjY5LjQsMjI4NCwxMjY5LjQsMjI4NHoiLz4KPHBhdGggY2xhc3M9InN0MCIgZD0iTTIyMDAuNyw3MDguNmMtNjcuMi0xMTcuMS01NDYuMSwzMS42LTEwNzAsMzMycy04OTMuNSw2MzguOS04MjYuMyw3NTUuOXM1NDYuMS0zMS42LDEwNzAtMzMyCglTMjI2Ny44LDgyNS42LDIyMDAuNyw3MDguNkwyMjAwLjcsNzA4LjZ6IE0zNjYuNCwxNzgwLjRjLTI1LjctMy4yLTE5LjktMjQuNC0xOS45LTI0LjRjMzEuNi04OS43LDEzMi0xODMuMiwxMzItMTgzLjIKCWMyNDkuNC0yNjguNCw5MTMuOC02MTkuNyw5MTMuOC02MTkuN2M1NDIuNS0yNTIuNCw3MTEuMS0yNDEuOCw3NTMuOC0yMzBjOS4xLDIuNSwxNSwxMS4yLDE0LDIwLjZjLTUuMSw1Ni0xMDQuMiwxNTctMTA0LjIsMTU3CgljLTMwOS4xLDMwOC42LTY1Ny44LDQ5Ni44LTY1Ny44LDQ5Ni44Yy0yOTMuOCwxODAuNS02NjEuOSwzMTQuMS02NjEuOSwzMTQuMUM0NTYsMTgxMi42LDM2Ni40LDE3ODAuNCwzNjYuNCwxNzgwLjRMMzY2LjQsMTc4MC40CglMMzY2LjQsMTc4MC40eiIvPgo8cGF0aCBjbGFzcz0ic3QwIiBkPSJNMjE5OC40LDE4MDAuNGM2Ny43LTExNi44LTMwMC45LTQ1Ni44LTgyMy03NTkuNVMzNzQuNCw1ODcuOCwzMDYuOCw3MDQuN3MzMDAuOSw0NTYuOCw4MjMuMyw3NTkuNQoJUzIxMzAuNywxOTE3LjQsMjE5OC40LDE4MDAuNHogTTM1MS42LDc0OS44Yy0xMC0yMy43LDExLjEtMjkuNCwxMS4xLTI5LjRjOTMuNS0xNy42LDIyNC43LDIyLjYsMjI0LjcsMjIuNgoJYzM1Ny4yLDgxLjMsOTk0LDQ4MC4yLDk5NCw0ODAuMmM0OTAuMywzNDMuMSw1NjUuNSw0OTQuMiw1NzYuOCw1MzcuMWMyLjQsOS4xLTIuMiwxOC42LTEwLjcsMjIuNGMtNTEuMSwyMy40LTE4OC4xLTExLjUtMTg4LjEtMTEuNQoJYy00MjIuMS0xMTMuMi03NTkuNi0zMjAuNS03NTkuNi0zMjAuNWMtMzAzLjMtMTYzLjYtNjAzLjItNDE1LjMtNjAzLjItNDE1LjNjLTIyNy45LTE5MS45LTI0NS0yODUuNC0yNDUtMjg1LjRMMzUxLjYsNzQ5Ljh6Ii8+CjxjaXJjbGUgY2xhc3M9InN0MSIgY3g9IjEyNTAiIGN5PSIxMjUwIiByPSIxMjguNiIvPgo8ZWxsaXBzZSBjbGFzcz0ic3QxIiBjeD0iMTc3Ny4zIiBjeT0iNzU2LjIiIHJ4PSI3NC42IiByeT0iNzcuMiIvPgo8ZWxsaXBzZSBjbGFzcz0ic3QxIiBjeD0iNTUzIiBjeT0iMTAxOC41IiByeD0iNzQuNiIgcnk9Ijc3LjIiLz4KPGVsbGlwc2UgY2xhc3M9InN0MSIgY3g9IjEwOTguMiIgY3k9IjE5NjUiIHJ4PSI3NC42IiByeT0iNzcuMiIvPgo8L3N2Zz4K
