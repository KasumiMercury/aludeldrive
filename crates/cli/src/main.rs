use std::process::ExitCode;

use aludel_api::aludel::v1::{GetServerInfoRequest, ServerServiceClient};
use connectrpc::client::{ClientConfig, HttpClient};

const DEFAULT_SERVER: &str = "http://127.0.0.1:8080";
const USAGE: &str = "usage: aludel-cli [--server URL] version";

#[tokio::main]
async fn main() -> ExitCode {
    match run(std::env::args().skip(1).collect()).await {
        Ok(()) => ExitCode::SUCCESS,
        Err(err) => {
            eprintln!("aludel-cli: {err}");
            ExitCode::FAILURE
        }
    }
}

async fn run(args: Vec<String>) -> Result<(), String> {
    let mut server = DEFAULT_SERVER.to_string();
    let mut rest = args.into_iter();
    let command = loop {
        match rest.next().as_deref() {
            Some("--server") => server = rest.next().ok_or(USAGE)?,
            Some(cmd) => break cmd.to_string(),
            None => return Err(USAGE.to_string()),
        }
    };

    match command.as_str() {
        "version" => version(&server).await,
        _ => Err(format!("unknown command {command:?}\n{USAGE}")),
    }
}

async fn version(server: &str) -> Result<(), String> {
    let uri: http::Uri = server.parse().map_err(|e| format!("invalid server URL {server:?}: {e}"))?;
    let client = ServerServiceClient::new(HttpClient::plaintext(), ClientConfig::new(uri));
    let info = client
        .get_server_info(GetServerInfoRequest::default())
        .await
        .map_err(|e| format!("GetServerInfo: {e}"))?
        .into_owned();

    let client_version = env!("CARGO_PKG_VERSION");
    println!("client {client_version}");
    println!("server {}", info.version);
    if info.version != client_version {
        eprintln!("warning: client and server versions differ; update to the same version");
    }
    Ok(())
}
