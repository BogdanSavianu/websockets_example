package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	clientID := flag.String("id", "", "Client ID (e.g., client1, client2)")
	localAddr := flag.String("ip", "", "Local IP address to bind to (e.g., 127.0.0.2)")
	flag.Parse()

	if *clientID == "" {
		fmt.Println("Usage: go run . -id=clientX [-ip=127.0.0.X]")
		fmt.Println("Example: go run . -id=client1 -ip=127.0.0.2")
		os.Exit(1)
	}

	serverAddr := "ws://localhost:9001"
	conn := connectToServer(serverAddr, *localAddr)
	defer conn.Close()

	go readMessages(conn)
	writeMessages(conn, *clientID)
}
