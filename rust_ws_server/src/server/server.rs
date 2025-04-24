use tokio::net::TcpListener;
use tokio::sync::broadcast;
use tokio_tungstenite::accept_async;
use tracing::{info, error};

use crate::handler::handle_connection;

pub async fn run(addr: &str) {
    let listener = TcpListener::bind(addr).await.expect("Failed to bind");
    let (tx, _rx) = broadcast::channel(100);
    
    loop {
        match listener.accept().await {
            Ok((stream, addr)) => {
                info!("New connection from {}", addr);
                let tx = tx.clone();
                let rx = tx.subscribe();
                tokio::spawn(async move {
                    match accept_async(stream).await {
                        Ok(ws_stream) => {
                            info!("WebSocket handshake successful with {}", addr);
                            handle_connection(ws_stream, tx, rx).await;
                        }
                        Err(e) => error!("WebSocket handshake failed: {}", e),
                    }
                });
            }
            Err(e) => error!("Failed to accept connection: {}", e),        }
}}
