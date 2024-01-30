package main

import (
	"github.com/gin-gonic/gin"
)

func ChatHandlerFunc(recieve OnRecieveMessage) func(c *gin.Context) {
	return func(c *gin.Context) {
		m := c.MustGet(requestMessageKey).(*Message)

		r, err := recieve(m.ChatID, m.Content)

		if err != nil {
			c.JSON(500, Error("unable to process request"))
			return
		}

		c.Set(responseMessageKey, r)
	}
}
