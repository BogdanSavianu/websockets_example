package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func writeMessages(conn *websocket.Conn, clientID string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[%s] Enter message to send:\n", clientID)
	for scanner.Scan() {
		text := scanner.Text()
		message := fmt.Sprintf("[%s] %s", clientID, text)
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Write error: ", err)
			return
		}
		fmt.Printf("[%s] Enter message to send:\n", clientID)
	}
}
