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
	serviceConnection := services.NewServiceConnection()

	userService := api.NewUserServiceServer("localhost:1234")
	userServiceImpl := services.NewUserService(serviceConnection, s.database)
	userService.SetServerImpl(userServiceImpl)

	chatService := api.NewChatServiceServer("localhost:1235")
	chatServiceImpl := services.NewChatService(s.database)
	chatService.SetServerImpl(chatServiceImpl)

	calendarService := api.NewCalendarServiceServer("localhost:1236")
	calendarServiceImpl := services.NewCalendarService(serviceConnection, s.database)
	calendarService.SetServerImpl(calendarServiceImpl)

	chatService.SubscribeToNewConnectionEvent(chatServiceImpl.HandleNewConnection)
	calendarService.SubscribeToNewConnectionEvent(calendarServiceImpl.HandleNewConnection)
	userService.SubscribeToDisconnectionEvent(userServiceImpl.HandleDisconnect)

	s.server.AddServer(userService)
	s.server.AddServer(chatService)
	s.server.AddServer(calendarService)
}

func (s *Server) Serve() {
	s.server.Serve()
}

func (s *Server) Shutdown() {
	s.database.Close()
	s.server.Stop()
}
