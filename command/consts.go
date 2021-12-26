package command

type CommandID int

const EndMessageByte = '\n'

const (
	StartConnection CommandID = iota
	LogInUser
	RegisterUser
	ActiveUsers
	SendMessage
	Quit
)

type SereverStaus uint32

const (
	OK SereverStaus = iota
	Fail
)
