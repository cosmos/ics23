.PHONY: protoc test

# make sure we turn on go modules
export GO111MODULE := on

# PROTOC_FLAGS := -I=.. -I=./vendor -I=$(GOPATH)/src
PROTOC_FLAGS := -I=.. -I=$(GOPATH)/src

test:
	go test .

protoc:
#	@go mod vendor
	buf generate --template ../buf.gen.yaml --path ..

install-proto-dep:
	@echo "Installing protoc-gen-gocosmos..."
	@go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos
