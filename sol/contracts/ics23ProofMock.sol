// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;
import {Proof}  from "./ics23Proof.sol";
import {Ics23ExistenceProof, Ics23ProofSpec, Ics23CommitmentProof} from "./proofs.sol";

contract ProofMock {
    function calculateExistenceProofRoot(Ics23ExistenceProof.Data memory proof) public pure returns(bytes memory) {
        (bytes memory res, Proof.CalculateRootError eCode) = Proof.calculateRoot(proof);
        require(eCode == Proof.CalculateRootError.None); // dev: expand this require to check error code
        return res;
    }
    function checkAgainstSpec(Ics23ExistenceProof.Data memory proof, Ics23ProofSpec.Data memory spec) public pure {
        Proof.CheckAgainstSpecError cCode = Proof.checkAgainstSpec(proof, spec);
        require(cCode == Proof.CheckAgainstSpecError.None); // dev: expand this require to check error code
    }
    function calculateCommitmentProofRoot(Ics23CommitmentProof.Data memory proof) public pure returns(bytes memory) {
        (bytes memory res, Proof.CalculateRootError eCode) = Proof.calculateRoot(proof);
        require(eCode == Proof.CalculateRootError.None); // dev: expand this require to check error code
        return res;
    }
    function protobufDecodeCommitmentProof(bytes memory protoMsg) public pure returns(Ics23CommitmentProof.Data memory) {
        Ics23CommitmentProof.Data memory res = Ics23CommitmentProof.decode(protoMsg);
        return res;
    }
}
