package main

import (
	"Chat/RemoteServer/database"
	"bufio"
	"common"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
	"sync"
)

type Handler func([]byte) common.CommandResponce

type Server struct {
	// key - (ip address)user id, value - connection
	cache    *sync.Map
	listener net.Listener
	database database.Database
	handlers map[common.CommandType]Handler
}

func NewServer() *Server {
	fmt.Println("Start server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("New Server err:", err)
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	fmt.Println(addrs)

	s := &Server{
		cache:    new(sync.Map),
		listener: ln,
		handlers: make(map[common.CommandType]Handler),
		database: database.NewDatabase(),
	}

	s.addHandlers()
	// s.database.Connect()

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
	s.listener.Close()
}

func (s *Server) processConnection(conn net.Conn) {
	fmt.Println("Accept connection:", conn.RemoteAddr().String())
	s.cache.Store(conn.RemoteAddr().String(), conn)
	defer conn.Close()
	defer s.cache.Delete(conn.RemoteAddr().String())

	for {
		reader := bufio.NewReader(conn)
		tp := textproto.NewReader(reader)

		data, err := tp.ReadLine()

		if err != nil {
			fmt.Println("processConnection error", err)
			continue
		}

		fmt.Println("processConnection message", data)
		s.handleCommand(data, conn)
	}
}

func (s *Server) handleCommand(payload string, conn net.Conn) {
	decodedPayload, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded text: %s\n", decodedPayload)
	var c common.Command

	err = json.Unmarshal(decodedPayload, &c)
	if err != nil {
		fmt.Println("handleCommand unmarshal error", err)
	}

	var responce common.CommandResponce
	if handler, ok := s.handlers[c.Type]; ok {
		responce = handler(c.Payload)
	} else {
		fmt.Println("Unknown command", c.Type)
		responce.Error = fmt.Sprintf("Unknown command %v", c.Type)
	}

	s.send(conn, responce)
}

func (s *Server) addHandlers() {
	s.handlers[common.LogIn] = s.LogIn
	s.handlers[common.Register] = s.RegisterUser
	s.handlers[common.SendMessage] = s.SendMessage
}

func (s *Server) send(conn net.Conn, responce common.CommandResponce) {
	bytes, err := json.Marshal(responce)

	if err != nil {
		fmt.Println("send error", err)
		return
	}

	payload := base64.StdEncoding.EncodeToString(bytes)
	payload = payload + "\n"
	conn.Write([]byte(payload))
}
