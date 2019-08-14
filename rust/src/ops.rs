extern crate protobuf;
extern crate hex;
extern crate sha2;
extern crate sha3;
extern crate ripemd160;

use std::convert::TryInto;
use std::result;
use std::vec::Vec;
use sha2::{Digest, Sha256, Sha512};
use sha3::{Sha3_512};
use ripemd160::{Ripemd160};

use crate::proofs::{HashOp, InnerOp, LengthOp, LeafOp};

pub type Result<T> = result::Result<T, &'static str>;
pub type Hash = Vec<u8>;

pub fn apply_inner(inner: &InnerOp, child: Hash) -> Result<Hash> {
    if child.len() == 0 {
        return Err("Missing child hash");
    }
    let mut image = inner.prefix.clone();
    image.extend(child);
    image.extend(&inner.suffix);
    return do_hash(inner.hash, &image);
}

// apply_leaf will take a key, value pair and a LeafOp and return a LeafHash
pub fn apply_leaf(leaf: &LeafOp, key: &[u8], value: &[u8]) -> Result<Hash> {
    let pkey = prepare_leaf_data(leaf.prehash_key, leaf.length, key)?;
    let pval = prepare_leaf_data(leaf.prehash_value, leaf.length, value)?;
    let mut hash = leaf.prefix.clone();
    hash.extend(pkey);
    hash.extend(pval);
    return do_hash(leaf.hash, &hash);
}

fn prepare_leaf_data(prehash: HashOp, length: LengthOp, data: &[u8]) -> Result<Hash> {
    if data.len() == 0 {
        return Err("Input to prepare_leaf_data missing");
    }
    let h = do_hash(prehash, data)?;
    return do_length(length, &h);
}

fn do_hash(hash: HashOp, data: &[u8]) -> Result<Hash> {
    match hash {
        HashOp::NO_HASH => Ok(Hash::from(data)),
        HashOp::SHA256 => Ok(Hash::from(Sha256::digest(data).as_slice())),
        HashOp::SHA512 => Ok(Hash::from(Sha512::digest(data).as_slice())),
        HashOp::KECCAK => Ok(Hash::from(Sha3_512::digest(data).as_slice())),
        HashOp::RIPEMD160 => Ok(Hash::from(Ripemd160::digest(data).as_slice())),
        HashOp::BITCOIN => Ok(Hash::from(Ripemd160::digest(Sha256::digest(data).as_slice()).as_slice())),
    }
}

fn do_length(length: LengthOp, data: &[u8]) -> Result<Hash> {
    match length {
        LengthOp::NO_PREFIX => Ok(Hash::from(data)),
        LengthOp::REQUIRE_32_BYTES => if data.len() == 32 { Ok(Hash::from(data)) } else { Err("Invalid length") },
        LengthOp::REQUIRE_64_BYTES => if data.len() == 64 { Ok(Hash::from(data)) } else { Err("Invalid length") },
        LengthOp::VAR_PROTO => {
            let mut len = proto_len(data.len());
            len.extend(data);
            Ok(len)
        },
        _ => Err("Unsupported LengthOp"),
    }
}

fn proto_len(length: usize) -> Hash {
    // TODO: remove unwrap
    let size: u64 = length.try_into().unwrap();
    let mut len = Hash::new();
    let mut str = protobuf::CodedOutputStream::vec(&mut len);
    str.write_uint64_no_tag(size).unwrap();
    str.flush().unwrap();
    len
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn hashing_food() {
        let hash = do_hash(HashOp::NO_HASH, &"food".as_bytes()).unwrap();
        assert_eq!(hash, hex::decode("666f6f64").unwrap());

        let hash = do_hash(HashOp::SHA256, &"food".as_bytes()).unwrap();
        assert_eq!(hash, hex::decode("c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b").unwrap());

        let hash = do_hash(HashOp::SHA512, &"food".as_bytes()).unwrap();
        assert_eq!(hash, hex::decode("c235548cfe84fc87678ff04c9134e060cdcd7512d09ed726192151a995541ed8db9fda5204e72e7ac268214c322c17787c70530513c59faede52b7dd9ce64331").unwrap());

        let hash = do_hash(HashOp::RIPEMD160, &"food".as_bytes()).unwrap();
        assert_eq!(hash, hex::decode("b1ab9988c7c7c5ec4b2b291adfeeee10e77cdd46").unwrap());

        let hash = do_hash(HashOp::BITCOIN, &"food".as_bytes()).unwrap();
        assert_eq!(hash, hex::decode("0bcb587dfb4fc10b36d57f2bba1878f139b75d24").unwrap());
    }

    #[test]
    fn length_prefix() {
        let prefixed = do_length(LengthOp::NO_PREFIX, &"food".as_bytes()).unwrap();
        assert_eq!(prefixed, hex::decode("666f6f64").unwrap());

        let prefixed = do_length(LengthOp::VAR_PROTO, &"food".as_bytes()).unwrap();
        assert_eq!(prefixed, hex::decode("04666f6f64").unwrap());
    }

    #[test]
    fn apply_leaf_hash() {
        let mut leaf = LeafOp::new();
        leaf.set_hash(HashOp::SHA256);
        let key = &"foo".as_bytes();
        let val = &"bar".as_bytes();
        let hash = hex::decode("c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2").unwrap();
        assert_eq!(hash, apply_leaf(&leaf, key, val).unwrap());
    }

    #[test]
    fn apply_leaf_hash_512() {
        let mut leaf = LeafOp::new();
        leaf.set_hash(HashOp::SHA512);
        let key = &"f".as_bytes();
        let val = &"oobaz".as_bytes();
        let hash = hex::decode("4f79f191298ec7461d60136c60f77c2ae8ddd85dbf6168bb925092d51bfb39b559219b39ae5385ba04946c87f64741385bef90578ea6fe6dac85dbf7ad3f79e1").unwrap();
        assert_eq!(hash, apply_leaf(&leaf, key, val).unwrap());
    }

    #[test]
    fn apply_leaf_hash_length() {
        let mut leaf = LeafOp::new();
        leaf.set_hash(HashOp::SHA256);
        leaf.set_length(LengthOp::VAR_PROTO);
        let key = &"food".as_bytes();
        let val = &"some longer text".as_bytes();
        let hash = hex::decode("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265").unwrap();
        assert_eq!(hash, apply_leaf(&leaf, key, val).unwrap());
    }

    #[test]
    fn apply_leaf_prehash_length() {
        let mut leaf = LeafOp::new();
        leaf.set_hash(HashOp::SHA256);
        leaf.set_prehash_value(HashOp::SHA256);
        leaf.set_length(LengthOp::VAR_PROTO);
        let key = &"food".as_bytes();
        let val = &"yet another long string".as_bytes();
        let hash = hex::decode("87e0483e8fb624aef2e2f7b13f4166cda485baa8e39f437c83d74c94bedb148f").unwrap();
        assert_eq!(hash, apply_leaf(&leaf, key, val).unwrap());
    }

    #[test]
    fn apply_inner_prefix_suffix() {
        let mut inner = InnerOp::new();
        inner.set_hash(HashOp::SHA256);
        inner.set_prefix(hex::decode("0123456789").unwrap());
        inner.set_suffix(hex::decode("deadbeef").unwrap());
        let child = hex::decode("00cafe00").unwrap();

        // echo -n 012345678900cafe00deadbeef | xxd -r -p | sha256sum
        let expected = hex::decode("0339f76086684506a6d42a60da4b5a719febd4d96d8b8d85ae92849e3a849a5e").unwrap();
        assert_eq!(expected, apply_inner(&inner, child).unwrap());
    }

    #[test]
    fn apply_inner_prefix_only() {
        let mut inner = InnerOp::new();
        inner.set_hash(HashOp::SHA256);
        inner.set_prefix(hex::decode("00204080a0c0e0").unwrap());
        let child = hex::decode("ffccbb997755331100").unwrap();

        // echo -n 00204080a0c0e0ffccbb997755331100 | xxd -r -p | sha256sum
        let expected = hex::decode("45bece1678cf2e9f4f2ae033e546fc35a2081b2415edcb13121a0e908dca1927").unwrap();
        assert_eq!(expected, apply_inner(&inner, child).unwrap());
    }


    #[test]
    fn apply_inner_suffix_only() {
        let mut inner = InnerOp::new();
        inner.set_hash(HashOp::SHA256);
        inner.set_suffix(Hash::from(" just kidding!"));
        let child = Hash::from("this is a sha256 hash, really....");

        // echo -n 'this is a sha256 hash, really.... just kidding!'  | sha256sum
        let expected = hex::decode("79ef671d27e42a53fba2201c1bbc529a099af578ee8a38df140795db0ae2184b").unwrap();
        assert_eq!(expected, apply_inner(&inner, child).unwrap());
    }
}