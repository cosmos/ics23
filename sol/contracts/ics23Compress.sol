// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {InnerOp, ExistenceProof, NonExistenceProof, CommitmentProof, CompressedBatchEntry, CompressedBatchProof, CompressedExistenceProof, BatchEntry, BatchProof} from "./proofs.sol";
import {SafeCast} from "OpenZeppelin/openzeppelin-contracts@4.2.0/contracts/utils/math/SafeCast.sol";

library Compress {
    /**
      @notice will return a BatchProof if the input is CompressedBatchProof. Otherwise it will return the input.
      This is safe to call multiple times (idempotent)
    */
    function decompress(CommitmentProof.Data memory proof) internal pure returns(CommitmentProof.Data memory) {
        //CompressedBatchProof.isNil() does not work
        if (CompressedBatchProof._empty(proof.compressed) == true){
            return proof;
        }
        return CommitmentProof.Data({
            exist: ExistenceProof.nil(),
            nonexist: NonExistenceProof.nil(),
            batch: BatchProof.Data({
                entries: decompress(proof.compressed)
            }),
            compressed: CompressedBatchProof.nil()
        });
    }

    function decompress(CompressedBatchProof.Data memory proof) private pure returns(BatchEntry.Data[] memory) {
        BatchEntry.Data[] memory entries = new BatchEntry.Data[](proof.entries.length);
        for(uint i = 0; i < proof.entries.length; i++) {
            entries[i] = decompressEntry(proof.entries[i], proof.lookup_inners);
        }
        return entries;
    }

    function decompressEntry(
        CompressedBatchEntry.Data memory entry,
        InnerOp.Data[] memory lookup
    ) private pure returns(BatchEntry.Data memory) {
        //CompressedExistenceProof.isNil does not work
        if (CompressedExistenceProof._empty(entry.exist) == false) {
            return BatchEntry.Data({
                exist: decompressExist(entry.exist, lookup),
                nonexist: NonExistenceProof.nil()
            });
        }
        return BatchEntry.Data({
            exist: ExistenceProof.nil(),
            nonexist: NonExistenceProof.Data({
                key: entry.nonexist.key,
                left: decompressExist(entry.nonexist.left, lookup),
                right: decompressExist(entry.nonexist.right, lookup)
            })
        });
    }

    function decompressExist(
        CompressedExistenceProof.Data memory proof,
        InnerOp.Data[] memory lookup
    ) private pure returns(ExistenceProof.Data memory) {
        if (CompressedExistenceProof._empty(proof)) {
            return ExistenceProof.nil();
        }
        ExistenceProof.Data memory decoProof = ExistenceProof.Data({
            key: proof.key,
            value: proof.value,
            leaf: proof.leaf,
            path : new InnerOp.Data[](proof.path.length)
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
