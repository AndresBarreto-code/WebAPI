package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Handle("POST", "/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Listen()
}
