package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// AddWearableData godoc
// @Security ApiKeyAuth
// @Summary Add wearable data
// @Description Adds wearable data for a user
// @Tags WearableData
// @Param id path string true "User ID"
// @Param body body health.AddWearableDataReq true "Request body for adding wearable data"
// @Success 200 {object} health.AddWearableDataResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /wearable/data/{id} [post]
func (h *Handler) AddWearableData(ctx *gin.Context) {
	var warable health.AddWearableDataReq

	if err := ctx.ShouldBindJSON(&warable); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}
	id := ctx.Param("id")

	resp, err := h.Wearable.AddWearableData(ctx, &health.AddWearableDataReq{UserId: id, DeviceType: warable.DeviceType, DataType: warable.DataType, DataValue: warable.DataValue})
	if err != nil {
		h.Logger.Error("Error Adding user Werable data: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


// GetWearableData godoc
// @Security ApiKeyAuth
// @Summary Get wearable data
// @Description Retrieves all wearable data for a user
// @Tags WearableData
// @Param user_id path string true "User ID"
// @Success 200 {object} health.GetWearableDataResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /wearable/data/{user_id} [get]
func (h *Handler) GetWearableData(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.Wearable.GetWearableData(ctx, &health.GetWearableDataReq{UserId: id})
	if err != nil {
		h.Logger.Error("Error Get Medical record Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


// GetWearableDataById godoc
// @Security ApiKeyAuth
// @Summary Get wearable data by ID
// @Description Retrieves a specific wearable data entry by its ID
// @Tags WearableData
// @Param id path string true "Wearable Data ID"
// @Success 200 {object} health.GetWearableDataByIdResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /wearable/data/id/{id} [get]
func (h *Handler) GetWearableDataById(ctx *gin.Context) {
	id := ctx.Param("id")

    resp, err := h.Wearable.GetWearableDataById(ctx, &health.GetWearableDataByIdReq{Id: id})
    if err!= nil {
        h.Logger.Error("Error Get user Wearable data: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}


// UpdateWearableData godoc
// @Security ApiKeyAuth
// @Summary Update wearable data
// @Description Updates a specific wearable data entry
// @Tags WearableData
// @Param id path string true "Wearable Data ID"
// @Param body body health.UpdateWearableDataReq true "Request body for updating wearable data"
// @Success 200 {object} health.UpdateWearableDataResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /wearable/data/{id} [put]
func (h *Handler) UpdateWearableData(ctx *gin.Context) {
	var warable health.UpdateWearableDataReq

    if err := ctx.ShouldBindJSON(&warable); err!= nil {
        h.Logger.Error("Error binding JSON: ", "error", err)
        ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
        return
    }
    id := ctx.Param("id")

    resp, err := h.Wearable.UpdateWearableData(ctx, &health.UpdateWearableDataReq{Id: id, DeviceType: warable.DeviceType, DataType: warable.DataType, DataValue: warable.DataValue})
    if err!= nil {
        h.Logger.Error("Error Updating user Wearable data: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	ctx.JSON(http.StatusOK, resp)
}


// DeleteWearableData godoc
// @Security ApiKeyAuth
// @Summary Delete wearable data
// @Description Deletes a specific wearable data entry by its ID
// @Tags WearableData
// @Param id path string true "Wearable Data ID"
// @Success 200 {object} health.DeleteWearableDataResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /wearable/data/{id} [delete]
func (h *Handler) DeleteWearableData(ctx *gin.Context) {
	id := ctx.Param("id")

    resp, err := h.Wearable.DeleteWearableData(ctx, &health.DeleteWearableDataReq{Id: id})
    if err!= nil {
        h.Logger.Error("Error deleting user Wearable data: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}