NIGHTLY_VERSION=nightly

.DEFAULT_GOAL := help

.PHONY: all setup check-panic-conflict check-cargo-build-std check-wasm help

all: ## Run all checks
	$(MAKE) check-panic-conflict
	$(MAKE) check-cargo-build-std
	$(MAKE) check-wasm

setup: ## Setup the required nightly toolchain and the wasm32 target
	rustup install $(NIGHTLY_VERSION)
	rustup target    add wasm32-unknown-unknown --toolchain $(NIGHTLY_VERSION)
	rustup component add rust-src               --toolchain $(NIGHTLY_VERSION)

check-panic-conflict: ## Check for `no_std` compliance by installing a panic handler, and any other crate importing `std` will cause a conflict. Runs on default target.
	cargo build \
		--no-default-features \
		--features panic-handler

check-cargo-build-std: ## Check for `no_std` compliance using Cargo nightly's `build-std` feature. Runs on the target `x86_64-unknown-linux-gnu`.
	rustup run $(NIGHTLY_VERSION) -- \
		cargo build -Z build-std=core,alloc \
		--no-default-features \
		--target x86_64-unknown-linux-gnu

check-wasm: ## Check for WebAssembly and `no_std` compliance by building on the target `wasm32-unknown-unknown` and installing a panic handler.
	rustup run $(NIGHTLY_VERSION) -- \
		cargo build \
		--features panic-handler \
		--target wasm32-unknown-unknown

help: ## Show this help message
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

