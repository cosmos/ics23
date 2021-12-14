// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {CommitmentProof} from "./proofs.sol";
import {Compress} from "./ics23Compress.sol";


contract CompressMock {
    function decompress(CommitmentProof.Data memory proof) public pure returns(CommitmentProof.Data memory) {
        return Compress.decompress(proof);
    }
}
