package main

import (
	"bufio"
	"command"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

type server struct {
	listener      net.Listener
	connections   []net.Conn
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
		connections:   make([]net.Conn, 0),
		CommandReader: bufio.NewReader(nil),
		DB:            NewDB(),
		SQL_DB:        NewDataBase(),
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

func (s *server) handleCommand(c command.Command, conn net.Conn) string {
	fmt.Println("command ", c.ID)
	switch c.ID {
	case command.StartConnection:
		s.connections = append(s.connections, conn)
	case command.LogInUser:
		payload := command.UserLoginPayload{}
		response := command.Response{}
		if err := json.Unmarshal((c.Payload), &payload); err != nil {
			response.SetError(err.Error())
			return string(response.Marshal())
		}
		record, err := s.DB.Select(conn.RemoteAddr().String())
		if err == nil {
			if payload.Password != record.Password || payload.Email != record.Email {
				response.SetError("Invalid password or email")
				return string(response.Marshal())
			}
			response.SetPayload("Welcome " + record.NickName)
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
			IP:       conn.RemoteAddr().String(),
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
		return string(response.Marshal())
	case command.Quit:
		s.closeClientConnection(conn.RemoteAddr().String())
	case command.ActiveUsers:
		result := ""
		response := command.Response{}
		for _, c := range s.connections {
			record, err := s.DB.Select(c.RemoteAddr().String())
			if err != nil {
				fmt.Println("err", err)
				continue
			}
			result += record.NickName + ","
		}
		response.SetPayload(result)
		return string(response.Marshal())
	case command.SendMessage:
		sender, _ := s.DB.Select(conn.RemoteAddr().String())
		message := sender.NickName + ": " + string(c.Payload)
		for _, client := range s.connections {
			if client.RemoteAddr().String() != sender.IP {
				response := command.Response{}
				response.SetPayload(message)
				client.Write(response.Marshal())
			}
		}
	}

	return "\n"
}

func (s *server) closeClientConnection(connAddr string) {
	removeElement := func(slice []net.Conn, index int) []net.Conn {
		return append(slice[:index], slice[index+1:]...)
	}
	for index, c := range s.connections {
		if c.RemoteAddr().String() == connAddr {
			s.connections = removeElement(s.connections, index)
		}
	}
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
	// s.SQL_DB.AddRecord(Record{IP: "1232", Email: "sad@asd", Password: "1234", NickName: "user"})

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("cannot accept")
		}

		go s.handleRequest(conn)
	}
}
