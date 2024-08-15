package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	kafka "api-gateway/kafka/producer"
	"api-gateway/models"
	"encoding/json"
	"fmt"
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	writerKafka, err := kafka.NewKafkaProducerInit([]string{"kafka:9092"})
	if err != nil {
		h.Logger.Error("Error initializing Kafka producer", "error", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer writerKafka.Close()

	msgBytes, err := json.Marshal(&req)
	if err != nil {
		h.Logger.Error("Error marshaling request to JSON", "error", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = writerKafka.Producermessage("health", msgBytes)
	if err != nil {
		h.Logger.Error("Error producing Kafka message", "error", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	message := fmt.Sprintf("Generated health recommendations for user %s Description : %s, Priority : %d, RecommendationType : %s, ", req.UserId, req.Description, req.Priority, req.RecommendationType)

	fmt.Println(req.UserId)
	resp, err := h.User.CreateNotifications(c, &user.CreateNotificationsReq{UserId: req.UserId, Message: message})
	if err!= nil {
        h.Logger.Error("Error creating notifications: ", "error", err)
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }
	fmt.Println(resp.Id)

	h.Logger.Info("GenerateHealthRecommendations finished successfully")
	c.JSON(http.StatusOK, models.Success{Message: "Health recommendations generated successfully ",})
}

// GetRealtimeHealthMonitoring godoc
// @Security ApiKeyAuth
// @Summary Get real-time health monitoring data
// @Description Retrieves real-time health monitoring data for a user
// @Tags HealthCheck
// @Accept       json
// @Produce      json
// @Success 200 {object} models.GetRealtimeHealthMonitoringRes "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/health/getRealtimeHealthMonitoring/{user_id} [get]
func (h *Handler) GetRealtimeHealthMonitoring(ctx *gin.Context) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("User ID not found in context")
		ctx.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)


	resp, err := h.Health.GetRealtimeHealthMonitoring(ctx, &health.GetRealtimeHealthMonitoringReq{UserId: id})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
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
// @Param date query string true "Date in format YYYY-MM-DD"
// @Success 200 {object} health.GetDailyHealthSummaryRes "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/health/getDailyHealthSummary/{date} [get]
func (h *Handler) GetDailyHealthSummary(ctx *gin.Context) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("User ID not found in context")
		ctx.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	
	// user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	// fmt.Println(id)
	// fmt.Println("Qara")
	// fmt.Println(user)
	// if err != nil {
	// 	h.Logger.Error("Error getting user profile", "error", err)
	// 	ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
	// 	return
	// }

	date := ctx.Query("date")

	resp, err := h.Health.GetDailyHealthSummary(ctx, &health.GetDailyHealthSummaryReq{
		UserId:    id,
		Date:      date,
	})
	if err != nil {
		h.Logger.Error("Error getting daily health summary", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
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
// @Param start_date query string true "Date in format YYYY-MM-DD"
// @Param end_date query string true "Date in format YYYY-MM-DD"
// @Success 200 {object} health.GetWeeklyHealthSummaryRes "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/health/getWeeklyHealthSummary/{start_date}/{end_date} [get]
func (h *Handler) GetWeeklyHealthSummary(ctx *gin.Context) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("User ID not found in context")
		ctx.JSON(http.StatusBadRequest, models.Error{Message: "User ID not found in context"})
		return
	}
	id := userId.(string)
	fmt.Println(id)

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	
	startdate := ctx.Param("start_date")
    enddate := ctx.Param("end_date")

	resp, err := h.Health.GetWeeklyHealthSummary(ctx, &health.GetWeeklyHealthSummaryReq{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserId:    id,
		StartDate: startdate,
		EndDate:   enddate})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
