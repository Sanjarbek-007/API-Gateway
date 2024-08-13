package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddWearableData godoc
// @Security ApiKeyAuth
// @Summary Add wearable data
// @Description Adds wearable data for a user
// @Tags WearableData
// @Accept       json
// @Produce      json
// @Param body body health.AddWearableDataReq true "Request body for adding wearable data"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/add [post]
func (h *Handler) AddWearableData(ctx *gin.Context) {
	var warable health.AddWearableDataReq

	if err := ctx.ShouldBindJSON(&warable); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	_, err := h.Wearable.AddWearableData(ctx, &health.AddWearableDataReq{UserId: warable.UserId, DeviceType: warable.DeviceType, DataType: warable.DataType, DataValue: warable.DataValue})
	if err != nil {
		h.Logger.Error("Error Adding user Werable data: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{Message: "Wearable data added successfully"})
}


// GetWearableData godoc
// @Security ApiKeyAuth
// @Summary Get wearable data
// @Description Retrieves all wearable data for a user
// @Tags WearableData
// @Accept       json
// @Produce      json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Warable "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/get/{user_id} [get]
func (h *Handler) GetWearableData(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, models.Error{Message: "User ID not found in token"})
        return
    }

    id := userID.(string)

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err!= nil {
        h.Logger.Error("Error getting user profile: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	resp, err := h.Wearable.GetWearableData(ctx, &health.GetWearableDataReq{UserId: id, FirstName: user.FirstName, LastName: user.LastName})
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
// @Accept       json
// @Produce      json
// @Param id path string true "Wearable Data ID"
// @Success 200 {object} health.GetWearableDataByIdRes "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/getById/{id} [get]
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
// @Accept       json
// @Produce      json
// @Param id path string true "Wearable Data ID"
// @Param body body health.UpdateWearableDataReq true "Request body for updating wearable data"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/update [put]
func (h *Handler) UpdateWearableData(ctx *gin.Context) {
	var warable health.UpdateWearableDataReq

    if err := ctx.ShouldBindJSON(&warable); err!= nil {
        h.Logger.Error("Error binding JSON: ", "error", err)
        ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
        return
    }
    id := ctx.Param("id")

    _, err := h.Wearable.UpdateWearableData(ctx, &health.UpdateWearableDataReq{Id: id, DeviceType: warable.DeviceType, DataType: warable.DataType, DataValue: warable.DataValue})
    if err!= nil {
        h.Logger.Error("Error Updating user Wearable data: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	ctx.JSON(http.StatusOK, models.Success{Message: "Wearable data updated successfully"})
}


// DeleteWearableData godoc
// @Security ApiKeyAuth
// @Summary Delete wearable data
// @Description Deletes a specific wearable data entry by its ID
// @Tags WearableData
// @Accept       json
// @Produce      json
// @Param id path string true "Wearable Data ID"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/delete/{id} [delete]
func (h *Handler) DeleteWearableData(ctx *gin.Context) {
	id := ctx.Param("id")

    _, err := h.Wearable.DeleteWearableData(ctx, &health.DeleteWearableDataReq{Id: id})
    if err!= nil {
        h.Logger.Error("Error deleting user Wearable data: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, models.Success{Message: "Wearable data deleted successfully"})
}