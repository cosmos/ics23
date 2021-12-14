// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {LeafOp, CompressedBatchProof, ExistenceProof, NonExistenceProof, BatchEntry, BatchProof, ProofSpec, InnerOp, InnerSpec, CommitmentProof} from "./proofs.sol";
import {Ops} from "./ics23Ops.sol";
import {SafeCast} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/SafeCast.sol";
import {BytesLib} from "GNSPS/solidity-bytes-utils@0.8.0/contracts/BytesLib.sol";
import {Compress} from "./ics23Compress.sol";
import {Ops} from "./ics23Ops.sol";

library Proof{
    bytes constant empty = new bytes(0);

    enum VerifyExistenceError{
        None,
        KeyNotMatching,
        ValueNotMatching,
        CheckSpec,
        CalculateRoot,
        RootNotMatching
    }
    // ExistenceProof
    function verify(ExistenceProof.Data memory proof, ProofSpec.Data memory spec, bytes memory commitmentRoot,bytes memory key, bytes memory value) internal pure returns(VerifyExistenceError) {
        //require(BytesLib.equal(proof.key, key)); // dev: Provided key doesn't match proof
        bool keyMatch = BytesLib.equal(proof.key, key);
        if (keyMatch == false) return VerifyExistenceError.KeyNotMatching;
        //require(BytesLib.equal(proof.value, value)); // dev: Provided value doesn't match proof
        bool valueMatch = BytesLib.equal(proof.value, value);
        if (valueMatch == false) return VerifyExistenceError.ValueNotMatching;
        CheckAgainstSpecError cCode =  checkAgainstSpec(proof, spec);
        if (cCode != CheckAgainstSpecError.None) return VerifyExistenceError.CheckSpec;
        (bytes memory root, CalculateRootError rCode) = calculateRoot(proof);
        if (rCode != CalculateRootError.None) return VerifyExistenceError.CalculateRoot;
        //require(BytesLib.equal(root, commitmentRoot)); // dev: Calculcated root doesn't match provided root
        bool rootMatch = BytesLib.equal(root, commitmentRoot);
        if (rootMatch == false) return VerifyExistenceError.RootNotMatching;

        return VerifyExistenceError.None;
    }

    enum CalculateRootError {
        None,
        LeafNil,
        LeafOp,
        PathOp,
        BatchEntriesLength,
        BatchEntryEmpty,
        EmptyProof
    }
    function calculateRoot(ExistenceProof.Data memory proof) internal pure returns(bytes memory, CalculateRootError) {
        //require(LeafOp.isNil(proof.leaf) == false); // dev: Existence Proof needs defined LeafOp
        if (LeafOp.isNil(proof.leaf)) return (empty, CalculateRootError.LeafNil);
        (bytes memory root, Ops.ApplyLeafOpError lCode) = Ops.applyOp(proof.leaf, proof.key, proof.value);
        if (lCode != Ops.ApplyLeafOpError.None) return (empty, CalculateRootError.LeafOp);
        for (uint i = 0; i < proof.path.length; i++) {
            Ops.ApplyInnerOpError iCode;
            (root, iCode) = Ops.applyOp(proof.path[i], root);
            if (iCode != Ops.ApplyInnerOpError.None) return (empty, CalculateRootError.PathOp);
        }

        return (root, CalculateRootError.None);
    }

    enum CheckAgainstSpecError{
        None,
        EmptyLeaf,
        OpsCheckAgainstSpec,
        InnerOpsDepthTooShort,
        InnerOpsDepthTooLong
    }
    function checkAgainstSpec(ExistenceProof.Data memory proof, ProofSpec.Data memory spec) internal pure returns(CheckAgainstSpecError) {
        // LeafOp.isNil does not work
        //require(LeafOp._empty(proof.leaf) == false); // dev: Existence Proof needs defined LeafOp
        if (LeafOp._empty(proof.leaf)) return CheckAgainstSpecError.EmptyLeaf;
        Ops.CheckAgainstSpecError cCode = Ops.checkAgainstSpec(proof.leaf, spec);
        if (cCode != Ops.CheckAgainstSpecError.None) return CheckAgainstSpecError.OpsCheckAgainstSpec;
        if (spec.min_depth > 0) {
            bool innerOpsDepthTooShort = proof.path.length < SafeCast.toUint256(int256(spec.min_depth));
            //require(innerOpsDepthTooShort == false); // dev: InnerOps depth too short
            if (innerOpsDepthTooShort) return CheckAgainstSpecError.InnerOpsDepthTooShort;
        }
        if (spec.max_depth > 0) {
            bool innerOpsDepthTooLong = proof.path.length > SafeCast.toUint256(int256(spec.max_depth));
            //require(innerOpsDepthTooLong == false); // dev: InnerOps depth too long
            if (innerOpsDepthTooLong) return CheckAgainstSpecError.InnerOpsDepthTooLong;
        }
        for(uint i = 0; i < proof.path.length; i++) {
            Ops.CheckAgainstSpecError cCode = Ops.checkAgainstSpec(proof.path[i], spec);
            if (cCode != Ops.CheckAgainstSpecError.None) return CheckAgainstSpecError.OpsCheckAgainstSpec;
        }
    }

    enum VerifyNonExistenceError {
        None,
        VerifyLeft,
        VerifyRight,
        LeftAndRightKeyEmpty,
        RightKeyRange,
        LeftKeyRange,
        RightProofLeftMost,
        LeftProofRightMost,
        IsLeftNeighbor
    }
    // NonExistenceProof
    function verify(NonExistenceProof.Data memory proof, ProofSpec.Data memory spec, bytes memory commitmentRoot,bytes memory key) internal pure returns(VerifyNonExistenceError) {
        bytes memory leftKey;
        bytes memory rightKey;
        // ExistenceProof.isNil does not work
        if (ExistenceProof._empty(proof.left) == false) {
            VerifyExistenceError eCode = verify(proof.left, spec, commitmentRoot, proof.left.key, proof.left.value);
            if (eCode != VerifyExistenceError.None) return VerifyNonExistenceError.VerifyLeft;

            leftKey = proof.left.key;
        }
        if (ExistenceProof._empty(proof.right) == false) {
            VerifyExistenceError eCode = verify(proof.right, spec, commitmentRoot, proof.right.key, proof.right.value);
            if (eCode != VerifyExistenceError.None) return VerifyNonExistenceError.VerifyRight;

            rightKey = proof.right.key;
        }
        // If both proofs are missing, this is not a valid proof
        //require(leftKey.length > 0 || rightKey.length > 0); // dev: both left and right proofs missing
        if (leftKey.length == 0 && rightKey.length == 0) return VerifyNonExistenceError.LeftAndRightKeyEmpty;
        // Ensure in valid range
        if (rightKey.length > 0 && Ops.compare(key, rightKey) >= 0) {
            //require(Ops.compare(key, rightKey) < 0); // dev: key is not left of right proof
            return VerifyNonExistenceError.RightKeyRange;
        }
        if (leftKey.length > 0 && Ops.compare(key, leftKey) <= 0) {
            //require(Ops.compare(key, leftKey) > 0); // dev: key is not right of left proof
            return VerifyNonExistenceError.LeftKeyRange;
        }
        if (leftKey.length == 0) {
            //require(isLeftMost(spec.inner_spec, proof.right.path)); // dev: left proof missing, right proof must be left-most
            if(isLeftMost(spec.inner_spec, proof.right.path) == false) return VerifyNonExistenceError.RightProofLeftMost;
        } else if (rightKey.length == 0) {
            //require(isRightMost(spec.inner_spec, proof.left.path)); // dev: isRightMost: right proof missing, left proof must be right-most
            if (isRightMost(spec.inner_spec, proof.left.path) == false) return VerifyNonExistenceError.LeftProofRightMost;
        } else {
            //require(isLeftNeighbor(spec.inner_spec, proof.left.path, proof.right.path)); // dev: isLeftNeighbor: right proof missing, left proof must be right-most
            bool isLeftNeigh = isLeftNeighbor(spec.inner_spec, proof.left.path, proof.right.path);
            if (isLeftNeigh == false) return VerifyNonExistenceError.IsLeftNeighbor;
        }

        return VerifyNonExistenceError.None;
    }

    function calculateRoot(NonExistenceProof.Data memory proof) internal pure returns(bytes memory, CalculateRootError) {
        if (ExistenceProof._empty(proof.left) == false) {
            return calculateRoot(proof.left);
        }
        if (ExistenceProof._empty(proof.right) == false) {
            return calculateRoot(proof.right);
        }
        //revert(); // dev: Nonexistence proof has empty Left and Right proof
        return (empty, CalculateRootError.EmptyProof);
    }

    // commitment proof
    function calculateRoot(CommitmentProof.Data memory proof) internal pure returns(bytes memory, CalculateRootError) {
        if (ExistenceProof._empty(proof.exist) == false) {
            return calculateRoot(proof.exist);
        }
        if (NonExistenceProof._empty(proof.nonexist) == false) {
            return calculateRoot(proof.nonexist);
        }
        if (BatchProof._empty(proof.batch) == false) {
            //require(proof.batch.entries.length > 0); // dev: batch proof has no entry
            if (proof.batch.entries.length == 0) return (empty, CalculateRootError.BatchEntriesLength);
            //require(BatchEntry._empty(proof.batch.entries[0]) == false); // dev: batch proof has empty entry
            if (BatchEntry._empty(proof.batch.entries[0])) return (empty, CalculateRootError.BatchEntryEmpty);
            if (ExistenceProof._empty(proof.batch.entries[0].exist) == false) {
                return calculateRoot(proof.batch.entries[0].exist);
            }
            if (NonExistenceProof._empty(proof.batch.entries[0].nonexist) == false) {
                return calculateRoot(proof.batch.entries[0].nonexist);
            }
        }
        if (CompressedBatchProof._empty(proof.compressed) == false) {
            return calculateRoot(Compress.decompress(proof));
        }
        //revert(); // dev: calculateRoot(CommitmentProof) empty proof
        return (empty, CalculateRootError.EmptyProof);
    }


    // private
    function isLeftMost(InnerSpec.Data memory spec, InnerOp.Data[] memory path) private pure returns(bool) {
        (uint minPrefix, uint maxPrefix, uint suffix, GetPaddingError gCode) = getPadding(spec, 0);
        if (gCode != GetPaddingError.None) return false;
        for (uint i = 0; i < path.length; i++) {
            if (hasPadding(path[i], minPrefix, maxPrefix, suffix) == false){
                return false;
            }
        }
        return true;
    }

    function isRightMost(InnerSpec.Data memory spec, InnerOp.Data[] memory path) private pure returns(bool){
        uint last = spec.child_order.length - 1;
        (uint minPrefix, uint maxPrefix, uint suffix, GetPaddingError gCode) = getPadding(spec, last);
        if (gCode != GetPaddingError.None) return false;
        for (uint i = 0; i < path.length; i++) {
            if (hasPadding(path[i], minPrefix, maxPrefix, suffix) == false){
                return false;
            }
        }

        return true;
    }

    function isLeftStep(InnerSpec.Data memory spec, InnerOp.Data memory left, InnerOp.Data memory right) private pure returns(bool){
        (uint leftIdx, OrderFromPaddingError lCode) = orderFromPadding(spec, left);
        if (lCode != OrderFromPaddingError.None) return false;
        (uint rightIdx, OrderFromPaddingError rCode) = orderFromPadding(spec, right);
        if (lCode != OrderFromPaddingError.None) return false;
        if (rCode != OrderFromPaddingError.None) return false;

        return rightIdx == leftIdx + 1;
    }

    function isLeftNeighbor(InnerSpec.Data memory spec, InnerOp.Data[] memory left, InnerOp.Data[] memory right) private pure returns(bool) {
        uint leftIdx = left.length - 1;
        uint rightIdx = right.length - 1;
        while (leftIdx >= 0 && rightIdx >= 0) {
            if (BytesLib.equal(left[leftIdx].prefix, right[rightIdx].prefix) &&
                BytesLib.equal(left[leftIdx].suffix, right[rightIdx].suffix)) {
                leftIdx -= 1;
            rightIdx -= 1;
            continue;
            }
            break;
        }
        if (isLeftStep(spec, left[leftIdx], right[rightIdx]) == false) {
            return false;
        }
        // slicing does not work for ``memory`` types
        if (isRightMost(spec, sliceInnerOps(left, 0, leftIdx)) == false){
            return false;
        }
        if (isLeftMost(spec, sliceInnerOps(right, 0, rightIdx)) == false) {
            return false;
        }
        return true;
    }

    enum OrderFromPaddingError {
        None,
        NotFound,
        GetPadding
    }
    function orderFromPadding(InnerSpec.Data memory spec, InnerOp.Data memory op) private pure returns(uint, OrderFromPaddingError) {
        uint256 maxBranch = spec.child_order.length;
        for(uint branch = 0; branch < maxBranch; branch++) {
            (uint minp, uint maxp, uint suffix, GetPaddingError gCode) = getPadding(spec, branch);
            if (gCode != GetPaddingError.None) return (0, OrderFromPaddingError.GetPadding);
            if (hasPadding(op, minp, maxp, suffix) == true) return (branch, OrderFromPaddingError.None);
        }
        //revert(); // dev: Cannot find any valid spacing for this node
        return (0, OrderFromPaddingError.NotFound);
    }

    enum GetPaddingError {
        None,
        GetPosition
    }
    function getPadding(InnerSpec.Data memory spec, uint branch) private pure returns(uint minPrefix, uint maxPrefix, uint suffix, GetPaddingError) {
        uint uChildSize = SafeCast.toUint256(spec.child_size);
        (uint idx, GetPositionError gCode) = getPosition(spec.child_order, branch);
        if (gCode != GetPositionError.None) return (0, 0, 0, GetPaddingError.GetPosition);
        uint prefix = idx * uChildSize;
        minPrefix = prefix + SafeCast.toUint256(spec.min_prefix_length);
        maxPrefix = prefix + SafeCast.toUint256(spec.max_prefix_length);
        suffix = (spec.child_order.length - 1 - idx) * uChildSize;

        return (minPrefix, maxPrefix, suffix, GetPaddingError.None);
    }

    enum GetPositionError {
        None,
        BranchLength,
        NoFound
    }
    function getPosition(int32[] memory order, uint branch) private pure returns(uint, GetPositionError) {
        //require(branch < order.length); // dev: invalid branch
        if (branch >= order.length) return (0, GetPositionError.BranchLength);
        for (uint i = 0; i < order.length; i++) {
            if (SafeCast.toUint256(order[i]) == branch) return (i, GetPositionError.None);
        }
        //revert(); // dev: branch not found in order
        return (0, GetPositionError.NoFound);
    }

    function hasPadding(InnerOp.Data memory op, uint minPrefix, uint maxPrefix, uint suffix) private pure returns(bool) {
        if (op.prefix.length < minPrefix) return false;
        if (op.prefix.length > maxPrefix) return false;
        return op.suffix.length == suffix;
    }

    function sliceInnerOps(InnerOp.Data[] memory array, uint start, uint end) private pure returns(InnerOp.Data[] memory) {
        InnerOp.Data[] memory slice = new InnerOp.Data[](end-start);
        for (uint i = start; i < end; i++) {
            slice[i] = array[i];
        }
        return slice;
    }
}
