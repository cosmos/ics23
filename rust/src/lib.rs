extern crate protobuf;

mod proofs;

#[cfg(test)]
mod tests {
    use crate::proofs;

    use protobuf::Message;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn proto_encode() {
        let mut leaf = proofs::LeafOp::new();
        leaf.set_hash(proofs::HashOp::SHA256);
        leaf.set_prehash_key(proofs::HashOp::SHA256);
        leaf.set_prefix(vec![0, 1, 2].into());
        assert_eq!(leaf.get_hash(), proofs::HashOp::SHA256);
        assert_eq!(leaf.get_prehash_value(), proofs::HashOp::NO_HASH);

        // encode it
        let ser = leaf.write_to_bytes().expect(&"cannot serialize");
        assert_eq!(ser.len(), 9);

        // decode it
        let mut parsed = proofs::LeafOp::new();
        parsed.merge_from_bytes(&ser).expect(&"cannot parse");

        // ensure we got it
        assert_eq!(leaf.get_prehash_key(), parsed.get_prehash_key());
        assert_eq!(leaf.get_prefix(), parsed.get_prefix());
    }
}

