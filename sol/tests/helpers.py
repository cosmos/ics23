import json
from brownie.convert import to_bytes


def loadFile(f, contract):
    case = json.load(f)
    refData = {
            "rootHash": to_bytes(case.get("root", b''), "bytes"),
            "key": to_bytes(case.get("key", b''), "bytes"),
            "value": None if case.get("value", "") == "" else to_bytes(case["value"], "bytes"),
            }
    proof = contract.protobufDecodeCommitmentProof(to_bytes(case.get("proof", b''), "bytes"))
    return proof, refData


def loadBatchFile(f, contract):
    case = json.load(f)
    refData = {
            "rootHash": to_bytes(case.get("root", b''), "bytes"),
            "items": [(to_bytes(item["key"], "bytes"), 
                None if item.get("value", "") == "" else to_bytes(item["value"], "bytes"))
                for item in case["items"]
                ]
            }
    print(to_bytes(case.get("proof", b''), "bytes"))
    proof = contract.protobufDecodeCommitmentProof(to_bytes(case.get("proof", b''), "bytes"))
    return proof, refData


def loadJsonTestCases(fname):
    with open(fname) as f:
        testcases = json.load(f)
        cases = [ (name, testcases[name]) for name in testcases]
        return cases
