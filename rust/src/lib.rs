// Include the `proofs` module, which is generated from proofs.proto.
pub mod proofs {
    include!(concat!(env!("OUT_DIR"), "/proofs.rs"));
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
