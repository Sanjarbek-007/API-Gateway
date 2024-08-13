package handler

import (
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddMedicalReport godoc
// @Security ApiKeyAuth
// @Summary Add medical report
// @Description Adds a medical report for a user
// @Tags MedicalReport
// @Param id path string true "User ID"
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
	id := ctx.Param("id")

	_, err := h.Mecdical.AddMedicalReport(ctx, &health.AddMedicalReportReq{UserId: id, RecordType: record.RecordType, RecordDate: record.RecordDate, Description: record.Description, DoctorId: record.DoctorId, Attachments: record.Attachments})
	if err != nil {
		h.Logger.Error("Error Adding user medical record: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{Message: "Medical report added successfully"})
}


// GetMedicalReport godoc
// @Security ApiKeyAuth
// @Summary Get medical reports
// @Description Retrieves all medical reports for a user
// @Tags MedicalReport
// @Param user_id path string true "User ID"
// @Success 200 {object} models.MedicalReport "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/get/{user_id} [get]
func (h *Handler) GetMedicalReport(ctx *gin.Context) {
	id := ctx.Param("user_id")

	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err!= nil {
        h.Logger.Error("Error getting user profile: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	resp, err := h.Mecdical.GetMedicalReport(ctx, &health.GetMedicalReportReq{
		UserId: id,
		FirstName: user.FirstName,
	    LastName: user.LastName,
    })
	if err != nil {
		h.Logger.Error("Error Get Medical record Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
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
		ctx.JSON(500, models.Error{Message: "Internal server error"})
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
// @Param id path string true "Report ID"
// @Param body body health.UpdateMedicalReportReq true "Request body for updating a medical report"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/medicalReport/update [put]
func (h *Handler) UpdateMedicalReport(ctx *gin.Context) {
	var record health.UpdateMedicalReportReq

    if err := ctx.ShouldBindJSON(&record); err!= nil {
        h.Logger.Error("Error binding JSON: ", "error", err)
        ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
        return
    }
    id := ctx.Param("id")

    _, err := h.Mecdical.UpdateMedicalReport(ctx, &health.UpdateMedicalReportReq{Id: id, RecordType: record.RecordType, Description: record.Description, DoctorId: record.DoctorId, Attachments: record.Attachments})
    if err!= nil {
        h.Logger.Error("Error Updating user medical record: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
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
    if err!= nil {
        h.Logger.Error("Error deleting user medical record: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, models.Success{Message: "Medical report deleted successfully"})
}