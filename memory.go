package main

import "sync"

type ChatMessage struct {
	Content string `json:"content"`
	Roll    string `json:"roll"`
}

type ChatCache struct {
	mu       sync.RWMutex
	messages map[string][]ChatMessage
}

func NewChatCache() *ChatCache {
	return &ChatCache{
		messages: make(map[string][]ChatMessage),
	}
}

func (c *ChatCache) SaveMessage(msg ChatMessage, chatID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.messages[chatID] = append(c.messages[chatID], msg)
}

func (c *ChatCache) GetMessages(chatID string) ([]ChatMessage, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	messages, ok := c.messages[chatID]
	return messages, ok
}
