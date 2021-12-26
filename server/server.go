package main

import (
	db "Chat/server/database"
	"bufio"
	"command"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
)

type connectionInfo struct {
	connection net.Conn
	Name       string
}

type server struct {
	listener    net.Listener
	Connections *sync.Map
	DB          db.DBInterface
}

func NewServer() *server {
	fmt.Println("Start server...")
	ln, _ := net.Listen("tcp", ":8081")
	return &server{
		listener:    ln,
		Connections: new(sync.Map),
		DB:          db.NewDataBase(),
	}
}

func (s *server) addConnection(key string, value connectionInfo) {
	s.Connections.Store(key, value)
}

func (s *server) send(conn net.Conn, payload []byte) {
	conn.Write(append(payload, '\n'))
}

func (s *server) broadcast(payload []byte, except string) {
	s.Connections.Range(func(k, v interface{}) bool {
		if info, ok := v.(connectionInfo); ok {
			conn := info.connection
			if conn.RemoteAddr().String() != except {
				s.send(conn, payload)
			}
		} else {
			fmt.Println("Cannot cast to connectionInfo")
		}
		return true
	})
}

func (s *server) process(addr string) {
	c, err := s.Connections.Load(addr)
	if !err {
		fmt.Println("cannot get item from map")
		return
	}
	conn := c.(connectionInfo).connection
	fmt.Println("Accept cnn:", conn.RemoteAddr().String())
	defer conn.Close()

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		c := command.Command{}
		err = json.Unmarshal([]byte(msg), &c)
		if err == nil {
			s.handleCommand(c, conn)
		} else {
			fmt.Println("unmarshal error", err)
		}
	}
}

func (s *server) handleCommand(c command.Command, conn net.Conn) {
	fmt.Println("command", c)
	switch c.ID {
	case command.StartConnection:
	case command.LogInUser:
		payload := command.UserLoginPayload{}
		response := command.Response{}
		if err := json.Unmarshal((c.Payload), &payload); err != nil {
			response.SetError(err.Error())
			s.send(conn, response.Marshal())
		}
		record, err := s.DB.Select(payload.Email)
		if err == nil {
			if payload.Password != record.Password {
				response.SetError("Invalid password or email")
				s.send(conn, response.Marshal())
			}
			response.SetPayload("Welcome " + record.NickName)
			s.addConnection(conn.RemoteAddr().String(), connectionInfo{Name: record.NickName, connection: conn})
			fmt.Println("add connection", s.Connections)
			s.send(conn, response.Marshal())
			return
		}
		response.SetError(err.Error())
		s.send(conn, response.Marshal())
	case command.RegisterUser:
		payload := command.UserLoginPayload{}
		response := command.Response{}
		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			response.SetError(err.Error())
			s.send(conn, response.Marshal())
			return
		}
		record := db.Record{
			Email:    payload.Email,
			Password: payload.Password,
			NickName: payload.NickName,
		}

		if !s.DB.IsEmailUnique(record.Email) {
			response.SetError("this email is already used")
			s.send(conn, response.Marshal())
			return
		}

		s.DB.AddRecord(record)
		response.SetPayload("Welcome " + payload.NickName)
		s.addConnection(conn.RemoteAddr().String(), connectionInfo{Name: record.NickName, connection: conn})
		s.send(conn, response.Marshal())
	case command.Quit:
		s.Connections.Delete(conn.RemoteAddr().String())
	case command.ActiveUsers:
		result := ""
		response := command.Response{}
		s.Connections.Range(func(k, v interface{}) bool {
			if info, ok := v.(connectionInfo); ok {
				result += info.Name + ","
			} else {
				fmt.Println("Cannot cast to connectionInfo")
			}
			return true
		})

		result = strings.TrimSuffix(result, ",")
		response.SetPayload(result)
		s.send(conn, response.Marshal())
	case command.SendMessage:
		sender := func() string {
			value, status := s.Connections.Load(conn.RemoteAddr().String())
			if !status {
				fmt.Println("cannot find user info for conn.RemoteAddr().String()")
				return "unknown"
			}

			if info, ok := value.(connectionInfo); ok {
				return info.Name
			}
			fmt.Println("cannot cast to connectionInfo")
			return ""
		}()
		message := "> " + sender + ": " + string(c.Payload)
		response := command.Response{}
		response.SetPayload(message)
		s.broadcast(response.Marshal(), conn.RemoteAddr().String())
	}
}

func (s *server) Destroy() {
	s.listener.Close()
	s.DB.Close()
}

func (s *server) Run() {
	for {
		conn, _ := s.listener.Accept()
		s.addConnection(conn.RemoteAddr().String(), connectionInfo{connection: conn})
		go s.process(conn.RemoteAddr().String())
	}
}
