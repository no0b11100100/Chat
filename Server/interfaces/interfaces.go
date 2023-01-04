package interfaces

import api "Chat/Server/api"

type UserServiceDatabase interface {
	IsEmailUnique(string) bool
	ValidateUser(email, password string) (bool, string)
	RegisterUser(*api.SignUp) (bool, string)
	AddUserToChat(userID string, chatID string)
}

type ChatServiceDatabase interface {
	GetUserChats(string) api.Chats
	GetMessages(string, string, api.Direction) []*api.Message
	GetChatParticipants(string) []api.UserID
}
