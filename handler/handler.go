package handler

import (
	"net/http"

	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"github.com/Marcus-Nastasi/gopportunities/usecases"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger = config.NewLogger("Handler")
)

func GetHandler(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"get": "all"})
}

func GetSingleHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"get": id})
}

func CreateHandler(ctx *gin.Context) {
	var opportunitie schemas.Opening
	err := ctx.BindJSON(&opportunitie)
	if err != nil {
		logger.Errorf("Error binding json: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	o, err := usecases.CreateOpportunitie(opportunitie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	ctx.JSON(201, o)
}

func UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"update": id})
}

func DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, map[string]string{"delete": id})
}
