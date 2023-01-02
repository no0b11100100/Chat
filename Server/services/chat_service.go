package services

import (
	"Chat/Server/api"
	log "Chat/Server/logger"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type ChatService struct {
	api.UnimplementedChatServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (chat *ChatService) GetChats(_ context.Context, userID *api.UserID) (*api.Chats, error) {
	var chats api.Chats
	return &chats, nil
}

func (chat *ChatService) GetMessages(_ context.Context, messageChan *api.MessageChan) (*api.Messages, error) {
	log.Info.Printf("GetMessages %+v\n", *messageChan)
	result := &api.Messages{}
	return result, nil
}

func (chat *ChatService) ReadMessage(context.Context, *api.ReadMessage) (*api.Status, error) {
	return nil, nil
}

func (chat *ChatService) EditChat(context.Context, *api.ChatData) (*api.Status, error) {
	return nil, nil
}

func (chat *ChatService) ChatChanged(*empty.Empty, api.Chat_ChatChangedServer) error {
	return nil
}

func (chat *ChatService) SendMessage(_ context.Context, message *api.ExchangedMessage) (*api.Status, error) {
	log.Info.Printf("SendMessage %+v\n", *message)
	return &api.Status{Status: "OK"}, nil
}

func (chat *ChatService) RecieveMessage(_ *empty.Empty, stream api.Chat_RecieveMessageServer) error {
	log.Info.Println("RecieveMessage")

	return nil
}
