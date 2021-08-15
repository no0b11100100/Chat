package commands

import "encoding/json"

type CommandID int

const (
	LogIn CommandID = iota
	LogInUser
	LogInGuest
	AllActiveUsers
	Send
	Quit
)

type Command struct {
	ID      CommandID       `json:"id"`
	Payload json.RawMessage `json:"payload"`
}
