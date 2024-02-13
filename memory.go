package gochat

import (
	"log"
	"sync"
)

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

func (c *ChatCache) OnGetMessages(chatId string) ([]interface{}, error) {
	log.Println("all messages", c.messages)
	log.Println("getting messages for chatId: ", chatId)
	messages, _ := c.GetMessages(chatId)
	msList := make([]interface{}, len(messages))
	log.Printf("[%s]messages: %v", chatId, messages)
	// convert messages to interface
	for i, m := range messages {
		msList[i] = m
	}
	return msList, nil
}

func (c *ChatCache) SaveChatPair(chatId string, content string, response string) {
	c.SaveMessage(
		ChatMessage{
			Content: content,
			Roll:    assistantRoll,
		},
		chatId,
	)
	c.SaveMessage(
		ChatMessage{
			Content: response,
			Roll:    userRoll,
		},
		chatId,
	)

	log.Printf("[%s]messages: %v", chatId, c.messages[chatId])
}
