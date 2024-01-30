package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cs := ChatService{
		OnRecieveMessage: func(chatId string, content string) (interface{}, error) {
			return Success("hello"), nil
		},
	}

	cs.Bind(r.Group("/api"))

	r.Run() // listen and serve on
}
