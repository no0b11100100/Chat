package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
)

type ChatService struct {
	database interfaces.ChatServiceDatabase

	messageNotificationCh chan api.Message
}

func NewChatService(database interfaces.ChatServiceDatabase) *ChatService {
	return &ChatService{
		database:              database,
		messageNotificationCh: make(chan api.Message),
	}
}

func (chat *ChatService) SendMessage(_ api.ServerContext, message api.Message) api.Status {
	chat.messageNotificationCh <- message
	return api.Status{Status: api.OK}
}

func (chat *ChatService) GetUserChats(_ api.ServerContext, userID string) []api.Chat {
	return []api.Chat{} //chat.database.GetUserChats(userID)
}

func (chat *ChatService) GetChatMessages(_ api.ServerContext, chatID string) []api.Message {
	return []api.Message{} //chat.database.GetMessages(chatID)
}
