package domain

import "github.com/gin-gonic/gin"

type Message struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func MessageSuccess(c *gin.Context, data interface{}, message string, status int) {
	res := Message{
		Status: status,
		Message: message,
		Data: data,
	}

	c.JSON(status, res)
}

func MessageError(c *gin.Context, message string, status int) {
	res := Message{
		Status: status,
		Message: message,
		Data: nil,
	}

	c.JSON(status, res)
}