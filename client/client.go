package main

import (
	"bufio"
	"command"
	"context"
	"encoding/json"
	"fmt"
	"os/signal"

	"net"
	"os"
	"strings"
)

type client struct {
	Connection   net.Conn
	UserReader   *bufio.Reader
	ServerReader *bufio.Reader
	ch           chan string
}

func NewClient(conn net.Conn) *client {
	return &client{
		Connection:   conn,
		UserReader:   bufio.NewReader(os.Stdin),
		ServerReader: bufio.NewReader(conn),
		ch:           make(chan string),
	}
}

const (
	LogIn       = "/logIn"
	Register    = "/register"
	Quit        = "/quit"
	Activeusers = "/activeUsers"
)

func (c *client) handleInput(input string) {
	input = strings.TrimSuffix(input, "\n")
	if strings.HasPrefix(input, LogIn) {
		data := strings.Split(input[len(LogIn)+1:], " ")
		commandPayload := command.UserLoginPayload{
			Email:    data[0],
			Password: data[1],
		}
		command := command.Command{
			ID:      command.LogInUser,
			Payload: commandPayload.Marshal(),
		}
		if payload, err := json.Marshal(command); err == nil {
			c.send(payload)
		}
	} else if strings.HasPrefix(input, Register) {
		data := strings.Split(input[len(Register)+1:], " ")
		commandPayload := command.UserLoginPayload{
			Email:    data[0],
			Password: data[1],
			NickName: data[2],
		}
		command := command.Command{
			ID:      command.RegisterUser,
			Payload: commandPayload.Marshal(),
		}
		if payload, err := json.Marshal(command); err == nil {
			c.send(payload)
		}
	} else if strings.TrimSuffix(input, "\n") == Quit {
		command := command.Command{ID: command.Quit}
		if payload, err := json.Marshal(command); err == nil {
			c.send(payload)
		}
		c.Connection.Close()
		os.Exit(1)
	} else if strings.TrimSuffix(input, "\n") == Activeusers {
		command := command.Command{ID: command.ActiveUsers}
		if payload, err := json.Marshal(command); err == nil {
			c.send(payload)
		}
	} else {
		// send message
		fmt.Println("user message", input)
		command := command.Command{ID: command.SendMessage, Payload: []byte(input)}
		payload, err := json.Marshal(command)
		if err == nil {
			c.send(payload)
		} else {
			fmt.Println(err)
		}
	}
}

func (c *client) readFromUser() {
	for {
		fmt.Print("Enter command:")
		text, _ := c.UserReader.ReadString('\n')
		c.handleInput(text)
	}
}

func (c *client) readResponse() {
	// for {
	text, _ := c.ServerReader.ReadString('\n')
	c.handleResponse(text)
	// }
}

func (c *client) handleResponse(payload string) {
	if payload == "" || payload == "\n" {
		return
	}

	response := command.Response{}
	if err := json.Unmarshal([]byte(payload), &response); err == nil {
		fmt.Println(string(response.Payload))
	}
}

func (c *client) send(data []byte) {
	fmt.Fprintf(c.Connection, string(data)+"\n")
}

func (c *client) handleInterrupt() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ct := make(chan os.Signal, 1)
	signal.Notify(ct, os.Interrupt)
	defer func() {
		signal.Stop(ct)
		fmt.Println("handled")
		c.handleInput(Quit)
		cancel()
	}()

	select {
	case <-ct:
		cancel()
	case <-ctx.Done():
	}
}

func (c *client) sendStartUp() {
	command := command.Command{ID: command.StartConnection}
	if payload, err := json.Marshal(command); err == nil {
		c.send(payload)
	}
}

func (c *client) Run() {
	c.sendStartUp()
	go c.readFromUser()
	go c.handleInterrupt()

	for {
		c.readResponse()
	}
}
