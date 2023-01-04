package main

import (
	"Chat/Server/database"
	"net"
	"os"

	"Chat/Server/api"
	log "Chat/Server/logger"
	"Chat/Server/services"

	"google.golang.org/grpc"
)

type Server struct {
	database database.Database
	listener net.Listener
	server   *grpc.Server
}

func NewServer() *Server {
	log.Info.Println("Start server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Error.Println("New Server err:", err)
		os.Exit(1)
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	log.Info.Println(addrs)

	s := &Server{
		database: database.NewDatabase(),
		listener: ln,
	}

	s.server = grpc.NewServer()
	s.database.Connect()
	s.initServices()

	return s
}

func (s *Server) initServices() {
	userService := services.NewUserService(s.database)
	chatService := services.NewChatService()

	api.RegisterChatServer(s.server, chatService)
	api.RegisterUserServer(s.server, userService)
}

func (s *Server) Serve() {
	err := s.server.Serve(s.listener)

	if err != nil {
		log.Error.Println(err)
	}

}

func (s *Server) Shutdown() {
	s.database.Close()
	s.server.GracefulStop()
}
