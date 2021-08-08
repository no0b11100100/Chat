package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"strconv"
	"strings"
)

type server struct {
	listener net.Listener
	connections []net.Conn
}

func NewServer() *server {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("cannot create server")
	}

	return &server{
		listener: ln,
		connections: make([]net.Conn, 0),
	}
}

func (s* server) Close() {
	for _, c := range s.connections {
		c.Close()
	}
	s.listener.Close()
}

func (s *server) isAlive() {
	for _ = range time.Tick(10*time.Second) {
		fmt.Println("alive", strconv.Itoa(len(s.connections)))
	}
}

func (s *server) Run() {
	go s.isAlive()
	for {
		conn, err := s.listener.Accept()
		fmt.Print("Run\n")

		if err != nil {
			fmt.Println("cannot accept")
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Recieved message: ", string(message))

		conn.Write([]byte("Response\n"));

		if strings.HasPrefix(string(message), "start client") {
			s.connections = append(s.connections, conn)
		}
	}
}