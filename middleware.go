package main

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context) {
	var m Message
	if err := c.BindJSON(&m); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Error("unable to process request"))
		return
	}
	a := m.Content
	log.Println(a, &a)
	c.Set(requestMessageKey, &a)
	c.Set(chatIdKey, m.ChatID)
	c.Next()
}

func AfterResponseMiddlewareFunc(save OnAfterResponse) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		if r, exists := c.Get(responseMessageKey); exists {
			rm, _ := c.MustGet(requestMessageKey).(*string)
			ci, _ := c.MustGet(chatIdKey).(string)
			resm, _ := r.(string)

			log.Println("Chat id: ", ci, "Request message: ", *rm, "Response message: ", resm)

			save(ci, *rm, resm)
			c.JSON(http.StatusOK, r)
		}
	}
}

func LogMessagesMiddleware(c *gin.Context) {
	c.Next()

	ci, _ := c.MustGet(chatIdKey).(string)
	rm, _ := c.MustGet(requestMessageKey).(*string)
	r, exists := c.Get(responseMessageKey)

	if exists {
		log.Println("Chat id: ", ci, "Request message: ", *rm, "Response message: ", r)
	} else {
		log.Println("Chat id: ", ci, "Request message: ", *rm, "No reponse message")
	}
}
