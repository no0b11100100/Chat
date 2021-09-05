package main

// https://dev-gang.ru/article/golang-prostoy-server-tcp-i-tcp-klient/

import (
	"fmt"
	"net"
)

func main() {

	// Подключаемся к сокету
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	// defer conn.Close()
	if err != nil {
		fmt.Print("error", err)
	}

	c := NewClient(conn)

	c.Run()
}
