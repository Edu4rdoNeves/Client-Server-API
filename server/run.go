package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	server := gin.Default()
	Routes(server)

	server.Run(":" + os.Getenv("SERVER_PORT"))
}
