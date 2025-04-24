package main

import (
	"log"
	"github.com/gorilla/websocket"
)

func connectToServer(url string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
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
