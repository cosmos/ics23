use std::env;
use std::io::Result;
use std::path::PathBuf;
use std::vec::Vec;

fn main() -> Result<()> {
    let args: Vec<_> = env::args().collect();
    let root = if args.len() > 1 {
        &args[1]
    } else {
        env!("CARGO_MANIFEST_DIR")
    };

    println!("Root: {root}");

    let root = PathBuf::from(root).join("..").join("..").canonicalize()?;
    let input = root.join("proto/cosmos/ics23/v1/proofs.proto");
    let out_dir = root.join("rust/src");
    let descriptor_path = out_dir.join("proto_descriptor.bin");

    println!("Input: {}", input.display());
    println!("Output: {}", out_dir.display());
    println!("Descriptor: {}", descriptor_path.display());

    prost_build::Config::new()
        .out_dir(&out_dir)
        .format(true)
        // Needed for pbjson_build to generate the serde implementations
        .file_descriptor_set_path(&descriptor_path)
        .compile_well_known_types()
        // As recommended in pbjson_types docs
        .extern_path(".google.protobuf", "::pbjson_types")
        .compile_protos(&[input], &[root])?;

    // Finally, build pbjson Serialize, Deserialize impls:
    let descriptor_set = std::fs::read(descriptor_path)?;

    pbjson_build::Builder::new()
        .register_descriptors(&descriptor_set)?
        .out_dir(&out_dir)
        .build(&[".cosmos.ics23.v1"])?;

    Ok(())
}
