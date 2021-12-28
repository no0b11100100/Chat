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
	connection     net.Conn
	Name           string
	ActiveRoomName string
}

type Room struct {
	Name          string
	Pariticipants []connectionInfo
}

func (r *Room) addNewParticipant(user connectionInfo) {
	for _, p := range r.Pariticipants {
		if p.connection.RemoteAddr() == user.connection.RemoteAddr() {
			break
		}
	}

	r.Pariticipants = append(r.Pariticipants, user)
}

func (r *Room) removeUser(user string) {
	remove := func(slice []connectionInfo, s int) []connectionInfo {
		return append(slice[:s], slice[s+1:]...)
	}

	for index, p := range r.Pariticipants {
		if p.connection.RemoteAddr().String() == user {
			r.Pariticipants = remove(r.Pariticipants, index)
		}
	}
}

func (r *Room) sendMessage(msg []byte, sender string) {
	for _, p := range r.Pariticipants {
		if p.connection.RemoteAddr().String() == sender {
			continue
		}
		p.connection.Write(append(msg, '\n'))
	}
}

type server struct {
	listener    net.Listener
	Connections *sync.Map
	DB          db.DBInterface
	Rooms       map[string]*Room
}

func NewServer() *server {
	fmt.Println("Start server...")
	ln, _ := net.Listen("tcp", ":8081")
	return &server{
		listener:    ln,
		Connections: new(sync.Map),
		DB:          db.NewDataBase(),
		Rooms: map[string]*Room{
			"common": &Room{Name: "common"},
		}, //make(map[string]*Room),
	}

}

func (s *server) addConnection(key string, value *connectionInfo) {
	s.Connections.Store(key, value)
}

func (s *server) send(conn net.Conn, payload []byte) {
	conn.Write(append(payload, '\n'))
}

func (s *server) broadcast(payload []byte, except string) {
	s.Connections.Range(func(k, v interface{}) bool {
		if info, ok := v.(*connectionInfo); ok {
			conn := info.connection
			if conn.RemoteAddr().String() != except && info.ActiveRoomName == "" {
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
	conn := c.(*connectionInfo).connection
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
			s.addConnection(conn.RemoteAddr().String(), &connectionInfo{Name: record.NickName, connection: conn})
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
		s.addConnection(conn.RemoteAddr().String(), &connectionInfo{Name: record.NickName, connection: conn})
		s.send(conn, response.Marshal())
	case command.Quit:
		s.Connections.Delete(conn.RemoteAddr().String())
	case command.ActiveUsers:
		response := command.Response{}
		result := ""
		value, status := s.Connections.Load(conn.RemoteAddr().String())
		if !status {
			response.SetError("server error")
			s.send(conn, response.Marshal())
			return
		}

		user, ok := value.(*connectionInfo)
		if !ok {
			response.SetError("server error")
			s.send(conn, response.Marshal())
			return
		}

		roomName := user.ActiveRoomName

		s.Connections.Range(func(k, v interface{}) bool {
			if info, ok := v.(*connectionInfo); ok && info.ActiveRoomName == roomName {
				result += info.Name + ","
			} else {
				fmt.Println("Cannot cast to connectionInfo")
			}
			return true
		})

		result = strings.TrimSuffix(result, ",")
		response.SetPayload(result)
		s.send(conn, response.Marshal())
	case command.JoinToRoom:
		payload := command.RoomInfo{}
		response := command.Response{}

		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			response.SetError(err.Error())
			s.send(conn, response.Marshal())
			return
		}

		value, status := s.Connections.Load(conn.RemoteAddr().String())
		if !status {
			response.SetError("invalid user")
			s.send(conn, response.Marshal())
			return
		}

		info, ok := value.(*connectionInfo)
		if !ok {
			response.SetError("server error")
			s.send(conn, response.Marshal())
			return
		}

		if _, ok := s.Rooms[payload.Name]; !ok {
			response.SetError("wrong room name")
			s.send(conn, response.Marshal())
			return
		}

		info.ActiveRoomName = payload.Name

		s.Rooms[payload.Name].addNewParticipant(*info)
		response.SetPayload("Welcome in " + payload.Name)
		s.send(conn, response.Marshal())

		response = command.Response{}
		response.SetPayload(info.Name + " joined")
		s.Rooms[payload.Name].sendMessage(response.Marshal(), conn.RemoteAddr().String())
	case command.CreateRoom:
		payload := command.RoomInfo{}
		response := command.Response{}

		if err := json.Unmarshal(c.Payload, &payload); err != nil {
			response.SetError(err.Error())
			s.send(conn, response.Marshal())
			return
		}

		if _, ok := s.Rooms[payload.Name]; ok {
			response.SetError("this room already exists")
			s.send(conn, response.Marshal())
			return
		}

		s.Rooms[payload.Name] = &Room{Name: payload.Name, Pariticipants: make([]connectionInfo, 0, 1)}

		value, status := s.Connections.Load(conn.RemoteAddr().String())
		if !status {
			response.SetError("invalid user")
			s.send(conn, response.Marshal())
			return
		}

		info, ok := value.(*connectionInfo)

		if !ok {
			return
		}

		info.ActiveRoomName = payload.Name

		s.Rooms[payload.Name].addNewParticipant(*info)
		response.SetPayload("Welcome in " + payload.Name)
		s.send(conn, response.Marshal())

		fmt.Println(s.Rooms)
	case command.LeaveRoom:
		response := command.Response{}

		value, status := s.Connections.Load(conn.RemoteAddr().String())
		if !status {
			response.SetError("invalid user")
			s.send(conn, response.Marshal())
			return
		}

		info, ok := value.(*connectionInfo)

		if !ok {
			return
		}

		if info.ActiveRoomName == "" {
			response.SetError("you are not in room")
			s.send(conn, response.Marshal())
			return
		}

		s.Rooms[info.ActiveRoomName].removeUser(conn.RemoteAddr().String())
		response.SetPayload("You leave " + info.ActiveRoomName)
		s.send(conn, response.Marshal())

		response = command.Response{}
		response.SetPayload(info.Name + " leave the room")
		s.Rooms[info.ActiveRoomName].sendMessage(response.Marshal(), "")
		info.ActiveRoomName = ""
	case command.SendMessage:
		value, status := s.Connections.Load(conn.RemoteAddr().String())

		if !status {
			return
		}

		sender, ok := value.(*connectionInfo)
		if !ok {
			return
		}

		message := "> " + sender.Name + ": " + string(c.Payload)
		response := command.Response{}
		response.SetPayload(message)

		if sender.ActiveRoomName == "" {
			s.broadcast(response.Marshal(), conn.RemoteAddr().String())
		} else {
			s.Rooms[sender.ActiveRoomName].sendMessage(response.Marshal(), conn.RemoteAddr().String())
		}
	}
}

func (s *server) Destroy() {
	s.listener.Close()
	s.DB.Close()
}

func (s *server) Run() {
	for {
		conn, _ := s.listener.Accept()
		s.addConnection(conn.RemoteAddr().String(), &connectionInfo{connection: conn})
		go s.process(conn.RemoteAddr().String())
	}
}
