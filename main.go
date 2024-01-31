package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cs := ChatService{
		OnRecieveMessage: func(chatId string, content string) (interface{}, error) {
			return Success("hello"), nil
		},
		OnStream: sampleHandler,
	}

	cs.Bind(r.Group("/api"))

	r.Run() // listen and serve on
}

func sampleHandler(messageChannel *chan any, eventNameChannel *chan string, DoneChannel *chan bool) {
	for i := 0; i < 5; i++ {
		// message := fmt.Sprintf("Message %d", i+1)
		message := struct {
			Message string `json:"message"`
			Count   int    `json:"count"`
		}{
			Message: fmt.Sprintf("Message %d", i+1),
			Count:   i + 1,
		}
		// Send the message to the client
		*messageChannel <- message
		*eventNameChannel <- "message"

		// Introduce a delay to simulate some processing
		time.Sleep(1 * time.Second)
	}
	// Close the channel when the messages are sent
	*DoneChannel <- true
}
