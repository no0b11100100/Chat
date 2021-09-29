package command

import "encoding/json"

type CommandID int

const (
	StartConnection CommandID = iota
	LogInUser
	RegisterUser
	ActiveUsers
	SendMessage
	Quit
)

type Command struct {
	ID      CommandID `json:"id"`
	Payload []byte    `json:"payload"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

type SereverStaus uint32

const (
	OK SereverStaus = iota
	Fail
)

type Response struct {
	Status  SereverStaus `json:"status"`
	Payload string       `json:"payload,omitempty"`
}

func (p *UserLoginPayload) Marshal() []byte {
	if payload, err := json.Marshal(p); err == nil {
		return append(payload, '\n')
	}
	return []byte{'\n'}
}

func (r *Response) Marshal() []byte {
	if payload, err := json.Marshal(r); err == nil {
		return append(payload, '\n')
	}
	return []byte{'\n'}
}

func (r *Response) SetError(errorString string) {
	r.Status = Fail
	r.Payload = errorString
}

func (r *Response) SetPayload(payload string) {
	r.Status = OK
	r.Payload = payload
}
