package main

import (
	"Chat/RemoteServer/common"
	log "Chat/RemoteServer/common/logger"
	"Chat/RemoteServer/database"
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
	"sync"
)

type responseType int

const (
	noResponse responseType = iota
	normal
	broadcast
)

type Handler func([]byte, string) (*common.CommandResponce, responseType)

type Server struct {
	// key - (ip address)user id, value - connection
	cache    *sync.Map
	listener net.Listener
	database database.Database
	handlers map[common.CommandType]Handler
	mu       *sync.Mutex
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
		mu:       new(sync.Mutex),
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
	s.mu.Lock()
	defer s.mu.Unlock()
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

	log.Info.Println("Command", c.Type)
	var responce *common.CommandResponce
	var t responseType
	if handler, ok := s.handlers[c.Type]; ok {
		responce, t = handler(c.Payload, c.ID)
	} else {
		fmt.Println("Unknown command", c.Type)
		responce.Command.Status = common.UnknownCommand
	}

	switch t {
	case normal:
		s.send(conn, *responce)
	case broadcast:
		s.send(conn, *responce)
		s.broadcastChat("", c.Payload)
	default:
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

	s.notify(conn, bytes)
}

func (s *Server) notify(conn net.Conn, bytes []byte) {
	log.Info.Println("Notify", conn.RemoteAddr().String())
	payload := base64.StdEncoding.EncodeToString(bytes)
	payload = payload + "\n"
	conn.Write([]byte(payload))
}

func (s *Server) broadcastChat(_ string, commandPayload []byte) {
	notification := common.CommandResponce{Type: common.Notification, Command: common.Command{Status: common.OK, Type: common.NotifyMessageCommand}}
	// /////
	// msg := api.ExchangedMessage{ChatId: "1", Message: &api.Message{MessageJson: string([]byte(`{"message": "notification"}`))}}
	// commandPayload, _ = json.Marshal(msg)
	// /////
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
