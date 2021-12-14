from brownie import accounts, ICS23Mock, ProofMock
from brownie.convert import to_bytes
import pytest
import specs
import os
import helpers


IAVL_TEST_DIR = "../testdata/iavl"
TENDERMINT_TEST_DIR = "../testdata/tendermint"
cases_fname = ["exist_left.json","exist_right.json", "exist_middle.json", "nonexist_left.json", "nonexist_right.json", "nonexist_middle.json"]


@pytest.fixture
def ics23():
    return accounts[0].deploy(ICS23Mock)


@pytest.fixture
def proof():
    return accounts[0].deploy(ProofMock)


@pytest.mark.parametrize("fname", [os.path.join(IAVL_TEST_DIR, fname) for fname in cases_fname])
def test_iavl_vectors(ics23, proof, fname):
    proofSpec = specs.iavl
    with open(fname) as f:
        comProof, refData = helpers.loadFile(f, proof)
        calculatedRoot = to_bytes(proof.calculateCommitmentProofRoot(comProof), "bytes")
        assert refData["rootHash"] == calculatedRoot, completeFname
        if not refData["value"]:
            ics23.verifyNonMembership(proofSpec, refData["rootHash"], comProof, refData["key"])
        else:
            ics23.verifyMembership(proofSpec, refData["rootHash"], comProof, refData["key"], refData["value"])


@pytest.mark.parametrize("fname", [os.path.join(TENDERMINT_TEST_DIR, fname) for fname in cases_fname])
def test_tendermint_vectors(ics23, proof, fname):
    proofSpec = specs.tendermint
    with open(fname) as f:
        comProof, refData = helpers.loadFile(f, proof)
        calculatedRoot = to_bytes(proof.calculateCommitmentProofRoot(comProof), "bytes")
        assert refData["rootHash"] == calculatedRoot, completeFname
        if not refData["value"]:
            ics23.verifyNonMembership(proofSpec, refData["rootHash"], comProof, refData["key"])
        else:
            ics23.verifyMembership(proofSpec, refData["rootHash"], comProof, refData["key"], refData["value"])
