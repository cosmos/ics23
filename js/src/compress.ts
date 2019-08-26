import { ics23 } from "./generated/codecimpl";

export function compress(
  proof: ics23.ICommitmentProof
): ics23.ICommitmentProof {
  if (!proof.batch) {
    return proof;
  }
  // TODO
  return { compressed: compress_batch(proof.batch) };
}

export function decompress(
  proof: ics23.ICommitmentProof
): ics23.ICommitmentProof {
  if (!proof.compressed) {
    return proof;
  }
  // TODO
  return { batch: decompress_batch(proof.compressed) };
}

// tslint:disable:readonly-array
function compress_batch(proof: ics23.IBatchProof): ics23.ICompressedBatchProof {
  const centries: ics23.ICompressedBatchEntry[] = [];
  const lookup: ics23.IInnerOp[] = [];
  const registry = new Map<Uint8Array, number>();

  for (const entry of proof.entries!) {
    if (!!entry.exist) {
      const centry = { exist: compress_exist(entry.exist, lookup, registry) };
      centries.push(centry);
    } else if (!!entry.nonexist) {
      const non = entry.nonexist;
      const centry = {
        nonexist: {
          key: non.key,
          left: compress_exist(non.left, lookup, registry),
          right: compress_exist(non.right, lookup, registry)
        }
      };
      centries.push(centry);
    } else {
      throw new Error("Unexpected batch entry during compress");
    }
  }

  return {
    entries: centries,
    lookupInners: lookup
  };
}

function compress_exist(
  exist: ics23.IExistenceProof | null | undefined,
  lookup: ics23.IInnerOp[],
  registry: Map<Uint8Array, number>
): ics23.ICompressedExistenceProof | undefined {
  if (!exist) {
    return undefined;
  }

  const path = exist.path!.map(inner => {
    const sig = ics23.InnerOp.encode(inner).finish();
    let idx = registry.get(sig);
    if (idx === undefined) {
      idx = lookup.length;
      lookup.push(inner);
      registry.set(sig, idx);
    }
    return idx;
  });

  return {
    key: exist.key,
    value: exist.value,
    leaf: exist.leaf,
    path
  };
}

function decompress_batch(
  proof: ics23.ICompressedBatchProof
): ics23.IBatchProof {
  const lookup = proof.lookupInners!;
  const entries = proof.entries!.map(comp => {
    if (!!comp.exist) {
      return { exist: decompress_exist(comp.exist, lookup) };
    } else if (!!comp.nonexist) {
      const non = comp.nonexist;
      return {
        nonexist: {
          key: non.key,
          left: decompress_exist(non.left, lookup),
          right: decompress_exist(non.right, lookup)
        }
      };
    } else {
      throw new Error("Unexpected batch entry during compress");
    }
  });
  return {
    entries
  };
}

function decompress_exist(
  exist: ics23.ICompressedExistenceProof | null | undefined,
  lookup: ReadonlyArray<ics23.IInnerOp>
): ics23.IExistenceProof | undefined {
  if (!exist) {
    return undefined;
  }
  const { key, value, leaf, path } = exist;
  const newPath = (path || []).map(idx => lookup[idx]);
  return { key, value, leaf, path: newPath };
}
