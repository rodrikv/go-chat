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
