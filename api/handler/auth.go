package handler

import (
	"api-gateway/api/token"
	"api-gateway/genproto/user"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get All Users
// @Description Get all users for admin
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} models.GetProfileRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /users/getUserProfile/{:id} [get]
func (h *Handler) GetUserProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.User.GetUserProfile(ctx, &user.GetProfileReq{UserId: id})
	if err!= nil {
        h.Logger.Error("Error getting user profile: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update User By ID
// @Description Update user by id for admin
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param UpdateUser body models.UpdateProfileReq true "Update user request"
// @Success 200 {object} models.Update
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/users/{:id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userUpdate models.UpdateProfileReq

	if err := ctx.ShouldBindJSON(&userUpdate); err!= nil {
        h.Logger.Error("Error binding JSON: ", "error", err)
        ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
        return
    }

	resp, err := h.User.UpdateProfile(ctx, &user.UpdateProfileReq{
		Id: id,
		Email: userUpdate.Email,
		FirstName: userUpdate.FirstName,
		LastName: userUpdate.LastName,
		DateOfBirth: userUpdate.DateOfBirth,
        Gender: userUpdate.Gender,
	})
	if err!= nil {
        h.Logger.Error("Error updating user: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get User By Email
// @Description Get user by Email for admin
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "user email"
// @Success 200 {object} models.GetProfileRes
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/users/{:email} [get]
func (h *Handler) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
    resp, err := h.User.GetUSerByEmail(ctx, &user.GetUSerByEmailReq{
        Email: email,
    })
	
    if err!= nil {
        h.Logger.Error("Error getting user by email: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete User
// @Description Delete user by token
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 204 {object} models.Success
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /api/users/{:id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	value, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error getting claims",
		})
	}
	claims, ok := value.(*token.Claims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error getting claims",
		})
		return
	}

	resp, err := h.User.DeleteUser(ctx, &user.UserId{Id: claims.ID})
	if err != nil {
		h.Logger.Error("Error deleting user", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error deleting user",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
