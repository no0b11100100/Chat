package interfaces

import "Chat/Client/Server/common"

type NotificationObserver interface {
	Handle(common.Command)
}

type RemoteCommunicator interface {
	AddNotificationObserver(NotificationObserver)
	Sender
}

type Sender interface {
	Send(common.Command) chan common.Command
}

type Server interface {
	Serve()
}

type Service interface {
	IsHandledTopic(common.CommandType) bool
	HandleTopic(common.Command)
}
