use std::path::PathBuf;
use std::process::Command;

fn main() {
    let root = PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("../..");
    let descriptor = PathBuf::from(std::env::var("OUT_DIR").unwrap()).join("aludel.binpb");

    println!("cargo:rerun-if-changed={}", root.join("proto").display());
    println!("cargo:rerun-if-changed={}", root.join("buf.yaml").display());

    let status = Command::new("buf")
        .arg("build")
        .arg(&root)
        .arg("-o")
        .arg(&descriptor)
        .status()
        .expect("buf must be on PATH");
    assert!(status.success(), "buf build failed");

    connectrpc_build::Config::new()
        .descriptor_set(&descriptor)
        .files(&["aludel/v1/server.proto"])
        .include_file("_connectrpc.rs")
        .compile()
        .unwrap();
}
