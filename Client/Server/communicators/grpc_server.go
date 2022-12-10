package communicators

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	log "Chat/Client/Server/common/logger"
	"Chat/Client/Server/interfaces"
	"Chat/Client/Server/services"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	services []interfaces.Service
	connect  func()
}

func NewGRPCServer(addr string, communicator interfaces.RemoteCommunicator) *gRPCServer {
	server := &gRPCServer{}
	communicator.AddNotificationObserver(server)
	chatService := services.NewChatService(communicator)
	server.services = append(server.services, chatService)
	userService := services.NewUserService(communicator)
	server.services = append(server.services, userService)

	gprcServer := grpc.NewServer()
	api.RegisterChatServer(gprcServer, chatService)
	api.RegisterUserServer(gprcServer, userService)
	l, err := net.Listen("tcp", addr)
	server.connect = func() {
		if err := gprcServer.Serve(l); err != nil {
			log.Error.Println(err)
		}
	}

	if err != nil {
		log.Error.Println(err)
		return nil
	}

	return server
}

func (client *gRPCServer) Serve() {
	client.connect()
}

func (client *gRPCServer) Handle(c common.Command) {
	for _, service := range client.services {
		if service.IsHandledTopic(c.Type) {
			service.HandleTopic(c)
		}
	}
}
