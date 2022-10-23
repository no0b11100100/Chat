package common

type CommandType int

const (
	SignIn CommandType = iota
	SignUp
	SignOut

	GetUserChatsCommand
	GetMessagesCommand
	SendMessageCommand
	NotifyMessageCommand

	Max
)

type CommandStatus int

const (
	SignInOK CommandStatus = iota
	SignUpOK
	SignInError
	SignUpInvalidEmail
	ServerError
	OK
	UnknownCommand
)

type ResponseType int

const (
	Response = iota
	Notification
)
