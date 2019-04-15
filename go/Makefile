.PHONY: protoc test

# PROTOC_FLAGS := -I=.. -I=./vendor -I=$(GOPATH)/src
PROTOC_FLAGS := -I=.. -I=$(GOPATH)/src

test:
	go test .

protoc:
	protoc --gogofaster_out=. $(PROTOC_FLAGS) ../proofs.proto

