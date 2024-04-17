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

echo "Cleaning API directory"
(cd go/api; find ./ -type f \( -iname \*.pulsar.go -o -iname \*.pb.go -o -iname \*.pb.gw.go \) -delete; find . -empty -type d -delete; cd ../..)

echo "Generating API module"
(cd proto; buf generate --template buf.gen.pulsar.yaml)

# move files to appropriate location (go/api)
cp -r go/api/github.com/cosmos/ics23/go/* go/api/
rm -rf go/api/github.com

# tidy go/api module
(cd go/api; go mod tidy)