extern crate protoc_rust;

use protoc_rust::Customize;

fn main() {
	protoc_rust::run(protoc_rust::Args {
	    out_dir: "../src",
	    input: &["../../proofs.proto"],
	    includes: &["../.."],
	    customize: Customize {
	      ..Default::default()
	    },
	}).expect("protoc");
}


// fn main() {
//     prost_build::compile_protos(&["../../proofs.proto"],
//                                 &["../.."]).unwrap();
// }