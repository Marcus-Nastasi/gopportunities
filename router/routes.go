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

		v1.GET("",func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"get": "all"})
		})
		v1.GET("/:id",func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(200, map[string]string{"get": id})
		})
		v1.POST("",func(ctx *gin.Context) {
			ctx.JSON(201, map[string]string{"post": "create"})
		})
		v1.PATCH("/:id",func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(200, map[string]string{"update": id})
		})
		v1.DELETE("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(200, map[string]string{"delete": id})
		})
	}
}
