package main

import (
	"github.com/Marcus-Nastasi/gopportunities/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	router.Router(server)
	server.Run(":8000")
}
