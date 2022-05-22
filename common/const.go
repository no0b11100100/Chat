package common

type CommandType int

const (
	LogIn CommandType = iota
	Register
	SendMessage
	LeaveChat
	CreateChat
	GetMessages
	GetParticipants
	AddUserToChat
)
