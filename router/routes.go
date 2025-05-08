package router

import "github.com/gin-gonic/gin"

func routes(server *gin.Engine) {
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "pong"})
	})

	v1 := server.Group("/v1/api")
	{
		v1.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"status": "ok"})
		})
	}
}
