// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {BatchProof, CompressedBatchProof, CommitmentProof, ProofSpec, ExistenceProof, NonExistenceProof} from "./proofs.sol";
import {Compress} from "./ics23Compress.sol";
import {Proof} from "./ics23Proof.sol";
import {Ops} from "./ics23Ops.sol";
import {BytesLib} from "GNSPS/solidity-bytes-utils@0.8.0/contracts/BytesLib.sol";

library Ics23  {

    enum VerifyMembershipError {
        None,
        ExistenceProofIsNil,
        ProofVerify
    }
    // verifyMembership, throws an exception in case anything goes wrong
    function verifyMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, bytes memory key, bytes memory value) internal pure returns(VerifyMembershipError){
        CommitmentProof.Data memory decoProof = Compress.decompress(proof);
        ExistenceProof.Data memory exiProof = getExistProofForKey(decoProof, key);
        //require(ExistenceProof.isNil(exiProof) == false); // dev: getExistProofForKey not available
        if (ExistenceProof.isNil(exiProof)) return VerifyMembershipError.ExistenceProofIsNil;
        Proof.VerifyExistenceError vCode = Proof.verify(exiProof, spec, commitmentRoot, key, value);
        if (vCode != Proof.VerifyExistenceError.None) return VerifyMembershipError.ProofVerify;

        return VerifyMembershipError.None;
    }

    enum VerifyNonMembershipError {
        None,
        NonExistenceProofIsNil,
        ProofVerify
    }
    function verifyNonMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, bytes memory key) internal pure returns(VerifyNonMembershipError) {
        CommitmentProof.Data memory decoProof = Compress.decompress(proof);
        NonExistenceProof.Data memory nonProof = getNonExistProofForKey(decoProof, key);
        //require(NonExistenceProof.isNil(nonProof) == false); // dev: getNonExistProofForKey not available
        if (NonExistenceProof.isNil(nonProof)) return VerifyNonMembershipError.NonExistenceProofIsNil;
        Proof.VerifyNonExistenceError vCode =  Proof.verify(nonProof, spec, commitmentRoot, key);
        if (vCode != Proof.VerifyNonExistenceError.None) return VerifyNonMembershipError.ProofVerify;

        return VerifyNonMembershipError.None;
    }
/* -- temporarily disabled as they are not covered by unit tests
    struct BatchItem {
        bytes key;
        bytes value;
    }
    function batchVerifyMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, BatchItem[] memory items ) internal pure {
        CommitmentProof.Data memory decoProof = Compress.decompress(proof);
        for (uint i = 0; i < items.length; i++) {
            verifyMembership(spec, commitmentRoot, decoProof, items[i].key, items[i].value);
        }
    }

    function batchVerifyNonMembership(ProofSpec.Data memory spec, bytes memory commitmentRoot, CommitmentProof.Data memory proof, bytes[] memory keys ) internal pure {
        CommitmentProof.Data memory decoProof = Compress.decompress(proof);
        for (uint i = 0; i < keys.length; i++) {
            verifyNonMembership(spec, commitmentRoot, decoProof, keys[i]);
        }
    }
*/

    // private
    function getExistProofForKey(CommitmentProof.Data memory proof, bytes memory key) private pure returns(ExistenceProof.Data memory) {
        if (ExistenceProof.isNil(proof.exist) == false){
            if (BytesLib.equal(proof.exist.key, key) == true) {
                return proof.exist;
            }
        } else if(BatchProof.isNil(proof.batch) == false) {
            for (uint i = 0; i < proof.batch.entries.length; i++) {
                if (ExistenceProof.isNil(proof.batch.entries[i].exist) == false && 
                    BytesLib.equal(proof.batch.entries[i].exist.key, key)) {
                    return proof.batch.entries[i].exist;
                }
            }
        }
        return ExistenceProof.nil();
    }

    function getNonExistProofForKey(CommitmentProof.Data memory proof, bytes memory key) private pure returns(NonExistenceProof.Data memory) {
        if (NonExistenceProof.isNil(proof.nonexist) == false) {
            if (isLeft(proof.nonexist.left, key) && isRight(proof.nonexist.right, key)) {
                return proof.nonexist;
            }
        } else if (BatchProof.isNil(proof.batch) == false) {
            for (uint i = 0; i < proof.batch.entries.length; i++) {
                if (NonExistenceProof.isNil(proof.batch.entries[i].nonexist) == false && 
                    isLeft(proof.batch.entries[i].nonexist.left, key) && 
                    isRight(proof.batch.entries[i].nonexist.right, key)) {
                    return proof.batch.entries[i].nonexist;
                }
            }
        }
        return NonExistenceProof.nil();
    }

    function isLeft(ExistenceProof.Data memory left, bytes memory key) private pure returns(bool) {
        // ExistenceProof.isNil does not work
        return ExistenceProof._empty(left) || Ops.compare(left.key, key) < 0;
    }

    function isRight(ExistenceProof.Data memory right, bytes memory key) private pure returns(bool) {
        // ExistenceProof.isNil does not work
        return ExistenceProof._empty(right) || Ops.compare(right.key, key) > 0;
    }
}
