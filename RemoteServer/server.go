package main

import (
	"Chat/RemoteServer/commands"
	"Chat/RemoteServer/database"
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"sync"
)

type Handler func(string)

type Server struct {
	// key - user id, value - connection
	Cache    *sync.Map
	listener net.Listener
	database database.Database
	handlers map[commands.CommandType]Handler
}

func NewServer() *Server {
	fmt.Println("Start server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("New Server err:", err)
	}

	s := &Server{
		Cache:    new(sync.Map),
		listener: ln,
		handlers: make(map[commands.CommandType]Handler),
		database: database.NewDatabase(),
	}

	s.addHandlers()
	s.database.Connect()

	return s
}

func (s *Server) Serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
		}

		go s.processConnection(conn)
	}
}

func (s *Server) Shutdown() {
	s.database.Close()
}

func (s *Server) processConnection(conn net.Conn) {
	fmt.Println("Accept cnn:", conn.RemoteAddr().String())
	for {
		reader := bufio.NewReader(conn)
		tp := textproto.NewReader(reader)

		data, err := tp.ReadLine()
		tp.ReadContinuedLine()

		if err != nil {
			fmt.Println("processConnection error", err)
		}

		fmt.Println("processConnection message", string(data))
	}
}

func (s *Server) handleCommand(c commands.Command) {}

func (s *Server) addHandlers() {
	s.handlers[commands.LogIn] = s.LogIn
	s.handlers[commands.Register] = s.RegisterUser
	s.handlers[commands.SendMessage] = s.SendMessage
}
