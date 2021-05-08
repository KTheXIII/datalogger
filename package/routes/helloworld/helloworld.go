package helloworld

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message" binding:"required"`
}

type MessageReply struct {
	Reply string `json:"reply"`
}

func Register(router *gin.Engine) {
	router.GET("/hello-world", HelloWorld)
	router.POST("/hello-world", PostWorld)
}

func HelloWorld(c *gin.Context) {
	message := Message{Message: "Hello, World!"}
	c.JSON(http.StatusOK, message)
}

func PostWorld(c *gin.Context) {
	var messageJSON Message
	if err := c.BindJSON(&messageJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messag": "Error format!",
		})
		return
	}
	reply := MessageReply{Reply: messageJSON.Message}
	c.IndentedJSON(http.StatusOK, reply)
}
