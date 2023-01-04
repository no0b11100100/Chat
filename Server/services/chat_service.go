package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
	log "Chat/Server/logger"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type ChatService struct {
	api.UnimplementedChatServer
	database interfaces.ChatServiceDatabase

	messageNotificationCh chan api.ExchangedMessage
	userStreams           map[string]api.Chat_RecieveMessageServer //userId : stream
}

func NewChatService(database interfaces.ChatServiceDatabase) *ChatService {
	return &ChatService{
		database:              database,
		messageNotificationCh: make(chan api.ExchangedMessage),
		userStreams:           make(map[string]api.Chat_RecieveMessageServer),
	}
}

func (chat *ChatService) GetChats(_ context.Context, userID *api.UserID) (*api.Chats, error) {
	chats := chat.database.GetUserChats(userID.UserId)
	return &chats, nil
}

func (chat *ChatService) GetMessages(_ context.Context, messageChan *api.MessageChan) (*api.Messages, error) {
	log.Info.Printf("GetMessages %+v\n", *messageChan)
	messages := chat.database.GetMessages(messageChan.ChatId, messageChan.MessageId, messageChan.Direction)
	return messages, nil
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
	chat.messageNotificationCh <- *message
	return &api.Status{Status: "OK"}, nil
}

func (chat *ChatService) RecieveMessage(userID *api.UserID, stream api.Chat_RecieveMessageServer) error {
	log.Info.Println("RecieveMessage")
	chat.userStreams[userID.UserId] = stream
	for message := range chat.messageNotificationCh {
		log.Info.Println("ReciveMessage", message)
		participants := chat.database.GetChatParticipants(message.ChatId)
		for _, participant := range participants {
			stream, ok := chat.userStreams[participant.ChatId]
			if !ok {
				continue
			}
			if err := stream.Send(&message); err != nil {
				log.Error.Println(err)
			}
		}
	}
	return nil
}
