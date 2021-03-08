pragma solidity ^0.5.3;

contract Proofs {
    enum HashOp{NO_HASH, SHA256, SHA512, KECCAK, RIPEMD160, BITCOIN}
    enum LengthOp{NO_PREFIX, VAR_PROTO, VAR_RLP, FIXED32_BIG, FIXED32_LITTLE, FIXED64_BIG, FIXED64_LITTLE, REQUIRE_32_BYTES, REQUIRE_64_BYTES}

    struct ExistenceProof {
        bytes key;
        bytes value;
        LeafOp leaf;
        InnerOp[] path;
    }

    struct LeafOp {
        HashOp hash;
        HashOp prehash_key;
        HashOp prehash_value;
        LengthOp len;
        bytes prefix;
    }

    struct InnerOp {
        HashOp hash;
        bytes prefix;
        bytes suffix;
    }

    LeafOp iavlSpec = LeafOp(
        HashOp.SHA256,
        HashOp.NO_HASH,
        HashOp.SHA256,
        LengthOp.VAR_PROTO,
        hex"00"
    );
    
    function decodeExistenceProof( 
      bytes memory key, bytes memory value,
      bytes memory leafOpEncoded, bytes memory pathEncodedEncoded
    )
    internal pure returns (ExistenceProof memory) {
      (HashOp leafHash, HashOp prehashKey, HashOp prehashValue, LengthOp len, bytes memory leafPrefix) = abi.decode(leafOpEncoded, (HashOp, HashOp, HashOp, LengthOp, bytes));
      LeafOp memory leaf = LeafOp(leafHash, prehashKey, prehashValue, len, leafPrefix);
      bytes[] memory pathEncoded = abi.decode(pathEncodedEncoded, (bytes[]));
      InnerOp[] memory path = new InnerOp[](pathEncoded.length);
      for (uint i = 0; i < pathEncoded.length; i++) {
        (HashOp hash, bytes memory prefix, bytes memory suffix) = abi.decode(pathEncoded[i], (HashOp, bytes, bytes));
        path[i] = InnerOp(hash, prefix, suffix);
      }
      
      return ExistenceProof(key, value, leaf, path);
    }

    function encodeLeafOp(HashOp hash, HashOp phkey, HashOp phval, LengthOp len, bytes memory prefix) public pure returns (bytes memory) {
      return abi.encode(hash, phkey, phval, len, prefix);
    }

    function encodeInnerOp(HashOp hash, bytes memory prefix, bytes memory suffix) public pure returns (bytes memory) {
      return abi.encode(hash, prefix, suffix);
    }

    function doHashOrNoop(HashOp op, bytes memory preimage) public pure returns (bytes memory) {
        if (op == HashOp.NO_HASH) {
            return preimage;
        }
        return doHash(op, preimage);
    }

    function doHash(HashOp op, bytes memory preimage) public pure returns (bytes memory) {
        if (op == HashOp.KECCAK) {
            return abi.encodePacked(keccak256(preimage));
        }
        if (op == HashOp.SHA256) {
            return abi.encodePacked(sha256(preimage));
        }
        if (op == HashOp.RIPEMD160) {
            return abi.encodePacked(ripemd160(preimage));
        }
        if (op == HashOp.BITCOIN) {
            return abi.encodePacked(ripemd160(abi.encodePacked(sha256(preimage))));
        }
        revert("invalid or unsupported hash operation");
    }

    function doLength(LengthOp op, bytes memory data) public pure returns (bytes memory) {
        if (op == LengthOp.NO_PREFIX) {
            return data;
        }
        if (op == LengthOp.VAR_PROTO) {
            uint l = data.length;
            if (l >= 1<<7) {
                return abi.encodePacked(uint8(l&0x7f|0x80), uint8(l>>=7), data);
            }
            return abi.encodePacked(uint8(l), data);
        }
        if (op == LengthOp.REQUIRE_32_BYTES) {
            require(data.length == 32);
            return data;
        }
        if (op == LengthOp.REQUIRE_64_BYTES) {
            require(data.length == 64);
            return data;
        }
        revert("invalid or unsupported length operation");
    }

    function prepareLeafData(HashOp hashop, LengthOp lengthop, bytes memory data) public pure returns (bytes memory) {
        bytes memory hashed = doHashOrNoop(hashop, data);
        bytes memory result = doLength(lengthop, hashed);
        return result;
    }

    function applyLeaf(LeafOp memory op, bytes memory key, bytes memory value) internal pure returns (bytes memory) {
        require(key.length != 0);
        require(value.length != 0);

        bytes memory pkey = prepareLeafData(op.prehash_key, op.len, key);
        bytes memory pvalue = prepareLeafData(op.prehash_value, op.len, value);

        bytes memory data = abi.encodePacked(op.prefix, pkey, pvalue);
        return doHash(op.hash, data);
    }

    function applyInner(InnerOp memory op, bytes memory child) internal pure returns (bytes memory) {
        require(child.length == 0);
        bytes memory preimage = abi.encodePacked(op.prefix, child, op.suffix);
        return doHash(op.hash, preimage);
    }

    function hasprefix(bytes memory s, bytes memory prefix) public pure returns (bool) {
        for (uint i = 0; i < prefix.length; i++) {
            if (s[i] != prefix[i]) {
                return false;
            }
        }
        return true;
    }

    function checkAgainstSpec(ExistenceProof memory proof, LeafOp memory spec) internal pure {
        require(proof.leaf.hash == spec.hash);
        require(proof.leaf.prehash_key == spec.prehash_key);
        require(proof.leaf.prehash_value == spec.prehash_value);
        require(proof.leaf.len == spec.len);
        require(hasprefix(proof.leaf.prefix, spec.prefix));

        for (uint i = 0; i < proof.path.length; i++) {
            require(hasprefix(proof.path[i].prefix, spec.prefix));
        }
    }

    function calculate(ExistenceProof memory proof) internal pure returns (bytes memory) {
        bytes memory res = applyLeaf(proof.leaf, proof.key, proof.value);
        for (uint i = 0; i < proof.path.length; i++) {
            res = applyInner(proof.path[i], res);
        }
        return res;
    }


}
