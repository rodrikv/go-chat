package main

import "github.com/gin-gonic/gin"

func MessagesHandlerFunc(getMessages OnGetMessages) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		ms, err := getMessages(id)

		if err != nil {
			c.JSON(500, Error("unable to process request"))
			return
		}

		c.JSON(200, gin.H{
			"messages": ms,
		})
		return
	}
}

func ChatHandlerFunc(recieve OnRecieveMessage) func(c *gin.Context) {
	return func(c *gin.Context) {
		rm := c.MustGet(requestMessageKey).(*string)
		ci := c.MustGet(chatIdKey).(string)

		r, err := recieve(ci, *rm)

		if err != nil {
			c.JSON(500, Error("unable to process request"))
			return
		}

		c.Set(responseMessageKey, r)
	}
}
