// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {
    Ics23InnerOp, Ics23ExistenceProof, Ics23NonExistenceProof, Ics23CommitmentProof, Ics23CompressedBatchEntry, Ics23CompressedBatchProof,
    Ics23CompressedExistenceProof, Ics23BatchEntry, Ics23BatchProof
} from "./proofs.sol";
import {SafeCast} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/SafeCast.sol";

library Compress {
    /**
      @notice will return a Ics23BatchProof if the input is CompressedBatchProof. Otherwise it will return the input.
      This is safe to call multiple times (idempotent)
    */
    function decompress(
        Ics23CommitmentProof.Data memory proof
    ) internal pure returns(Ics23CommitmentProof.Data memory, DecompressEntryError) {
        //Ics23CompressedBatchProof.isNil() does not work
        if (Ics23CompressedBatchProof._empty(proof.compressed) == true){
            return (proof, DecompressEntryError.None);
        }
        (Ics23BatchEntry.Data[] memory entries, DecompressEntryError erCode) = decompress(proof.compressed);
        if (erCode != DecompressEntryError.None) return (Ics23CommitmentProof.nil(), erCode);
        Ics23CommitmentProof.Data memory retVal;
        retVal.exist = Ics23ExistenceProof.nil();
        retVal.nonexist = Ics23NonExistenceProof.nil();
        retVal.compressed = Ics23CompressedBatchProof.nil();
        retVal.batch.entries = entries;
        return (retVal, DecompressEntryError.None);
    }

    function decompress(
        Ics23CompressedBatchProof.Data memory proof
    ) private pure returns(Ics23BatchEntry.Data[] memory, DecompressEntryError) {
        Ics23BatchEntry.Data[] memory entries = new Ics23BatchEntry.Data[](proof.entries.length);
        for(uint i = 0; i < proof.entries.length; i++) {
            (Ics23BatchEntry.Data memory entry, DecompressEntryError erCode) = decompressEntry(proof.entries[i], proof.lookup_inners);
            if (erCode != DecompressEntryError.None) return (entries, erCode);
            entries[i] = entry;
        }
        return (entries, DecompressEntryError.None);
    }

    enum DecompressEntryError{
        None,
        ExistDecompress,
        LeftDecompress,
        RightDecompress
    }
    function decompressEntry(
        Ics23CompressedBatchEntry.Data memory entry,
        Ics23InnerOp.Data[] memory lookup
    ) private pure returns(Ics23BatchEntry.Data memory, DecompressEntryError) {
        //Ics23CompressedExistenceProof.isNil does not work
        if (Ics23CompressedExistenceProof._empty(entry.exist) == false) {
            (Ics23ExistenceProof.Data memory exist, DecompressExistError existErCode) = decompressExist(entry.exist, lookup);
            if (existErCode != DecompressExistError.None) return(Ics23BatchEntry.nil(), DecompressEntryError.ExistDecompress);
            return (Ics23BatchEntry.Data({
                exist: exist,
                nonexist: Ics23NonExistenceProof.nil()
            }), DecompressEntryError.None);
        }
        (Ics23ExistenceProof.Data memory left, DecompressExistError leftErCode) = decompressExist(entry.nonexist.left, lookup);
        if (leftErCode != DecompressExistError.None) return(Ics23BatchEntry.nil(), DecompressEntryError.LeftDecompress);
        (Ics23ExistenceProof.Data memory right, DecompressExistError rightErCode) = decompressExist(entry.nonexist.right, lookup);
        if (rightErCode != DecompressExistError.None) return(Ics23BatchEntry.nil(), DecompressEntryError.RightDecompress);
        return (Ics23BatchEntry.Data({
            exist: Ics23ExistenceProof.nil(),
            nonexist: Ics23NonExistenceProof.Data({
                key: entry.nonexist.key,
                left: left,
                right: right
            })
        }), DecompressEntryError.None);
    }

    enum DecompressExistError{
        None,
        PathLessThanZero,
        StepGreaterOrEqualToLength
    }
    function decompressExist(
        Ics23CompressedExistenceProof.Data memory proof,
        Ics23InnerOp.Data[] memory lookup
    ) private pure returns(Ics23ExistenceProof.Data memory, DecompressExistError) {
        if (Ics23CompressedExistenceProof._empty(proof)) {
            return (Ics23ExistenceProof.nil(), DecompressExistError.None);
        }
        Ics23ExistenceProof.Data memory decoProof = Ics23ExistenceProof.Data({
            key: proof.key,
            value: proof.value,
            leaf: proof.leaf,
            path : new Ics23InnerOp.Data[](proof.path.length)
        });
        for (uint i = 0; i < proof.path.length; i++) {
            //require(proof.path[i] >= 0); // dev: proof.path < 0
            if (proof.path[i] < 0) return (Ics23ExistenceProof.nil(), DecompressExistError.PathLessThanZero);
            uint step = SafeCast.toUint256(proof.path[i]);
            //require(step < lookup.length); // dev: step >= lookup.length
            if (step >= lookup.length) return (Ics23ExistenceProof.nil(), DecompressExistError.StepGreaterOrEqualToLength);
            decoProof.path[i] = lookup[step];
        }
        return (decoProof, DecompressExistError.None);
    }
}
