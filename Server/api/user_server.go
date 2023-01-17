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
type UserInfo struct {
	UserID   string   `json:"UserID",omitempty bson:"userid"`
	Name     string   `json:"Name",omitempty bson:"name"`
	NickName string   `json:"NickName",omitempty bson:"nickname"`
	Photo    string   `json:"Photo",omitempty bson:"photo"`
	Chats    []string `json:"Chats",omitempty bson:"chats"`
	Email    string   `json:"Email",omitempty bson:"email"`
	Password string   `json:"Password",omitempty bson:"password"`
}

type UserInfoTagger struct {
	UserID   string
	Name     string
	NickName string
	Photo    string
	Chats    string
	Email    string
	Password string
}

func UserInfoTags() UserInfoTagger {
	return UserInfoTagger{
		UserID:   strings.ToLower("UserID"),
		Name:     strings.ToLower("Name"),
		NickName: strings.ToLower("NickName"),
		Photo:    strings.ToLower("Photo"),
		Chats:    strings.ToLower("Chats"),
		Email:    strings.ToLower("Email"),
		Password: strings.ToLower("Password"),
	}
}

type Response struct {
	Info          UserInfo       `json:"Info",omitempty bson:"info"`
	Status        ResponseStatus `json:"Status",omitempty bson:"status"`
	StatusMessage string         `json:"StatusMessage",omitempty bson:"statusmessage"`
}

type ResponseTagger struct {
	Info          string
	Status        string
	StatusMessage string
}

func ResponseTags() ResponseTagger {
	return ResponseTagger{
		Info:          strings.ToLower("Info"),
		Status:        strings.ToLower("Status"),
		StatusMessage: strings.ToLower("StatusMessage"),
	}
}

type SignIn struct {
	Email    string `json:"Email",omitempty bson:"email"`
	Password string `json:"Password",omitempty bson:"password"`
}

type SignInTagger struct {
	Email    string
	Password string
}

func SignInTags() SignInTagger {
	return SignInTagger{
		Email:    strings.ToLower("Email"),
		Password: strings.ToLower("Password"),
	}
}

type SignUp struct {
	Name              string `json:"Name",omitempty bson:"name"`
	NickName          string `json:"NickName",omitempty bson:"nickname"`
	Email             string `json:"Email",omitempty bson:"email"`
	Password          string `json:"Password",omitempty bson:"password"`
	ConfirmedPassword string `json:"ConfirmedPassword",omitempty bson:"confirmedpassword"`
	Photo             string `json:"Photo",omitempty bson:"photo"`
}

type SignUpTagger struct {
	Name              string
	NickName          string
	Email             string
	Password          string
	ConfirmedPassword string
	Photo             string
}

func SignUpTags() SignUpTagger {
	return SignUpTagger{
		Name:              strings.ToLower("Name"),
		NickName:          strings.ToLower("NickName"),
		Email:             strings.ToLower("Email"),
		Password:          strings.ToLower("Password"),
		ConfirmedPassword: strings.ToLower("ConfirmedPassword"),
		Photo:             strings.ToLower("Photo"),
	}
}

// server
type UserServiceConnectionCallback = func(string, UserServiceNotifier)

type UserServiceServerImpl interface {
	SignIn(ServerContext, SignIn) Response
	SignUp(ServerContext, SignUp) Response
}

type UserServiceNotifier interface {
}

type UserServiceNotificator struct {
	conn net.Conn
}

type UserServiceServer struct {
	listener              net.Listener
	impl                  UserServiceServerImpl
	notifierObservers     []UserServiceConnectionCallback
	disconectionObservers []disconnectionCallback
}

func NewUserServiceServer(addr string) *UserServiceServer {
	ln, _ := net.Listen("tcp", addr)
	server := &UserServiceServer{
		listener:              ln,
		notifierObservers:     make([]UserServiceConnectionCallback, 0),
		disconectionObservers: make([]disconnectionCallback, 0),
	}

	return server
}

func (s *UserServiceServer) Stop() {
	if err := s.listener.Close(); err != nil {
		fmt.Println("UserServiceServer Stop error:", err)
	}
}

func (s *UserServiceServer) SetServerImpl(impl UserServiceServerImpl) {
	s.impl = impl
}

func (s *UserServiceServer) Serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
		}

		go s.processConnection(conn)
	}
}

func (s *UserServiceServer) processConnection(conn net.Conn) {
	fmt.Println("UserServiceServer Accept connection:", conn.RemoteAddr().String())
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

func (s *UserServiceServer) handleCommand(payload string, conn net.Conn) {
	recievedMessage := MessageData{}
	json.Unmarshal([]byte(payload), &recievedMessage)
	switch recievedMessage.Endpoint {
	case "UserService.SignIn":
		args := make([]json.RawMessage, 0)
		json.Unmarshal(recievedMessage.Payload, &args)
		var index int

		var data SignIn
		json.Unmarshal(args[index], &data)
		index++

		serverContex := ServerContext{ConnectionAddress: conn.RemoteAddr().String()}
		response := s.impl.SignIn(serverContex, data)
		bytes, _ := json.Marshal(response)
		messageToSend := recievedMessage
		messageToSend.Payload = json.RawMessage(bytes)
		responseData, _ := json.Marshal(messageToSend)
		conn.Write(responseData)
	case "UserService.SignUp":
		args := make([]json.RawMessage, 0)
		json.Unmarshal(recievedMessage.Payload, &args)
		var index int

		var data SignUp
		json.Unmarshal(args[index], &data)
		index++

		serverContex := ServerContext{ConnectionAddress: conn.RemoteAddr().String()}
		response := s.impl.SignUp(serverContex, data)
		bytes, _ := json.Marshal(response)
		messageToSend := recievedMessage
		messageToSend.Payload = json.RawMessage(bytes)
		responseData, _ := json.Marshal(messageToSend)
		conn.Write(responseData)
	}
}

// Events
func (s *UserServiceServer) SubscribeToNewConnectionEvent(observer UserServiceConnectionCallback) {
	s.notifierObservers = append(s.notifierObservers, observer)
}

func (s *UserServiceServer) SubscribeToDisconnectionEvent(observer disconnectionCallback) {
	s.disconectionObservers = append(s.disconectionObservers, observer)
}

func (s *UserServiceServer) emitNewConnectionEvent(conn net.Conn) {
	notificator := &UserServiceNotificator{conn: conn}

	for _, callback := range s.notifierObservers {
		callback(conn.RemoteAddr().String(), notificator)
	}
}

func (s *UserServiceServer) emitDisconnectionEvent(conn net.Conn) {
	for _, callback := range s.disconectionObservers {
		callback(conn.RemoteAddr().String())
	}
}
