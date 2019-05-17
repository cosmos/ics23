import * as $protobuf from "protobufjs";
/** Namespace proofs. */
export namespace proofs {
  /** HashOp enum. */
  enum HashOp {
    NO_HASH = 0,
    SHA256 = 1,
    SHA512 = 2,
    KECCAK = 3,
    RIPEMD160 = 4,
    BITCOIN = 5
  }

  /**
   * LengthOp defines how to process the key and value of the LeafOp
   * to include length information. After encoding the length with the given
   * algorithm, the length will be prepended to the key and value bytes.
   * (Each one with it's own encoded length)
   */
  enum LengthOp {
    NO_PREFIX = 0,
    VAR_PROTO = 1,
    VAR_RLP = 2,
    FIXED32_BIG = 3,
    FIXED32_LITTLE = 4,
    FIXED64_BIG = 5,
    FIXED64_LITTLE = 6,
    REQUIRE_32_BYTES = 7,
    REQUIRE_64_BYTES = 8
  }

  /** Properties of an ExistenceProof. */
  interface IExistenceProof {
    /** ExistenceProof key */
    key?: Uint8Array | null;

    /** ExistenceProof value */
    value?: Uint8Array | null;

    /** ExistenceProof leaf */
    leaf?: proofs.ILeafOp | null;

    /** ExistenceProof path */
    path?: proofs.IInnerOp[] | null;
  }

  /**
   * ExistenceProof takes a key and a value and a set of steps to perform on it.
   * The result of peforming all these steps will provide a "root hash", which can
   * be compared to the value in a header.
   *
   * Since it is computationally infeasible to produce a hash collission for any of the used
   * cryptographic hash functions, if someone can provide a series of operations to transform
   * a given key and value into a root hash that matches some trusted root, these key and values
   * must be in the referenced merkle tree.
   *
   * The only possible issue is maliablity in LeafOp, such as providing extra prefix data,
   * which should be controlled by a spec. Eg. with lengthOp as NONE,
   * prefix = FOO, key = BAR, value = CHOICE
   * and
   * prefix = F, key = OOBAR, value = CHOICE
   * would produce the same value.
   *
   * With LengthOp this is tricker but not impossible. Which is why the "leafPrefixEqual" field
   * in the ProofSpec is valuable to prevent this mutability. And why all trees should
   * length-prefix the data before hashing it.
   */
  class ExistenceProof implements IExistenceProof {
    /**
     * Constructs a new ExistenceProof.
     * @param [properties] Properties to set
     */
    constructor(properties?: proofs.IExistenceProof);

    /** ExistenceProof key. */
    public key: Uint8Array;

    /** ExistenceProof value. */
    public value: Uint8Array;

    /** ExistenceProof leaf. */
    public leaf?: proofs.ILeafOp | null;

    /** ExistenceProof path. */
    public path: proofs.IInnerOp[];

    /**
     * Creates a new ExistenceProof instance using the specified properties.
     * @param [properties] Properties to set
     * @returns ExistenceProof instance
     */
    public static create(
      properties?: proofs.IExistenceProof
    ): proofs.ExistenceProof;

    /**
     * Encodes the specified ExistenceProof message. Does not implicitly {@link proofs.ExistenceProof.verify|verify} messages.
     * @param message ExistenceProof message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(
      message: proofs.IExistenceProof,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Encodes the specified ExistenceProof message, length delimited. Does not implicitly {@link proofs.ExistenceProof.verify|verify} messages.
     * @param message ExistenceProof message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(
      message: proofs.IExistenceProof,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Decodes an ExistenceProof message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns ExistenceProof
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(
      reader: $protobuf.Reader | Uint8Array,
      length?: number
    ): proofs.ExistenceProof;

    /**
     * Decodes an ExistenceProof message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns ExistenceProof
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(
      reader: $protobuf.Reader | Uint8Array
    ): proofs.ExistenceProof;

    /**
     * Verifies an ExistenceProof message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): string | null;

    /**
     * Creates an ExistenceProof message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns ExistenceProof
     */
    public static fromObject(object: {
      [k: string]: any;
    }): proofs.ExistenceProof;

    /**
     * Creates a plain object from an ExistenceProof message. Also converts values to other types if specified.
     * @param message ExistenceProof
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(
      message: proofs.ExistenceProof,
      options?: $protobuf.IConversionOptions
    ): { [k: string]: any };

    /**
     * Converts this ExistenceProof to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
  }

  /** Properties of a LeafOp. */
  interface ILeafOp {
    /** LeafOp hash */
    hash?: proofs.HashOp | null;

    /** LeafOp prehashKey */
    prehashKey?: proofs.HashOp | null;

    /** LeafOp prehashValue */
    prehashValue?: proofs.HashOp | null;

    /** LeafOp length */
    length?: proofs.LengthOp | null;

    /** LeafOp prefix */
    prefix?: Uint8Array | null;
  }

  /**
   * LeafOp represents the raw key-value data we wish to prove, and
   * must be flexible to represent the internal transformation from
   * the original key-value pairs into the basis hash, for many existing
   * merkle trees.
   *
   * key and value are passed in. So that the signature of this operation is:
   * leafOp(key, value) -> output
   *
   * To process this, first prehash the keys and values if needed (ANY means no hash in this case):
   * hkey = prehashKey(key)
   * hvalue = prehashValue(value)
   *
   * Then combine the bytes, and hash it
   * output = hash(prefix || length(hkey) || hkey || length(hvalue) || hvalue)
   */
  class LeafOp implements ILeafOp {
    /**
     * Constructs a new LeafOp.
     * @param [properties] Properties to set
     */
    constructor(properties?: proofs.ILeafOp);

    /** LeafOp hash. */
    public hash: proofs.HashOp;

    /** LeafOp prehashKey. */
    public prehashKey: proofs.HashOp;

    /** LeafOp prehashValue. */
    public prehashValue: proofs.HashOp;

    /** LeafOp length. */
    public length: proofs.LengthOp;

    /** LeafOp prefix. */
    public prefix: Uint8Array;

    /**
     * Creates a new LeafOp instance using the specified properties.
     * @param [properties] Properties to set
     * @returns LeafOp instance
     */
    public static create(properties?: proofs.ILeafOp): proofs.LeafOp;

    /**
     * Encodes the specified LeafOp message. Does not implicitly {@link proofs.LeafOp.verify|verify} messages.
     * @param message LeafOp message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(
      message: proofs.ILeafOp,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Encodes the specified LeafOp message, length delimited. Does not implicitly {@link proofs.LeafOp.verify|verify} messages.
     * @param message LeafOp message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(
      message: proofs.ILeafOp,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Decodes a LeafOp message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns LeafOp
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(
      reader: $protobuf.Reader | Uint8Array,
      length?: number
    ): proofs.LeafOp;

    /**
     * Decodes a LeafOp message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns LeafOp
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(
      reader: $protobuf.Reader | Uint8Array
    ): proofs.LeafOp;

    /**
     * Verifies a LeafOp message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): string | null;

    /**
     * Creates a LeafOp message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns LeafOp
     */
    public static fromObject(object: { [k: string]: any }): proofs.LeafOp;

    /**
     * Creates a plain object from a LeafOp message. Also converts values to other types if specified.
     * @param message LeafOp
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(
      message: proofs.LeafOp,
      options?: $protobuf.IConversionOptions
    ): { [k: string]: any };

    /**
     * Converts this LeafOp to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
  }

  /** Properties of an InnerOp. */
  interface IInnerOp {
    /** InnerOp hash */
    hash?: proofs.HashOp | null;

    /** InnerOp prefix */
    prefix?: Uint8Array | null;

    /** InnerOp suffix */
    suffix?: Uint8Array | null;
  }

  /**
   * InnerOp represents a merkle-proof step that is not a leaf.
   * It represents concatenating two children and hashing them to provide the next result.
   *
   * The result of the previous step is passed in, so the signature of this op is:
   * innerOp(child) -> output
   *
   * The result of applying InnerOp should be:
   * output = op.hash(op.prefix || child || op.suffix)
   *
   * where the || operator is concatenation of binary data,
   * and child is the result of hashing all the tree below this step.
   *
   * Any special data, like prepending child with the length, or prepending the entire operation with
   * some value to differentiate from leaf nodes, should be included in prefix and suffix.
   * If either of prefix or suffix is empty, we just treat it as an empty string
   */
  class InnerOp implements IInnerOp {
    /**
     * Constructs a new InnerOp.
     * @param [properties] Properties to set
     */
    constructor(properties?: proofs.IInnerOp);

    /** InnerOp hash. */
    public hash: proofs.HashOp;

    /** InnerOp prefix. */
    public prefix: Uint8Array;

    /** InnerOp suffix. */
    public suffix: Uint8Array;

    /**
     * Creates a new InnerOp instance using the specified properties.
     * @param [properties] Properties to set
     * @returns InnerOp instance
     */
    public static create(properties?: proofs.IInnerOp): proofs.InnerOp;

    /**
     * Encodes the specified InnerOp message. Does not implicitly {@link proofs.InnerOp.verify|verify} messages.
     * @param message InnerOp message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(
      message: proofs.IInnerOp,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Encodes the specified InnerOp message, length delimited. Does not implicitly {@link proofs.InnerOp.verify|verify} messages.
     * @param message InnerOp message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(
      message: proofs.IInnerOp,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Decodes an InnerOp message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns InnerOp
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(
      reader: $protobuf.Reader | Uint8Array,
      length?: number
    ): proofs.InnerOp;

    /**
     * Decodes an InnerOp message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns InnerOp
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(
      reader: $protobuf.Reader | Uint8Array
    ): proofs.InnerOp;

    /**
     * Verifies an InnerOp message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): string | null;

    /**
     * Creates an InnerOp message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns InnerOp
     */
    public static fromObject(object: { [k: string]: any }): proofs.InnerOp;

    /**
     * Creates a plain object from an InnerOp message. Also converts values to other types if specified.
     * @param message InnerOp
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(
      message: proofs.InnerOp,
      options?: $protobuf.IConversionOptions
    ): { [k: string]: any };

    /**
     * Converts this InnerOp to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
  }

  /** Properties of a ProofSpec. */
  interface IProofSpec {
    /** ProofSpec leafSpec */
    leafSpec?: proofs.ILeafOp | null;
  }

  /**
   * ProofSpec defines what the expected parameters are for a given proof type.
   * This can be stored in the client and used to validate any incoming proofs.
   *
   * verify(ProofSpec, Proof) -> Proof | Error
   *
   * As demonstrated in tests, if we don't fix the algorithm used to calculate the
   * LeafHash for a given tree, there are many possible key-value pairs that can
   * generate a given hash (by interpretting the preimage differently).
   * We need this for proper security, requires client knows a priori what
   * tree format server uses. But not in code, rather a configuration object.
   */
  class ProofSpec implements IProofSpec {
    /**
     * Constructs a new ProofSpec.
     * @param [properties] Properties to set
     */
    constructor(properties?: proofs.IProofSpec);

    /** ProofSpec leafSpec. */
    public leafSpec?: proofs.ILeafOp | null;

    /**
     * Creates a new ProofSpec instance using the specified properties.
     * @param [properties] Properties to set
     * @returns ProofSpec instance
     */
    public static create(properties?: proofs.IProofSpec): proofs.ProofSpec;

    /**
     * Encodes the specified ProofSpec message. Does not implicitly {@link proofs.ProofSpec.verify|verify} messages.
     * @param message ProofSpec message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encode(
      message: proofs.IProofSpec,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Encodes the specified ProofSpec message, length delimited. Does not implicitly {@link proofs.ProofSpec.verify|verify} messages.
     * @param message ProofSpec message or plain object to encode
     * @param [writer] Writer to encode to
     * @returns Writer
     */
    public static encodeDelimited(
      message: proofs.IProofSpec,
      writer?: $protobuf.Writer
    ): $protobuf.Writer;

    /**
     * Decodes a ProofSpec message from the specified reader or buffer.
     * @param reader Reader or buffer to decode from
     * @param [length] Message length if known beforehand
     * @returns ProofSpec
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decode(
      reader: $protobuf.Reader | Uint8Array,
      length?: number
    ): proofs.ProofSpec;

    /**
     * Decodes a ProofSpec message from the specified reader or buffer, length delimited.
     * @param reader Reader or buffer to decode from
     * @returns ProofSpec
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    public static decodeDelimited(
      reader: $protobuf.Reader | Uint8Array
    ): proofs.ProofSpec;

    /**
     * Verifies a ProofSpec message.
     * @param message Plain object to verify
     * @returns `null` if valid, otherwise the reason why it is not
     */
    public static verify(message: { [k: string]: any }): string | null;

    /**
     * Creates a ProofSpec message from a plain object. Also converts values to their respective internal types.
     * @param object Plain object
     * @returns ProofSpec
     */
    public static fromObject(object: { [k: string]: any }): proofs.ProofSpec;

    /**
     * Creates a plain object from a ProofSpec message. Also converts values to other types if specified.
     * @param message ProofSpec
     * @param [options] Conversion options
     * @returns Plain object
     */
    public static toObject(
      message: proofs.ProofSpec,
      options?: $protobuf.IConversionOptions
    ): { [k: string]: any };

    /**
     * Converts this ProofSpec to JSON.
     * @returns JSON object
     */
    public toJSON(): { [k: string]: any };
  }
}
