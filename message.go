package main

type ChatMessageRetriever interface {
	GetMessages(chatID string) ([]Message, error)
}
