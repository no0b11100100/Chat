package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

type client struct {
	Connection net.Conn
	Username string
	Password string
	Status string
}

func (c *client) sendStartUpRequest() {
	fmt.Fprintf(c.Connection, "start client\n")
}

func (c *client) Run() {
	for {
		// Прослушиваем ответ
		message, _ := bufio.NewReader(c.Connection).ReadString('\n')
		fmt.Print("Message from server: "+message)

		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Print("User input: ", string(text))
		// Отправляем в socket
		fmt.Fprintf(c.Connection, text + "\n")
		// c.Connection.Write([]byte(text))
	  }
}