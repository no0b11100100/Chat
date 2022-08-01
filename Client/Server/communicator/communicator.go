package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/app"
	"Chat/Client/Server/common"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Communicator struct {
	baseService  *BaseService
	chatService  *ChatService
	remoteServer *app.RemoteServer
	notification common.ChannelType
}

func NewCommunicator() *Communicator {
	ch := make(common.ChannelType)
	remoteServer := app.NewRemoteServer(ch)
	c := &Communicator{baseService: NewBaseService(remoteServer), chatService: NewChatService(remoteServer), remoteServer: remoteServer, notification: ch}

	return c
}

func (c *Communicator) runGRPCClient() {
	gprcServer := grpc.NewServer()
	api.RegisterBaseServer(gprcServer, c.baseService)
	api.RegisterChatServer(gprcServer, c.chatService)
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := gprcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func (c *Communicator) Serve() {
	go c.remoteServer.Serve()
	c.runGRPCClient()
}
