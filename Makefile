DOCKER := $(shell which docker)
HTTPS_GIT := https://github.com/cosmos/ics23.git

##### GO commands #####

include go/Makefile


##### Protobuf #####
protoVer=0.11.2
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

##### Rust #####
rustVer=1.65-slim
rustImageName=rust:$(rustVer)
rustImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(rustImageName)

proto-all: proto-format proto-lint proto-gen-go proto-gen-rust

proto-gen-go:
	@echo "Generating Protobuf definitons for Go"
	@$(protoImage) sh ./scripts/protocgen_go.sh

proto-gen-rust:
	@echo "Generating Protobuf definitons for Rust"
	@$(rustImage) sh ./scripts/protocgen_rust.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=master

.PHONY: proto-all proto-gen-go proto-gen-rust proto-format proto-lint proto-check-breaking
