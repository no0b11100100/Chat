package main

import (
	"Chat/Client/Server/common"
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
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
	conn, err := net.Dial("tcp", "localhost:8081") // "172.17.0.2:8081"
	if err != nil {
		fmt.Println(err)
	}
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

		fmt.Println(string(payload))

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
	info := common.User{Email: "email@test.com", Password: "cGFzc3dvcmQ="}
	bytes, _ := json.Marshal(info)
	cmd := common.Command{Type: common.LogIn, Payload: []byte(bytes)}
	result, _ := json.Marshal(cmd)
	payload := base64.StdEncoding.EncodeToString(result)

	return payload, nil
}

func (c *client) readConsole() {
	for {
		fmt.Print("Enter command: ")

		fmt.Print(">")
		out := string([]byte(`{"value":10}`))
		msg, err := c.handleCommand(out)
		time.Sleep(10 * time.Second)
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

func main() {
	client := NewClient()
	client.Run()
	defer client.Destroy()
}
