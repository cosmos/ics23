.PHONY: protoc test

# make sure we turn on go modules
export GO111MODULE := on

# PROTOC_FLAGS := -I=.. -I=./vendor -I=$(GOPATH)/src
PROTOC_FLAGS := -I=.. -I=$(GOPATH)/src

test:
	go test .

protoc:
#	@go mod vendor
	protoc --gogofaster_out=. $(PROTOC_FLAGS) ../proofs.proto

