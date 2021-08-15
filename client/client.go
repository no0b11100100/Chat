package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type CommandID int

const (
	LogIn CommandID = iota
	LogInUser
	LogInGuest
	AllActiveUsers
	Send
	Quit
)

type Command struct {
	ID      CommandID       `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

type client struct {
	Connection net.Conn
	Username   string
	Password   string
	Status     string
}

func (c *client) sendStartUpRequest() {
	payload, _ := json.Marshal(Command{ID: LogIn})
	fmt.Fprintf(c.Connection, string(payload)+"\n")
}

func (c *client) handleInput(input string) Command {
	if strings.HasPrefix(input, "Quit") {
		return Command{ID: Quit}
	}

	return Command{}
}

func (c *client) Run() {
	for {
		// Прослушиваем ответ
		message, _ := bufio.NewReader(c.Connection).ReadString('\n')
		fmt.Print("Message from server: " + message)

		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Print("User input: ", string(text))

		command := c.handleInput(text)

		// Отправляем в socket
		fmt.Fprintf(c.Connection, text+"\n")
	}
}
