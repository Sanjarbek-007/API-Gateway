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