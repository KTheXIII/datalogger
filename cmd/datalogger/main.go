package main

import (
	"datalogger/package/routes/helloworld"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create gin engine
	router := gin.Default()

	// Register the path
	helloworld.Register(router)

	// Start listening with default port 8080
	router.Run("127.0.0.1:8080")
}
