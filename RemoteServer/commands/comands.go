package commands

type CommandType int

const (
	LogIn CommandType = iota
	Register
	SendMessage
	JoinToChat
	LeaveChat
	CreateChat
	GetChatMessages
	GetChatParticipants
)

type Command struct {
	Type    CommandType `json:"type"`
	Payload string      `json:"payload"`
}
