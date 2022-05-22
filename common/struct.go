package common

type CommandResponce struct {
	Error   string  `json:"error"`
	Command Command `json:"command"`
}

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name,omitempty"`
	NickName string `json:"nickname,omitempty"`
}

type Command struct {
	Type    CommandType `json:"type"`
	Payload []byte      `json:"payload,omitempty"`
}
