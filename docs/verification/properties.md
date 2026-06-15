# ICS23 Verification — Property Catalogue (Phase 0)

Status: living document. Companion to `docs/rfc/001-formal-verification.md` and
the Lean development under `lean/`.

This catalogue enumerates the properties the formal model must establish and the
invariants the verifier relies on. Each property links to where it is enforced
in the code and where it lives in the Lean model. It also records the regression
corpus: concrete malicious inputs the model is required to reject.

## Sources to mine

- `docs/audits/ICS-23 - Zellic Audit Report.pdf` — **transcribed** into the
  regression corpus (`lean/Ics23/Corpus.lean`, "Zellic audit findings" section):
  all eight findings recorded, with machine-checked model witnesses for the ones
  that have a decidable counterpart (3.1 sorted-keys → `SortedTree` hypothesis,
  3.2 `MaxPrefixLength` → `innerWFB`, 3.4 zero `child_size`, 3.5 proof-size depth
  gate, 3.6 out-of-bounds `child_order` → total `getPadding`) and notes for the
  implementation-only/out-of-scope ones (3.3 TS, 3.7 compressed batch, 3.8
  stricter checks). See "Zellic finding map" below.
- The 2020 Informal Systems audit of the original `confio/ics23`.
- The Cosmos "Dragonberry" advisory (Oct 2022) — proof-forgery class in the
  Cosmos proof-verification stack; the canonical motivation.
- The ICS-023 specification (vector commitments).

## Modeling assumptions (must stay honest)

These are the places where the Lean model deliberately differs from the Rust.
Each is justified as *sound* (the model accepts at least every proof the code
accepts), so a soundness theorem about the model transfers to the code.

1. **Optional fields collapsed.** `ProofSpec.{leaf,inner}_spec` and
   `ExistenceProof.leaf` are `Option` in proto/Rust; a `None` triggers an
   immediate reject. The model makes them non-optional, dropping inputs the
   code rejects. Enlarges the accepted set ⇒ sound to omit.
2. **IAVL prefix checks omitted.** `ensure_leaf_prefix` / `ensure_inner_prefix`
   (`rust/src/api.rs`) only fire when the spec equals the IAVL spec and can only
   reject. Omitting them enlarges the accepted set ⇒ sound to omit. (A later
   increment may add them to also reason about IAVL prefix structure.)
3. **Hash family abstract.** `do_hash` is modeled as an arbitrary
   `HashFn : HashOp → Bytes → Bytes` with the single law `noHash = id`. No
   collision-resistance assumption is made; soundness theorems *exhibit* a
   collision instead.
4. **Spec integers are `Int`.** Faithful to the `i32` proto fields, so malformed
   (e.g. negative) specs are representable and excluded by `WellFormed` rather
   than by Lean typing. (Wrap/overflow of the `i32`/`usize` casts in the Rust is
   a separate, code-level concern handled by Kani in Phase 3, not here.)

## Properties

### A. Existence binding (soundness) — headline

A well-formed spec and a fixed root cannot bind one key to two different values
without a hash collision.

- Enforced by: `verify_existence` + `check_existence_spec` +
  `calculate_existence_root_for_spec` (`rust/src/verify.rs`).
- Model: proved — byte-level `Ics23.existence_binding` (Existence.lean) and
  honest-root `membership_sound` / `membership_sound_tendermint` (Tree.lean) /
  `membership_sound_smt` (SmtTree.lean). See "Remaining obligations" items 0–1.
- Rests on:
  - **A1 Leaf-encoding injectivity.** `prefix ++ enc(prehash key) ++ enc(prehash value)`
    parses uniquely into `(key, value)` — via a length-determining `LengthOp`
    (VarProto/Fixed*/Require*) or fixed-length prehash images on both fields.
    Captured by `leafDelimitingB`. Violated ⇒ a forger can re-split bytes into a
    different `(key, value)`.
  - **A2 Leaf/inner domain separation.** No accepted inner-op preimage carries
    the leaf prefix (`ensure_inner`: `!has_prefix(leaf_spec.prefix, inner.prefix)`).
    Requires a nonempty leaf prefix (`wellFormedB`). Violated ⇒ a leaf hash can
    be reinterpreted as an inner hash (depth confusion).
  - **A3 Positional unambiguity.** `min_prefix_length`, `max_prefix_length`,
    `child_size`, suffix length, and `max < min + child_size` together fix a
    child's position in its parent. Captured by `innerWFB` and `ensure_inner`.
  - **A4 Inner preimage injectivity in the child.** Same op ⇒ equal images force
    equal children. Proved: `Ics23.innerImage_inj`.

### B. Non-existence soundness

A well-formed spec and a fixed root cannot simultaneously admit a non-existence
proof for key `k` and an existence proof for `k`.

- Enforced by: `verify_non_existence`, `ensure_left_most`, `ensure_right_most`,
  `ensure_left_neighbor`, `left_branches_are_empty`, `right_branches_are_empty`
  (`rust/src/verify.rs`).
- Model: proved in the honest-root tree models, total over verifier-accepted
  shapes — `nonexistence_sound_tree_total` / `_tendermint_total`
  (TreeNonExist.lean) and `nonexistence_sound_smt_total` (SmtNonExist.lean,
  including `empty_child` placeholder handling and hashed-key order). See
  "Remaining obligations" items 2–2a. Arbitrary `child_order` permutations
  remain out of scope (all shipped specs are binary `[0, 1]`).
- Note: when `prehash_key_before_comparison` is set (SMT/JMT), the ordering is
  over hashed keys, so the guarantee is non-existence of the *hashed* key.

### C. Spec well-formedness certificate — proved

A decidable `WellFormed` predicate capturing exactly A1–A3's side conditions,
plus machine-checked certificates that the shipped specs satisfy it.

- Model: `wellFormedB` / `WellFormed`; certificates `iavl_wellFormed`,
  `tendermint_wellFormed`, `smt_wellFormed` (proved by `decide`).
- Standalone value: a checkable answer to "is this custom `ProofSpec` safe?"

### D. Implementation safety (Rust) — Phase 3 (Kani)

Panic freedom, no lossy/overflowing integer casts (notably the `i32`/`i64`/`usize`
casts in `ensure_inner`, `ensure_inner_prefix`, and compressed-batch index
handling), termination, input-bounded allocation. Not in the Lean model.

Landed (`rust/src/kani_proofs.rs`, `#[cfg(kani)]`, run by `cargo kani` / CI):
**five harnesses verified** by CBMC — the `ensure_inner` prefix-bound `i32`
arithmetic and the `get_padding` products are overflow-free under well-formed
bounds; the `left_branches_are_empty` / `right_branches_are_empty` (binary) slice
accesses are always in bounds (the Dragonberry-class out-of-range-slice surface);
and `decompress_exist`'s `lookup.get(x as usize)` compressed-batch index handling
is panic-free for any attacker-controlled `i32` index. Noted finding: an
adversarial spec with an enormous `child_size` could overflow the `i32`
prefix-bound product; the harnesses pin the safe precondition. Remaining:
panic-freedom of the `Result`-returning entry points (`do_length`, `proto_len`)
needs formatting stubs to keep CBMC tractable.

### E. Go/Rust acceptance equivalence — Phase 2a (differential oracle)

The two implementations accept exactly the same tuples. Covered by differential
testing of the executable model against both, not by proof.

## Regression corpus (model must reject)

Concrete malicious proofs, each targeting one invariant. To be encoded as Lean
`example … = false` and as fixtures for the Phase 2a oracle.

- [ ] **A1-split:** NoPrefix length with variable-length prehash (e.g. NoHash on
      both fields) — re-split `key ++ value` to forge a different pair under one
      root. (A spec with this shape must be rejected by `WellFormed`.)
- [ ] **A2-depthconfusion:** an inner op whose prefix begins with the leaf
      prefix; must fail `ensure_inner`.
- [ ] **A3-childsize:** inner suffix length not a multiple of `child_size`; must
      fail `ensure_inner`.
- [ ] **A3-prefixwindow:** inner prefix length outside `[min, max + maxLeftChild]`.
- [ ] **C-negative:** spec with non-positive `child_size` or
      `max_prefix_length ≥ min_prefix_length + child_size`; must fail `WellFormed`.
- [ ] **depth-bounds:** `min_depth ≠ 0` with path length outside `[min, max]`.
- [x] **Zellic findings** — transcribed; see the map below.

### Zellic finding map (audit → model)

The July 2024 Zellic assessment (`docs/audits/`) reported eight findings. Each is
recorded in `lean/Ics23/Corpus.lean` under "Zellic audit findings", with a
machine-checked witness where the model has a decidable counterpart:

| # | Finding (severity) | Model treatment |
|---|---|---|
| 3.1 | Nonexistence soundness depends on sorted keys (Critical) | The `SortedTree`/`SortedTreeS` hypothesis of Theorem B (modeling note F4). Witness: the Zellic Fig. 3.1 tree is a member-bearing tree machine-checked **not** `SortedTree`. |
| 3.2 | Forging with `MaxPrefixLength ≥ MinPrefixLength + ChildSize` (Critical) | `innerWFB` requires `max < min + child_size`; the split-pinning lemmas rely on it. Witness: the boundary spec fails `wellFormedB` and the op fails `ensureInner`. |
| 3.3 | TypeScript skips IAVL prefix checks (Critical) | Implementation-specific (TS removed upstream). Model is the Rust/Go surface; omits IAVL prefix checks as a *sound over-approximation* (modeling assumption 2). Note only. |
| 3.4 | Zero division in `ensure_inner` (High) | `child_size = 0` fails `innerWFB`/`ensureInner` (model is `%`-total, no panic). Witness: `ensureInner … negChildSizeSpec = false`. Rust panic-freedom → Kani. |
| 3.5 | Unrestricted proof size (Medium) | Liveness/DoS, outside soundness. Model has the depth gate; shipped specs disable it (`min_depth = 0`), faithfully reflecting the finding. Witness: over-deep path fails `checkExistenceSpec` once armed. |
| 3.6 | Panic on out-of-bounds `ChildOrder` (Low) | Model's `getPadding` is total (returns `none`, never panics); theorems stated for binary `[0,1]`. Witness: `getPadding {childOrder := [0,2]} 1 = none`. |
| 3.7 | Panic in `decompressExist` (Low) | Compressed/batched proofs removed upstream; out of model scope. Note only. |
| 3.8 | IAVL/Tendermint checks could be stricter (Informational) | The honest-root tree models bake the stricter node shape into `WFTree`/`WFTreeI`/`WFTreeS`. Note only. |

## Findings (surfaced by the verification work)

- **F4 — non-existence soundness requires the store key-sortedness invariant
  (it is a hypothesis, not enforced by the verifier).** `verify_non_existence`
  checks `key_for_comparison(l.key) < key_for_comparison(key) <
  key_for_comparison(r.key)` and that `l`, `r` are *position*-adjacent leaves
  (`ensure_left_neighbor` etc.). But it never checks that leaf positions track
  key order. So nothing in the verifier alone rules out an existence proof
  placing `key`'s leaf at an unrelated position while its key value sits between
  the neighbors — i.e. `nonexistence_sound` is **false without** assuming the
  tree's leaves are sorted by key (left-to-right). ICS23 states this requirement
  informally ("Supported merkle stores must be lexicographically ordered to
  maintain soundness"); the formalization makes it precise: Theorem B must carry
  a `KeySorted root` hypothesis (any two existence proofs to `root` have key
  order iff position order). The `Order.lean` position lemmas are the start of
  stating that invariant; with it, the proof is: `key` between `l`,`r` in key
  order ⇒ (sortedness) between them in position order ⇒ contradicts adjacency.
- **F5 — the adjacency step is hash-structural, not pure lex (proof-obligation
  refinement).** The adjacency lemma (`ensure_left_neighbor` ⇒ no leaf position
  strictly between the neighbors) is *not* provable from the `lexLt` order alone:
  a position `C ++ [0] ++ (k+1 ones)` is lex-greater than the left neighbor
  `C ++ [0] ++ (k ones)` yet still less than the right neighbor `C ++ [1] ++
  (zeros)`. "No position between" holds only because the left neighbor's leaf is
  **terminal** (no deeper-right leaf) — a fact about the tree's hash structure,
  not the position order. A proof placing a leaf "past" the rightmost terminal
  leaf forces an *inner* hash where the neighbor has a *leaf* hash, reducing to
  `leaf_inner_domain_collision` (already proved). So Theorem B's remaining core is
  the *integration* of the `Order.lean` position model with the hash-level
  domain-separation machinery; the `lexLt` order theory is necessary scaffolding
  but not sufficient alone.


- **F1 — i32 prefix-bound overflow (malformed spec).** `ensure_inner` computes
  `max_prefix_length + (child_order.len()-1)*child_size` in `i32`; an adversarial
  spec with an enormous `child_size` overflows it. Not reachable with the shipped
  specs; the Kani harness pins the safe precondition.
- **F3 — positional ambiguity ⇒ general binding needs more than collision
  resistance (analysis, not a live exploit).** Working the differing-path case of
  Theorem A surfaced this. For a binary spec with `min_prefix_length =
  max_prefix_length = p` (Tendermint, SMT), a node preimage
  `P = prefix(p) ‖ child(cs) ‖ sibling(cs)` is accepted by `ensure_inner` under
  *two* readings: as a **left**-child step (`|prefix| = p`, child `= P[p..p+cs]`,
  suffix `= P[p+cs..]`) and as a **right**-child step (`|prefix| = p+cs`, child
  `= P[p+cs..p+2cs]`, suffix empty) — both satisfy the prefix window
  `[min, max + (n-1)·cs] = [p, p+cs]` and the `suffix % cs = 0` check. So the same
  node hash does not, by the spec checks alone, determine which half is the
  recursive child.
  - Consequence: the clean "two accepted existence proofs that disagree ⇒ exhibit
    a hash *collision*" reduction does **not** go through for differing tree
    positions against an *arbitrary* hash `H`. An adversary exploiting the
    ambiguity would need the sibling bytes to themselves be a valid
    subtree/leaf hash of a different value — a **preimage** problem, not a
    collision. Against an arbitrary `H` that is not ruled out; the honest model
    must therefore either (a) treat hashes as injective/opaque tokens (a
    structured/"free" hash model) or assume preimage resistance, or (b) re-include
    the per-store prefix structure the model abstracts away — notably IAVL's
    `ensure_inner_prefix`, which encodes `height/size/version` and pins the
    position explicitly. The **same-shape** binding theorem (already proved) sits
    below this obstacle because it fixes the path structure.
  - **The IAVL prefix check does *not* fix it (machine-checked).** Re-including
    `ensure_inner_prefix` (height/size/version + remaining-byte count) was the
    natural candidate to pin position. It does not: for the IAVL preimage
    `hsv ‖ 0x20 ‖ A ‖ 0x20 ‖ B`, the left reading leaves `remaining = 1` and the
    right reading leaves `remaining = 34`, and the check explicitly admits *both*
    (see `lean/Ics23/IavlPrefix.lean`, `native_decide` witnesses). So the
    ambiguity is intrinsic to the prefix/suffix construction across all shipped
    specs, IAVL included. Binding genuinely rests on **preimage resistance** (the
    sibling bytes would have to be a valid subtree hash of a different value),
    which the abstract-hash model does not capture.
  - Action: this is why `existence_binding` (general) is still open. The correct
    fix is a hash-model refinement — model node hashes as opaque/injective tokens
    (a symbolic Merkle model) so sibling bytes cannot be re-read as a subtree
    hash — or weaken the theorem's conclusion to "collision OR preimage". This is
    a foundational reformulation, not a tactic tweak. No evidence of a live
    exploit against the shipped stores; flagged for maintainer review.
- **F2 — left/right empty-branch asymmetry (ternary+ specs).**
  `right_branches_are_empty` guards `suffix.len() == child_size` (one child) but
  then reads `suffix[i*child_size..]` for `i in 0..right_branches`. For
  `child_order.len() > 2`, `right_branches` can exceed 1, so the `i = 1` access is
  out of bounds — a panic. The left side correctly sizes the prefix by
  `left_branches*child_size`. No supported store is ternary (merk is unsupported),
  so this is latent, not currently exploitable. Surfaced while writing the Kani
  harnesses; the binary-case harness is verified.

## Proof status (Lean)

Model is complete (existence + non-existence). Proved with no `sorry`:
Theorem C (all three spec certificates); A1 leaf-encoding injectivity for both
`NoPrefix` (fixed-prehash) and `VarProto` (varint self-delimiting) shapes;
the path-fold backbone (`applyInner_inj`, `applyPath_sameops_inj`,
`applyPath_eqlen_merge`); **same-shape existence binding for all three shipped
specs** (`existence_binding_sameshape{,_noPrefix,_varProto}`); and
**equal-length existence binding** (`existence_binding_eqlen`) — same key, same
leaf op, equal path *depth* but arbitrary differing inner ops ⇒ `HashCollision ∨
PositionalAmbiguity`. Non-existence padding / empty-branch logic is exercised by
the corpus.

### Remaining obligations

0. **Theorem A — honest-root form, DONE (`lean/Ics23/Tree.lean`).** The
   *strongest* existence result: a Merkle tree model (`MTree`, `rootHash`,
   `TreeMember`) and `membership_sound` — an existence proof verifying against
   `root = rootHash t` for a real (Tendermint-shaped) tree implies *genuine
   membership* in `t`, up to a hash collision, with **no ambiguity arm**. The F3
   positional ambiguity is fully resolved by `split_pins`: the verifier's
   `suffix % child_size` check pins a node split to exactly the two genuine
   children, so an accepted proof follows actual tree structure. Instantiated:
   `membership_sound_tendermint` (structural side conditions by `decide`; joint
   leaf injectivity now *proved* in `leafInj_tendermint`/`leafInj_varProto`, so
   `FixedHash` is the only remaining clean crypto assumption). This is Option 1
   delivered for existence — it removes the byte-level disjunction's ambiguity
   arms entirely.

1. **Theorem A — DONE (existence binding, byte-level).** Fully proved, no `sorry`:
   `existence_binding_shaped` (general, production-spec shape) concludes the
   honest three-way disjunction `HashCollision ∨ PositionalAmbiguity ∨
   LeafAmbiguity`; `existence_binding_sameleaf` gives the stronger two-way
   conclusion for the honest same-leaf case (a key's leaf op is determined in a
   real tree). The two ambiguity arms are real machine-checkable obstructions
   (F3 and its leaf-level analogue); collapsing them to a bare collision is what
   the symbolic-Merkle model would add — not required for the honest theorem.
2. **Theorem B — honest-root form, DONE, total
   (`lean/Ics23/TreeNonExist.lean`).** Non-existence soundness in the same
   honest-root tree model as item 0, covering **every proof shape the verifier
   accepts**: a verifying non-existence proof for `key` plus a verifying
   existence proof for `key` against `root = rootHash t` of a key-sorted
   (`SortedTree`) well-formed tree yields a `HashCollision`.
   - *Two-sided* (`nonexistence_sound_tree`): `ensure_left_neighbor`
     decomposes (`ensureLeftNeighbor_spec`) into a shared root path, a
     left-step divergence, and right-most/left-most remainders;
     `neighbor_divergence` (byte→tree navigation induction) pins both
     bracketing proofs to a real divergence node `N`, so the absent key would
     sit in the gap between `maxKey N.left` and `minKey N.right` —
     `node_gap_no_member` rules that out in a sorted tree. The F4/F5
     padding-vs-placeholder subtlety is discharged by
     `ensureRightMost_suffix_nil` / `ensureLeftMost_suffix_cs` (needs
     `emptyChild.length ≠ childSize`, true for Tendermint's `[]`).
   - *One-sided* (`nonexistence_sound_tree_leftOnly` / `_rightOnly`): a
     left-only proof's `ensure_right_most` path makes its neighbor the
     rightmost leaf (`reaches_max`), so `maxKey t < key` contradicts
     `member_le_maxKey`; mirror with `reaches_min` / `minKey_le_member`.
   - *Total* (`nonexistence_sound_tree_total`): case split over the four
     neighbor shapes (no-neighbor proofs never verify,
     `verifyNonExistence_none`). Instantiated:
     `nonexistence_sound_tree_tendermint{,_total}` — side conditions by
     `decide`, leaf injectivity proved, `FixedHash` the only assumption.
   The *abstract-root* `nonexistence_sound` (NonExistSound.lean) remains the
   one deliberate `sorry`: an opaque root carries no tree structure, so
   nothing connects byte-order to tree position (exactly F3/F4) — adjacency is
   inherently a statement about a real tree, making the honest-root model the
   natural domain, not a shortcut.
2a. **Theorems A and B for the SMT/JMT spec — DONE
   (`lean/Ics23/SmtTree.lean`, `lean/Ics23/SmtNonExist.lean`).** The `MTree`
   model cannot represent sparse trees (empty subtrees hash to the constant
   `empty_child` placeholder, not a hash image), and SMT's 32-byte
   `emptyChild` *is* digest-sized, so the F4/F5 bridge above does not apply.
   `SMTree` adds an `.empty` constructor with `rootHashS .empty = empty_child`,
   and one new explicitly-stated assumption, `EmptyChildFree` (no exhibited
   SHA-256 preimage of the all-zero placeholder — the standard sparse-merkle
   assumption; without it a prover knowing `H(x) = 0^32` could graft a fake
   empty subtree).
   - *Theorem A* (`membership_sound_smt`): `reachesS` reuses `split_pins`
     byte-level; a fold can never terminate inside an empty subtree because
     fold values are hash images (`applyPath_ne_emptyChild`).
   - *Theorem B* (`nonexistence_sound_smt_total`, with two-sided and
     one-sided forms): the placeholder arm of `ensure_right_most` /
     `ensure_left_most` is *embraced* instead of excluded —
     `ensureRightMost_step_smt` shows each step's suffix is `[]` (genuine
     right step) or exactly `emptyChild`, which pins the skipped sibling to a
     genuinely empty subtree (`rootHashS_ec_empty`); mirrored on the prefix
     side for left-most. `reaches_maxS`/`reaches_minS` land on the rightmost/
     leftmost *nonempty* leaf (Option-valued `maxKeyS`/`minKeyS`), and
     `neighbor_divergence_smt` + `node_gap_no_memberS` close the gap argument.
     Key order is on **comparison keys** (`keyForComparison` = SHA-256 of the
     raw key for `smt_spec`, honoring `prehash_key_before_comparison`):
     `SortedTreeS` is sortedness of the hashed key space, exactly how a real
     SMT/JMT arranges its leaves. Side conditions by `decide`; remaining
     hypotheses: `FixedHash`, `EmptyChildFree`, sortedness.
2b. **Theorems A and B for the IAVL spec — DONE
   (`lean/Ics23/IavlTree.lean`, `lean/Ics23/IavlNonExist.lean`).** With this,
   honest-root Theorems A and B hold for **all three shipped specs**. IAVL
   breaks the Tendermint shape differently: variable 4–12 byte inner prefixes
   (height/size/version varints) violate `min = max`, and the 33-byte amino
   child slots (`0x20` length byte + 32-byte digest) make digest length ≠
   `child_size`. An IAVL node is the existing `MTree` shape with a 1-byte
   `mid` (`WFTreeI`), and the pinning argument is replaced by
   `split_pins_var`: the verifier's `suffix % child_size = 0` check together
   with `max_prefix_length < child_size` (12 < 33 — the `max < min +
   child_size` well-formedness family doing its job) pins the suffix length
   to `0` or `child_size`, each resolving the whole split by plain append
   injectivity, with *no* prefix-length pinning. The order machinery and
   padding bridges reuse TreeNonExist.lean verbatim (IAVL's `emptyChild = []`
   has length 0 ≠ 33, so the placeholder arm is unreachable as in
   Tendermint). `membership_sound_iavl`, `nonexistence_sound_iavl_total`:
   side conditions by `decide`, `FixedHash` the only remaining assumption.
3. **Transcribe Zellic findings** into Properties / corpus — **DONE.** All eight
   recorded in `Corpus.lean` with machine-checked witnesses where decidable; see
   the "Zellic finding map" above.
4. **Batch/compressed** verification — model + decide whether in proof scope
   (RFC open question 3). (Zellic 3.7 concerns this; removed upstream.)

## Open items (cross-phase)

- Phase 2a differential oracle — **DONE for the shared vectors.**
  `rust/examples/lean_testdata.rs` decodes every `testdata/{iavl,tendermint,
  smt}/` vector's protobuf `CommitmentProof` and generates
  `lean/Ics23/TestVectors.lean`: the decoded proofs as Lean literals plus a
  `native_decide` acceptance check per vector mirroring the Rust call, so
  `lake build` re-verifies all 18 shared vectors against the model with real
  SHA-256. CI regenerates the file and fails on drift.
  **First catch:** on its first run the oracle rejected all six IAVL vectors —
  the model's `Sha256.pad` used truncating `Nat` subtraction and silently
  dropped the spill-over padding block for message lengths `≡ 56..62 (mod
  64)`. The three embedded validation vectors and every Tendermint/SMT
  preimage miss that window; IAVL's 57-byte leaf preimages land in it. Fixed
  with boundary regressions in `Sha256.lean`. (Scope note: `Sha256.lean` is
  the *executable* layer only — the soundness theorems are abstract over
  `HashFn` and were never affected.)
  Possible extension: also run the negative/malformed vectors
  (`TestCheckAgainstSpecData.json` etc.) through the model.
- Phase 3 Kani harnesses for Rust panic/overflow safety.
