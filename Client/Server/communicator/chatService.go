package communicator

import (
	"Chat/Client/Server/api"
	"context"
)

type ChatService struct {
	api.UnimplementedChatServer
}

func (c *ChatService) GetUserChats(context.Context, *api.UserID) (*api.UserChats, error) {
	return nil, nil
}

func (c *ChatService) GetChatInfo(context.Context, *api.ChatID) (*api.ChatInformation, error) {
	return nil, nil
}

func (c *ChatService) GetParticipantInfo(context.Context, *api.UserID) (*api.ParticipantInfo, error) {

	return nil, nil
}

func (c *ChatService) GetMessages(context.Context, *api.MessageChan) (*api.Messages, error) {
	return nil, nil
}

func (c *ChatService) MessageExchange(api.Chat_MessageExchangeServer) error {
	return nil
}
