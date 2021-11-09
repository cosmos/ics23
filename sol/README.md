## how to compile solidity go bindings to run unit tests
- be sure to have a local version of abigen aligned with the go-ethereum package version declared in the go.mod
```bash
$ grep go-ethereum go.mod
	github.com/ethereum/go-ethereum v1.10.0

```
and
```bash
$ abigen --version
abigen version 1.10.0-stable

```
- be sure to clone locally the solidity dependencies - check version tags in solidity code
```bash
$ git clone --depth 1 --branch <version-tag> https://github.com:OpenZeppelin/openzeppelin-contracts.git
$ git clone --depth 1 --branch <verion-tag> https://github.com/GNSPS/solidity-bytes-utils
```
- compile solidity code with the proper alias substitutions
```bash
$ solc @openzeppelin/contracts@<version-tag>/=<local-path-to-openzeppelin-repo>/openzeppelin-contracts/contracts/ \
     @solidity-bytes-utils@<version-tag>=<local-path-to-solidity-bytes-util-repo>/solidity-bytes-utils \
     --optimize \
     ics23Ops.sol ics23Proof.sol ics23Compress.sol ics23.sol --combined-json abi,ast,bin > ics23.json
```
- generate go bindings from abi and binaries
```bash
$ abigen  --pkg ics23_sol --out ics23.go --combined-json ics23.json
```

- everything is set up to run tests
```bash
$ go test ./...
```
