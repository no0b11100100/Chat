package main

import (
	"Chat/Server/database"
	"net"
	"os"

	"Chat/Server/api"
	log "Chat/Server/logger"
	"Chat/Server/services"
)

type Server struct {
	database database.Database
	listener net.Listener
	server   *api.Server
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

	s.server = api.NewServer()
	s.database.Connect()
	s.initServices()

	return s
}

func (s *Server) initServices() {
	userService := api.NewUserServiceServer("localhost:8080")
	userServiceImpl := services.NewUserService(s.database)
	userService.SetServerImpl(userServiceImpl)

	chatService := api.NewChatServiceServer("localhost:8081")
	chatServiceImpl := services.NewChatService(s.database)
	chatService.SetServerImpl(chatServiceImpl)

	s.server.AddServer(userService)
	s.server.AddServer(chatService)
}

func (s *Server) Serve() {
	s.server.Serve()
}

func (s *Server) Shutdown() {
	s.database.Close()
	s.server.Stop()
}
