package gochat

type Message struct {
	Content string `json:"content"`
	ChatID  string `json:"chat_id"`
}
