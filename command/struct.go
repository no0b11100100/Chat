package command

import "encoding/json"

type Command struct {
	ID      CommandID `json:"id"`
	Payload []byte    `json:"payload"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

type Response struct {
	Status  SereverStaus `json:"status"`
	Payload string       `json:"payload,omitempty"`
}

func (p *UserLoginPayload) Marshal() []byte {
	if payload, err := json.Marshal(p); err == nil {
		return payload
	}
	return []byte{}
}

func (r *Response) Marshal() []byte {
	if payload, err := json.Marshal(r); err == nil {
		return payload
	}
	return []byte{}
}

func (r *Response) SetError(errorString string) {
	r.Status = Fail
	r.Payload = errorString
}

func (r *Response) SetPayload(payload string) {
	r.Status = OK
	r.Payload = payload
}
