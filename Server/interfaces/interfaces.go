package interfaces

import (
	"Chat/Server/api"
)

type UserServiceDatabase interface {
	IsEmailUnique(string) bool
	ValidateUser(email, password string) (bool, string)
	RegisterUser(api.SignUp) (bool, string)
	AddUserToChat(userID string, chatID string)
}

type ChatServiceDatabase interface {
	GetUserChats(string) []api.Chat
	GetMessages(string) []api.Message
	// GetChatParticipants(string) []string
}
