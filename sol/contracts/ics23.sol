// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {
    Ics23BatchProof, Ics23CompressedBatchProof, Ics23CommitmentProof, Ics23ProofSpec, Ics23ExistenceProof, Ics23NonExistenceProof
} from "./proofs.sol";
import {Compress} from "./ics23Compress.sol";
import {Proof} from "./ics23Proof.sol";
import {Ops} from "./ics23Ops.sol";
import {BytesLib} from "GNSPS/solidity-bytes-utils@0.8.0/contracts/BytesLib.sol";

library Ics23  {

    enum VerifyMembershipError {
        None,
        ExistenceProofIsNil,
        ProofVerify,
        Decompress
    }
    function verifyMembership(
        Ics23ProofSpec.Data memory spec,
        bytes memory commitmentRoot,
        Ics23CommitmentProof.Data memory proof,
        bytes memory key,
        bytes memory value
    ) internal pure returns(VerifyMembershipError){
        (Ics23CommitmentProof.Data memory decoProof, Compress.DecompressEntryError erCode) = Compress.decompress(proof);
        if (erCode != Compress.DecompressEntryError.None) return VerifyMembershipError.Decompress;
        Ics23ExistenceProof.Data memory exiProof = getExistProofForKey(decoProof, key);
        //require(Ics23ExistenceProof.isNil(exiProof) == false); // dev: getExistProofForKey not available
        if (Ics23ExistenceProof.isNil(exiProof)) return VerifyMembershipError.ExistenceProofIsNil;
        Proof.VerifyExistenceError vCode = Proof.verify(exiProof, spec, commitmentRoot, key, value);
        if (vCode != Proof.VerifyExistenceError.None) return VerifyMembershipError.ProofVerify;

        return VerifyMembershipError.None;
    }

    enum VerifyNonMembershipError {
        None,
        NonExistenceProofIsNil,
        ProofVerify,
        Decompress
    }

    function verifyNonMembership(
        Ics23ProofSpec.Data memory spec,
        bytes memory commitmentRoot,
        Ics23CommitmentProof.Data memory proof,
        bytes memory key
    ) internal pure returns(VerifyNonMembershipError) {
        (Ics23CommitmentProof.Data memory decoProof, Compress.DecompressEntryError erCode) = Compress.decompress(proof);
        if (erCode != Compress.DecompressEntryError.None) return VerifyNonMembershipError.Decompress;
        Ics23NonExistenceProof.Data memory nonProof = getNonExistProofForKey(decoProof, key);
        //require(Ics23NonExistenceProof.isNil(nonProof) == false); // dev: getNonExistProofForKey not available
        if (Ics23NonExistenceProof.isNil(nonProof)) return VerifyNonMembershipError.NonExistenceProofIsNil;
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
    function getExistProofForKey(
        Ics23CommitmentProof.Data memory proof,
        bytes memory key
    ) private pure returns(Ics23ExistenceProof.Data memory) {
        if (Ics23ExistenceProof.isNil(proof.exist) == false){
            if (BytesLib.equal(proof.exist.key, key) == true) {
                return proof.exist;
            }
        } else if(Ics23BatchProof.isNil(proof.batch) == false) {
            for (uint i = 0; i < proof.batch.entries.length; i++) {
                if (Ics23ExistenceProof.isNil(proof.batch.entries[i].exist) == false && 
                    BytesLib.equal(proof.batch.entries[i].exist.key, key)) {
                    return proof.batch.entries[i].exist;
                }
            }
        }
        return Ics23ExistenceProof.nil();
    }

    function getNonExistProofForKey(
        Ics23CommitmentProof.Data memory proof,
        bytes memory key
    ) private pure returns(Ics23NonExistenceProof.Data memory) {
        if (Ics23NonExistenceProof.isNil(proof.nonexist) == false) {
            if (isLeft(proof.nonexist.left, key) && isRight(proof.nonexist.right, key)) {
                return proof.nonexist;
            }
        } else if (Ics23BatchProof.isNil(proof.batch) == false) {
            for (uint i = 0; i < proof.batch.entries.length; i++) {
                if (Ics23NonExistenceProof.isNil(proof.batch.entries[i].nonexist) == false && 
                    isLeft(proof.batch.entries[i].nonexist.left, key) && 
                    isRight(proof.batch.entries[i].nonexist.right, key)) {
                    return proof.batch.entries[i].nonexist;
                }
            }
        }
        return Ics23NonExistenceProof.nil();
    }

    function isLeft(Ics23ExistenceProof.Data memory left, bytes memory key) private pure returns(bool) {
        // Ics23ExistenceProof.isNil does not work
        return Ics23ExistenceProof._empty(left) || Ops.compare(left.key, key) < 0;
    }

    function isRight(Ics23ExistenceProof.Data memory right, bytes memory key) private pure returns(bool) {
        // Ics23ExistenceProof.isNil does not work
        return Ics23ExistenceProof._empty(right) || Ops.compare(right.key, key) > 0;
    }
}
