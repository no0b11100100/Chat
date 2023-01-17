// Code generated by goqface. DO NOT EDIT.
package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

//enums

// structs
type Message struct {
	MessageJSON string `json:"MessageJSON",omitempty bson:"messagejson"`
	ChatID      string `json:"ChatID",omitempty bson:"chatid"`
	SenderID    string `json:"SenderID",omitempty bson:"senderid"`
}

type MessageTagger struct {
	MessageJSON string
	ChatID      string
	SenderID    string
}

func MessageTags() MessageTagger {
	return MessageTagger{
		MessageJSON: strings.ToLower("MessageJSON"),
		ChatID:      strings.ToLower("ChatID"),
		SenderID:    strings.ToLower("SenderID"),
	}
}

type Chat struct {
	ChatID        string    `json:"ChatID",omitempty bson:"chatid"`
	Title         string    `json:"Title",omitempty bson:"title"`
	SecondLine    string    `json:"SecondLine",omitempty bson:"secondline"`
	LastMessage   string    `json:"LastMessage",omitempty bson:"lastmessage"`
	UnreadedCount int       `json:"UnreadedCount",omitempty bson:"unreadedcount"`
	Cover         string    `json:"Cover",omitempty bson:"cover"`
	Participants  []string  `json:"Participants",omitempty bson:"participants"`
	Messages      []Message `json:"Messages",omitempty bson:"messages"`
}

type ChatTagger struct {
	ChatID        string
	Title         string
	SecondLine    string
	LastMessage   string
	UnreadedCount string
	Cover         string
	Participants  string
	Messages      string
}

func ChatTags() ChatTagger {
	return ChatTagger{
		ChatID:        strings.ToLower("ChatID"),
		Title:         strings.ToLower("Title"),
		SecondLine:    strings.ToLower("SecondLine"),
		LastMessage:   strings.ToLower("LastMessage"),
		UnreadedCount: strings.ToLower("UnreadedCount"),
		Cover:         strings.ToLower("Cover"),
		Participants:  strings.ToLower("Participants"),
		Messages:      strings.ToLower("Messages"),
	}
}

// server
type ChatServiceConnectionCallback = func(string, ChatServiceNotifier)

type ChatServiceServerImpl interface {
	SendMessage(ServerContext, Message) ResponseStatus
	GetUserChats(ServerContext, string) []Chat
	GetChatMessages(ServerContext, string) []Message
}

type ChatServiceNotifier interface {
	RecieveMessage(Message)
}

type ChatServiceNotificator struct {
	conn net.Conn
}

func (n *ChatServiceNotificator) RecieveMessage(message Message) {
	args := make([]json.RawMessage, 0)
	var bytes []byte
	bytes, _ = json.Marshal(message)
	args = append(args, json.RawMessage(bytes))

	bytes, _ = json.Marshal(args)

	messageToSend := MessageData{}
	messageToSend.Endpoint = "ChatService.RecieveMessage"
	messageToSend.Payload = json.RawMessage(bytes)
	messageToSend.Type = Notification

	payloadToSend, _ := json.Marshal(messageToSend)

	n.conn.Write(payloadToSend)
}

type ChatServiceServer struct {
	listener              net.Listener
	impl                  ChatServiceServerImpl
	notifierObservers     []ChatServiceConnectionCallback
	disconectionObservers []disconnectionCallback
}

func NewChatServiceServer(addr string) *ChatServiceServer {
	ln, _ := net.Listen("tcp", addr)
	server := &ChatServiceServer{
		listener:              ln,
		notifierObservers:     make([]ChatServiceConnectionCallback, 0),
		disconectionObservers: make([]disconnectionCallback, 0),
	}

	return server
}

func (s *ChatServiceServer) Stop() {
	if err := s.listener.Close(); err != nil {
		fmt.Println("ChatServiceServer Stop error:", err)
	}
}

func (s *ChatServiceServer) SetServerImpl(impl ChatServiceServerImpl) {
	s.impl = impl
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
	fmt.Println("ChatServiceServer Accept connection:", conn.RemoteAddr().String())
	s.emitNewConnectionEvent(conn)
	defer conn.Close()
	defer func() { s.emitDisconnectionEvent(conn) }()

	for {
		reader := bufio.NewReader(conn)
		tp := textproto.NewReader(reader)

		data, err := tp.ReadLine()

		if err != nil {
			fmt.Println("processConnection error", err)
			break
		}

		fmt.Println("processConnection message", data)
		s.handleCommand(data, conn)
	}
}

func (s *ChatServiceServer) handleCommand(payload string, conn net.Conn) {
	recievedMessage := MessageData{}
	json.Unmarshal([]byte(payload), &recievedMessage)
	switch recievedMessage.Endpoint {
	case "ChatService.SendMessage":
		args := make([]json.RawMessage, 0)
		json.Unmarshal(recievedMessage.Payload, &args)
		var index int

		var message Message
		json.Unmarshal(args[index], &message)
		index++

		serverContex := ServerContext{ConnectionAddress: conn.RemoteAddr().String()}
		response := s.impl.SendMessage(serverContex, message)
		bytes, _ := json.Marshal(response)
		messageToSend := recievedMessage
		messageToSend.Payload = json.RawMessage(bytes)
		responseData, _ := json.Marshal(messageToSend)
		conn.Write(responseData)
	case "ChatService.GetUserChats":
		args := make([]json.RawMessage, 0)
		json.Unmarshal(recievedMessage.Payload, &args)
		var index int

		var userID string
		json.Unmarshal(args[index], &userID)
		index++

		serverContex := ServerContext{ConnectionAddress: conn.RemoteAddr().String()}
		response := s.impl.GetUserChats(serverContex, userID)
		bytes, _ := json.Marshal(response)
		messageToSend := recievedMessage
		messageToSend.Payload = json.RawMessage(bytes)
		responseData, _ := json.Marshal(messageToSend)
		conn.Write(responseData)
	case "ChatService.GetChatMessages":
		args := make([]json.RawMessage, 0)
		json.Unmarshal(recievedMessage.Payload, &args)
		var index int

		var chatID string
		json.Unmarshal(args[index], &chatID)
		index++

		serverContex := ServerContext{ConnectionAddress: conn.RemoteAddr().String()}
		response := s.impl.GetChatMessages(serverContex, chatID)
		bytes, _ := json.Marshal(response)
		messageToSend := recievedMessage
		messageToSend.Payload = json.RawMessage(bytes)
		responseData, _ := json.Marshal(messageToSend)
		conn.Write(responseData)
	}
}

// Events
func (s *ChatServiceServer) SubscribeToNewConnectionEvent(observer ChatServiceConnectionCallback) {
	s.notifierObservers = append(s.notifierObservers, observer)
}

func (s *ChatServiceServer) SubscribeToDisconnectionEvent(observer disconnectionCallback) {
	s.disconectionObservers = append(s.disconectionObservers, observer)
}

func (s *ChatServiceServer) emitNewConnectionEvent(conn net.Conn) {
	notificator := &ChatServiceNotificator{conn: conn}

	for _, callback := range s.notifierObservers {
		callback(conn.RemoteAddr().String(), notificator)
	}
}

func (s *ChatServiceServer) emitDisconnectionEvent(conn net.Conn) {
	for _, callback := range s.disconectionObservers {
		callback(conn.RemoteAddr().String())
	}
}
