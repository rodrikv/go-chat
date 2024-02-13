package gochat

type ChatMessageRetriever interface {
	GetMessages(chatID string) ([]Message, error)
}
