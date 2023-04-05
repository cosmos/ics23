extern crate prost_build;

use std::env;
use std::vec::Vec;

fn main() {
    let args: Vec<_> = env::args().collect();
    let mut root = "../..";
    if args.len() > 1 {
        root = &args[1];
    }

    let out_dir: &str = &format!("{}{}", root, "/rust/src");
    let input: &str = &format!("{}{}", root, "/proto/cosmos/ics23/v1/proofs.proto");

    prost_build::Config::new()
        .out_dir(&out_dir)
        .format(true)
        .type_attribute(".", "#[derive(Eq)]")
        .compile_protos(&[input], &[root])
        .unwrap();
}
