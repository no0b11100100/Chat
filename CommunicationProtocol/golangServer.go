package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
)

// type NotifierInfo interface {
// 	AddProperty(string, interface{})
// 	SetProperty(string, interface{})
// 	GetProperty(string) interface{}
// }

type Notifier interface {
	// NotifierInfo
	NotifyMessage(Message)
}

type Notificator struct {
	conn net.Conn
}

type connectionCallback = func(string, Notifier)
type disconnectionCallback = func(string)

func (n *Notificator) NotifyMessage(value Message) {
	//TODO: convert to response value and send back to client
	n.conn.Write(value.toBytes())
}

type ChatServiceServerImpl interface {
	SendMessage(string, Message)
	GetUserChats(string, UserID) []Chats
}

type ChatServiceServer struct {
	listener              net.Listener
	impl                  ChatServiceServerImpl
	notifierObservers     []connectionCallback
	disconectionObservers []disconnectionCallback
}

type RequestData struct {
	Endpoint string
	Topic    int64
	Payload  string
}

func (r *RequestData) toBytes() []byte {
	return []byte{}
}

func NewServerBase(addr string) *ChatServiceServer {
	ln, _ := net.Listen("tcp", addr)
	server := &ChatServiceServer{
		listener: ln,
	}

	return server
}

func (s *ChatServiceServer) Serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
		}

		go s.processConnection(conn)
	}
}

func (s *ChatServiceServer) processConnection(conn net.Conn) {
	fmt.Println("Accept connection:", conn.RemoteAddr().String())
	s.emitNewConnectionEvent(conn)
	defer conn.Close()
	defer func() { s.emitDisconnectionEvent(conn) }()

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

func (s *ChatServiceServer) handleCommand(payload string, conn net.Conn) {
	data := RequestData{}
	json.Unmarshal([]byte(payload), &data)
	switch data.Endpoint {
	case "ChatService.sendMessage":
		value := Message{}
		json.Unmarshal(data.Payload, &value)
		s.impl.SendMessage(conn.RemoteAddr().String(), value)
	case "ChatService.getUserChats":
		value := UserID{}
		json.Unmarshal(data.Payload, &value)
		result := s.impl.GetUserChats(conn.RemoteAddr().String(), value)
		bytes, _ := json.Marshal(result)
		response := RequestData{Endpoint: data.Endpoint, Topic: data.Topic, Payload: string(bytes)}
		conn.Write(response.toBytes())
	}
}

// Events
func (s *ChatServiceServer) SubscribeToNewConnectionEvent(observer connectionCallback) {
	s.notifierObservers = append(s.notifierObservers, observer)
}

func (s *ChatServiceServer) SubscribeToDisconnectionEvent(observer disconnectionCallback) {
	s.disconectionObservers = append(s.disconectionObservers, observer)
}

func (s *ChatServiceServer) emitNewConnectionEvent(conn net.Conn) {
	notificator := &Notificator{conn: conn}

	for _, callback := range s.notifierObservers {
		callback(conn.RemoteAddr().String(), notificator)
	}
}

func (s *ChatServiceServer) emitDisconnectionEvent(conn net.Conn) {
	for _, callback := range s.disconectionObservers {
		callback(conn.RemoteAddr().String())
	}
}
