package main

import (
	"log"
	"net"
	"github.com/gorilla/websocket"
)

func connectToServer(url string, localAddr string) *websocket.Conn {
	if localAddr == "" {
		log.Printf("Connecting to %s (default source address)", url)
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Fatal("Dial error: ", err)
		}
		return conn
	}

	dialer := &websocket.Dialer{
		NetDialContext: (&net.Dialer{
			LocalAddr: &net.TCPAddr{
				IP: net.ParseIP(localAddr),
			},
		}).DialContext,
	}
	
	log.Printf("Connecting from %s to %s", localAddr, url)
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
	return conn
}

func readMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			return
		}
		log.Printf("Received: %s", message)
	}
}
