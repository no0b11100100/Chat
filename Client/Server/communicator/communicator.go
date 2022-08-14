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
	userService  *UserService
	chatService  *ChatService
	remoteServer *app.RemoteServer
	notification common.ChannelType
}

func NewCommunicator() *Communicator {
	ch := make(common.ChannelType)
	chatServiceNotification := make(common.ChannelType)
	remoteServer := app.NewRemoteServer(ch)
	c := &Communicator{userService: NewUserService(remoteServer), chatService: NewChatService(remoteServer, chatServiceNotification), remoteServer: remoteServer, notification: ch}
	go c.handleServerNotifications(ch, chatServiceNotification)

	return c
}

func (c *Communicator) handleServerNotifications(ch common.ChannelType, chatServiceChan common.ChannelType) {
	for {
		for data := range ch {
			if c.chatService.IsHandledTopic(data.Type) {
				chatServiceChan <- data
			}
		}
	}
}

func (c *Communicator) runGRPCClient() {
	gprcServer := grpc.NewServer()
	api.RegisterUserServer(gprcServer, c.userService)
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
