package main

func main() {
	server := NewServer()
	server.Serve()
	defer server.Shutdown()
}
