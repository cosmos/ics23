// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {LeafOp, InnerOp, PROOFS_PROTO_GLOBAL_ENUMS, ProofSpec} from "./proofs.sol";
import {ProtoBufRuntime} from "./ProtoBufRuntime.sol";
import {SafeCast} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/SafeCast.sol";
import {Math} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/Math.sol";
import {BytesLib} from "GNSPS/solidity-bytes-utils@0.8.0/contracts/BytesLib.sol";

library Ops {
    bytes constant empty = new bytes(0);

    enum ApplyLeafOpError {
        None,
        KeyLength,
        ValueLength,
        DoHash,
        PrepareLeafData
    }
    /**
    @notice calculates the leaf hash given the key and value being proven
    @return VerifyExistenceError enum giving indication of where error happened, None if verification succeded
        */
    function applyOp(
        LeafOp.Data memory leafOp,
        bytes memory key,
        bytes memory value
    ) internal pure returns(bytes memory, ApplyLeafOpError) {
        //require(key.length > 0); // dev: Leaf op needs key
        if (key.length == 0) return (empty, ApplyLeafOpError.KeyLength);
        //require(value.length > 0); // dev: Leaf op needs value
        if (value.length == 0) return (empty, ApplyLeafOpError.ValueLength);
        (bytes memory pKey, PrepareLeafDataError pCode1) = prepareLeafData(leafOp.prehash_key, leafOp.length, key);
        if (pCode1 != PrepareLeafDataError.None) return (empty, ApplyLeafOpError.PrepareLeafData);
        (bytes memory pValue, PrepareLeafDataError pCode2) = prepareLeafData(leafOp.prehash_value, leafOp.length, value);
        if (pCode2 != PrepareLeafDataError.None) return (empty, ApplyLeafOpError.PrepareLeafData);
        bytes memory data = abi.encodePacked(leafOp.prefix, pKey, pValue);
        (bytes memory hashed, DoHashError hCode) = doHash(leafOp.hash, data);
        if (hCode != DoHashError.None) return (empty, ApplyLeafOpError.DoHash);
        return(hashed, ApplyLeafOpError.None);
    }

    enum PrepareLeafDataError {
        None,
        DoHash,
        DoLengthOp
    }
    function prepareLeafData(
        PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp,
        PROOFS_PROTO_GLOBAL_ENUMS.LengthOp lenOp,
        bytes memory data
    ) internal pure returns(bytes memory, PrepareLeafDataError) {
        (bytes memory hased, DoHashError hCode) = doHashOrNoop(hashOp, data);
        if (hCode != DoHashError.None)return (empty, PrepareLeafDataError.DoHash);
        (bytes memory res, DoLengthOpError lCode) = doLengthOp(lenOp, hased);
        if (lCode != DoLengthOpError.None) return (empty, PrepareLeafDataError.DoLengthOp);

        return (res, PrepareLeafDataError.None);
    }

    enum CheckAgainstSpecError{
        None,
        Hash,
        PreHashKey,
        PreHashValue,
        Length,
        MinPrefixLength,
        HasPrefix,
        MaxPrefixLength
    }
    /**
    @notice will verify the LeafOp is in the format defined in spec
    */
    function checkAgainstSpec(LeafOp.Data memory leafOp, ProofSpec.Data memory spec) internal pure returns(CheckAgainstSpecError) {
        //require (leafOp.hash == spec.leaf_spec.hash); // dev: checkAgainstSpec for LeafOp - Unexpected HashOp
        if (leafOp.hash != spec.leaf_spec.hash) return CheckAgainstSpecError.Hash;
        //require(leafOp.prehash_key == spec.leaf_spec.prehash_key); // dev: checkAgainstSpec for LeafOp - Unexpected PrehashKey
        if (leafOp.prehash_key != spec.leaf_spec.prehash_key) return CheckAgainstSpecError.PreHashKey;
        //require(leafOp.prehash_value == spec.leaf_spec.prehash_value); // dev: checkAgainstSpec for LeafOp - Unexpected PrehashValue");
        if (leafOp.prehash_value != spec.leaf_spec.prehash_value) return CheckAgainstSpecError.PreHashValue;
        //require(leafOp.length == spec.leaf_spec.length); // dev: checkAgainstSpec for LeafOp - Unexpected lengthOp
        if (leafOp.length != spec.leaf_spec.length) return CheckAgainstSpecError.Length;
        bool hasprefix = hasPrefix(leafOp.prefix, spec.leaf_spec.prefix);
        //require(hasprefix); // dev: checkAgainstSpec for LeafOp - Leaf Prefix doesn't start with
        if (hasprefix == false) return CheckAgainstSpecError.HasPrefix;

        return CheckAgainstSpecError.None;
    }

    enum ApplyInnerOpError {
        None,
        ChildLength,
        DoHash
    }
    /**
    @notice apply will calculate the hash of the next step, given the hash of the previous step
    */
    function applyOp(InnerOp.Data memory innerOp, bytes memory child ) internal pure returns(bytes memory, ApplyInnerOpError) {
        //require(child.length > 0); // dev: Inner op needs child value
        if (child.length == 0) return (empty, ApplyInnerOpError.ChildLength);
        bytes memory preImage = abi.encodePacked(innerOp.prefix, child, innerOp.suffix);
        (bytes memory hashed, DoHashError code) = doHash(innerOp.hash, preImage);
        if (code != DoHashError.None) return (empty, ApplyInnerOpError.DoHash);

        return (hashed, ApplyInnerOpError.None);
    }

    /**
    @notice will verify the InnerOp is in the format defined in spec
    */
    function checkAgainstSpec(InnerOp.Data memory innerOp, ProofSpec.Data memory spec) internal pure returns(CheckAgainstSpecError) {
        //require(innerOp.hash == spec.inner_spec.hash); // dev: checkAgainstSpec for InnerOp - Unexpected HashOp
        if (innerOp.hash != spec.inner_spec.hash) return CheckAgainstSpecError.Hash;
        uint256 minPrefixLength = SafeCast.toUint256(spec.inner_spec.min_prefix_length);
        //require(innerOp.prefix.length >= minPrefixLength); // dev: InnerOp prefix too short;
        if (innerOp.prefix.length < minPrefixLength)  return CheckAgainstSpecError.MinPrefixLength;
        bytes memory leafPrefix = spec.leaf_spec.prefix;
        bool hasprefix = hasPrefix(innerOp.prefix, leafPrefix);
        //require(hasprefix == false); // dev: Inner Prefix starts with wrong value
        if (hasprefix) return CheckAgainstSpecError.HasPrefix;
        uint256 childSize = SafeCast.toUint256(spec.inner_spec.child_size);
        uint256 maxLeftChildBytes = (spec.inner_spec.child_order.length - 1) * childSize;
        uint256 maxPrefixLength = SafeCast.toUint256(spec.inner_spec.max_prefix_length);
        //require(innerOp.prefix.length <= maxPrefixLength + maxLeftChildBytes); // dev: InnerOp prefix too long
        if (innerOp.prefix.length > maxPrefixLength + maxLeftChildBytes)  return CheckAgainstSpecError.MaxPrefixLength;

        return CheckAgainstSpecError.None;
    }

    /**
    @notice will return the preimage untouched if hashOp == NONE, otherwise, perform doHash
    */
    function doHashOrNoop(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, bytes memory preImage) internal pure returns(bytes memory, DoHashError) {
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.NO_HASH) {
            return (preImage, DoHashError.None);
        }
        return doHash(hashOp, preImage);
    }

    enum DoHashError {
        None,
        Sha512,
        Sha512_256,
        Unsupported
    }
    /**
    @notice will preform the specified hash on the preimage. If hashOp == NONE,
     */
    function doHash(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, bytes memory preImage) internal pure returns(bytes memory, DoHashError) {
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA256) {
            return (abi.encodePacked(sha256(preImage)), DoHashError.None);
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.KECCAK) {
            return (abi.encodePacked(keccak256(preImage)), DoHashError.None);
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.RIPEMD160) {
            return (abi.encodePacked(ripemd160(preImage)), DoHashError.None);
        }
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.BITCOIN) {
            bytes memory tmp = abi.encodePacked(sha256(preImage));
            return (abi.encodePacked(ripemd160(tmp)), DoHashError.None);
        }
        //require(hashOp != PROOFS_PROTO_GLOBAL_ENUMS.HashOp.Sha512); // dev: SHA512 not supported
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA512) {
            return (empty, DoHashError.Sha512);
        }
        //require(hashOp != PROOFS_PROTO_GLOBAL_ENUMS.HashOp.Sha512_256); // dev: SHA512_256 not supported
        if (hashOp == PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA512_256) {
            return (empty, DoHashError.Sha512_256);
        }
        //revert(); // dev: Unsupported hashOp
        return (empty, DoHashError.Unsupported);
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
    enum DoLengthOpError {
        None,
        Require32DataLength,
        Require64DataLength,
        Unsupported
    }
    /**
      @notice will calculate the proper prefix and return it prepended doLengthOp(op, data) -> length(data) || data
     */
    function doLengthOp(PROOFS_PROTO_GLOBAL_ENUMS.LengthOp lenOp, bytes memory data) private pure returns(bytes memory, DoLengthOpError) {
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.NO_PREFIX) {
            return (data, DoLengthOpError.None);
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.VAR_PROTO) {
            uint256 sz = ProtoBufRuntime._sz_varint(data.length);
            bytes memory encoded = new bytes(sz);
            ProtoBufRuntime._encode_varint(data.length, 32, encoded);
            return (abi.encodePacked(encoded, data), DoLengthOpError.None);
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.REQUIRE_32_BYTES) {
            //require(data.length == 32); // dev: data.length != 32
            if (data.length != 32) return (empty, DoLengthOpError.Require32DataLength);

            return (data, DoLengthOpError.None);
        }
        if (lenOp == PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.REQUIRE_64_BYTES) {
            //require(data.length == 64); // dev: data.length != 64"
            if (data.length != 64) return (empty, DoLengthOpError.Require64DataLength);

            return (data, DoLengthOpError.None);
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
            return (abi.encodePacked(littleE, data), DoLengthOpError.None);
        }
        //revert(); // dev: Unsupported lenOp
        return (empty, DoLengthOpError.Unsupported);
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
