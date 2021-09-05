package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type client struct {
	Connection   net.Conn
	UserReader   *bufio.Reader
	ServerReader *bufio.Reader
}

func NewClient(conn net.Conn) *client {
	return &client{
		Connection:   conn,
		UserReader:   bufio.NewReader(os.Stdin),
		ServerReader: bufio.NewReader(conn),
	}
}

func (c *client) handleInput(input string) {
	input = strings.TrimSuffix(input, "\n")
	switch input {
	case "/quit":
		HandleQuit(c.Connection, c.ServerReader)
		c.Connection.Close()
		os.Exit(1)
	case "/activeUsers":
		HandleActiveusers(c.Connection, c.ServerReader)
	case "/send":
		HandleSendMessage(c.Connection, c.UserReader, c.ServerReader)
		// case "/changePassword":
		// case "/changeNickName":
		// case "/changeEmail":
	}
}

func (c *client) Run() {
	HandleStartUp(c.Connection, c.ServerReader)
	HandleLogIn(c.Connection, c.UserReader, c.ServerReader)

	go func() {
		for {
			text, _ := c.ServerReader.ReadString('\n')
			if text != "" {
				fmt.Print("Message from ", text)
			}
		}
	}()

	for {
		fmt.Print("Enter command: ")
		text, _ := c.UserReader.ReadString('\n')
		fmt.Print("User input: ", string(text))
		c.handleInput(text)
	}
}
