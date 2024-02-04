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
	c.Set(requestMessageKey, &m.Content)
	c.Set(chatIdKey, m.ChatID)
	c.Next()
}

func AfterResponseMiddlewareFunc(save OnAfterResponse) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		if r, exists := c.Get(responseMessageKey); exists {
			rm, _ := c.MustGet(requestMessageKey).(*string)
			ci, _ := c.MustGet(chatIdKey).(string)
			save(ci, *rm, r)
			c.JSON(http.StatusOK, r)
		}
	}
}
