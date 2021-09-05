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

func (c *client) logIn() {
	fmt.Print("Hello\nWhat action would you do?\nLogIn\nGuest\nRegister\n")
	HandleLogIn(c.Connection, c.UserReader, c.ServerReader)
}

func (c *client) sendStartUpRequest() {
	HandleStartUp(c.Connection)
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
		// case "/send":
		// case "/changePassword":
		// case "/changeNickName":
		// case "/changeEmail":
	}
}

func (c *client) Run() {
	c.sendStartUpRequest()
	c.logIn()

	for {
		fmt.Print("Enter command: ")
		text, _ := c.UserReader.ReadString('\n')
		fmt.Print("User input: ", string(text))
		c.handleInput(text)
	}
}
