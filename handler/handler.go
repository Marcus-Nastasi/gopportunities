package handler

import (
	"net/http"

	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/handler/input"
	"github.com/Marcus-Nastasi/gopportunities/handler/mappers"
	"github.com/Marcus-Nastasi/gopportunities/handler/output"
	"github.com/Marcus-Nastasi/gopportunities/usecases"
	"github.com/gin-gonic/gin"
)

var logger *config.Logger = config.NewLogger("Handler")

type Handler struct {
	getOpportunities *usecases.GetOpportunitiesUsecase
	getOpportunitie *usecases.GetOpportunitieUsecase
	createOpportunitie *usecases.CreateOpportunitieUsecase
	updateOpportunitie *usecases.UpdateOpportunitieUsecase
	deleteOpportunitie *usecases.DeleteOpportunitieUsecase
}

func NewHandler(
	getOpportunities *usecases.GetOpportunitiesUsecase,
	getOpportunitie *usecases.GetOpportunitieUsecase,
	createOpportunitie *usecases.CreateOpportunitieUsecase,
	updateOpportunitie *usecases.UpdateOpportunitieUsecase,
	deleteOpportunitie *usecases.DeleteOpportunitieUsecase,
) *Handler {
	return &Handler{
		getOpportunities: getOpportunities,
		getOpportunitie: getOpportunitie,
		createOpportunitie: createOpportunitie,
		updateOpportunitie: updateOpportunitie,
		deleteOpportunitie: deleteOpportunitie,
	}
}

func (h *Handler) GetHandler(ctx *gin.Context) {
	o, err := h.getOpportunities.GetOpportunities()
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if len(o) == 0 {
		ctx.JSON(http.StatusNotFound, map[string]string{})
	}
	res := make([]output.OpportunitieResponse, 0, len(o))
	for i := range o {
		res = append(res, mappers.MapFromDomain(&o[i]))
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetSingleHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	o, err := h.getOpportunitie.GetOpportunitie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if o.ID == 0 {
		ctx.JSON(http.StatusNotFound, map[string]string{})
	}
	ctx.JSON(http.StatusOK, mappers.MapFromDomain(&o))
}

func (h *Handler) CreateHandler(ctx *gin.Context) {
	var opportunitie input.OpportunitieRequest
	err := ctx.BindJSON(&opportunitie)
	if err != nil {
		logger.Errorf("Error binding json: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	o, err := h.createOpportunitie.CreateOpportunitie(mappers.MapToDomain(&opportunitie))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, mappers.MapFromDomain(&o))
}

func (h *Handler) UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var err error
	var o input.OpportunitieRequest
	err = ctx.BindJSON(&o)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	updated, err := h.updateOpportunitie.UpdateOpportunitie(id, mappers.MapToDomain(&o))
	ctx.JSON(http.StatusOK, mappers.MapFromDomain(&updated))
}

func (h *Handler) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	o, err := h.deleteOpportunitie.DeleteOpportunitie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, mappers.MapFromDomain(&o))
}
