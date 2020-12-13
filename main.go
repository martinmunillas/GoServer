package main

func main() {
	server := NewServer(":5000")
	server.Handle("/", "GET", server.AddMiddleware(HandleRoot, Log()))
	server.Handle("/auth", "POST", server.AddMiddleware(HandleAuth, CheckAuth(true), Log()))
	server.Handle("/user", "GET", server.AddMiddleware(PostUser, Log()))
	server.Handle("/user", "POST", server.AddMiddleware(PostUser, Log()))
	server.Listen()
}
