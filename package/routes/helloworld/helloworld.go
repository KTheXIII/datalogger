package helloworld

import (
	"fmt"
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
	router.GET("/hello", HelloWorld)
	router.POST("/hello", PostWorld)
	router.POST("/hello-text", PostTextWorld)
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

func PostTextWorld(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	fmt.Println(reqBody)

	c.String(http.StatusOK, "Hello, World")
}
