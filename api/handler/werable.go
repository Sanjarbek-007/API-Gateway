package handler

import (
	"api-gateway/genproto/health"
	kafka "api-gateway/kafka/producer"
	"api-gateway/models"
	"encoding/json"
	"fmt"
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
	userId, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("User ID not found in context")
		ctx.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in token"})
		return
	}
	id := userId.(string)

	if err := ctx.ShouldBindJSON(&warable); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: err.Error()})
		return
	}

	warable.UserId = id

	writerKafka, err := kafka.NewKafkaProducerInit([]string{"kafka:9092"})
	if err != nil {
		h.Logger.Error("Error initializing Kafka producer", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}
	defer writerKafka.Close()

	msgBytes, err := json.Marshal(&warable)
	if err != nil {
		h.Logger.Error("Error marshaling request to JSON", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	err = writerKafka.Producermessage("werable", msgBytes)
	if err != nil {
		h.Logger.Error("Error producing Kafka message", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	h.Logger.Info("GenerateHealthRecommendations finished successfully")
	ctx.JSON(http.StatusOK, models.Success{Message: "Recommendations generated successfully"})
}

// GetWearableData godoc
// @Security ApiKeyAuth
// @Summary Get wearable data
// @Description Retrieves all wearable data for a user
// @Tags WearableData
// @Accept       json
// @Produce      json
// @Success 200 {object} models.Warable "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/get [get]
func (h *Handler) GetWearableData(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Error{Message: "User ID not found in token"})
		return
	}

	id := userID.(string)
	fmt.Println(id)

	// user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	// if err != nil {
	// 	fmt.Println(err, "------------------------")
	// 	h.Logger.Error("Error getting user profile: ", "error", err)
	// 	ctx.JSON(500, models.Error{Message: err.Error()})
	// 	return
	// }
	// fmt.Println(user.FirstName, user.LastName)
	resp, err := h.Wearable.GetWearableData(ctx, &health.GetWearableDataReq{UserId: id})
	if err != nil {
		fmt.Println(err, "+++++++++++++++++++++++++++++++++++")
		h.Logger.Error("Error Get Medical record Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
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
	fmt.Println(id)

	resp, err := h.Wearable.GetWearableDataById(ctx, &health.GetWearableDataByIdReq{Id: id})
	if err != nil {
		h.Logger.Error("Error Get user Wearable data: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
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
// @Param body body health.UpdateWearableDataReq true "Request body for updating wearable data"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/wearable/update/ [put]
func (h *Handler) UpdateWearableData(ctx *gin.Context) {
	var warable health.UpdateWearableDataReq

	if err := ctx.ShouldBindJSON(&warable); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	_, err := h.Wearable.UpdateWearableData(ctx, &health.UpdateWearableDataReq{Id: warable.Id, DeviceType: warable.DeviceType, DataType: warable.DataType, DataValue: warable.DataValue})
	if err != nil {
		h.Logger.Error("Error Updating user Wearable data: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
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
	if err != nil {
		h.Logger.Error("Error deleting user Wearable data: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{Message: "Wearable data deleted successfully"})
}
