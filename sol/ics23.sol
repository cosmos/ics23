pragma solidity >=0.8.0;

contract ICS23 {
    // Data structures and helper functions

    enum HashOp{NO_HASH, SHA256, SHA512, KECCAK, RIPEMD160, BITCOIN}
    enum LengthOp{NO_PREFIX, VAR_PROTO, VAR_RLP, FIXED32_BIG, FIXED32_LITTLE, FIXED64_BIG, FIXED64_LITTLE, REQUIRE_32_BYTES, REQUIRE_64_BYTES}

    struct ExistenceProof {
        bool valid;
        bytes key;
        bytes value;
        LeafOp leaf;
        InnerOp[] path;
    }

    struct NonExistenceProof {
        bool valid;
        bytes key;
        ExistenceProof left;
        ExistenceProof right;
    }

    struct LeafOp {
        bool valid;
        HashOp hash;
        HashOp prehash_key;
        HashOp prehash_value;
        LengthOp len;
        bytes prefix;
    }

    struct InnerOp {
        bool valid;
        HashOp hash;
        bytes prefix;
        bytes suffix;
    }

    struct ProofSpec {
        LeafOp leafSpec;
        InnerSpec innerSpec;
        uint maxDepth;
        uint minDepth;
    }

    struct InnerSpec {
        uint[] childOrder;
        uint childSize;
        uint minPrefixLength;
        uint maxPrefixLength;
        bytes emptyChild;
        HashOp hash;
    }

    LeafOp iavlSpec = LeafOp(
        true,
        HashOp.SHA256,
        HashOp.NO_HASH,
        HashOp.SHA256,
        LengthOp.VAR_PROTO,
        hex"00"
    );
   
    enum Ordering {LT, EQ, GT}

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

    function applyLeaf(LeafOp memory op, bytes memory key, bytes memory value) public pure returns (bytes memory) {
        require(key.length != 0);
        require(value.length != 0);

        bytes memory pkey = prepareLeafData(op.prehash_key, op.len, key);
        bytes memory pvalue = prepareLeafData(op.prehash_value, op.len, value);

        bytes memory data = abi.encodePacked(op.prefix, pkey, pvalue);
        return doHash(op.hash, data);
    }

    function applyInner(InnerOp memory op, bytes memory child) public pure returns (bytes memory) {
        require(child.length != 0);
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

    function checkAgainstSpec(ExistenceProof memory proof, LeafOp memory spec) public pure returns (bool) {
        if (!(
          proof.leaf.hash == spec.hash &&
          proof.leaf.prehash_key == spec.prehash_key &&
          proof.leaf.prehash_value == spec.prehash_value &&
          proof.leaf.len == spec.len &&
          hasprefix(proof.leaf.prefix, spec.prefix)
        )) {
          return false;
        }

        for (uint i = 0; i < proof.path.length; i++) {
          if (hasprefix(proof.path[i].prefix, spec.prefix)) {
            return false; 
          }
        }

        return true;
    }

    function calculate(ExistenceProof memory proof) public pure returns (bytes memory) {
        bytes memory res = applyLeaf(proof.leaf, proof.key, proof.value);
        for (uint i = 0; i < proof.path.length; i++) {
            res = applyInner(proof.path[i], res);
        }
        return res;
    }
  
    function verifyExistence(ExistenceProof memory proof, LeafOp memory spec, bytes memory root, bytes memory key, bytes memory value) public pure returns (bool) {
      return (
        checkAgainstSpec(proof, spec) &&
        equalBytes(key, proof.key) &&
        equalBytes(value, proof.value) &&
        equalBytes(calculate(proof), root)
      );
    }

    function verifyNonExistence(NonExistenceProof memory proof, ProofSpec memory spec, bytes memory root, bytes memory key) public pure returns (bool) {
      if (!proof.left.valid && !proof.right.valid) {
        return false;
      }
      if (proof.left.valid) {
        if (!verifyExistence(proof.left, spec.leafSpec, root, proof.left.key, proof.left.value)) {
          return false;
        }
        if (compareBytes(key, proof.right.key) != Ordering.LT) {
          return false;
        }
      }
      if (proof.right.valid) {
        if (!verifyExistence(proof.right, spec.leafSpec, root, proof.right.key, proof.right.value)) {
          return false;
        }
        if (compareBytes(key, proof.left.key) != Ordering.GT) {
          return false;
        }
      }

      if (!proof.left.valid) {
        if (!isLeftMost(spec.innerSpec, proof.right.path)) {
          return false;
        }
      } else if (!proof.right.valid) {
        if (!isRightMost(spec.innerSpec, proof.left.path)) {
          return false;
        }
      } else {
        if (!isLeftNeighbor(spec.innerSpec, proof.left.path, proof.right.path)) {
          return false;
        }
      }

      return true;
    }

    function orderFromPadding(InnerSpec memory spec, InnerOp memory op) public pure returns (uint) {
      uint maxBranch = spec.childOrder.length;
      
      for (uint branch = 0; branch < maxBranch; branch++) {
        (uint minPrefix, uint maxPrefix, uint suffix) = getPadding(spec, branch);
        if (hasPadding(op, minPrefix, maxPrefix, suffix)) {
          return branch;
        }
      }

      revert();
    }

    function isLeftStep(InnerSpec memory spec, InnerOp memory left, InnerOp memory right) public pure returns (bool) {
      uint leftidx = orderFromPadding(spec, left);
      uint rightidx = orderFromPadding(spec, right);
      return leftidx + 1 == rightidx;
    }

    function isLeftMost(InnerSpec memory spec, InnerOp[] memory path) public pure returns (bool) {
      (uint minPrefix, uint maxPrefix, uint suffix) = getPadding(spec, 0);
      
      for (uint i = 0; i < path.length; i++) {
        if (!hasPadding(path[i], minPrefix, maxPrefix, suffix)) {
          return false;
        }
      }
      return true;
    }

    function isRightMost(InnerSpec memory spec, InnerOp[] memory path) public pure returns (bool) {
      uint last = spec.childOrder.length - 1;
      (uint minPrefix, uint maxPrefix, uint suffix) = getPadding(spec, last);

      for (uint i = 0; i < path.length; i++) {
        if (!hasPadding(path[i], minPrefix, maxPrefix, suffix)) {
          return false;
        }
      }
      return true;

    }

    function isLeftNeighbor(InnerSpec memory spec, InnerOp[] memory left, InnerOp[] memory right) public pure returns (bool) {
      uint top = 1; 
      for ( ;
        equalBytes(left[left.length-top].prefix, right[right.length-top].prefix) && equalBytes(left[left.length-top].suffix, right[right.length-top].suffix); 
        top++) {}

      InnerOp memory topLeft = left[left.length-top];
      InnerOp memory topRight = right[right.length-top];

      if (!isLeftStep(spec, topLeft, topRight)) {
        return false;
      }

      if (!isRightMost(spec, left)) {
        return false;
      }

      if (!isLeftMost(spec, right)) {
        return false;
      }

      return true;
    }

    function getPosition(uint[] memory order, uint branch) public pure returns (uint) {
      require(branch < order.length);
      for (uint i = 0; i < order.length; i++) {
        if (order[i] == branch) {
          return i;
        }
      }
      revert();
    }

    function hasPadding(InnerOp memory op, uint minPrefix, uint maxPrefix, uint suffix) public pure returns (bool) {
      if (op.prefix.length < minPrefix) {
        return false;
      }
      if (op.prefix.length > maxPrefix) {
        return false;
      } 
      return op.suffix.length == suffix;
    }

    function getPadding(InnerSpec memory spec, uint branch) public pure returns (uint, uint, uint) {
      uint idx = getPosition(spec.childOrder, branch);
      
      uint prefix = idx * spec.childSize;
      uint minPrefix = prefix + spec.minPrefixLength;
      uint maxPrefix = prefix + spec.maxPrefixLength;

      uint suffix = (spec.childOrder.length - 1 - idx) * spec.childSize;
    
      return (minPrefix, maxPrefix, suffix);
    }

    function equalBytes(bytes memory bz1, bytes memory bz2) public pure returns (bool) {
      if (bz1.length != bz2.length) {
        return false;
      }
      for (uint i = 0; i < bz1.length; i++) {
        if (bz1[i] != bz2[i]) {
          return false;
        }
      }
      return true;
    }
   
    function compareBytes(bytes memory bz1, bytes memory bz2) public pure returns (Ordering) {
      for (uint i = 0; i < bz1.length && i < bz2.length; i++) {
        if (bz1[i] < bz2[i]) {
          return Ordering.LT;
        }
        if (bz1[i] > bz2[i]) {
          return Ordering.GT;
        }
      }
      if (bz1.length == bz2.length) {
        return Ordering.EQ;
      }
      if (bz1.length < bz2.length) {
        return Ordering.LT;
      }
      if (bz1.length > bz2.length) {
        return Ordering.GT;
      }
      revert("should not reach this line");
    }

    // ICS23 interface implementation

    // verifyMembership is synonym for verifyExistence
    function verifyMembership(LeafOp memory spec, bytes memory root, ExistenceProof memory proof, bytes memory key, bytes memory value) public pure returns (bool) {
      return verifyExistence(proof, spec, root, key, value);
    }

    function verifyNonMembership(ProofSpec memory spec, bytes memory root, NonExistenceProof memory proof, bytes memory key) public pure returns (bool) {
      return verifyNonExistence(proof, spec, root, key);
    }
}
