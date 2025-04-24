package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func writeMessages(conn *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter message to send:")
	for scanner.Scan() {
		text := scanner.Text()
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("Write error: ", err)
			return
		}
	}
}
