package common

type CommandType int

const (
	SignIn CommandType = iota
	SignOut
	Max
)

type CommandStatus int

const (
	OK CommandStatus = iota
	SignInError
	SignOutInvalidEmail
	ServerError
	UnknownCommand
)

type ResponseType int

const (
	Response = iota
	Notification
)
