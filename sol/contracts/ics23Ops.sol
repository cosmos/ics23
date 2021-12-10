// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {LeafOp, InnerOp, PROOFS_PROTO_GLOBAL_ENUMS, ProofSpec} from "./proofs.sol";
import {ProtoBufRuntime} from "./ProtoBufRuntime.sol";
import {SafeCast} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/SafeCast.sol";
import {Math} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/Math.sol";
import {BytesLib} from "GNSPS/solidity-bytes-utils@0.8.0/contracts/BytesLib.sol";

library Ops {

    // LeafOp operations
    function applyOp(LeafOp.Data memory leafOp, bytes memory key, bytes memory value) internal pure returns(bytes memory) {
        require(key.length > 0); // dev: Leaf op needs key
        require(value.length > 0); // dev: Leaf op needs value
        bytes memory pKey = prepareLeafData(leafOp.prehash_key, leafOp.length, key);
        bytes memory pValue = prepareLeafData(leafOp.prehash_value, leafOp.length, value);
        bytes memory data = abi.encodePacked(leafOp.prefix, pKey, pValue);
        return doHash(leafOp.hash, data);
    }
    function prepareLeafData(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, PROOFS_PROTO_GLOBAL_ENUMS.LengthOp lenOp, bytes memory data) internal pure returns(bytes memory) {
        bytes memory hdata = doHashOrNoop(hashOp, data);
        return doLengthOp(lenOp, hdata);
    }
    function checkAgainstSpec(LeafOp.Data memory leafOp, ProofSpec.Data memory spec) internal pure {
        require (leafOp.hash == spec.leaf_spec.hash); // dev: checkAgainstSpec for LeafOp - Unexpected HashOp
        require(leafOp.prehash_key == spec.leaf_spec.prehash_key); // dev: checkAgainstSpec for LeafOp - Unexpected PrehashKey
        require(leafOp.prehash_value == spec.leaf_spec.prehash_value); // dev: checkAgainstSpec for LeafOp - Unexpected PrehashValue");
        require(leafOp.length == spec.leaf_spec.length); // dev: checkAgainstSpec for LeafOp - Unexpected lengthOp
        bool hasprefix = hasPrefix(leafOp.prefix, spec.leaf_spec.prefix);
        require(hasprefix); // dev: checkAgainstSpec for LeafOp - Leaf Prefix doesn't start with
    }

    // InnerOp operations
    function applyOp(InnerOp.Data memory innerOp, bytes memory child ) internal pure returns(bytes memory) {
        require(child.length > 0); // dev: Inner op needs child value
        bytes memory preImage = abi.encodePacked(innerOp.prefix, child, innerOp.suffix);
        return doHash(innerOp.hash, preImage);
    }
    function checkAgainstSpec(InnerOp.Data memory innerOp, ProofSpec.Data memory spec) internal pure {
        require(innerOp.hash == spec.inner_spec.hash); // dev: checkAgainstSpec for InnerOp - Unexpected HashOp
        uint256 minPrefixLength = SafeCast.toUint256(spec.inner_spec.min_prefix_length);
        require(innerOp.prefix.length >= minPrefixLength); // dev: InnerOp prefix too short;
        bytes memory leafPrefix = spec.leaf_spec.prefix;
        bool hasprefix = hasPrefix(innerOp.prefix, leafPrefix);
        require(hasprefix == false); // dev: Inner Prefix starts with wrong value
        uint256 childSize = SafeCast.toUint256(spec.inner_spec.child_size);
        uint256 maxLeftChildBytes = (spec.inner_spec.child_order.length - 1) * childSize;
        uint256 maxPrefixLength = SafeCast.toUint256(spec.inner_spec.max_prefix_length);
        require(innerOp.prefix.length <= maxPrefixLength + maxLeftChildBytes); // dev: InnerOp prefix too long
    }

    function doHashOrNoop(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, bytes memory preImage) internal pure returns(bytes memory) {
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.NO_HASH) {
            return preImage;
        }
        return doHash(hashOp, preImage);
    }

    function doHash(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, bytes memory preImage) internal pure returns(bytes memory) {
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA256) {
            return abi.encodePacked(sha256(preImage));
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.KECCAK) {
            return abi.encodePacked(keccak256(preImage));
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.RIPEMD160) {
            return abi.encodePacked(ripemd160(preImage));
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.BITCOIN) {
            bytes memory tmp = abi.encodePacked(sha256(preImage));
            return abi.encodePacked(ripemd160(tmp));
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA512) {
            revert(); // dev: SHA512 not supported
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA512_256) {
            revert(); // dev: SHA512_256 not supported
        }
        revert(); // dev: Unsupported hashOp
    }

    function compare(bytes memory a, bytes memory b) internal pure returns(int) {
        uint256 minLen = Math.min(a.length, b.length);
        for (uint i = 0; i < minLen; i++) {
            if (uint8(a[i]) < uint8(b[i])) {
                return -1;
            } else if (uint8(a[i]) > uint8(b[i])) {
                return 1;
            }
        }
        if (a.length > minLen) {
            return 1;
        }
        if (b.length > minLen) {
            return -1;
        }
        return 0;
    }

    // private
    function doLengthOp(PROOFS_PROTO_GLOBAL_ENUMS.LengthOp lenOp, bytes memory data) private pure returns(bytes memory) {
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.NO_PREFIX) {
            return data;
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.VAR_PROTO) {
            uint256 sz = ProtoBufRuntime._sz_varint(data.length);
            bytes memory encoded = new bytes(sz);
            ProtoBufRuntime._encode_varint(data.length, 32, encoded);
            return abi.encodePacked(encoded, data);
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.REQUIRE_32_BYTES) {
            require(data.length == 32); // dev: data.length != 32
            return data;
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.REQUIRE_64_BYTES) {
            require(data.length == 64); // dev: data.length != 64"
            return data;
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.FIXED32_LITTLE) {
            uint32 size = SafeCast.toUint32(data.length);
            // maybe some assembly here to make it faster
            bytes4 sizeB = bytes4(size);
            bytes memory littleE = new bytes(4);
            //unfolding for loop is cheaper
            littleE[0] = sizeB[3];
            littleE[1] = sizeB[2];
            littleE[2] = sizeB[1];
            littleE[3] = sizeB[0];
            return abi.encodePacked(littleE, data);
        }
        revert(); // dev: Unsupported lenOp
    }

    function hasPrefix(bytes memory element, bytes memory prefix) private pure returns (bool) {
        if (prefix.length == 0) {
            return true;
        }
        if (prefix.length > element.length) {
            return false;
        }
        bytes memory slice = BytesLib.slice(element, 0, prefix.length);
        return BytesLib.equal(prefix, slice);
    }
}




contract Ops_UnitTest {
    function applyLeafOp(LeafOp.Data memory leaf, bytes memory key, bytes memory value) public pure returns(bytes memory) {
        return Ops.applyOp(leaf, key, value);
    }
    function checkAgainstLeafOpSpec(LeafOp.Data memory op, ProofSpec.Data memory spec) public pure {
        return Ops.checkAgainstSpec(op, spec);
    }
    function applyInnerOp(InnerOp.Data memory inner,bytes memory child) public pure returns(bytes memory) {
        return Ops.applyOp(inner, child);
    }
    function checkAgainstInnerOpSpec(InnerOp.Data memory op, ProofSpec.Data memory spec) public pure {
        return Ops.checkAgainstSpec(op, spec);
    }
    function doHash(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, bytes memory preImage) public pure returns(bytes memory) {
        return Ops.doHash(hashOp, preImage);
    }
}
