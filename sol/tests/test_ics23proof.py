from brownie import accounts, Proof_UnitTest
from brownie.convert import to_bytes, to_string
import brownie
import pytest
import convert
import base64
import helpers


@pytest.fixture
def proof():
    return accounts[0].deploy(Proof_UnitTest)


@pytest.mark.parametrize("name,case", helpers.loadJsonTestCases("../testdata/TestExistenceProofData.json"))
def test_existence_proof(proof, name, case):
    exProof = convert.existenceProof(case["Proof"])
    if case["IsErr"]:
        with brownie.reverts():
            res = proof.calculateExistenceProofRoot(exProof)
    else:
        res = proof.calculateExistenceProofRoot(exProof)
        expected = base64.b64decode(case["Expected"])
        assert to_bytes(res, "bytes") == expected, name


@pytest.mark.parametrize("name,case", helpers.loadJsonTestCases("../testdata/TestCheckAgainstSpecData.json"))
def test_checkagainstspec(proof, name, case):
    exProof = convert.existenceProof(case["Proof"])
    spec = convert.proofSpec(case["Spec"])
    if case["IsErr"]:
        with brownie.reverts():
            proof.checkAgainstSpec(exProof, spec)
    else:
        proof.checkAgainstSpec(exProof, spec)
