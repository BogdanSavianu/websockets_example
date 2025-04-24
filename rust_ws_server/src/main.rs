use tracing_subscriber;

mod server;
mod handler;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();
    tracing::info!("Starting WebSocket server on ws://127.0.0.1:9001");
    server::run("127.0.0.1:9001").await;
}
