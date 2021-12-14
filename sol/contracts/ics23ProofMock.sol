// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;
import {Proof}  from "./ics23Proof.sol";
import { ExistenceProof, ProofSpec, CommitmentProof} from "./proofs.sol";

contract ProofMock {
    function calculateExistenceProofRoot(ExistenceProof.Data memory proof) public pure returns(bytes memory) {
        (bytes memory res, Proof.CalculateRootError eCode) = Proof.calculateRoot(proof);
        require(eCode == Proof.CalculateRootError.None); // dev: expand this require to check error code
        return res;
    }
    function checkAgainstSpec(ExistenceProof.Data memory proof, ProofSpec.Data memory spec) public pure {
        Proof.CheckAgainstSpecError cCode = Proof.checkAgainstSpec(proof, spec);
        require(cCode == Proof.CheckAgainstSpecError.None); // dev: expand this require to check error code
    }
    function calculateCommitmentProofRoot(CommitmentProof.Data memory proof) public pure returns(bytes memory) {
        (bytes memory res, Proof.CalculateRootError eCode) = Proof.calculateRoot(proof);
        require(eCode == Proof.CalculateRootError.None); // dev: expand this require to check error code
        return res;
    }
    function protobufDecodeCommitmentProof(bytes memory msg) public pure returns(CommitmentProof.Data memory) {
        CommitmentProof.Data memory res = CommitmentProof.decode(msg);
        return res;
    }
}
