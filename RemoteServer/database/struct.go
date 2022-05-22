package database

type Participant struct {
	ID       string `json:"id"`
	Name     string `json:"name,omiemptye"`
	Nickname string `json:"nickname,omiempty"`
	Photo    string `json:"photo,omiempty"`
}

type User struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Chats    []string `json:"chats,omiempty"` // chat ids
	Participant
}

type Message struct {
	Text string `json:"message"`
}

type Chat struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Cover        string    `json:"cover"`
	Participants []string  `json:"participants"` // user ids
	Messages     []Message `json:"messages"`
}
