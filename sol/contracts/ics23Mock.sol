// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {BatchProof, CompressedBatchProof, CommitmentProof, ProofSpec, ExistenceProof, NonExistenceProof} from "./proofs.sol";
import {Ics23} from "./ics23.sol";


contract ICS23Mock {
    function verifyMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, bytes memory key, bytes memory value) public pure {
        Ics23.VerifyMembershipError vCode = Ics23.verifyMembership(spec, commitmentRoot, proof, key, value);
        require(vCode == Ics23.VerifyMembershipError.None);
    }
    function verifyNonMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, bytes memory key) public pure {
        Ics23.VerifyNonMembershipError vCode = Ics23.verifyNonMembership(spec, commitmentRoot, proof, key);
        require(vCode == Ics23.VerifyNonMembershipError.None);
    }
/* -- temporarily disabled as they are not covered by unit tests
    function batchVerifyMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, Ics23.BatchItem[] memory items ) public pure {
        Ics23.batchVerifyMembership(spec, commitmentRoot, proof, items);
    }
    function batchVerifyNonMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, bytes[] memory keys ) public pure {
        Ics23.batchVerifyNonMembership(spec, commitmentRoot, proof, keys);
    }
*/
}

