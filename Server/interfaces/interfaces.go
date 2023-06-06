package interfaces

import (
	"Chat/Server/api"
)

type UserServiceDatabase interface {
	IsEmailUnique(string) bool
	ValidateUser(email, password string) bool
	RegisterUser(api.SignUp) bool
	AddUserToChat(userID string, chatID string)
}

type ChatServiceDatabase interface {
	GetUserChats(string) []api.Chat
	GetMessages(string) []api.Message
	AddMessage(msg api.Message) error
	// GetChatParticipants(string) []string
}

type CalendarServiceDatabase interface {
	AddMeeting(userID string, meeting api.Meeting) error
	GetMeetings(userID string, startDay string, endDay string) []api.Meeting
}
