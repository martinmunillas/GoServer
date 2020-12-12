package main

func main() {
	server := NewServer(":5000")
	server.Listen()
}
