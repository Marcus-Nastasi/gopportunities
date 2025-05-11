package router

import (
	"github.com/Marcus-Nastasi/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func Initialize(h *handler.Handler) {
	server := gin.Default()
	routes(server, h)
	server.Run(":8000")
}
