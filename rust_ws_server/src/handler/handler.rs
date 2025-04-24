use tokio::sync::broadcast::{Sender, Receiver};
use tokio_tungstenite::WebSocketStream;
use tokio_tungstenite::tungstenite::Message;
use futures_util::{StreamExt, SinkExt};
use tokio::net::TcpStream;
use tracing::{info, error};

pub async fn handle_connection(
    ws_stream: WebSocketStream<TcpStream>,
    tx: Sender<String>,
    mut rx: Receiver<String>,
) {
    let (mut write, mut read) = ws_stream.split();

    let mut write_task = tokio::spawn(async move {
        while let Ok(msg) = rx.recv().await {
            if write.send(Message::Text(msg)).await.is_err() {
                error!("Failed to send message to client");
                break;
            }
        }
    });

    let mut read_task = tokio::spawn(async move {
        while let Some(Ok(msg)) = read.next().await {
            if let Message::Text(text) = msg {
                info!("Received message from client: {}", text);
                if tx.send(text).is_err() {
                    error!("Broadcast failed");
                    break;
                }
            }
        }
    });

    tokio::select! {
        _ = (&mut write_task) => read_task.abort(),
        _ = (&mut read_task) => write_task.abort(),
    }
}
