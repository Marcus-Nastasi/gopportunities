package router

import (
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "pong"})
	})

	v1 := server.Group("/api/v1")
	{
		v1.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"status": "ok"})
		})
	}
}
