export function fromHex(hexstring: string): Uint8Array {
  if (hexstring.length % 2 !== 0) {
    throw new Error("hex string length must be a multiple of 2");
  }

  const listOfInts: number[] = [];
  for (let i = 0; i < hexstring.length; i += 2) {
    const hexByteAsString = hexstring.substr(i, 2);
    if (!hexByteAsString.match(/[0-9a-f]{2}/i)) {
      throw new Error("hex string contains invalid characters");
    }
    listOfInts.push(parseInt(hexByteAsString, 16));
  }
  return new Uint8Array(listOfInts);
}

export function toAscii(input: string): Uint8Array {
  const toNums = (str: string): number[] =>
    str.split("").map((x: string) => {
      const charCode = x.charCodeAt(0);
      // 0x00–0x1F control characters
      // 0x20–0x7E printable characters
      // 0x7F delete character
      // 0x80–0xFF out of 7 bit ascii range
      if (charCode < 0x20 || charCode > 0x7e) {
        throw new Error(
          "Cannot encode character that is out of printable ASCII range: " +
            charCode
        );
      }
      return charCode;
    });
  return Uint8Array.from(toNums(input));
}
