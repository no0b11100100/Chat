package api

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
)

type BaseServer struct {
	ln net.Listener
}

func (b *BaseServer) Serve() {
	b.ln, _ = net.Listen("tcp", "localhost:1230")
	fmt.Println("Start base service server")
	for {
		conn, _ := b.ln.Accept()
		go func() {
			defer conn.Close()

			reader := bufio.NewReader(conn)
			tp := textproto.NewReader(reader)

			data, err := tp.ReadLine()

			if err != nil {
				fmt.Println("processConnection error", err)
				return
			}

			fmt.Println("processConnection message", data)
			// gen id
			connectionID := "1"
			fmt.Println("Generated connection id", connectionID)
			conn.Write([]byte(connectionID))
		}()
	}
}

func (b *BaseServer) Stop() {
	b.ln.Close()
}

type ServerContext struct {
	ConnectionID      string
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
	ConnectionID string          `json:"connectionid"`
	Endpoint     string          `json:"endpoint"`
	Topic        string          `json:"topic"`
	Payload      json.RawMessage `json:"payload"`
	Type         MessageType     `json:"type"`
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
	s := &Server{servers: make([]ServerImpl, 0)}
	s.AddServer(&BaseServer{})
	return s
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
