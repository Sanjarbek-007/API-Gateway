package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	kafka "api-gateway/kafka/producer"
	"api-gateway/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateHealthRecommendations godoc
// @Security ApiKeyAuth
// @Summary Generate health recommendations
// @Description Generates health recommendations for a user
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Param body body health.GenerateHealthRecommendationsReq true "Request body for generating health recommendations"
// @Success 200 {object} map[string]string "Successful operation"
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 500 {object} map[string]string "Server error"
// @Router /api/health/generate [post]
func (h *Handler) GenerateHealthRecommendations(c *gin.Context) {
	h.Logger.Info("GenerateHealthRecommendations called")

	var req health.GenerateHealthRecommendationsReq
	if err := c.BindJSON(&req); err != nil {
		h.Logger.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid data"})
		return
	}

	writerKafka, err := kafka.NewKafkaProducerInit([]string{"localhost:9092"})
	if err != nil {
		h.Logger.Error("Error initializing Kafka producer", "error", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Server error"})
		return
	}
	defer writerKafka.Close()

	msgBytes, err := json.Marshal(&req)
	if err != nil {
		h.Logger.Error("Error marshaling request to JSON", "error", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Server error"})
		return
	}

	err = writerKafka.Producermessage("health", msgBytes)
	if err != nil {
		h.Logger.Error("Error producing Kafka message", "error", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"message": "Server error"})
		return
	}

	h.Logger.Info("GenerateHealthRecommendations finished successfully")
	c.JSON(http.StatusOK, map[string]string{"message": "Recommendations generated successfully"})
}

// GetRealtimeHealthMonitoring godoc
// @Security ApiKeyAuth
// @Summary Get real-time health monitoring data
// @Description Retrieves real-time health monitoring data for a user
// @Tags HealthCheck
// @Accept       json
// @Produce      json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.GetRealtimeHealthMonitoringRes "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/health/getRealtimeHealthMonitoring/{user_id} [get]
func (h *Handler) GetRealtimeHealthMonitoring(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.Health.GetRealtimeHealthMonitoring(ctx, &health.GetRealtimeHealthMonitoringReq{UserId: id})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetDailyHealthSummary godoc
// @Security ApiKeyAuth
// @Summary Get daily health summary
// @Description Retrieves the daily health summary for a user
// @Tags HealthCheck
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        body body models.GetDailyHealthSummaryReq true "Request body for getting daily health summary"
// @Success 200 {object} health.GetDailyHealthSummaryRes "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/health/getDailyHealthSummary/{id} [post]
func (h *Handler) GetDailyHealthSummary(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err != nil {
		h.Logger.Error("Error getting user profile", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
		return
	}

	var getDailySummary health.GetDailyHealthSummaryReq

	if err := ctx.ShouldBindJSON(&getDailySummary); err != nil {
		h.Logger.Error("Error binding JSON", "error", err)
		ctx.JSON(http.StatusBadRequest, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Health.GetDailyHealthSummary(ctx, &health.GetDailyHealthSummaryReq{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserId:    id,
		Date:      getDailySummary.Date,
	})
	if err != nil {
		h.Logger.Error("Error getting daily health summary", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}


// GetWeeklyHealthSummary godoc
// @Security ApiKeyAuth
// @Summary Get weekly health summary
// @Description Retrieves the weekly health summary for a user
// @Tags HealthCheck
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Param body body health.GetWeeklyHealthSummaryReq true "Request body for weekly health summary"
// @Success 200 {object} health.GetWeeklyHealthSummaryRes "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/health/weekly/{id} [post]
func (h *Handler) GetWeeklyHealthSummary(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	var getWeeklySummary health.GetWeeklyHealthSummaryReq

	if err := ctx.ShouldBindJSON(&getWeeklySummary); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Health.GetWeeklyHealthSummary(ctx, &health.GetWeeklyHealthSummaryReq{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserId:    id,
		StartDate: getWeeklySummary.StartDate,
		EndDate:   getWeeklySummary.EndDate})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
