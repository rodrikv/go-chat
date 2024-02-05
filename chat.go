package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uussoop/ginstream"
)

type ChatService struct {
	OnGetMessages    OnGetMessages
	OnRecieveMessage OnRecieveMessage
	OnAfterResponse  OnAfterResponse
	OnStream         func(MessageChannel *chan any, EventNameChannel *chan string, DoneChannel *chan bool, inputMessage *string)
	ContentPath      string

	TimeOut time.Duration
}

func (cs *ChatService) Bind(r *gin.RouterGroup) {
	cc := NewChatCache()
	if cs.TimeOut == 0 {
		cs.TimeOut = 10 * time.Second
	}
	chatGroup := r.Group("/chat")
	{
		chatGroup.Use(ReadBody)
		chatGroup.Use(LogMessagesMiddleware)
		if cs.OnAfterResponse == nil {
			cs.OnAfterResponse = cc.SaveChatPair
		}

		chatGroup.Use(AfterResponseMiddlewareFunc(cs.OnAfterResponse))
		HandlerConf := ginstream.GeneralPurposeHandlerType{
			StreamHandlerFunc:    cs.OnStream,
			NonStreamHandlerFunc: cs.OnRecieveMessage,

			Timeout:           60 * time.Second,
			InputName:         &requestMessageKey,
			OutputName:        &responseMessageKey,
			StreamMessagePath: &cs.ContentPath,
		}

		if cs.OnRecieveMessage != nil {
			chatGroup.POST("/", ginstream.GeneralPurposeHandler(HandlerConf))
		}
	}
	messageGroup := r.Group("/messages")
	{
		if cs.OnGetMessages == nil {
			messageGroup.GET("/:id", MessagesHandlerFunc(cc.OnGetMessages))
		} else {
			messageGroup.GET("/:id", MessagesHandlerFunc(cs.OnGetMessages))
		}
	}
}

func sampleNonstreamHandler(
	*string,
) any {
	return struct {
		message string
		status  int
	}{
		message: "hi",
		status:  123,
	}
}

type OnGetMessages func(string) ([]interface{}, error)
type OnRecieveMessage func(content *string) any
type OnAfterResponse func(chatId string, content string, response string)
