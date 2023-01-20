package main

func main() {
	server := NewServer()
	defer server.Shutdown()
	server.Serve()
}
