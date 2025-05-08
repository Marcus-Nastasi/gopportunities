package handler

import "github.com/gin-gonic/gin"

func GetHandler(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"get": "all"})
}

func GetSingleHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"get": id})
}

func CreateHandler(ctx *gin.Context) {
	ctx.JSON(201, map[string]string{"post": "create"})
}

func UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"update": id})
}

func DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"delete": id})
}
