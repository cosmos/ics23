import { ics23 } from "./generated/codecimpl";

export function compress(
  proof: ics23.ICommitmentProof
): ics23.ICommitmentProof {
  if (!proof.batch) {
    return proof;
  }
  return { compressed: compressBatch(proof.batch) };
}

export function decompress(
  proof: ics23.ICommitmentProof
): ics23.ICommitmentProof {
  if (!proof.compressed) {
    return proof;
  }
  return { batch: decompressBatch(proof.compressed) };
}

function compressBatch(proof: ics23.IBatchProof): ics23.ICompressedBatchProof {
  const centries: ics23.ICompressedBatchEntry[] = [];
  const lookup: ics23.IInnerOp[] = [];
  const registry = new Map<Uint8Array, number>();

  for (const entry of proof.entries!) {
    if (entry.exist) {
      const centry = { exist: compressExist(entry.exist, lookup, registry) };
      centries.push(centry);
    } else if (entry.nonexist) {
      const non = entry.nonexist;
      const centry = {
        nonexist: {
          key: non.key,
          left: compressExist(non.left, lookup, registry),
          right: compressExist(non.right, lookup, registry),
        },
      };
      centries.push(centry);
    } else {
      throw new Error("Unexpected batch entry during compress");
    }
  }

  return {
    entries: centries,
    lookupInners: lookup,
  };
}

function compressExist(
  exist: ics23.IExistenceProof | null | undefined,
  lookup: ics23.IInnerOp[],
  registry: Map<Uint8Array, number>
): ics23.ICompressedExistenceProof | undefined {
  if (!exist) {
    return undefined;
  }

  const path = exist.path!.map((inner) => {
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
    path,
  };
}

function decompressBatch(
  proof: ics23.ICompressedBatchProof
): ics23.IBatchProof {
  const lookup = proof.lookupInners!;
  const entries = proof.entries!.map((comp) => {
    if (comp.exist) {
      return { exist: decompressExist(comp.exist, lookup) };
    } else if (comp.nonexist) {
      const non = comp.nonexist;
      return {
        nonexist: {
          key: non.key,
          left: decompressExist(non.left, lookup),
          right: decompressExist(non.right, lookup),
        },
      };
    } else {
      throw new Error("Unexpected batch entry during compress");
    }
  });
  return {
    entries,
  };
}

function decompressExist(
  exist: ics23.ICompressedExistenceProof | null | undefined,
  lookup: readonly ics23.IInnerOp[]
): ics23.IExistenceProof | undefined {
  if (!exist) {
    return undefined;
  }
  const { key, value, leaf, path } = exist;
  const newPath = (path || []).map((idx) => lookup[idx]);
  return { key, value, leaf, path: newPath };
}
