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
    function decompress(Ics23CommitmentProof.Data memory proof) internal pure returns(Ics23CommitmentProof.Data memory) {
        //Ics23CompressedBatchProof.isNil() does not work
        if (Ics23CompressedBatchProof._empty(proof.compressed) == true){
            return proof;
        }
        return Ics23CommitmentProof.Data({
            exist: Ics23ExistenceProof.nil(),
            nonexist: Ics23NonExistenceProof.nil(),
            batch: Ics23BatchProof.Data({
                entries: decompress(proof.compressed)
            }),
            compressed: Ics23CompressedBatchProof.nil()
        });
    }

    function decompress(Ics23CompressedBatchProof.Data memory proof) private pure returns(Ics23BatchEntry.Data[] memory) {
        Ics23BatchEntry.Data[] memory entries = new Ics23BatchEntry.Data[](proof.entries.length);
        for(uint i = 0; i < proof.entries.length; i++) {
            entries[i] = decompressEntry(proof.entries[i], proof.lookup_inners);
        }
        return entries;
    }

    function decompressEntry(
        Ics23CompressedBatchEntry.Data memory entry,
        Ics23InnerOp.Data[] memory lookup
    ) private pure returns(Ics23BatchEntry.Data memory) {
        //Ics23CompressedExistenceProof.isNil does not work
        if (Ics23CompressedExistenceProof._empty(entry.exist) == false) {
            return Ics23BatchEntry.Data({
                exist: decompressExist(entry.exist, lookup),
                nonexist: Ics23NonExistenceProof.nil()
            });
        }
        return Ics23BatchEntry.Data({
            exist: Ics23ExistenceProof.nil(),
            nonexist: Ics23NonExistenceProof.Data({
                key: entry.nonexist.key,
                left: decompressExist(entry.nonexist.left, lookup),
                right: decompressExist(entry.nonexist.right, lookup)
            })
        });
    }

    function decompressExist(
        Ics23CompressedExistenceProof.Data memory proof,
        Ics23InnerOp.Data[] memory lookup
    ) private pure returns(Ics23ExistenceProof.Data memory) {
        if (Ics23CompressedExistenceProof._empty(proof)) {
            return Ics23ExistenceProof.nil();
        }
        Ics23ExistenceProof.Data memory decoProof = Ics23ExistenceProof.Data({
            key: proof.key,
            value: proof.value,
            leaf: proof.leaf,
            path : new Ics23InnerOp.Data[](proof.path.length)
        });
        for (uint i = 0; i < proof.path.length; i++) {
            require(proof.path[i] >= 0); // dev: proof.path < 0
            uint step = SafeCast.toUint256(proof.path[i]);
            require(step < lookup.length); // dev: step >= lookup.length
            decoProof.path[i] = lookup[step];
        }
        return decoProof;
    }
}
