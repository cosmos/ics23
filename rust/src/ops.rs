use anyhow::{bail, ensure};
use core::convert::TryInto;

use crate::helpers::{Hash, Result};
use crate::host_functions::HostFunctionsProvider;
use crate::ics23::{HashOp, InnerOp, LeafOp, LengthOp};

pub fn apply_inner<H: HostFunctionsProvider>(inner: &InnerOp, child: &[u8]) -> Result<Hash> {
    ensure!(!child.is_empty(), "Missing child hash");
    let mut image = inner.prefix.clone();
    image.extend(child);
    image.extend(&inner.suffix);
    Ok(do_hash::<H>(inner.hash(), &image))
}

// apply_leaf will take a key, value pair and a LeafOp and return a LeafHash
pub fn apply_leaf<H: HostFunctionsProvider>(
    leaf: &LeafOp,
    key: &[u8],
    value: &[u8],
) -> Result<Hash> {
    let mut hash = leaf.prefix.clone();
    let prekey = prepare_leaf_data::<H>(leaf.prehash_key(), leaf.length(), key)?;
    hash.extend(prekey);
    let preval = prepare_leaf_data::<H>(leaf.prehash_value(), leaf.length(), value)?;
    hash.extend(preval);
    Ok(do_hash::<H>(leaf.hash(), &hash))
}

fn prepare_leaf_data<H: HostFunctionsProvider>(
    prehash: HashOp,
    length: LengthOp,
    data: &[u8],
) -> Result<Hash> {
    ensure!(!data.is_empty(), "Input to prepare_leaf_data missing");
    let h = do_hash::<H>(prehash, data);
    do_length(length, &h)
}

fn do_hash<H: HostFunctionsProvider>(hash: HashOp, data: &[u8]) -> Hash {
    match hash {
        HashOp::NoHash => Hash::from(data),
        HashOp::Sha256 => Hash::from(H::sha2_256(data)),
        HashOp::Sha512 => Hash::from(H::sha2_512(data)),
        HashOp::Keccak => Hash::from(H::sha3_512(data)),
        HashOp::Ripemd160 => Hash::from(H::ripemd160(data)),
        HashOp::Bitcoin => Hash::from(H::ripemd160(&H::sha2_256(data)[..])),
        HashOp::Sha512256 => Hash::from(H::sha2_512_truncated(data)),
        HashOp::Blake3 => Hash::from(H::blake3hash(data)),
    }
}

pub(crate) fn do_length(length: LengthOp, data: &[u8]) -> Result<Hash> {
    match length {
        LengthOp::NoPrefix => {}
        LengthOp::Require32Bytes => ensure!(data.len() == 32, "Invalid length"),
        LengthOp::Require64Bytes => ensure!(data.len() == 64, "Invalid length"),
        LengthOp::VarProto => {
            let mut len = proto_len(data.len())?;
            len.extend(data);
            return Ok(len);
        }
        LengthOp::Fixed32Little => {
            let mut len = (data.len() as u32).to_le_bytes().to_vec();
            len.extend(data);
            return Ok(len);
        }
        _ => bail!("Unsupported LengthOp {:?}", length),
    }
    // if we don't error above or return custom string, just return item untouched (common case)
    Ok(data.to_vec())
}

pub(crate) fn proto_len(length: usize) -> Result<Hash> {
    let size: u64 = length.try_into().map_err(anyhow::Error::msg)?;
    let mut len = Hash::new();
    prost::encoding::encode_varint(size, &mut len);
    Ok(len)
}

#[cfg(test)]
mod tests {
    use super::*;

    use crate::host_functions::host_functions_impl::HostFunctionsManager;

    use alloc::vec;
    use alloc::vec::Vec;

    fn decode(input: &str) -> Vec<u8> {
        hex::decode(input).unwrap()
    }

    #[test]
    fn hashing_food() {
        let hash = do_hash::<HostFunctionsManager>(HashOp::NoHash, b"food");
        assert!(hash == hex::decode("666f6f64").unwrap(), "no hash fails");

        let hash = do_hash::<HostFunctionsManager>(HashOp::Sha256, b"food");
        assert!(
            hash == decode("c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b"),
            "sha256 hash fails"
        );

        let hash = do_hash::<HostFunctionsManager>(HashOp::Sha512, b"food");
        assert!(hash == decode("c235548cfe84fc87678ff04c9134e060cdcd7512d09ed726192151a995541ed8db9fda5204e72e7ac268214c322c17787c70530513c59faede52b7dd9ce64331"), "sha512 hash fails");

        let hash = do_hash::<HostFunctionsManager>(HashOp::Ripemd160, b"food");
        assert!(
            hash == decode("b1ab9988c7c7c5ec4b2b291adfeeee10e77cdd46"),
            "ripemd160 hash fails"
        );

        let hash = do_hash::<HostFunctionsManager>(HashOp::Bitcoin, b"food");
        assert!(
            hash == decode("0bcb587dfb4fc10b36d57f2bba1878f139b75d24"),
            "bitcoin hash fails"
        );

        let hash = do_hash::<HostFunctionsManager>(HashOp::Sha512256, b"food");
        assert!(
            hash == decode("5b3a452a6acbf1fc1e553a40c501585d5bd3cca176d562e0a0e19a3c43804e88"),
            "sha512/256 hash fails"
        );

        let hash = do_hash::<HostFunctionsManager>(HashOp::Blake3, b"food");
        assert!(hash == decode("f775a8ccf8cb78cd1c63ade4e9802de4ead836b36cea35242accf31d2c6a3697"),);
    }

    #[test]
    fn length_prefix() {
        let prefixed = do_length(LengthOp::NoPrefix, b"food").unwrap();
        assert!(prefixed == decode("666f6f64"), "no prefix modifies data");

        let prefixed = do_length(LengthOp::VarProto, b"food").unwrap();
        assert!(
            prefixed == decode("04666f6f64"),
            "proto prefix returned {}",
            hex::encode(&prefixed),
        );

        let prefixed = do_length(LengthOp::Fixed32Little, b"food").unwrap();
        assert!(
            prefixed == decode("04000000666f6f64"),
            "proto prefix returned {}",
            hex::encode(&prefixed),
        );
    }

    #[test]
    fn apply_leaf_hash() {
        let leaf = LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: 0,
            prefix: vec![],
        };
        let key = b"foo";
        let val = b"bar";
        let hash = decode("c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2");
        assert!(
            hash == apply_leaf::<HostFunctionsManager>(&leaf, key, val).unwrap(),
            "unexpected leaf hash"
        );
    }

    #[test]
    fn apply_leaf_hash_512() {
        let leaf = LeafOp {
            hash: HashOp::Sha512.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: 0,
            prefix: vec![],
        };
        let key = b"f";
        let val = b"oobaz";
        let hash = decode("4f79f191298ec7461d60136c60f77c2ae8ddd85dbf6168bb925092d51bfb39b559219b39ae5385ba04946c87f64741385bef90578ea6fe6dac85dbf7ad3f79e1");
        assert!(
            hash == apply_leaf::<HostFunctionsManager>(&leaf, key, val).unwrap(),
            "unexpected leaf hash"
        );
    }

    #[test]
    fn apply_leaf_hash_length() {
        let leaf = LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: 0,
            length: LengthOp::VarProto.into(),
            prefix: vec![],
        };
        let key = b"food";
        let val = b"some longer text";
        let hash = decode("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265");
        assert!(
            hash == apply_leaf::<HostFunctionsManager>(&leaf, key, val).unwrap(),
            "unexpected leaf hash"
        );
    }

    #[test]
    fn apply_leaf_prehash_length() {
        let leaf = LeafOp {
            hash: HashOp::Sha256.into(),
            prehash_key: 0,
            prehash_value: HashOp::Sha256.into(),
            length: LengthOp::VarProto.into(),
            prefix: vec![],
        };
        let key = b"food";
        let val = b"yet another long string";
        let hash = decode("87e0483e8fb624aef2e2f7b13f4166cda485baa8e39f437c83d74c94bedb148f");
        assert!(
            hash == apply_leaf::<HostFunctionsManager>(&leaf, key, val).unwrap(),
            "unexpected leaf hash"
        );
    }

    #[test]
    fn apply_inner_prefix_suffix() {
        let inner = InnerOp {
            hash: HashOp::Sha256.into(),
            prefix: decode("0123456789"),
            suffix: decode("deadbeef"),
        };
        let child = decode("00cafe00");

        // echo -n 012345678900cafe00deadbeef | xxd -r -p | sha256sum
        let expected = decode("0339f76086684506a6d42a60da4b5a719febd4d96d8b8d85ae92849e3a849a5e");
        assert!(
            expected == apply_inner::<HostFunctionsManager>(&inner, &child).unwrap(),
            "unexpected inner hash"
        );
    }

    #[test]
    fn apply_inner_prefix_only() {
        let inner = InnerOp {
            hash: HashOp::Sha256.into(),
            prefix: decode("00204080a0c0e0"),
            suffix: vec![],
        };
        let child = decode("ffccbb997755331100");

        // echo -n 00204080a0c0e0ffccbb997755331100 | xxd -r -p | sha256sum
        let expected = decode("45bece1678cf2e9f4f2ae033e546fc35a2081b2415edcb13121a0e908dca1927");
        assert!(
            expected == apply_inner::<HostFunctionsManager>(&inner, &child).unwrap(),
            "unexpected inner hash"
        );
    }

    #[test]
    fn apply_inner_suffix_only() {
        let inner = InnerOp {
            hash: HashOp::Sha256.into(),
            prefix: vec![],
            suffix: b" just kidding!".to_vec(),
        };
        let child = b"this is a sha256 hash, really....".to_vec();

        // echo -n 'this is a sha256 hash, really.... just kidding!'  | sha256sum
        let expected = decode("79ef671d27e42a53fba2201c1bbc529a099af578ee8a38df140795db0ae2184b");
        assert!(
            expected == apply_inner::<HostFunctionsManager>(&inner, &child).unwrap(),
            "unexpected inner hash"
        );
    }
}
