package common

type CommandType int

const (
	SignIn CommandType = iota
	SignUp
	SignOut

	GetUserChatsCommand
	GetChatInfoCommand
	GetParticipantInfoCommand
	GetMessagesCommand

	Max
)

type CommandStatus int

const (
	SignInOK CommandStatus = iota
	SignUpOK
	SignInError
	SignUpInvalidEmail
	ServerError
	UnknownCommand
)

type ResponseType int

const (
	Response = iota
	Notification
)
