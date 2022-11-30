#!/usr/bin/env bash

set -e

apt-get update
apt-get install -y protobuf-compiler

cd rust/codegen
cargo run

