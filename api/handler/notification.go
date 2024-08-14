package handler

import (
	tokenn "api-gateway/api/token"
	"api-gateway/genproto/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllNotifications godoc
// @Security ApiKeyAuth
// @Summary GetAllNotifications
// @Description it will GetAllNotifications
// @Tags Notifications
// @Success 200 {object} user.GetNotificationsResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/notifications [get]
func (h *Handler) GetAllNotifications(c *gin.Context) {
	h.Logger.Info("GetAllNotifications called")
	req := user.GetNotificationsReq{}
	accessToken := c.GetHeader("Authorization")
	id, _, err := tokenn.GetUserInfoFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.Notification.GetAllNotifications(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Logger.Info("GetAllNotifications finished successfully")
	c.JSON(http.StatusOK, res)
}

// GetAndMarkNotificationAsRead godoc
// @Security ApiKeyAuth
// @Summary GetAndMarkNotificationAsRead
// @Description it will GetAndMarkNotificationAsRead
// @Tags Notifications
// @Success 200 {object} user.GetAndMarkNotificationAsReadRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/notifications/new [get]
func (h *Handler) GetAndMarkNotificationAsRead(c *gin.Context) {
	h.Logger.Info("GetAndMarkNotificationAsRead called")

	req := user.GetAndMarkNotificationAsReadReq{}

	accessToken := c.GetHeader("Authorization")
	
	id, _, err := tokenn.GetUserInfoFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = id
	res, err := h.Notification.GetAndMarkNotificationAsRead(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	h.Logger.Info("GetAndMarkNotificationAsRead finished successfully")
	c.JSON(200, res)
}
