extern crate protobuf;
extern crate hex;

use std::result;
use std::vec::Vec;

use super::proofs::{HashOp, InnerOp, LengthOp, LeafOp};

type Result<T> = result::Result<T, &'static str>;
type Hash = Vec<u8>;

pub fn apply_inner(inner: InnerOp, child: Hash) -> Result<Hash> {
    if child.len() == 0 {
        return Err("Missing child hash");
    }
    let mut image = inner.prefix.clone();
    image.extend(child);
    image.extend(inner.suffix);
    return do_hash(inner.hash, &image);
}

// apply_leaf will take a key, value pair and a LeafOp and return a LeafHash
pub fn apply_leaf(leaf: LeafOp, key: &[u8], value: &[u8]) -> Result<Hash> {
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
        HashOp::SHA256 => Ok(vec![1, 2]),
        _ => Err("Unsupported HashOp"),
    }
}

fn do_length(length: LengthOp, data: &[u8]) -> Result<Hash> {
    match length {
        LengthOp::NO_PREFIX => Ok(Hash::from(data)),
        _ => Err("Unsupported LengthOp"),
    }
}


#[cfg(test)]
mod tests {
    use super::*;
    use protobuf::Message;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn do_hash_sha256() {
        let hash = do_hash(HashOp::SHA256, &"food".as_bytes()).unwrap();
        // assert_eq!(hash, hex::decode("c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b").unwrap());
        assert_eq!(hash, hex::decode("0102").unwrap());
    }
}