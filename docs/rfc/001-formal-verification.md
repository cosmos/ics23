# RFC 001: Formal Verification of ICS23

- Status: Draft
- Author: Zaki Manian
- Created: 2026-06-10

## Abstract

This RFC proposes a program to formally verify the ICS23 proof verification
algorithm. The proposal centers on a mechanized model of the verifier with
machine-checked soundness proofs, connected to the shipped Go and Rust
implementations through differential testing, with optional code-level
refinement proofs as a stretch goal. The total scope is roughly 4-6 months
for one verification engineer plus a domain reviewer.

## Motivation

The ICS23 verifier is the security boundary of IBC state verification. It
runs inside light clients (ibc-go, ibc-rs, wasm light clients), and in that
setting the prover is the adversary: every byte of a `CommitmentProof` is
attacker-controlled. A soundness bug in the verifier allows an attacker to
convince a light client that an arbitrary key/value pair exists (or does not
exist) in a counterparty chain's state, which translates directly into theft
over IBC. The October 2022 "Dragonberry" advisory demonstrated that
proof-forgery bugs in the Cosmos proof verification stack are a realistic,
ecosystem-wide threat, not a theoretical one.

The code has been audited (most recently by Zellic; see `docs/audits/`), and
both implementations carry fuzz tests. Audits and fuzzing find bugs but
cannot establish their absence. The verification surface is unusually well
suited to formal methods:

- The hand-written core is small: about 1,100 lines of Rust
  (`rust/src/verify.rs`, `rust/src/ops.rs`, `rust/src/compress.rs`) and
  about 950 lines of Go (`go/proof.go`, `go/ops.go`, `go/ics23.go`).
- The logic is purely functional: no concurrency, no I/O, no system state.
- The intended security property has a crisp cryptographic statement
  (commitment binding up to hash collisions).

At the same time, the algorithm is genuinely subtle. Soundness depends on a
web of interacting side conditions: leaf/inner domain separation, prefix
length windows, child size constraints, lexicographic neighbor checks. The
non-existence logic in particular (`ensure_left_most`, `ensure_left_neighbor`,
`left_branches_are_empty`) is where historical bugs have concentrated. These
conditions are exactly the kind of thing a theorem prover is good at and a
human reviewer is bad at.

## Scope

### In scope

- `verify_existence` / `ExistenceProof.Verify`: spec conformance checks and
  root calculation.
- `verify_non_existence` / `NonExistenceProof.Verify`: neighbor ordering and
  edge-of-tree logic.
- The `ProofSpec` well-formedness conditions under which the above are sound,
  and concrete certification of the shipped `IavlSpec`, `TendermintSpec`, and
  `SmtSpec`.
- Batch and compressed-batch verification (`rust/src/compress.rs`,
  `go/proof.go` batch handling), treated as a second pass after the core,
  since they are thin wrappers plus index handling.
- Implementation-level safety of the Rust crate: panic freedom, integer
  cast/overflow safety, termination, allocation bounds on adversarial input.
- Go/Rust acceptance equivalence (consensus criticality: the two
  implementations must accept exactly the same proofs).

### Out of scope

- Prover-side code (IAVL, CometBFT merkle, JMT). Completeness against real
  provers is handled by differential testing, not proof, because the provers
  live in other repositories.
- The cryptographic hash functions themselves. Hashes are modeled as
  abstract functions; all soundness theorems are stated as constructive
  reductions to collision finding (see below), so no axioms about SHA-256
  et al. are introduced.
- The legacy `calculate_existence_root` path that runs without a
  `ProofSpec` (`rust/src/verify.rs`, `CalculateRoot` in Go). It skips the
  spec-dependent child-size check and should be documented as carrying
  weaker guarantees rather than verified.

## Properties

The properties below are stated informally; making them precise is Phase 0/1
work. Throughout, "well-formed spec" refers to the decidable predicate of
Theorem C.

**Theorem A (existence binding).** For a well-formed `ProofSpec` S and root
hash r: from any two accepted existence proofs under (S, r) that disagree on
their (key, value) pair, one can extract an explicit collision of the hash
function used by S.

This is a constructive reduction: the proof term is an algorithm that
produces the two colliding preimages. The supporting lemmas mirror checks
the code performs operationally:

1. *Leaf encoding injectivity.* The `LengthOp`-prefixed encoding of
   (prehash(key), prehash(value)) is injective for the length ops permitted
   by well-formed specs (notably `VAR_PROTO`), so distinct pairs produce
   distinct leaf preimages.
2. *Leaf/inner domain separation.* No acceptable inner-node preimage shares
   a prefix with the leaf prefix (enforced at `rust/src/verify.rs`
   `ensure_inner`: `!has_prefix(leaf_spec.prefix, inner.prefix)`), so a leaf
   hash cannot be reinterpreted as an inner hash or vice versa.
3. *Positional unambiguity.* The constraints relating `min_prefix_length`,
   `max_prefix_length`, `child_size`, and suffix length (including
   `max_prefix_length < min_prefix_length + child_size` and
   `suffix.len() % child_size == 0`) guarantee that an accepted `InnerOp`
   determines a unique child position, so two proofs cannot place the same
   child hash at different positions in the same node without a collision.

**Theorem B (non-existence soundness).** For a well-formed spec S and root
r: from an accepted non-existence proof for key k and an accepted existence
proof for k under (S, r), one can extract a hash collision.

This requires formalizing the tree semantics that an `InnerSpec` describes:
which leaves a given root can commit to, what "left-most", "right-most", and
"adjacent" mean under arbitrary `child_order` permutations, and how
`empty_child` participates for sparse merkle trees. When
`prehash_key_before_comparison` is set (JMT), the ordering relation is over
hashed keys, and the theorem statement must reflect that the guarantee is
non-existence of the *hashed* key.

**Theorem C (spec well-formedness certificate).** A decidable predicate
`WellFormed(spec)` capturing exactly the side conditions assumed by Theorems
A and B, together with computations certifying that `IavlSpec`,
`TendermintSpec`, and `SmtSpec` satisfy it.

This artifact has standalone value beyond this repository: it gives chain
and bridge developers a machine-checkable answer to "is this custom
ProofSpec safe to accept?", which today is answered by folklore.

**Property D (implementation safety).** The Rust crate, on arbitrary input:
never panics, performs no lossy or overflowing integer conversion (several
`i32`/`i64`/`usize` casts in `ensure_inner`, `ensure_inner_prefix`, and the
compressed-batch index lookups warrant machine checking), terminates, and
has allocation bounded by input size.

**Property E (cross-implementation equivalence).** The Go and Rust
implementations accept exactly the same (proof, spec, root, key, value)
tuples. A divergence is itself a security incident even if both
implementations are individually sound, because IBC counterparties running
different implementations would disagree about state.

## Approach

The recommended architecture is model-first: prove the algorithm correct
once in a proof assistant, then connect the proven model to both shipped
implementations. This covers Go and Rust with one proof effort and keeps
the proof artifact maintainable independently of implementation churn.

### Phase 0: Property mining (1-2 weeks)

Extract every claimed invariant from the Zellic report, prior audits, and
the Dragonberry post-mortem. Reconstruct historical bugs as concrete
malicious proofs and add them to a regression corpus. The formal model is
required to reject every element of this corpus; this validates that the
model is modeling the right thing before any theorem is attempted.

### Phase 1: Mechanized model and soundness proofs (6-10 weeks)

Write the verifier as executable functions in Lean 4, mirroring
`rust/src/verify.rs` and `rust/src/ops.rs` as closely as practical, then
prove Theorems A, B, and C. Lean 4 is recommended over Rocq or F* for
ecosystem momentum and contributor availability; Rocq is the alternative if
alignment with `coq-of-rust` (Phase 2b) is prioritized. The expected
schedule risk concentrates in Theorem B's tree semantics; Theorem A should
land first and is independently publishable.

### Phase 2: Connecting model to code (4-8 weeks)

**2a (baseline): differential oracle.** Compile the Lean model to an
executable and differentially test it against both the Rust and Go
implementations: structure-aware generation of valid proofs from real IAVL,
CometBFT, and JMT trees, plus mutation-based adversarial inputs, run at
fuzzing scale in CI. This simultaneously discharges Property E by
transitivity and detects drift between the proven model and the shipped
code. It is cheap, immediate, and remains valuable forever.

**2b (stretch): refinement proof.** Extract the actual Rust verification
core to a proof assistant using hax (Rust to F*/Coq, the methodology used
for libcrux) or `coq-of-rust`, and prove it refines the Phase 1 model. The
crate is `no_std`, small, and first-order, which makes it an unusually good
extraction candidate, but budget for friction around `prost`-generated
types and `anyhow` error plumbing (likely requiring a verification feature
gate or a thin shim around the core functions). This adds an estimated 2-3
months and can be decided after Phase 1 results are in.

### Phase 3: Code-level checks (2-4 weeks, parallel with Phase 1/2)

- Kani harnesses on the Rust crate for Property D: panic freedom and cast
  safety on `ensure_inner` / `ensure_inner_prefix` arithmetic and the
  compressed-batch index handling, plus bounded-depth exhaustive checks of
  small proofs.
- `cargo fuzz` targets mirroring `go/fuzz_test.go`, run continuously.
- No deductive verification of the Go implementation (Gobra is the only
  candidate and the return on effort is poor); Go correctness is covered by
  the Phase 2a oracle and native fuzzing.

### Phase 4: CI integration and maintenance (1-2 weeks)

- Lean proofs, Kani harnesses, and the differential fuzzer run in CI.
- A traceability document mapping each function in `verify.rs` / `proof.go`
  to its model counterpart and the lemmas that depend on it.
- A contribution policy: changes touching the verification core require a
  corresponding model update, enforced by CODEOWNERS on the relevant paths.

## Alternatives considered

**Direct deductive verification of the Rust crate (Verus or Creusot).**
Strongest possible code-level result, but Verus requires rewriting the crate
in its dialect (forking the artifact users actually depend on), Creusot
carries significant annotation and toolchain friction, and neither covers
the Go implementation, which remains the most widely deployed. Rejected as
the primary strategy; partially recovered via Phase 2b.

**TLA+/Quint modeling.** Wrong tool: there is no concurrency or temporal
behavior here. The problem is data-structure and cryptographic soundness,
which model checkers over state machines do not address.

**F* end-to-end from day one (libcrux-style).** Viable and proven for
cryptographic Rust, but F* expertise is scarce and long-term maintenance by
the Cosmos community is a real concern. Kept available through the hax
route in Phase 2b.

**More auditing and fuzzing only.** Cheaper, but cannot establish absence
of soundness bugs, and provides no certificate for evaluating new
ProofSpecs (Theorem C), which is where much of the forward-looking risk
lives as new tree types are onboarded.

## Risks

- **Tree semantics formalization (Theorem B)** is the research-shaped
  component: arbitrary `child_order` permutations and `empty_child`
  handling for SMTs have no off-the-shelf formalization. Mitigation:
  sequence Theorem A first; descope Theorem B to the three shipped specs'
  parameter shapes if full generality stalls.
- **Model/code divergence** is the standing threat to any model-first
  approach. Mitigation: the Phase 2a oracle runs continuously in CI, and
  Phase 4 policy couples code changes to model changes.
- **Extraction friction in Phase 2b** around `prost` and `anyhow`.
  Mitigation: treat 2b as optional; restructure the core behind a
  dependency-free internal API only if 2b is funded.
- **Maintainer bandwidth**: a proof artifact nobody can maintain decays
  into documentation. Mitigation: prefer Lean 4 for contributor pool;
  require the traceability document; keep the model small by scoping it to
  the verifier only.

## Deliverables

1. Property specification document with regression corpus (Phase 0).
2. Lean 4 model of the ICS23 verifier with machine-checked Theorems A, B,
   and C, including well-formedness certificates for `IavlSpec`,
   `TendermintSpec`, and `SmtSpec` (Phase 1).
3. Differential oracle harness (model vs. Rust vs. Go) integrated into CI
   (Phase 2a).
4. Kani harnesses and `cargo fuzz` targets for the Rust crate (Phase 3).
5. Traceability document and contribution policy (Phase 4).
6. (Stretch) Refinement proof connecting `ics23-rs` to the model via hax or
   `coq-of-rust` (Phase 2b).

## Timeline and resourcing

Phases 0 through 4 total roughly 4-6 months for one verification engineer
with proof-assistant experience, plus part-time review from an ICS23/IBC
domain expert. Phase 2b adds an estimated 2-3 months and a go/no-go
decision point after Phase 1. The work is well shaped for an external
engagement (candidate teams with relevant track records include Cryspen,
Formal Land, Informal Systems, and Galois) or an Interchain ecosystem grant.

## Open questions

1. Lean 4 vs. Rocq: primarily a question of who maintains the artifact and
   whether Phase 2b via `coq-of-rust` is prioritized.
2. Is Phase 2b (refinement proof) in scope, or does the differential oracle
   provide sufficient assurance for the cost?
3. Should batch/compressed proof verification be in the initial proof scope
   or deferred until the core theorems land?
4. Where should the model live: this repository (keeps coupling tight) or a
   sibling repository (keeps toolchains separate)?

## References

- ICS23 specification: https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments
- Zellic audit report: `docs/audits/ICS-23 - Zellic Audit Report.pdf`
- Dragonberry advisory (Cosmos SDK, October 2022):
  https://forum.cosmos.network/t/ibc-security-advisory-dragonberry/7702
- hax (Rust extraction to F*/Coq): https://github.com/cryspen/hax
- coq-of-rust: https://github.com/formal-land/coq-of-rust
- Kani (Rust model checker): https://github.com/model-checking/kani
- libcrux, as precedent for verified Rust via extraction:
  https://github.com/cryspen/libcrux
