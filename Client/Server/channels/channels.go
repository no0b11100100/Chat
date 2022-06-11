package channels

import (
	"Chat/Client/Server/common"
)

type Channels struct {
	UserChan chan string //log in, register, log out
	ChatChan chan string
}

func NewChannels() *Channels {
	return &Channels{
		make(chan string),
		make(chan string),
	}
}

func (c *Channels) Notify(command common.Command) {
	switch command.Type {
	case common.LogIn, common.Register:
		c.UserChan <- string(command.Payload)
	default:
		c.ChatChan <- string(command.Payload)
	}
}
