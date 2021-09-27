package main

import (
	"bufio"
	"command"

	// . "crypto/sha512"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type connectionInfo struct {
	connection net.Conn
	Name       string
}

type server struct {
	listener      net.Listener
	Connections   map[string]connectionInfo
	CommandReader *bufio.Reader
	DB            DBInterface
	clientIDs     uint64
}

func NewServer() *server {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("cannot create server")
	}

	return &server{
		listener:      ln,
		Connections:   make(map[string]connectionInfo),
		CommandReader: bufio.NewReader(nil),
		DB:            NewDataBase(), //NewDB(),
		clientIDs:     0,
	}
}

func (s *server) Close() {
	for _, c := range s.Connections {
		c.connection.Close()
	}
	s.listener.Close()
}

func (s *server) isAlive() {
	for range time.Tick(10 * time.Second) {
		fmt.Println("alive", strconv.Itoa(len(s.Connections)))
	}
}

func (s *server) handleCommand(c command.Command, conn net.Conn) string {
	fmt.Println("command ", c.ID, conn.RemoteAddr().String())
	switch c.ID {
	case command.StartConnection:
	case command.LogInUser:
		payload := command.UserLoginPayload{}
		response := command.Response{}
		if err := json.Unmarshal((c.Payload), &payload); err != nil {
			response.SetError(err.Error())
			return string(response.Marshal())
		}
		record, err := s.DB.Select(payload.Email)
		if err == nil {
			if payload.Password != record.Password {
				response.SetError("Invalid password or email")
				return string(response.Marshal())
			}
			response.SetPayload("Welcome " + record.NickName)
			s.Connections[conn.RemoteAddr().Network()] = connectionInfo{Name: record.NickName, connection: conn}
			return string(response.Marshal())
		}
		response.SetError(err.Error())
		return string(response.Marshal())
	case command.RegisterUser:
		payload := command.UserLoginPayload{}
		response := command.Response{}
		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			response.SetError(err.Error())
			return string(response.Marshal())
		}
		record := Record{
			Email:    payload.Email,
			Password: payload.Password,
			NickName: payload.NickName,
		}

		if !s.DB.IsEmailUnique(record.Email) {
			response.SetError("this email is already used")
			return string(response.Marshal())
		}

		s.DB.AddRecord(record)
		response.SetPayload("Welcome " + payload.NickName)
		s.Connections[conn.RemoteAddr().Network()] = connectionInfo{Name: record.NickName, connection: conn}
		return string(response.Marshal())
	case command.Quit:
		s.closeClientConnection(conn.RemoteAddr().Network())
	case command.ActiveUsers:
		result := ""
		response := command.Response{}
		for _, c := range s.Connections {
			result += c.Name + ","
		}
		result = strings.TrimSuffix(result, ",")
		response.SetPayload(result)
		return string(response.Marshal())
	case command.SendMessage:
		sender := s.Connections[conn.RemoteAddr().Network()].Name
		message := "> " + sender + ": " + string(c.Payload)
		for id, client := range s.Connections {
			if id != conn.RemoteAddr().Network() {
				response := command.Response{}
				response.SetPayload(message)
				client.connection.Write(response.Marshal())
			}
		}
	}

	return "\n"
}

func (s *server) closeClientConnection(connAddr string) {
	if _, ok := s.Connections[connAddr]; ok {
		s.Connections[connAddr].connection.Close()
		delete(s.Connections, connAddr)
		fmt.Println("removed")
	}
	fmt.Println("remove connection", strconv.Itoa(len(s.Connections)), connAddr, s.Connections)
}

func (s *server) handleRequest(conn net.Conn) {
	s.CommandReader.Reset(conn)
	for {
		message, _ := s.CommandReader.ReadString('\n')
		if message == "\n" || message == "" {
			continue
		}
		fmt.Print("Recieved message: ", string(message))

		cmd := command.Command{}
		err := json.Unmarshal([]byte(message), &cmd)
		if err != nil {
			fmt.Print("error parse command ", err)
			conn.Write([]byte("Error\n"))
			continue
		}

		response := s.handleCommand(cmd, conn)
		if response != "\n" {
			fmt.Println("send message:", response)
			conn.Write([]byte(response))
		} else {
			fmt.Println("message was not sent")
		}
	}
}

// func (s *server) checkConnection() {
// 	broadcast := func() {
// 		fmt.Println("start check")
// 		for _, c := range s.Connections {
// 			c.connection.Write([]byte("Check\n"))
// 		}
// 		fmt.Println("end check")
// 	}
// 	for {
// 		ticker := time.NewTicker(1 * time.Second)
// 		<-ticker.C
// 		broadcast()
// 	}
// }

func (s *server) Run() {
	go s.isAlive()
	defer s.DB.Close()
	// go s.checkConnection()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("cannot accept")
		}

		go s.handleRequest(conn)
	}
}
