package gochat

import "github.com/gin-gonic/gin"

func Error(text string) gin.H {
	return gin.H{
		"error": text,
	}
}

func Success(text string) Response {
	return Response{
		Message: text,
	}
}

type Response struct {
	Message string `json:"message"`
}
