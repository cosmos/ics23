// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {Ics23CommitmentProof, Ics23ExistenceProof, Ics23BatchEntry} from "./proofs.sol";
import {Compress} from "./ics23Compress.sol";


contract CompressMock {
    function decompress(Ics23CommitmentProof.Data memory proof) public pure returns(Ics23CommitmentProof.Data memory) {
        (Ics23CommitmentProof.Data memory proof, Compress.DecompressEntryError erCode) = Compress.decompress(proof);
        require(erCode == Compress.DecompressEntryError.None, "CompressMock"); // dev: expand this require to check error code
        // nil Data do not parse well when trying to bring it out of the EVM
        Ics23CommitmentProof.Data memory retV;
        if (Ics23CommitmentProof.isNil(proof)) return retV;
        if (Ics23ExistenceProof.isNil(proof.exist)) {
            retV.batch.entries = new Ics23BatchEntry.Data[](proof.batch.entries.length);
            for(uint i = 0; i < proof.batch.entries.length; i++) {
                if (Ics23ExistenceProof.isNil(proof.batch.entries[i].exist)) {
                    retV.batch.entries[i].nonexist = proof.batch.entries[i].nonexist;
                } else {
                    retV.batch.entries[i].exist = proof.batch.entries[i].exist;
                }
            }
        }
        return retV;
    }
}
