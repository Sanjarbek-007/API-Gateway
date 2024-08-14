package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	kafka "api-gateway/kafka/producer"
	"api-gateway/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AddMedicalReport godoc
// @Security ApiKeyAuth
// @Summary Add medical report
// @Description Adds a medical report for a user
// @Tags MedicalReport
// @Param body body health.AddMedicalReportReq true "Request body for adding a medical report"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/add [post]
func (h *Handler) AddMedicalReport(ctx *gin.Context) {
	var record health.AddMedicalReportReq

	if err := ctx.ShouldBindJSON(&record); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}

	resp, err := h.Mecdical.AddMedicalReport(ctx, &health.AddMedicalReportReq{UserId: record.UserId, RecordType: record.RecordType, RecordDate: record.RecordDate, Description: record.Description, DoctorId: record.DoctorId, Attachments: record.Attachments})
	if err != nil {
		h.Logger.Error("Error Adding user medical record: ", "error", err)
		ctx.JSON(500, err.Error())
		return
	}

	kajkareq := user.CreateNotificationsReq{UserId: record.UserId, Message: fmt.Sprintf("You have added a new medical report for %s", time.Now().String())}
	writerKafka, err := kafka.NewKafkaProducerInit([]string{"kafka:9092"})
	if err != nil {
		h.Logger.Error("Error initializing Kafka producer", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}
	defer writerKafka.Close()

	msgBytes, err := json.Marshal(&kajkareq)
	if err != nil {
		h.Logger.Error("Error marshaling request to JSON", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	err = writerKafka.WriteToNotification("notification", msgBytes)
	if err != nil {
		h.Logger.Error("Error producing Kafka message", "error", err)
		ctx.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	h.Logger.Info("Medical report finished successfully")

	ctx.JSON(http.StatusOK, resp.Id)
}

// GetMedicalReport godoc
// @Security ApiKeyAuth
// @Summary Get medical reports
// @Description Retrieves all medical reports for a user
// @Tags MedicalReport
// @Success 200 {object} health.GetMedicalReportRes "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/get [get]
func (h *Handler) GetMedicalReport(ctx *gin.Context) {
	fmt.Println("Helllllllllllllllllllllllllllllllllllllllllllllollllllllllllllllllllllllllllllllllllllll")
	userId, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Error{Message: "User ID not found in token"})
		return
	}
	id := userId.(string)
	fmt.Println("Helllllllllllllllllllllllllllllllllllllllllllllollllllllllllllllllllllllllllllllllllllll")
	fmt.Println(id)

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err != nil {
		h.Logger.Error("Error getting user profile: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}
	fmt.Println(user)

	resp, err := h.Mecdical.GetMedicalReport(ctx, &health.GetMedicalReportReq{
		UserId:    id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		h.Logger.Error("Error Get Medical record Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetMedicalReportById godoc
// @Security ApiKeyAuth
// @Summary Get medical report by ID
// @Description Retrieves a specific medical report by its ID
// @Tags MedicalReport
// @Param id path string true "Report ID"
// @Success 200 {object} health.GetMedicalReportByIdRes "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/getById/{id} [get]
func (h *Handler) GetMedicalReportById(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.Mecdical.GetMedicalReportById(ctx, &health.GetMedicalReportByIdReq{Id: id})
	if err != nil {
		h.Logger.Error("Error Get user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateMedicalReport godoc
// @Security ApiKeyAuth
// @Summary Update medical report
// @Description Updates a specific medical report
// @Tags MedicalReport
// @Accept       json
// @Produce      json
// @Param body body health.UpdateMedicalReportReq true "Request body for updating a medical report"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/update [put]
func (h *Handler) UpdateMedicalReport(ctx *gin.Context) {
	var record health.UpdateMedicalReportReq

	if err := ctx.ShouldBindJSON(&record); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: err.Error()})
		return
	}

	_, err := h.Mecdical.UpdateMedicalReport(ctx, &health.UpdateMedicalReportReq{Id: record.Id, RecordType: record.RecordType, Description: record.Description, DoctorId: record.DoctorId, Attachments: record.Attachments})
	if err != nil {
		h.Logger.Error("Error Updating user medical record: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{Message: "Medical report updated successfully"})
}

// DeleteMedicalReport godoc
// @Security ApiKeyAuth
// @Summary Delete medical report
// @Description Deletes a specific medical report by its ID
// @Tags MedicalReport
// @Accept       json
// @Produce      json
// @Param id path string true "Report ID"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/delete/{id} [delete]
func (h *Handler) DeleteMedicalReport(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := h.Mecdical.DeleteMedicalReport(ctx, &health.DeleteMedicalReportReq{Id: id})
	if err != nil {
		h.Logger.Error("Error deleting user medical record: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{Message: "Medical report deleted successfully"})
}
