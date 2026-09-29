# ICS23 Verification — Traceability

Maps each function in the verifier to its Lean model counterpart and the
machine-checked results about it. Keep this current: a change to a left-column
function requires updating the corresponding model definition and re-checking
the right-column results (enforced socially via review; CI rebuilds the proofs).

Legend: ✅ proved (no `sorry`) · 🟡 stated, proof in progress · 🔧 Kani-verified
(Rust) · 📋 exercised by the regression corpus.

## Ops (`rust/src/ops.rs` ↔ `lean/Ics23/Ops.lean`, `Varint.lean`)

| Rust | Lean | Results |
|------|------|---------|
| `apply_inner` | `applyInner` | ✅ `innerImage_inj`, `applyInner_inj` (Soundness.lean) |
| `apply_leaf` | `applyLeaf` | used by `existence_binding_sameshape` |
| `do_hash` | `HashFn` (abstract) | no hash assumption; collisions are exhibited |
| `do_length` | `doLength` | ✅ `doLength_noPrefix_inj`, `doLength_varProto_inj` |
| `prepare_leaf_data` | `prepareLeafData` | ✅ `prepareLeafData_inj` |
| `proto_len` | `varintEncode` | ✅ `varintEncode_append_inj` (self-delimiting); 🔧 `proto_len` overflow precondition |

## Existence (`rust/src/verify.rs` ↔ `lean/Ics23/Verify.lean`, `Soundness.lean`, `Existence.lean`)

| Rust | Lean | Results |
|------|------|---------|
| `verify_existence` | `verifyExistence` | ✅ `existence_binding_sameshape{,_noPrefix,_varProto}` (same-shape, all specs); 🟡 general `existence_binding` (differing-path) |
| `check_existence_spec` | `checkExistenceSpec` | 📋 depth-bound + leaf/inner checks |
| `calculate_existence_root_for_spec` | `calculateExistenceRoot` | ✅ `verifyExistence_root`, `calculateExistenceRoot_eq`; ✅ `applyPath_sameops_inj` |
| `ensure_leaf` | `ensureLeaf` | used by binding |
| `ensure_inner` | `ensureInner` | 📋 corpus (A2 domain sep, A3 child-size/prefix); 🔧 prefix-bound overflow (F1) |
| `has_prefix` | `hasPrefix` | ✅ `hasPrefix_refl` |

## Non-existence (`rust/src/verify.rs` ↔ `lean/Ics23/NonExist.lean`, `NonExistSound.lean`)

| Rust | Lean | Results |
|------|------|---------|
| `verify_non_existence` | `verifyNonExistence` | 🟡 `nonexistence_sound`; ✅ `verifyNonExistence_{left,right}`, `verifyNonExistence_neighbors_ordered` |
| `ensure_left_most` / `ensure_right_most` | `ensureLeftMost` / `ensureRightMost` | 📋 corpus (SMT padding) |
| `ensure_left_neighbor` | `ensureLeftNeighbor` | (modeled) |
| `is_left_step` | `isLeftStep` | 📋 corpus |
| `get_padding` / `has_padding` / `order_from_padding` | `getPadding` / `hasPadding` / `orderFromPadding` | 📋 corpus; 🔧 `get_padding` overflow |
| `left_branches_are_empty` | `leftBranchesAreEmpty` | 📋 corpus; 🔧 slice-in-bounds |
| `right_branches_are_empty` | `rightBranchesAreEmpty` | 📋 corpus; 🔧 slice-in-bounds (binary); finding F2 for ternary+ |
| lexicographic `Vec<u8>` ordering | `bytesLt` | ✅ `bytesLt_irrefl`, `bytesLt_ne`, `bytesLt_trans` |

## Specs (`rust/src/api.rs` ↔ `lean/Ics23/Specs.lean`)

| Rust | Lean | Results |
|------|------|---------|
| `iavl_spec` | `iavlSpec` | ✅ `iavl_wellFormed` (Theorem C) |
| `tendermint_spec` | `tendermintSpec` | ✅ `tendermint_wellFormed` |
| `smt_spec` | `smtSpec` | ✅ `smt_wellFormed` |
| `verify_membership` / `verify_non_membership` | (thin wrappers over the above) | not modeled separately |
| `ensure_leaf_prefix` / `ensure_inner_prefix` (IAVL) | abstracted away | sound abstraction (only restricts; see properties.md) |

## Executable / oracle (`lean/Ics23/Sha256.lean`, `Executable.lean`)

| Item | Lean | Results |
|------|------|---------|
| SHA-256 | `Sha256.hash` | ✅ validated against `rust/src/ops.rs` vectors |
| concrete verifier runs | `concreteHash` + demos | end-to-end verify + forgery rejection (existence and non-existence) by `native_decide` |

## Implementation safety (`rust/src/kani_proofs.rs`)

| Harness | Property | Status |
|---------|----------|--------|
| `ensure_inner_prefix_bound_no_overflow` | i32 prefix-bound overflow-free (well-formed) | 🔧 verified |
| `get_padding_no_overflow` | i32 padding products overflow-free | 🔧 verified |
| `left_branches_slice_in_bounds` | prefix slice in bounds | 🔧 verified |
| `right_branches_slice_in_bounds_binary` | suffix slice in bounds (binary) | 🔧 verified |

## Not yet covered

- General `existence_binding` (differing-path, A3) and `nonexistence_sound`
  (tree semantics) — the two open `sorry`s.
- Batch / compressed proofs (`compress.rs`, batch handling).
- Go implementation: covered by acceptance equivalence via the planned Phase 2a
  differential oracle, not by proof.
