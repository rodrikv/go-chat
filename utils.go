package gochat

import "log"

func saveMessage(cc *ChatCache, chatId string, content string, response interface{}) {
	cc.SaveMessage(
		ChatMessage{
			Content: content,
			Roll:    assistantRoll,
		},
		chatId,
	)

	r, _ := response.(Response)

	cc.SaveMessage(
		ChatMessage{
			Content: r.Message,
			Roll:    userRoll,
		},
		chatId,
	)

	log.Println(cc.GetMessages(chatId))
}
