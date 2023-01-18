package api

import (
	"context"
	"encoding/json"
)

type ServerContext struct {
	ConnectionAddress string
}
type disconnectionCallback = func(string)

// enums
type MessageType int

const (
	Request      = 0
	Notification = 1
)

// structs
type MessageData struct {
	Endpoint string          `json:"Endpoint"`
	Topic    string          `json:"Topic"`
	Payload  json.RawMessage `json:"Payload"`
	Type     MessageType     `json:"Type"`
}

type ServerImpl interface {
	Serve()
	Stop()
}

type Server struct {
	servers []ServerImpl
	cancel  context.CancelFunc
}

func NewServer() *Server {
	return &Server{servers: make([]ServerImpl, 0)}
}

func (s *Server) AddServer(server ServerImpl) {
	s.servers = append(s.servers, server)
}

func (s *Server) Serve() {
	var ctx context.Context
	ctx, s.cancel = context.WithCancel(context.Background())
	for _, server := range s.servers {
		go func(serv ServerImpl) {
			serv.Serve()
		}(server)
	}

	<-ctx.Done()
}

func (s *Server) Stop() {
	s.cancel()
}
