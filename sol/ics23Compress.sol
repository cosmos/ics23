// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;
import {InnerOp, ExistenceProof, NonExistenceProof, CommitmentProof, CompressedBatchEntry, CompressedBatchProof, CompressedExistenceProof, BatchEntry, BatchProof} from "./proofs.sol";
import {SafeCast} from "@openzeppelin/contracts@4.2.0/utils/math/SafeCast.sol";

library Compress {
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

    // private 
    function decompress(CompressedBatchProof.Data memory proof) private pure returns(BatchEntry.Data[] memory) {
        BatchEntry.Data[] memory entries = new BatchEntry.Data[](proof.entries.length);
        for(uint i = 0; i < proof.entries.length; i++) {
            entries[i] = decompressEntry(proof.entries[i], proof.lookup_inners);
        }
        return entries;
    }

    function decompressEntry(CompressedBatchEntry.Data memory entry, InnerOp.Data[] memory lookup) private pure returns(BatchEntry.Data memory) {
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

    function decompressExist(CompressedExistenceProof.Data memory proof, InnerOp.Data[] memory lookup) private pure returns(ExistenceProof.Data memory) {
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
            require(proof.path[i] >= 0, "proof.path < 0");
            uint step = SafeCast.toUint256(proof.path[i]);
            require(step < lookup.length, "step >= lookup.length");
            decoProof.path[i] = lookup[step];
        }
        return decoProof;
    }
}



contract Compress_UnitTest {
    function decompress(CommitmentProof.Data memory proof) public pure returns(CommitmentProof.Data memory) {
        CommitmentProof.Data memory result = Compress.decompress(proof);
        // abi encoder/decoder does not like XXXProof.nil() functions and the .Data struct they return
        if (CommitmentProof.isNil(result)) {
            CommitmentProof.Data memory tmp;
            return tmp;
        }
        if (ExistenceProof.isNil(result.exist)) {
            ExistenceProof.Data memory tmp;
            result.exist = tmp;
        }
        if (NonExistenceProof.isNil(result.nonexist)) {
            NonExistenceProof.Data memory tmp;
            result.nonexist = tmp;
        }
        if (CompressedBatchProof.isNil(result.compressed)) {
            CompressedBatchProof.Data memory tmp;
            result.compressed = tmp;
        }
        for (uint i = 0; i < result.batch.entries.length; i++) {
            if (ExistenceProof.isNil(result.batch.entries[i].exist)) {
                ExistenceProof.Data memory tmp;
                result.batch.entries[i].exist = tmp;
            }
            if (NonExistenceProof.isNil(result.batch.entries[i].nonexist)){
                NonExistenceProof.Data memory tmp;
                result.batch.entries[i].nonexist = tmp;
            }
        }
        return result;
    }
}
