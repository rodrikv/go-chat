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
	OnStream         func(MessageChannel *chan any, EventNameChannel *chan string, DoneChannel *chan bool)

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
		if cs.OnAfterResponse == nil {
			cs.OnAfterResponse = func(chatId, content string, response interface{}) {
				cc.SaveChatPair(chatId, content, response)
			}
		}
		path := "message"

		chatGroup.Use(AfterResponseMiddlewareFunc(cs.OnAfterResponse))
		HandlerConf := ginstream.GeneralPurposeHandlerType{
			StreamHandlerFunc:    cs.OnStream,
			NonStreamHandlerFunc: sampleNonstreamHandler,

			Timeout:           100 * time.Millisecond,
			InputName:         &requestMessageKey,
			OutputName:        &responseMessageKey,
			StreamMessagePath: &path,
		}

		if cs.OnRecieveMessage != nil {
			chatGroup.POST("/", ginstream.GeneralPurposeHandler(HandlerConf))
		}
		if cs.OnStream != nil {
			// chatGroup.POST("/stream", ginstream.StreamHandler(cs.OnStream, cs.TimeOut))
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
type OnRecieveMessage func(chatId string, content string) (interface{}, error)
type OnAfterResponse func(chatId string, content string, response interface{})
