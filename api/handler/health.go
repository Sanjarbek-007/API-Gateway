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
func (h *Handler) GenerateHealthRecommendations(ctx *gin.Context) {
	var generate health.GenerateHealthRecommendationsReq

	if err := ctx.ShouldBindJSON(&generate); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}
	id := ctx.Param("id")

	resp, err := h.Health.GenerateHealthRecommendations(ctx, &health.GenerateHealthRecommendationsReq{UserId: id, RecommendationType: generate.RecommendationType, Description: generate.Description, Priority: generate.Priority})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
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
// @Router /users/getUserProfile/ [get]
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

func (h *Handler) GetDailyHealthSummary(ctx *gin.Context) {
	id := ctx.Param("id")

	var getDailySummary health.GetDailyHealthSummaryReq

	if err := ctx.ShouldBindJSON(&getDailySummary); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Health.GetDailyHealthSummary(ctx, &health.GetDailyHealthSummaryReq{UserId: id, Date: getDailySummary.Date})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetWeeklyHealthSummary(ctx *gin.Context) {
	id := ctx.Param("id")

	var getWeeklySummary health.GetWeeklyHealthSummaryReq

	if err := ctx.ShouldBindJSON(&getWeeklySummary); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Health.GetWeeklyHealthSummary(ctx, &health.GetWeeklyHealthSummaryReq{UserId: id, StartDate: getWeeklySummary.StartDate, EndDate: getWeeklySummary.EndDate})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
