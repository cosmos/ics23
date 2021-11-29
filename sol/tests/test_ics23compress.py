from brownie import accounts, Compress_UnitTest, Proof_UnitTest
from brownie.convert import to_bytes
import pytest
import specs
import os
import helpers


IAVL_TEST_DIR = "../testdata/iavl"
TENDERMINT_TEST_DIR = "../testdata/tendermint"
cases_fname = ["batch_exist.json", "batch_nonexist.json"]


@pytest.fixture
def compress():
    return accounts[0].deploy(Compress_UnitTest)


@pytest.fixture
def proof():
    return accounts[0].deploy(Proof_UnitTest)


@pytest.mark.parametrize("fname", [os.path.join(IAVL_TEST_DIR, fname) for fname in cases_fname])
@pytest.mark.skip(reason="potential bug in ProtobufRuntime")
def test_iavl_decompress(compress, proof, fname):
    proofSpec = specs.iavl
    with open(fname) as f:
        comProof, refData = helpers.loadBatchFile(f, proof)


@pytest.mark.parametrize("fname", [os.path.join(TENDERMINT_TEST_DIR, fname) for fname in cases_fname])
@pytest.mark.skip(reason="potential bug in ProtobufRuntime")
def test_tendermint_decompress(compress, proof, fname):
    proofSpec = specs.tendermint
    with open(fname) as f:
        comProof, refData = helpers.loadBatchFile(f, proof)
