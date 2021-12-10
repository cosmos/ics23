// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {LeafOp, CompressedBatchProof, ExistenceProof, NonExistenceProof, BatchEntry, BatchProof, ProofSpec, InnerOp, InnerSpec, CommitmentProof} from "./proofs.sol";
import {Ops} from "./ics23Ops.sol";
import {SafeCast} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/SafeCast.sol";
import {BytesLib} from "GNSPS/solidity-bytes-utils@0.8.0/contracts/BytesLib.sol";
import {Compress} from "./ics23Compress.sol";
import {Ops} from "./ics23Ops.sol";

library Proof{
    // ExistenceProof
    function verify(ExistenceProof.Data memory proof, ProofSpec.Data memory spec, bytes memory commitmentRoot,bytes memory key, bytes memory value) internal pure {
        require(BytesLib.equal(proof.key, key)); // dev: Provided key doesn't match proof
        require(BytesLib.equal(proof.value, value)); // dev: Provided value doesn't match proof
        checkAgainstSpec(proof, spec);
        bytes memory root = calculateRoot(proof);
        require(BytesLib.equal(root, commitmentRoot)); // dev: Calculcated root doesn't match provided root
    }

    function calculateRoot(ExistenceProof.Data memory proof) internal pure returns(bytes memory) {
        require(LeafOp.isNil(proof.leaf) == false); // dev: Existence Proof needs defined LeafOp
        bytes memory root = Ops.applyOp(proof.leaf, proof.key, proof.value);
        for (uint i = 0; i < proof.path.length; i++) {
            root = Ops.applyOp(proof.path[i], root);
        }
        return root;
    }

    function checkAgainstSpec(ExistenceProof.Data memory proof, ProofSpec.Data memory spec) internal pure {
        // LeafOp.isNil does not work
        require(LeafOp._empty(proof.leaf) == false); // dev: Existence Proof needs defined LeafOp
        Ops.checkAgainstSpec(proof.leaf, spec);
        if (spec.min_depth > 0) {
            bool innerOpsDepthTooShort = proof.path.length >= SafeCast.toUint256(int256(spec.min_depth));
            require(innerOpsDepthTooShort); // dev: InnerOps depth too short
        }
        if (spec.max_depth > 0) {
            bool innerOpsDepthTooLong = proof.path.length <= SafeCast.toUint256(int256(spec.max_depth));
            require(innerOpsDepthTooLong); // dev: InnerOps depth too long
        }
        for(uint i = 0; i < proof.path.length; i++) {
            Ops.checkAgainstSpec(proof.path[i], spec);
        }
    }

    // NonExistenceProof
    function verify(NonExistenceProof.Data memory proof, ProofSpec.Data memory spec, bytes memory commitmentRoot,bytes memory key) internal pure {
        bytes memory leftKey;
        bytes memory rightKey;
        // ExistenceProof.isNil does not work
        if (ExistenceProof._empty(proof.left) == false) {
            verify(proof.left, spec, commitmentRoot, proof.left.key, proof.left.value);
            leftKey = proof.left.key;
        }
        if (ExistenceProof._empty(proof.right) == false) {
            verify(proof.right, spec, commitmentRoot, proof.right.key, proof.right.value);
            rightKey = proof.right.key;
        }
        // If both proofs are missing, this is not a valid proof
        require(leftKey.length > 0 || rightKey.length > 0); // dev: both left and right proofs missing
        // Ensure in valid range
        if (rightKey.length > 0) {
            require(Ops.compare(key, rightKey) < 0); // dev: key is not left of right proof
        }
        if (leftKey.length > 0) {
            require(Ops.compare(key, leftKey) > 0); // dev: key is not right of left proof
        }
        if (leftKey.length == 0) {
            require(isLeftMost(spec.inner_spec, proof.right.path)); // dev: left proof missing, right proof must be left-most
        } else if (rightKey.length == 0) {
            require(isRightMost(spec.inner_spec, proof.left.path)); // dev: isRightMost: right proof missing, left proof must be right-most
        } else {
            require(isLeftNeighbor(spec.inner_spec, proof.left.path, proof.right.path)); // dev: isLeftNeighbor: right proof missing, left proof must be right-most
        }
    }

    function calculateRoot(NonExistenceProof.Data memory proof) internal pure returns(bytes memory ) {
        if (ExistenceProof._empty(proof.left) == false) {
            return calculateRoot(proof.left);
        }
        if (ExistenceProof._empty(proof.right) == false) {
            return calculateRoot(proof.right);
        }
        revert(); // dev: Nonexistence proof has empty Left and Right proof
    }

    // commitment proof
    function calculateRoot(CommitmentProof.Data memory proof) internal pure returns(bytes memory) {
        if (ExistenceProof._empty(proof.exist) == false) {
            return calculateRoot(proof.exist);
        }
        if (NonExistenceProof._empty(proof.nonexist) == false) {
            return calculateRoot(proof.nonexist);
        }
        if (BatchProof._empty(proof.batch) == false) {
            require(proof.batch.entries.length > 0); // dev: batch proof has no entry
            require(BatchEntry._empty(proof.batch.entries[0]) == false); // dev: batch proof has empty entry
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
        revert(); // dev: calculateRoot(CommitmentProof) empty proof
    }


    // private
    function isLeftMost(InnerSpec.Data memory spec, InnerOp.Data[] memory path) private pure returns(bool) {
        (uint minPrefix, uint maxPrefix, uint suffix) = getPadding(spec, 0);
        for (uint i = 0; i < path.length; i++) {
            if (hasPadding(path[i], minPrefix, maxPrefix, suffix) == false){
                return false;
            }
        }
        return true;
    }

    function isRightMost(InnerSpec.Data memory spec, InnerOp.Data[] memory path) private pure returns(bool){
        uint last = spec.child_order.length - 1;
        (uint minPrefix, uint maxPrefix, uint suffix) = getPadding(spec, last);
        for (uint i = 0; i < path.length; i++) {
            if (hasPadding(path[i], minPrefix, maxPrefix, suffix) == false){
                return false;
            }
        }
        return true;
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

    function isLeftStep(InnerSpec.Data memory spec, InnerOp.Data memory left, InnerOp.Data memory right) private pure returns(bool){
        uint leftIdx = orderFromPadding(spec, left);
        uint rightIdx = orderFromPadding(spec, right);
        return rightIdx == leftIdx + 1;
    }

    function orderFromPadding(InnerSpec.Data memory spec, InnerOp.Data memory op) private pure returns(uint) {
        uint256 maxBranch = spec.child_order.length;
        for(uint branch = 0; branch < maxBranch; branch++) {
            (uint minp, uint maxp, uint suffix) = getPadding(spec, branch);
            if (hasPadding(op, minp, maxp, suffix) == true) {
                return branch;
            }
        }
        revert(); // dev: Cannot find any valid spacing for this node
    }

    function getPadding(InnerSpec.Data memory spec, uint branch) private pure returns(uint minPrefix, uint maxPrefix, uint suffix) {
        uint uChildSize = SafeCast.toUint256(spec.child_size);
        uint idx = getPosition(spec.child_order, branch);
        uint prefix = idx * uChildSize;
        minPrefix = prefix + SafeCast.toUint256(spec.min_prefix_length);
        maxPrefix = prefix + SafeCast.toUint256(spec.max_prefix_length);
        suffix = (spec.child_order.length - 1 - idx) * uChildSize;
    }

    function getPosition(int32[] memory order, uint branch) private pure returns(uint) {
        require(branch < order.length); // dev: invalid branch
        for (uint i = 0; i < order.length; i++) {
            if (SafeCast.toUint256(order[i]) == branch) {
                return i;
            }
        }
        revert(); // dev: branch not found in order
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




contract Proof_UnitTest {
    function calculateExistenceProofRoot(ExistenceProof.Data memory proof) public pure returns(bytes memory) {
        return Proof.calculateRoot(proof);
    }
    function checkAgainstSpec(ExistenceProof.Data memory proof, ProofSpec.Data memory spec) public pure {
        return Proof.checkAgainstSpec(proof, spec);
    }
    function calculateCommitmentProofRoot(CommitmentProof.Data memory proof) public pure returns(bytes memory) {
        return Proof.calculateRoot(proof);
    }
    function protobufDecodeCommitmentProof(bytes memory msg) public pure returns(CommitmentProof.Data memory) {
        CommitmentProof.Data memory res = CommitmentProof.decode(msg);
        return res;
    }
}
