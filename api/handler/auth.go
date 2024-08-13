package handler

import (
	"api-gateway/genproto/user"
	"api-gateway/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserProfile godoc
// @Security ApiKeyAuth
// @Summary Get user profile
// @Description Retrieves the profile information of a user by their ID
// @Tags User
// @Param id path string true "User ID"
// @Success 200 {object} models.GetProfileRes "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/user/profile/{id} [get]
func (h *Handler) GetUserProfile(ctx *gin.Context) {
    userID, exists := ctx.Get("user_id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, models.Error{Message: "User ID not found in token"})
        return
    }

    // Convert to string
    id := userID.(string)
    resp, err := h.User.GetUserProfile(ctx, &user.GetProfileReq{UserId: id})
    if err != nil {
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
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/user/updateUser/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userUpdate models.UpdateProfileReq

	if err := ctx.ShouldBindJSON(&userUpdate); err!= nil {
        h.Logger.Error("Error binding JSON: ", "error", err)
        ctx.JSON(400, models.Error{Message: "Invalid request parameters"})
        return
    }

	_, err := h.User.UpdateProfile(ctx, &user.UpdateProfileReq{
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

	ctx.JSON(http.StatusOK, models.Success{Message: "Profile updated successfully"})
}


// GetUserByEmail godoc
// @Security ApiKeyAuth
// @Summary Get user by email
// @Description Retrieves a user’s information by their email address
// @Tags User
// @Param email path string true "User Email"
// @Success 200 {object} user.FilterUsers "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/user/email/{email} [get]
func (h *Handler) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	fmt.Println(email)
    resp, err := h.User.GetUSerByEmail(ctx, &user.GetUSerByEmailReq{
        Email: email,
    })
	fmt.Println(err)
	
    if err!= nil {
        h.Logger.Error("Error getting user by email: ", "error", err)
        ctx.JSON(500, models.Error{Message: "Internal server error"})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}
