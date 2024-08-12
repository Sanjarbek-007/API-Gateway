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
// @Summary generate health recommendations
// @Description generates health recommendations for a user
// @Tags HealthCheck
// @Param id path string true "id"
// @Success 200 {object} string "message"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router health/GenerateHealthRecommendation [post]
func (h *Handler) GenerateHealthRecommendations(c *gin.Context) {
	h.Logger.Info("GenerateHealthRecommendations called")
	id := c.Param("id")
	req := &health.GenerateHealthRecommendationsReq{UserId: id}

	writerKafka, err := kafka.NewKafkaProducerInit([]string{"localhost:9092"})
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err.Error())
		return

	}
	defer writerKafka.Close()
	msgBytes, err := json.Marshal(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err.Error())
		return

	}

	err = writerKafka.Producermessage("health", msgBytes)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}

	h.Logger.Info("DeleteLanguage finished successfully")
	c.JSON(200, gin.H{"message": "Recommendations generated successfully"})
}


// GetRealtimeHealthMonitoring godoc
// @Security ApiKeyAuth
// @Summary Get real-time health monitoring data
// @Description Retrieves real-time health monitoring data for a user
// @Tags HealthCheck
// @Param id path string true "User ID"
// @Success 200 {object} health.GetRealtimeHealthMonitoringResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /health/realtime/{id} [get]
func (h *Handler) GetRealtimeHealthMonitoring(ctx *gin.Context) {
	id := ctx.Param("id")

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
// @Param id path string true "User ID"
// @Param body body health.GetDailyHealthSummaryReq true "Request body for daily health summary"
// @Success 200 {object} health.GetDailyHealthSummaryResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /health/daily/{id} [post]
func (h *Handler) GetDailyHealthSummary(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err!= nil {
        h.Logger.Error("Error getting user profile: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	var getDailySummary health.GetDailyHealthSummaryReq

	if err := ctx.ShouldBindJSON(&getDailySummary); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Health.GetDailyHealthSummary(ctx, &health.GetDailyHealthSummaryReq{FirstName: user.FirstName, LastName: user.LastName, UserId: id, Date: getDailySummary.Date})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetWeeklyHealthSummary godoc
// @Security ApiKeyAuth
// @Summary Get weekly health summary
// @Description Retrieves the weekly health summary for a user
// @Tags HealthCheck
// @Param id path string true "User ID"
// @Param body body health.GetWeeklyHealthSummaryReq true "Request body for weekly health summary"
// @Success 200 {object} health.GetWeeklyHealthSummaryResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /health/weekly/{id} [post]
func (h *Handler) GetWeeklyHealthSummary(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err!= nil {
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
		LastName: user.LastName,
		UserId: id, 
		StartDate: getWeeklySummary.StartDate, 
		EndDate: getWeeklySummary.EndDate})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}



	ctx.JSON(http.StatusOK,resp)
}
