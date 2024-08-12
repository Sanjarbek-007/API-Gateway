package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Generate Health Recommendations
// @Description Generate Health Recommendations for doctor
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} models.GetProfileRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users/getUserProfile/ [get]
func (h *Handler) AddLifeStyleData(ctx *gin.Context) {
	var life health.AddLifeStyleDataReq

	if err := ctx.ShouldBindJSON(&life); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}
	id := ctx.Param("id")

	resp, err := h.Lifestyle.AddLifeStyleData(ctx, &health.AddLifeStyleDataReq{UserId: id, DataType: life.DataType, DataValue: life.DataValue})
	if err != nil {
		h.Logger.Error("Error Adding user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


// @Summary Generate Health Recommendations
// @Description Generate Health Recommendations for doctor
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} models.GetProfileRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users/getUserProfile/{:user_id} [get]
func (h *Handler) GetLifeStyleData(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.Lifestyle.GetLifeStyleData(ctx, &health.GetLifeStyleDataReq{UserId: id})
	if err != nil {
		h.Logger.Error("Error Get user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetLifeStyleDataById(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.Lifestyle.GetLifeStyleDataById(ctx, &health.GetLifeStyleDataByIdReq{Id: id})
	if err != nil {
		h.Logger.Error("Error Get user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateLifeStyleData(ctx *gin.Context) {
	var update health.UpdateLifeStyleDataReq

	if err := ctx.ShouldBindJSON(&update); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Lifestyle.UpdateLifeStyleData(ctx, &health.UpdateLifeStyleDataReq{Id: update.Id, DataType: update.DataType, DataValue: update.DataValue})
	if err != nil {
		h.Logger.Error("Error Updating user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteLifeStyleData(ctx *gin.Context) {
	id := ctx.Param("id")

    resp, err := h.Lifestyle.DeleteLifeStyleData(ctx, &health.DeleteLifeStyleDataReq{Id: id})
    if err!= nil {
        h.Logger.Error("Error deleting user life Style: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}