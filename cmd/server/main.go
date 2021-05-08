package main

import (
	"datalogger/package/routes/helloworld"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")

	// Create gin engine
	router := gin.Default()

	// Register the path
	helloworld.Register(router)

	// Start listening with default port 8080
	router.Run()
}
