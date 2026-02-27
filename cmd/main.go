package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	todoRoutes(server)

	err := server.Run("localhost:8000")
	if err != nil {
		panic("error starting server")
	}
}
