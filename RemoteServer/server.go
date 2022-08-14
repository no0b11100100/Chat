package main

import (
	"Chat/RemoteServer/common"
	log "Chat/RemoteServer/common/logger"
	"Chat/RemoteServer/database"
	api "Chat/RemoteServer/structs"
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
	"sync"
)

type Handler func([]byte) *common.CommandResponce

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

	var responce *common.CommandResponce
	if handler, ok := s.handlers[c.Type]; ok {
		responce = handler(c.Payload)
	} else {
		fmt.Println("Unknown command", c.Type)
		responce.Command.Status = common.UnknownCommand
	}

	if responce != nil {
		s.send(conn, *responce)
	} else {
		log.Info.Println("Skip response for", c.Type)
	}
}

func (s *Server) addHandlers() {
	s.handlers[common.SignIn] = s.SignIn
	s.handlers[common.SignUp] = s.SignUp

	s.handlers[common.GetUserChatsCommand] = s.GetUserChats
	s.handlers[common.GetMessagesCommand] = s.GetMessages
	s.handlers[common.SendMessageCommand] = s.SendMessage
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

func (s *Server) notify(conn net.Conn, bytes []byte) {
	payload := base64.StdEncoding.EncodeToString(bytes)
	payload = payload + "\n"
	conn.Write([]byte(payload))
}

func (s *Server) broadcastChat(_ string, commandPayload []byte) {
	notification := common.CommandResponce{Type: common.Notification, Command: common.Command{Status: common.OK, Type: common.NotifyMessageCommand}}
	/////
	msg := api.ExchangedMessage{ChatId: "1", Message: &api.Message{MessageJson: string([]byte(`{"message": "notification"}`))}}
	commandPayload, _ = json.Marshal(msg)
	/////
	notification.Command.Payload = commandPayload
	payload, err := json.Marshal(notification)
	if err != nil {
		log.Warning.Println(err)
	}
	s.cache.Range(func(k, v interface{}) bool {
		if conn, ok := v.(net.Conn); ok {
			s.notify(conn, payload)
		} else {
			log.Error.Printf("Cannot cast %+T to net.conn", v)
		}

		return true
	})
}
