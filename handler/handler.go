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
	o, err := usecases.GetOpportunities()
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{"status": "not found or error"})
	}
	ctx.JSON(http.StatusOK, o)
}

func GetSingleHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	o, err := usecases.GetOpportunitie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{"status": "not found or error"})
	}
	ctx.JSON(http.StatusOK, o)
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
	ctx.JSON(http.StatusCreated, o)
}

func UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var err error
	var o schemas.Opening
	err = ctx.BindJSON(&o)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	o, err = usecases.UpdateOpportunitie(id, o)
	ctx.JSON(http.StatusOK, o)
}

func DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	o, err := usecases.DeleteOpportunitie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, o)
}
