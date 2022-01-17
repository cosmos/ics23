// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;

import {Ics23CommitmentProof} from "./proofs.sol";
import {Compress} from "./ics23Compress.sol";


contract CompressMock {
    function decompress(Ics23CommitmentProof.Data memory proof) public pure returns(Ics23CommitmentProof.Data memory) {
        return Compress.decompress(proof);
    }
}
