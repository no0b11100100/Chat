package communicator

import "Chat/Client/Server/common"

// type Communicator interface {
// 	api.BaseServer
// }

type RemoteServerInterface interface {
	Send(common.Command, common.ChannelType)
}
