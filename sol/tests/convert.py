from brownie.convert import to_bytes
import base64

def leafOp(leafOp):
    if leafOp is None:
        return (0, 0, 0, 0, b'')
    return (leafOp.get("hash", 0),
            leafOp.get("prehash_key", 0),
            leafOp.get("prehash_value", 0),
            leafOp.get("length", 0),
            base64.b64decode(leafOp.get("prefix", ""))
            )


def innerOp(innerOp):
    if innerOp is None:
        return (0, b'', b'')
    return (innerOp.get("hash", 0),
            base64.b64decode(innerOp.get("prefix", "")),
            base64.b64decode(innerOp.get("suffix", "")),
            )


def existenceProof(exProof):
    return (base64.b64decode(exProof.get("key", "")),
            base64.b64decode(exProof.get("value", "")),
            leafOp(exProof.get("leaf")),
            [innerOp(op_elem) for op_elem in exProof.get("path", [])],
            )


def innerSpec(innerSpec):
    if innerSpec is None:
        return([], 0, 0, 0, b'', 0)
    return ([el for el in innerSpec.get("child_order", [])],
            innerSpec.get("child_size", 0),
            innerSpec.get("min_prefix_length", 0),
            innerSpec.get("max_prefix_length", 0),
            base64.b64decode(innerSpec.get("empty_child", "")),
            innerSpec.get("hash", 0),
            )


def proofSpec(proofSpec):
    if proofSpec is None:
        return (leafOp(None), innerSpec(None), 0, 0)
    return (leafOp(proofSpec.get("leaf_spec", None)),
            innerSpec(proofSpec.get("inner_spec", None)),
            proofSpec.get("max_depth", 0),
            proofSpec.get("min_depth", 0),
            )
