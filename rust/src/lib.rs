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
        let parsed: proofs::LeafOp = protobuf::parse_from_bytes(&ser).expect(&"cannot parse");

        // ensure we got it
        assert_eq!(leaf.get_prehash_key(), parsed.get_prehash_key());
        assert_eq!(leaf.get_prefix(), parsed.get_prefix());
    }


    #[test]
    fn proto_nested_encoding() {
        let mut leaf = proofs::LeafOp::new();
        leaf.set_hash(proofs::HashOp::BITCOIN);
        leaf.set_prehash_key(proofs::HashOp::KECCAK);
        leaf.set_prefix(vec![7, 8, 9, 233].into());

        let mut inner = proofs::InnerOp::new();
        inner.set_hash(proofs::HashOp::SHA256);
        inner.set_prefix(vec![5, 6, 7, 1, 2, 3, 0, 0, 77]);
        inner.set_suffix(vec![6, 9, 12]);

        let mut exist = proofs::ExistenceProof::new();
        exist.set_key(vec![66, 69]);
        exist.set_value(vec![66, 69, 81]);
        exist.set_leaf(leaf);
        exist.set_path(protobuf::RepeatedField::from_vec(vec![inner]));

        // encode it
        let ser = exist.write_to_bytes().expect(&"cannot serialize");
        assert_eq!(ser.len(), 41);

        // decode it
        let mut parsed = proofs::ExistenceProof::new();
        parsed.merge_from_bytes(&ser).expect(&"cannot parse");

        // ensure we got it
        assert_eq!(proofs::HashOp::KECCAK, parsed.get_leaf().get_prehash_key());
        assert_eq!(vec![7, 8, 9, 233], parsed.get_leaf().get_prefix());
        assert_eq!(vec![66, 69], parsed.get_key());
        assert_eq!(1, parsed.get_path().len());
        assert_eq!(vec![5, 6, 7, 1, 2, 3, 0, 0, 77], parsed.get_path()[0].get_prefix());


    }

}

