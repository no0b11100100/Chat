package command

import "encoding/json"

type CommandID int

const (
	StartConnection CommandID = iota
	LogInUser
	GuestUser
	RegisterUser
	ConfirmPasswd
	ActiveUsers
	ChangeNickName
	ChangePasswd
	SendMessage
	Quit
)

type Command struct {
	ID      CommandID       `json:"id"`
	Error   string          `json:"error"`
	Payload json.RawMessage `json:"payload"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

type MessagePayload struct {
	Message string `json:"message"`
	// Recievers
}
