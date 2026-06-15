//! Kani proof harnesses for Property D (implementation safety): overflow and
//! index-bounds safety of the arithmetic in the verifier internals — the
//! slice-indexing in the non-existence padding logic is the Dragonberry-class
//! surface.
//!
//! Run with `cargo kani` (the `kani` cfg is set only under Kani, so these are
//! invisible to normal builds). See `docs/rfc/001-formal-verification.md`.
//!
//! These harnesses are intentionally over the *arithmetic*, mirrored from the
//! verifier, rather than the full functions: the `anyhow`/`Result` error paths
//! pull in formatting machinery that explodes CBMC's formula. Panic-freedom of
//! the `Result`-returning entry points (`do_length`, `proto_len`) is left to a
//! later pass that stubs formatting.

/// `do_length` is panic-free for `NoPrefix` and the fixed-width ops: the
/// `data.len() as u32 / u64` casts and `to_be/le_bytes` extends cannot panic.
/// (`VarProto` and `Require*` route through `anyhow` construction, which blows up
/// CBMC's formula; they are excluded here.)
#[kani::proof]
fn do_length_fixed_no_panic() {
    use crate::ics23::LengthOp;
    use crate::ops::do_length;
    let choice: u8 = kani::any();
    kani::assume(choice < 5);
    let op = match choice {
        0 => LengthOp::NoPrefix,
        1 => LengthOp::Fixed32Big,
        2 => LengthOp::Fixed32Little,
        3 => LengthOp::Fixed64Big,
        _ => LengthOp::Fixed64Little,
    };
    let len: usize = kani::any();
    kani::assume(len <= 4);
    let data = alloc::vec![0u8; len];
    let _ = do_length(op, &data);
}

/// `ensure_inner`'s prefix bound `max_prefix_length + (child_order.len()-1) *
/// child_size` is overflow-free in `i32` for well-formed bounds (cf. the
/// IAVL/Tendermint/SMT specs). A malformed spec with an enormous `child_size`
/// could overflow it — tracked as a finding; this pins the safe precondition.
#[kani::proof]
fn ensure_inner_prefix_bound_no_overflow() {
    let child_order_len: usize = kani::any();
    let child_size: i32 = kani::any();
    let max_prefix_length: i32 = kani::any();
    kani::assume((2..=16).contains(&child_order_len));
    kani::assume((1..=128).contains(&child_size));
    kani::assume((0..=4096).contains(&max_prefix_length));

    let max_left_child_bytes = (child_order_len as i32 - 1)
        .checked_mul(child_size)
        .expect("child_order/child_size product fits i32");
    let _bound = max_prefix_length
        .checked_add(max_left_child_bytes)
        .expect("inner prefix bound fits i32");
}

/// `get_padding` computes `prefix = idx * child_size` and
/// `suffix = child_size * (child_order.len() - 1 - idx)`. Both are overflow-free
/// in `i32` for well-formed bounds.
#[kani::proof]
fn get_padding_no_overflow() {
    let n: usize = kani::any();
    let idx: i32 = kani::any();
    let child_size: i32 = kani::any();
    kani::assume((2..=16).contains(&n));
    kani::assume((0..=(n as i32 - 1)).contains(&idx));
    kani::assume((1..=128).contains(&child_size));

    let _prefix = idx.checked_mul(child_size).expect("prefix offset fits i32");
    let _suffix = child_size
        .checked_mul(n as i32 - 1 - idx)
        .expect("suffix size fits i32");
}

/// The slice accesses in `left_branches_are_empty`,
/// `op.prefix[from .. from + child_size]` with `from = actual_prefix + i*child_size`
/// and `actual_prefix = prefix.len() - left_branches*child_size` (via
/// `checked_sub`), are always in bounds — so the padding scan never panics on an
/// out-of-range slice. This is the Dragonberry-relevant safety property.
#[kani::proof]
#[kani::unwind(17)]
fn left_branches_slice_in_bounds() {
    let prefix_len: usize = kani::any();
    let left_branches: usize = kani::any();
    let child_size: usize = kani::any();
    kani::assume((1..=64).contains(&child_size));
    kani::assume((1..=16).contains(&left_branches));

    if let Some(actual_prefix) = prefix_len.checked_sub(left_branches * child_size) {
        let mut i = 0usize;
        while i < left_branches {
            let from = actual_prefix + i * child_size;
            // models `op.prefix[from .. from + child_size]`
            assert!(
                from + child_size <= prefix_len,
                "left-branch slice in bounds"
            );
            i += 1;
        }
    }
}

/// `right_branches_are_empty` guards `op.suffix.len() == child_size` (a *single*
/// child) but then reads `op.suffix[i*child_size .. (i+1)*child_size]` for
/// `i in 0..right_branches`. For binary specs (`child_order.len() == 2`,
/// `right_branches <= 1`) the only access is `i = 0`, which is in bounds — this
/// covers all currently supported stores (IAVL/Tendermint/SMT).
///
/// NOTE (finding): for a spec with `child_order.len() > 2`, `right_branches`
/// can exceed 1 while the guard still admits a `child_size`-long suffix, so the
/// `i = 1` access would be out of bounds (a panic). The left/right empty-branch
/// checks are asymmetric: the left side sizes the prefix by
/// `left_branches * child_size`, the right side only checks a single
/// `child_size`. No supported store is ternary, but this is latent. See
/// `docs/verification/properties.md`.
#[kani::proof]
#[kani::unwind(2)]
fn right_branches_slice_in_bounds_binary() {
    let child_size: usize = kani::any();
    let suffix_len: usize = kani::any();
    let right_branches: usize = kani::any();
    kani::assume((1..=64).contains(&child_size));
    kani::assume(right_branches == 1); // binary spec: at most one right sibling
    kani::assume(suffix_len == child_size); // the `suffix.len() == child_size` guard

    let mut i = 0usize;
    while i < right_branches {
        let from = i * child_size;
        // models `op.suffix[from .. from + child_size]`
        assert!(
            from + child_size <= suffix_len,
            "right-branch slice in bounds (binary)"
        );
        i += 1;
    }
}

/// `decompress_exist` resolves a compressed path index with
/// `lookup.get(x as usize)` for an attacker-controlled `i32` x. This is
/// panic-free for any x and any lookup table: a negative or out-of-range index
/// wraps under `as usize` and `get` returns `None` (the caller maps that to an
/// empty path), so the compressed-batch decode cannot panic on the index.
#[kani::proof]
fn decompress_index_no_panic() {
    let len: usize = kani::any();
    kani::assume(len <= 4);
    let lookup: alloc::vec::Vec<u8> = alloc::vec![0u8; len];
    let x: i32 = kani::any();
    let _ = lookup.get(x as usize);
}
