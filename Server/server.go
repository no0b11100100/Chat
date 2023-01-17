package main

import (
	"Chat/Server/database"

	"Chat/Server/api"
	"Chat/Server/services"
)

type Server struct {
	database database.Database
	server   *api.Server
}

func NewServer() *Server {

	s := &Server{
		database: database.NewDatabase(),
	}

	s.server = api.NewServer()
	s.database.Connect()
	s.initServices()

	return s
}

func (s *Server) initServices() {
	userService := api.NewUserServiceServer("localhost:1234")
	userServiceImpl := services.NewUserService(s.database)
	userService.SetServerImpl(userServiceImpl)

	chatService := api.NewChatServiceServer("localhost:1235")
	chatServiceImpl := services.NewChatService(s.database)
	chatService.SetServerImpl(chatServiceImpl)

	chatService.SubscribeToNewConnectionEvent(chatServiceImpl.HandleNewConnection)

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
