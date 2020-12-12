package main

func main() {
	server := NewServer(":5000")
	server.Handle("/", HandleRoot)
	server.Listen()
}
