package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/ginstream"
)

type ChatService struct {
	OnGetMessages    OnGetMessages
	OnRecieveMessage OnRecieveMessage
	OnAfterResponse  OnAfterResponse
	OnStream         func(MessageChannel *chan any, EventNameChannel *chan string)

	TimeOut time.Duration
}

func (cs *ChatService) Bind(r *gin.RouterGroup) {
	if cs.TimeOut == 0 {
		cs.TimeOut = 10 * time.Second
	}

	chatGroup := r.Group("/chat")
	{
		chatGroup.Use(ReadBody)
		if cs.OnAfterResponse == nil {
			cc := NewChatCache()
			cs.OnAfterResponse = func(chatId string, content string, response interface{}) {
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
		}

		if cs.OnAfterResponse != nil {
			chatGroup.Use(AfterResponseMiddlewareFunc(cs.OnAfterResponse))
		}
		if cs.OnRecieveMessage != nil {
			chatGroup.POST("/", ChatHandlerFunc(cs.OnRecieveMessage))
		}
		if cs.OnStream != nil {
			chatGroup.POST("/stream", ginstream.StreamHandler(cs.OnStream, cs.TimeOut))
		}
		if cs.OnGetMessages != nil {
			chatGroup.GET("/messages/:id", MessagesHandlerFunc(cs.OnGetMessages))
		}
	}
}

type OnGetMessages func(string) ([]interface{}, error)
type OnRecieveMessage func(chatId string, content string) (interface{}, error)
type OnAfterResponse func(chatId string, content string, response interface{})
