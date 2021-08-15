package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
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

type connectionInfo struct {
	isActive   bool
	connection net.Conn
}

type server struct {
	listener    net.Listener
	connections []connectionInfo
}

func NewServer() *server {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("cannot create server")
	}

	return &server{
		listener:    ln,
		connections: make([]connectionInfo, 0),
	}
}

func (s *server) Close() {
	// for _, c := range s.connections {
	// c.Close()
	// }
	s.listener.Close()
}

func (s *server) isAlive() {
	for range time.Tick(10 * time.Second) {
		fmt.Println("alive", strconv.Itoa(len(s.connections)))
	}
}

func (s *server) handleCommand(c Command, conn net.Conn) string {
	switch c.ID {
	case LogIn:
		go s.addConnection(conn)
		return "Hello. Would you like to loggin as a guest?[Yes\\No]\n"
	case Quit:
		go s.removeConnection(conn)
		return ":(\n"
	default:
		return "\n"
	}
}

func (s *server) addConnection(conn net.Conn) {
	if len(s.connections) == 0 {
		s.connections = append(s.connections, connectionInfo{true, conn})
		return
	}

	for _, c := range s.connections {
		if c.connection.RemoteAddr() != conn.RemoteAddr() {
			s.connections = append(s.connections, connectionInfo{true, conn})
		}
	}
}

func (s *server) removeConnection(conn net.Conn) {
	if len(s.connections) == 0 {
		return
	}

	for index, c := range s.connections {
		if c.connection.RemoteAddr() == conn.RemoteAddr() {
			s.connections = s.connections[index : index+1]
		}
	}
}

func (s *server) handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {

		message, _ := reader.ReadString('\n')
		if message == "\n" {
			continue
		}
		fmt.Print("Recieved message: ", string(message))

		command := Command{}
		err := json.Unmarshal([]byte(message), &command)
		if err != nil {
			fmt.Print("error parse command", err)
			conn.Write([]byte("Error\n"))
			continue
		}

		responce := s.handleCommand(command, conn)
		conn.Write([]byte(responce))
	}
}

func (s *server) Run() {

	go s.isAlive()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("cannot accept")
		}

		go s.handleRequest(conn)
	}

}
