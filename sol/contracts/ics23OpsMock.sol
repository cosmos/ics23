// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {Ops} from "./ics23Ops.sol";
import {LeafOp, InnerOp, PROOFS_PROTO_GLOBAL_ENUMS, ProofSpec} from "./proofs.sol";


contract OpsMock {
    function applyLeafOp(LeafOp.Data memory leaf, bytes memory key, bytes memory value) public pure returns(bytes memory) {
        (bytes memory res, Ops.ApplyLeafOpError aCode) = Ops.applyOp(leaf, key, value);
        require(aCode == Ops.ApplyLeafOpError.None); // dev: expand this require to check error code
        return res;
    }
    function checkAgainstLeafOpSpec(LeafOp.Data memory op, ProofSpec.Data memory spec) public pure  {
        Ops.CheckAgainstSpecError cCode = Ops.checkAgainstSpec(op, spec);
        require(cCode == Ops.CheckAgainstSpecError.None); // dev: expand this require to check error code
    }
    function applyInnerOp(InnerOp.Data memory inner,bytes memory child) public pure returns(bytes memory) {
        (bytes memory res, Ops.ApplyInnerOpError aCode) = Ops.applyOp(inner, child);
        require(aCode == Ops.ApplyInnerOpError.None); // dev: expand this require to check error code
        return res;
    }
    function checkAgainstInnerOpSpec(InnerOp.Data memory op, ProofSpec.Data memory spec) public pure {
        Ops.CheckAgainstSpecError cCode = Ops.checkAgainstSpec(op, spec);
        require(cCode == Ops.CheckAgainstSpecError.None); // dev: expand this require to check error code
    }
    function doHash(PROOFS_PROTO_GLOBAL_ENUMS.HashOp hashOp, bytes memory preImage) public pure returns(bytes memory) {
        (bytes memory res, Ops.DoHashError hCode) = Ops.doHash(hashOp, preImage);
        require(hCode == Ops.DoHashError.None);  // dev: expand this require to check error code
        return res;
    }
}
