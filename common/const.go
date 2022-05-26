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
	Notify
	Max
)

type MessageType int

const (
	Normal MessageType = iota
	Info
	Forwarded
	Replied
)

type ChatType int

const (
	Private ChatType = iota
	Channel
)

// Notification fields
const (
	Chats        = "chats"
	Participants = "participants"
)
