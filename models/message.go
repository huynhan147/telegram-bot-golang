package models

type ParamSendMessage struct {
	Method    string `json:"method"`
	ChatId    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}
type User struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

type Message struct {
	ID     int    `json:"message_id"`
	Sender User   `json:"from"`
	Date   int    `json:"date"`
	Text   string `json:"text,omitempty"`
}
