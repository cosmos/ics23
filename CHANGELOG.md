# Changelog

# 0.8.0
The following functions have been made generic over a new trait `HostFunctionProvider`:
    
- [x] `calculate_existence_root` 
- [x] `verify_batch_membership` 
- [x] `verify_batch_non_membership` 
- [x] `verify_membership` 
- [x] `verify_non_membership`

For `wasm32-unknown-unknown` environments this trait allows you to delegate hashing functions to a native implementation  through host functions. For `std` you can simply use `ics23::HostFunctionManager` as this provides a default implementation of this trait.

# v0.7.0

This handles non-existence tests for empty branches properly. This
is needed for properly handling proofs on Tries, like the SMT being 
integrated with the Cosmos SDK.

This is used in ibc-go v3

# 0.6.x

This handles proofs for normal merkle trees, where every branch is full.
This works for tendermint merkle hashes and iavl hashes, and should work
for merk (nomic's db) proofs.

This was used in the original ibc release (cosmos sdk v0.40) and up until
ibc-go v2.