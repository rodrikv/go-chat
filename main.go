package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Next()
	})

	cs := ChatService{
		OnRecieveMessage: func(content *string) any {
			return Success("hello")
		},
		OnStream:    sampleHandler,
		ContentPath: "message",
	}

	cs.Bind(r.Group("/api"))

	r.Run()
}

func sampleHandler(messageChannel *chan any, eventNameChannel *chan string, DoneChannel *chan bool, inputMessage *string) {
	for i := 0; i < 5; i++ {
		message := struct {
			Message string `json:"message"`
			Count   int    `json:"count"`
		}{
			Message: fmt.Sprintf("Message %s", *inputMessage),
			Count:   i + 1,
		}
		*messageChannel <- message
		*eventNameChannel <- "message"

		time.Sleep(1 * time.Second)
	}
	*DoneChannel <- true
}
