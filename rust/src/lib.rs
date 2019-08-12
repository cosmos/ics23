extern crate protobuf;

mod proofs;

#[cfg(test)]
mod tests {
    use crate::proofs;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn proto_encode() {
        let mut leaf = proofs::LeafOp::new();
        leaf.set_hash(proofs::HashOp::SHA256);
        leaf.set_prehash_key(proofs::HashOp::SHA256);
        assert_eq!(leaf.get_hash(), proofs::HashOp::SHA256);
    }
}

