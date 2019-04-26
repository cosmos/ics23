import { proofs } from "./generated/codecimpl";

import { fromHex, toAscii } from "./helpers";
import { applyInner, applyLeaf, doHash } from "./ops";

describe("doHash", () => {
  it("sha256 hashes food", () => {
    // echo -n food | sha256sum
    const hash = doHash(proofs.HashOp.SHA256, toAscii("food"));
    expect(hash).toEqual(
      fromHex(
        "c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b"
      )
    );
  });

  it("ripemd160 hashes food", () => {
    // echo -n food | openssl dgst -rmd160 -hex | cut -d' ' -f2
    const hash = doHash(proofs.HashOp.RIPEMD160, toAscii("food"));
    expect(hash).toEqual(fromHex("b1ab9988c7c7c5ec4b2b291adfeeee10e77cdd46"));
  });

  it("'bitcoin' hashes food", () => {
    // echo -n c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b | xxd -r -p | openssl dgst -rmd160 -hex
    const hash = doHash(proofs.HashOp.BITCOIN, toAscii("food"));
    expect(hash).toEqual(fromHex("0bcb587dfb4fc10b36d57f2bba1878f139b75d24"));
  });
});

describe("applyLeaf", () => {
  it("hashes foobar", () => {
    const op: proofs.ILeafOp = { hash: proofs.HashOp.SHA256 };
    const key = toAscii("foo");
    const value = toAscii("bar");
    // echo -n foobar | sha256sum
    const expected = fromHex(
      "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes foobaz with sha-512", () => {
    const op: proofs.ILeafOp = { hash: proofs.HashOp.SHA512 };
    const key = toAscii("foo");
    const value = toAscii("baz");
    // echo -n foobaz | sha512sum
    const expected = fromHex(
      "4f79f191298ec7461d60136c60f77c2ae8ddd85dbf6168bb925092d51bfb39b559219b39ae5385ba04946c87f64741385bef90578ea6fe6dac85dbf7ad3f79e1"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes foobar (different breakpoint)", () => {
    const op: proofs.ILeafOp = { hash: proofs.HashOp.SHA256 };
    const key = toAscii("f");
    const value = toAscii("oobar");
    // echo -n foobar | sha256sum
    const expected = fromHex(
      "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes with length prefix", () => {
    const op: proofs.ILeafOp = {
      hash: proofs.HashOp.SHA256,
      length: proofs.LengthOp.VAR_PROTO
    };
    // echo -n food | xxd -ps
    const key = toAscii("food"); // 04666f6f64
    const value = toAscii("some longer text"); // 10736f6d65206c6f6e6765722074657874
    // echo -n 04666f6f6410736f6d65206c6f6e6765722074657874 | xxd -r -p | sha256sum -b
    const expected = fromHex(
      "b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes with prehash and length prefix", () => {
    const op: proofs.ILeafOp = {
      hash: proofs.HashOp.SHA256,
      length: proofs.LengthOp.VAR_PROTO,
      prehashValue: proofs.HashOp.SHA256
    };
    const key = toAscii("food"); // 04666f6f64
    // echo -n yet another long string | sha256sum
    const value = toAscii("yet another long string"); // 20a48c2d4f67b9f80374938535285ed285819d8a5a8fc1fccd1e3244e437cf290d
    // echo -n 04666f6f6420a48c2d4f67b9f80374938535285ed285819d8a5a8fc1fccd1e3244e437cf290d | xxd -r -p | sha256sum
    const expected = fromHex(
      "87e0483e8fb624aef2e2f7b13f4166cda485baa8e39f437c83d74c94bedb148f"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("requires key", () => {
    const op: proofs.ILeafOp = {
      hash: proofs.HashOp.SHA256
    };
    const key = toAscii("food");
    const value = toAscii("");
    expect(() => applyLeaf(op, key, value)).toThrow();
  });

  it("requires value", () => {
    const op: proofs.ILeafOp = {
      hash: proofs.HashOp.SHA256
    };
    const key = toAscii("");
    const value = toAscii("time");
    expect(() => applyLeaf(op, key, value)).toThrow();
  });
});

describe("applyInner", () => {
  it("hash child with prefix and suffix", () => {
    const op: proofs.IInnerOp = {
      hash: proofs.HashOp.SHA256,
      prefix: fromHex("0123456789"),
      suffix: fromHex("deadbeef")
    };
    const child = fromHex("00cafe00");
    // echo -n 012345678900cafe00deadbeef | xxd -r -p | sha256sum
    const expected = fromHex(
      "0339f76086684506a6d42a60da4b5a719febd4d96d8b8d85ae92849e3a849a5e"
    );
    expect(applyInner(op, child)).toEqual(expected);
  });

  it("requies child", () => {
    const op: proofs.IInnerOp = {
      hash: proofs.HashOp.SHA256,
      prefix: fromHex("0123456789"),
      suffix: fromHex("deadbeef")
    };
    expect(() => applyInner(op, fromHex(""))).toThrow();
  });

  it("hash child with only prefix", () => {
    const op: proofs.IInnerOp = {
      hash: proofs.HashOp.SHA256,
      prefix: fromHex("00204080a0c0e0")
    };
    const child = fromHex("ffccbb997755331100");
    // echo -n 00204080a0c0e0ffccbb997755331100 | xxd -r -p | sha256sum
    const expected = fromHex(
      "45bece1678cf2e9f4f2ae033e546fc35a2081b2415edcb13121a0e908dca1927"
    );
    expect(applyInner(op, child)).toEqual(expected);
  });

  it("hash child with only suffix", () => {
    const op: proofs.IInnerOp = {
      hash: proofs.HashOp.SHA256,
      suffix: toAscii(" just kidding!")
    };
    const child = toAscii("this is a sha256 hash, really....");
    // echo -n 'this is a sha256 hash, really.... just kidding!'  | sha256sum
    const expected = fromHex(
      "79ef671d27e42a53fba2201c1bbc529a099af578ee8a38df140795db0ae2184b"
    );
    expect(applyInner(op, child)).toEqual(expected);
  });
});
