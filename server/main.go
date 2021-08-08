package main

func main() {

	server := NewServer()
	server.Run()
  	defer server.Close()
}
