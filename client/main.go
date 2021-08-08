package main

// https://dev-gang.ru/article/golang-prostoy-server-tcp-i-tcp-klient/

import "net"

func main() {

  // Подключаемся к сокету
  	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	c := client{Connection: conn}
	c.sendStartUpRequest()
	c.Run()
	defer conn.Close()
}
