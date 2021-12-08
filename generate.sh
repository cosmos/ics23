#!/usr/bin/env bash

set -eo pipefail

protoc_gen_gocosmos() {
  go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@latest 2>/dev/null
}

protoc_gen_gocosmos

buf generate

mv ./proto/confio/ics23/v1/proofs.pb.go ./go
