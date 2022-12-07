#!/usr/bin/env bash

set -e

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./cosmos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
      buf generate --template buf.gen.gogo.yaml $file
  done
done

cd ..

cp -r github.com/cosmos/ics23/go/* ./go
rm -rf github.com
