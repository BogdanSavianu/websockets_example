package main

func main() {
	conn := connectToServer("ws://localhost:9001")
	defer conn.Close()

	go readMessages(conn)
	writeMessages(conn)
}
