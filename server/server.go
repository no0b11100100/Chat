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
	isActive   bool
	connection net.Conn
}

type server struct {
	listener      net.Listener
	connections   []connectionInfo
	CommandReader *bufio.Reader
	DB            DBInterface
}

func NewServer() *server {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("cannot create server")
	}

	return &server{
		listener:      ln,
		connections:   make([]connectionInfo, 0),
		CommandReader: bufio.NewReader(nil),
		DB:            NewDB(),
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
	fmt.Print("command ", c.ID, "\n")
	switch c.ID {
	case command.StartConnection:
		return "\n"
	case command.LogInUser:
		payload := command.UserLoginPayload{}
		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			return err.Error() + "\n"
		}
		record, err := s.DB.Select(conn.RemoteAddr().String())
		if err == nil {
			if payload.Password != record.Password || payload.Email != record.Email {
				return "Invalid data\n"
			}
			s.setUserStatus(true, conn.RemoteAddr().String())
			return "Welcome " + record.NickName + "\n"
		}
		return err.Error()
	case command.RegisterUser:
		payload := command.UserLoginPayload{}
		s.setUserStatus(true, conn.RemoteAddr().String())
		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			return err.Error() + "\n"
		}
		record := Record{
			IP:       conn.RemoteAddr().String(),
			Email:    payload.Email,
			Password: payload.Password,
			NickName: payload.NickName,
		}
		s.DB.AddRecord(record)
		s.connections = append(s.connections, connectionInfo{isActive: true, connection: conn})
		return "Welcome " + payload.NickName + "\n"
	case command.GuestUser:
		payload := command.UserLoginPayload{}
		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			return err.Error() + "\n"
		}

		record, err := s.DB.Select(conn.RemoteAddr().String())
		if err != nil {
			record = Record{
				IP:       conn.RemoteAddr().String(),
				NickName: "guest_1",
			}
			s.DB.AddRecord(record)
			s.connections = append(s.connections, connectionInfo{isActive: true, connection: conn})
		} else {
			s.setUserStatus(true, conn.RemoteAddr().String())
		}

		return "Welcome " + record.NickName + "\n"
	case command.Quit:
		s.setUserStatus(false, conn.RemoteAddr().String())
		record, err := s.DB.Select(conn.RemoteAddr().String())
		if err != nil {
			return "Bye\n"
		}
		return "Bye " + record.NickName + "\n"
	case command.ActiveUsers:
		result := ""
		for _, c := range s.connections {
			if c.isActive {
				record, err := s.DB.Select(c.connection.RemoteAddr().String())
				if err != nil {
					fmt.Println("err", err)
					continue
				}
				result += record.NickName + ","
			}
		}

		return strings.TrimSuffix(result, ",") + "\n"
	}

	return "\n"
}

func (s *server) setUserStatus(status bool, connection string) {
	for index, c := range s.connections {
		if c.connection.RemoteAddr().String() == connection {
			s.connections[index].isActive = status
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
		conn.Write([]byte(response))
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
