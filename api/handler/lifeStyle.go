package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Add lifestyle data
// @Description Adds lifestyle data for a user
// @Tags Lifestyle
// @Param id path string true "User ID"
// @Param body body health.AddLifeStyleDataReq true "Request body for adding lifestyle data"
// @Success 200 {object} health.AddLifeStyleDataResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /lifestyle/data/{id} [post]
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


// GetLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Get lifestyle data
// @Description Retrieves lifestyle data for a user
// @Tags Lifestyle
// @Param user_id path string true "User ID"
// @Success 200 {object} health.GetLifeStyleDataResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /lifestyle/data/{user_id} [get]
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

// GetLifeStyleDataById godoc
// @Security ApiKeyAuth
// @Summary Get lifestyle data by ID
// @Description Retrieves lifestyle data for a specific entry by its ID
// @Tags Lifestyle
// @Param id path string true "Data ID"
// @Success 200 {object} health.GetLifeStyleDataByIdResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /lifestyle/data/id/{id} [get]
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


// UpdateLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Update lifestyle data
// @Description Updates a specific lifestyle data entry
// @Tags Lifestyle
// @Param body body health.UpdateLifeStyleDataReq true "Request body for updating lifestyle data"
// @Success 200 {object} health.UpdateLifeStyleDataResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /lifestyle/data/update [put]
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


// DeleteLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Delete lifestyle data
// @Description Deletes a specific lifestyle data entry by its ID
// @Tags Lifestyle
// @Param id path string true "Data ID"
// @Success 200 {object} health.DeleteLifeStyleDataResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /lifestyle/data/{id} [delete]
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