package common

type Command struct {
	ID string `json:"id"`
	Type    CommandType   `json:"type"`
	Status  CommandStatus `json:"status"`
	Payload []byte        `json:"payload,omitempty"`
}

type CommandResponce struct {
	Type    ResponseType `json:"type"`
	Command Command      `json:"command"`
}

type User struct {
	ID       string `bson:"user_id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name,omitempty"`
	NickName string `bson:"nickname" json:"nickname,omitempty"`
}

type ChannelType chan Command
