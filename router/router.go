package router

import "github.com/gin-gonic/gin"

func Initialize() {
	server := gin.Default()
	routes(server)
	server.Run(":8000")
}
