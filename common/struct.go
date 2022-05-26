package common

type Command struct {
	Type    CommandType `json:"type"`
	Payload []byte      `json:"payload"`
}

type CommandResponce struct {
	Error   string  `json:"error"`
	Command Command `json:"command"`
}

type Participant struct {
	ID       string `bson:"user_id"`
	Name     string `bson:"name" json:"name,omitempty"`
	Nickname string `bson:"nickname" json:"nickname,omitempty"`
	Photo    string `bson:"photo" json:"photo,omitempty"`
}

type User struct {
	Email    string   `bson:"email" json:"email"`
	Password string   `bson:"password" json:"password"`
	Chats    []string `bson:"chats"` // chat ids
	Participant
}

type Message struct {
	ChatID string      `bson:"-" json:"chat_id"`
	Text   MessageType `bson:"text" json:"text"`
	Sender string      `bson:"sender" json:"user_id"`
	Type   int         `bson:"type" json:"type"`
}

type Chat struct {
	ID           string    `bson:"id" json:"id"`
	Type         ChatType  `bson:"type" json:"type"`
	Title        string    `bson:"title" json:"title"`
	Cover        string    `bson:"cover" json:"cover"`
	Participants []string  `bson:"participants" json:"participants"` // user ids
	Messages     []Message `bson:"messages" json:"messages"`
}

type NewChatUser struct {
	ChatID          string `json:"chat_id"`
	UserID          string `json:"user_id"`
	IsExportHistiry bool   `json:"isExportHistory"`
}

type Notification struct {
	Field   string `json:"field"`
	Paylaod []byte `json:"payload"`
}
