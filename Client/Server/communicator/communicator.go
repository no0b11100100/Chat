package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/app"
	"Chat/Client/Server/channels"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Communicator struct {
	baseService  *BaseService
	remoteServer *app.RemoteServer
	channels     *channels.Channels
}

func NewCommunicator() *Communicator {
	// ch := make(chan string)
	remoteServer := app.NewRemoteServer()
	channels := channels.NewChannels()
	c := &Communicator{NewBaseService(remoteServer, channels.UserChan), remoteServer, channels}

	return c
}

func (c *Communicator) runGRPCClient() {
	gprcServer := grpc.NewServer()
	api.RegisterBaseServer(gprcServer, c.baseService)
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := gprcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func (c *Communicator) Serve() {
	go c.remoteServer.Serve(c.channels)
	c.runGRPCClient()
}
