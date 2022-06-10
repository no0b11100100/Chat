package app

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/communicator"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LocalServer struct {
	UIClient           api.BaseServer
	RemoteServerClient interface{}
}

func NewLocalServer() *LocalServer {
	return &LocalServer{
		UIClient: communicator.NewCommunicator(),
	}
}

func (s *LocalServer) Serve() {
	gprcServer := grpc.NewServer()
	api.RegisterBaseServer(gprcServer, s.UIClient)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := gprcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
