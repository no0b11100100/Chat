package services

import (
	"Chat/Server/api"
	interfaces "Chat/Server/interfaces"
)

type ChatService struct {
	database interfaces.ChatServiceDatabase
	users    map[string]api.ChatServiceNotifier
}

func NewChatService(database interfaces.ChatServiceDatabase) *ChatService {
	s := &ChatService{
		database: database,
		users:    make(map[string]api.ChatServiceNotifier),
	}

	return s
}

func (chat *ChatService) SendMessage(ctx api.ServerContext, message api.Message) api.ResponseStatus {
	for _, notifier := range chat.users {
		//TODO: check chatID before notify
		notifier.RecieveMessage(message)
	}
	return api.OK
}

func (chat *ChatService) GetUserChats(_ api.ServerContext, userID string) []api.Chat {
	return chat.database.GetUserChats(userID)
}

func (chat *ChatService) GetChatMessages(_ api.ServerContext, chatID string) []api.Message {
	return chat.database.GetMessages(chatID)
}

func (chat *ChatService) HandleNewConnection(address string, notifier api.ChatServiceNotifier) {
	chat.users[address] = notifier
}

func (chat *ChatService) CallTo(api.ServerContext, string, string) api.CallStatus {
	return api.Connected
}

func (chat *ChatService) SendCallData(_ api.ServerContext, data api.CallData) {
	for _, notifier := range chat.users {
		//TODO: check chatID before notify
		notifier.NotifyCallData(data)
	}
}

func (chat *ChatService) HandleCallFrom(api.ServerContext, api.CallStatus) {}
