package models

type User struct {
	ID int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Chat struct {
	ID int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Type string `json:"type"`

}

type Message struct {
	ID int `json:"message_id"`
	Sender User `json:"from"`
	Date int `json:"date"`
	Text string `json:"text,omitempty"`

}
