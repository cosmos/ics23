from brownie import accounts, OpsMock
from brownie.convert import to_bytes, to_string
import brownie
import pytest
import convert
import base64
import helpers


@pytest.fixture
def ops():
    return accounts[0].deploy(OpsMock)


@pytest.mark.parametrize("name,case", helpers.loadJsonTestCases("../testdata/TestLeafOpData.json"))
def test_leafOp(ops, name, case):
    if "sha-512" in name:
        pytest.skip("sha-512 not supported in solidity")
    opCase = convert.leafOp(case["Op"])
    key = base64.b64decode(case["Key"] or "")
    value = base64.b64decode(case["Value"] or "")
    if case["IsErr"]:
        with brownie.reverts():
            res = ops.applyLeafOp(opCase, key, value)
    else:
        res = ops.applyLeafOp(opCase, key, value)
        expected = base64.b64decode(case["Expected"])
        assert to_bytes(res, "bytes") == expected, name


@pytest.mark.parametrize("name,case", helpers.loadJsonTestCases("../testdata/TestInnerOpData.json"))
def test_inner_op(ops, name, case):
    opCase = convert.innerOp(case["Op"])
    child = base64.b64decode(case["Child"] or "")
    if case["IsErr"]:
        with brownie.reverts():
            res = ops.applyInnerOp(opCase, child)
    else:
        res = ops.applyInnerOp(opCase, child)
        expected = base64.b64decode(case["Expected"])
        assert to_bytes(res, "bytes") == expected, name


@pytest.mark.parametrize("name,case", helpers.loadJsonTestCases("../testdata/TestDoHashData.json"))
def test_do_hash(ops, name, case):
    if "sha512" in name: # sha-512 not supported in solidity
        pytest.skip("sha-512 not supported in solidity")
    preImage = case["Preimage"].encode()
    res = ops.doHash(case["HashOp"], preImage)
    expected = to_bytes(case["ExpectedHash"], "bytes")
    assert to_bytes(res, "bytes") == expected, name


@pytest.mark.parametrize("name,case", helpers.loadJsonTestCases("../testdata/TestCheckLeafData.json"))
def test_check_leaf(ops, name, case):
    leafOp = convert.leafOp(case.get("Leaf", None))
    proofSpec = (convert.leafOp(case.get("Spec", None)), convert.innerSpec(None), 0, 0)
    if case["IsErr"]:
        with brownie.reverts():
            res = ops.checkAgainstLeafOpSpec(leafOp, proofSpec)
    else:
        res = ops.checkAgainstLeafOpSpec(leafOp, proofSpec)
