package handler

import (
	"api-gateway/api/token"
	"api-gateway/genproto/user"
	"api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// GetUserProfile godoc
// @Security ApiKeyAuth
// @Summary Get user profile
// @Description Retrieves the profile information of a user by their ID
// @Tags User
// @Param id path string true "User ID"
// @Success 200 {object} user.GetProfileResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /user/profile/{id} [get]
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


// UpdateUser godoc
// @Security ApiKeyAuth
// @Summary Update user profile
// @Description Updates the profile information of a user
// @Tags User
// @Param id path string true "User ID"
// @Param body body models.UpdateProfileReq true "Request body for updating user profile"
// @Success 200 {object} user.UpdateProfileResp "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /user/profile/{id} [put]
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


// GetUserByEmail godoc
// @Security ApiKeyAuth
// @Summary Get user by email
// @Description Retrieves a user’s information by their email address
// @Tags User
// @Param email path string true "User Email"
// @Success 200 {object} user.GetUSerByEmailResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /user/email/{email} [get]
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


// DeleteUser godoc
// @Security ApiKeyAuth
// @Summary Delete user
// @Description Deletes the authenticated user’s account
// @Tags User
// @Success 200 {object} user.DeleteUserResp "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /user/delete [delete]
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
