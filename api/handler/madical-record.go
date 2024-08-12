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
func (h *Handler) AddMedicalReport(ctx *gin.Context) {
	var record health.AddMedicalReportReq

	if err := ctx.ShouldBindJSON(&record); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
		return
	}
	id := ctx.Param("id")

	resp, err := h.Mecdical.AddMedicalReport(ctx, &health.AddMedicalReportReq{UserId: id, RecordType: record.RecordType, RecordDate: record.RecordDate, Description: record.Description, DoctorId: record.DoctorId, Attachments: record.Attachments})
	if err != nil {
		h.Logger.Error("Error Adding user medical record: ", "error", err)
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
func (h *Handler) GetMedicalReport(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.Mecdical.GetMedicalReport(ctx, &health.GetMedicalReportReq{UserId: id})
	if err != nil {
		h.Logger.Error("Error Get Medical record Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

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

func (h *Handler) UpdateMedicalReport(ctx *gin.Context) {
	var record health.UpdateMedicalReportReq

    if err := ctx.ShouldBindJSON(&record); err!= nil {
        h.Logger.Error("Error binding JSON: ", "error", err)
        ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
        return
    }
    id := ctx.Param("id")

    resp, err := h.Mecdical.UpdateMedicalReport(ctx, &health.UpdateMedicalReportReq{Id: id, RecordType: record.RecordType, Description: record.Description, DoctorId: record.DoctorId, Attachments: record.Attachments})
    if err!= nil {
        h.Logger.Error("Error Updating user medical record: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteMedicalReport(ctx *gin.Context) {
	id := ctx.Param("id")

    resp, err := h.Mecdical.DeleteMedicalReport(ctx, &health.DeleteMedicalReportReq{Id: id})
    if err!= nil {
        h.Logger.Error("Error deleting user medical record: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}