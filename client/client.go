package main

import (
	"bufio"
	"command"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"
)

type client struct {
	Connection   net.Conn
	UserReader   *bufio.Reader
	ServerReader *bufio.Reader
	state        State
	Chan         chan string
}

type State int

const (
	Init State = iota
	Ready
	Killed
)

func NewClient() *client {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	return &client{
		Connection:   conn,
		UserReader:   bufio.NewReader(os.Stdin),
		ServerReader: bufio.NewReader(conn),
		state:        Init,
		Chan:         make(chan string),
	}
}

func (c *client) Destroy() {
	close(c.Chan)
	c.Connection.Close()
}

func (c *client) readSock() {
	if c.Connection == nil {
		panic("Connection is nil")
	}
	for {
		payload, err := c.ServerReader.ReadString('\n')
		if err != nil || payload == "" || payload == "\n" {
			continue
		}

		response := command.Response{}
		if err := json.Unmarshal([]byte(payload), &response); err == nil {
			if c.state == Init && response.Status == command.OK {
				c.state = Ready
			}
			fmt.Println(string(response.Payload))
		}
	}
}

const (
	LogIn       = "/logIn"
	Register    = "/register"
	Quit        = "/quit"
	Activeusers = "/activeUsers"
	JoinToRoom  = "/join"
	LeaveRoom   = "/leave"
	CreateRoom  = "/create"
)

func (c *client) handleCommand(input string) (string, error) {
	input = strings.TrimSuffix(input, "\n")
	removeCommandNameFromInput := func(inputString string, commandName string) string {
		return strings.TrimPrefix(inputString[len(commandName):], " ")
	}
	if strings.HasPrefix(input, LogIn) {
		formatedInput := removeCommandNameFromInput(input, LogIn)
		data := strings.Split(formatedInput, " ")
		if len(data) != 2 {
			fmt.Println("Invalid data. Please try again")
			return "", errors.New("Invalid data")
		}
		commandPayload := command.UserLoginPayload{
			Email:    data[0],
			Password: data[1],
		}
		command := command.Command{
			ID:      command.LogInUser,
			Payload: commandPayload.Marshal(),
		}
		if payload, err := json.Marshal(command); err == nil {
			return string(payload), nil
		}
	} else if strings.HasPrefix(input, Register) {
		formatedInput := removeCommandNameFromInput(input, Register)
		data := strings.Split(formatedInput, " ")
		if len(data) != 3 {
			fmt.Println("Invalid data. Please try again", data, input)
			return "", errors.New("Invalid data")
		}
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
			return string(payload), nil
		}
	} else if input == Activeusers {
		if c.state == Init {
			fmt.Println("please log in or create account")
			return "", errors.New("Invalid data")
		}
		command := command.Command{ID: command.ActiveUsers}
		payload, err := json.Marshal(command)
		if err == nil {
			return string(payload), nil
		}
	} else if input == Quit {
		command := command.Command{ID: command.Quit}
		if payload, err := json.Marshal(command); err == nil {
			c.state = Killed
			return string(payload), nil
		}

	} else if strings.HasPrefix(input, JoinToRoom) {
		formatedInput := removeCommandNameFromInput(input, JoinToRoom)
		commandPayload := command.RoomInfo{Name: formatedInput}
		command := command.Command{
			ID:      command.JoinToRoom,
			Payload: commandPayload.Marshal(),
		}

		payload, err := json.Marshal(command)
		if err == nil {
			return string(payload), nil
		}
	} else if input == LeaveRoom {
		command := command.Command{
			ID: command.LeaveRoom,
		}

		payload, err := json.Marshal(command)
		if err == nil {
			return string(payload), nil
		}
	} else if strings.HasPrefix(input, CreateRoom) {
		formatedInput := removeCommandNameFromInput(input, CreateRoom)
		commandPayload := command.RoomInfo{Name: formatedInput}
		command := command.Command{
			ID:      command.CreateRoom,
			Payload: commandPayload.Marshal(),
		}

		payload, err := json.Marshal(command)
		if err == nil {
			return string(payload), nil
		}
	} else {
		if c.state == Init {
			fmt.Println("please log in or create account")
			return "", errors.New("Invalid data")
		}
		// send message
		fmt.Println("user message", input)
		command := command.Command{ID: command.SendMessage, Payload: []byte(input)}
		payload, err := json.Marshal(command)
		if err == nil {
			return string(payload), nil
		}
	}

	return "", errors.New("Error")
}

func (c *client) readConsole() {
	for {
		fmt.Print("Enter command: ")
		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Print(">")
		out := msg[:len(msg)-1]
		msg, err := c.handleCommand(out)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.Chan <- msg + "\n"
	}
}

func (c *client) handleInterrupt() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ct := make(chan os.Signal, 1)
	signal.Notify(ct, os.Interrupt)
	defer func() {
		signal.Stop(ct)
		fmt.Println("handled")
		msg, _ := c.handleCommand(Quit)
		cancel()
		c.Chan <- msg + "\n"
	}()

	select {
	case <-ct:
		cancel()
	case <-ctx.Done():
	}
}

func (c *client) Run() {
	go c.readConsole()
	go c.readSock()
	go c.handleInterrupt()

	for {
		val, ok := <-c.Chan
		if ok {
			out := []byte(val)
			fmt.Println("sent to server", val)
			_, err := c.Connection.Write(out)
			if err != nil {
				fmt.Println("Write error:", err.Error())
				continue
			}

			if c.state == Killed {
				os.Exit(1)
			}

		} else {
			time.Sleep(time.Second * 2)
		}

	}
}
