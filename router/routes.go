package router

import (
	"github.com/Marcus-Nastasi/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func routes(server *gin.Engine, h *handler.Handler) {
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "pong"})
	})

	v1 := server.Group("/v1/api")
	{
		v1.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]string{"status": "ok"})
		})

		v1.GET("", h.GetHandler)
		v1.GET("/:id", h.GetSingleHandler)
		v1.POST("", handler.CreateHandler)
		v1.PATCH("/:id", h.UpdateHandler)
		v1.DELETE("/:id", h.DeleteHandler)
	}
}
