package main

import (
	"bufio"
	"command"
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
	SQL_DB        DBInterface
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
		DB:            NewDB(),
		SQL_DB:        NewDataBase(),
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
	fmt.Println("command ", c.ID)
	switch c.ID {
	case command.StartConnection:
		s.Connections[conn.RemoteAddr().String()] = connectionInfo{connection: conn}
	case command.LogInUser:
		payload := command.UserLoginPayload{}
		response := command.Response{}
		if err := json.Unmarshal((c.Payload), &payload); err != nil {
			response.SetError(err.Error())
			return string(response.Marshal())
		}
		record, err := s.SQL_DB.Select(payload.Email)
		if err == nil {
			if payload.Password != record.Password || payload.Email != record.Email {
				response.SetError("Invalid password or email")
				return string(response.Marshal())
			}
			response.SetPayload("Welcome " + record.NickName)
			s.Connections[conn.RemoteAddr().String()] = connectionInfo{Name: record.NickName, connection: conn}
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
		s.Connections[conn.RemoteAddr().String()] = connectionInfo{Name: record.NickName, connection: conn}
		response.SetPayload("Welcome " + payload.NickName)
		return string(response.Marshal())
	case command.Quit:
		s.closeClientConnection(conn.RemoteAddr().String())
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
		sender := s.Connections[conn.RemoteAddr().String()].Name
		message := sender + ": " + string(c.Payload)
		for _, client := range s.Connections {
			if client.connection.RemoteAddr().String() != conn.RemoteAddr().String() {
				response := command.Response{}
				response.SetPayload(message)
				client.connection.Write(response.Marshal())
			}
		}
	}

	return "\n"
}

func (s *server) closeClientConnection(connAddr string) {
	s.Connections[connAddr].connection.Close()
	delete(s.Connections, connAddr)
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
		fmt.Println("message:", response)
		if response != "\n" {
			conn.Write([]byte(response))
		} else {
			fmt.Println("message was not sent")
		}
	}
}

func (s *server) Run() {
	go s.isAlive()
	defer s.SQL_DB.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("cannot accept")
		}

		go s.handleRequest(conn)
	}
}
