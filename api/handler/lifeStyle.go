package handler

import (
	tokenn "api-gateway/api/token"
	"api-gateway/genproto/health"
	"api-gateway/genproto/user"
	"api-gateway/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Add lifestyle data
// @Description Adds lifestyle data for a user
// @Tags Lifestyle
// @Accept       json
// @Produce      json
// @Param body body health.AddLifeStyleDataReq true "Request body for adding lifestyle data"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/lifestyle/addLifestyleData [post]
func (h *Handler) AddLifeStyleData(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	id, _, err := tokenn.GetUserInfoFromAccessToken(accessToken)
	if err != nil {
		fmt.Println("ok")
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, models.Error{Message: "Invalid access token"})
	}

	var life health.AddLifeStyleDataReq

	if err := ctx.ShouldBindJSON(&life); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: err.Error()})
		return
	}

	resp, err := h.Lifestyle.AddLifeStyleData(ctx, &health.AddLifeStyleDataReq{UserId: id, DataType: life.DataType, DataValue: life.DataValue})
	if err != nil {
		h.Logger.Error("Error Adding user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp.Id)
}


// GetLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Get lifestyle data
// @Description Retrieves lifestyle data for a user
// @Tags Lifestyle
// @Accept       json
// @Produce      json
// @Success 200 {object} models.GetLifeStyle "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/lifestyle/getAllLifestyleData [get]
func (h *Handler) GetLifeStyleData(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
    if !exists {
		ctx.JSON(http.StatusUnauthorized, models.Error{Message: "User ID not found in token"})
        return
    }
	
	fmt.Println("Ok", userID)
    id := userID.(string)
	user, err := h.User.GetUserById(ctx, &user.UserId{UserId: id})
	if err!= nil {
        h.Logger.Error("Error getting user profile: ", "error", err)
        ctx.JSON(500, models.Error{Message: err.Error()})
        return
    }
	fmt.Println(user)

	resp, err := h.Lifestyle.GetLifeStyleData(ctx, &health.GetLifeStyleDataReq{
		UserId: id,
		FirstName: user.FirstName,
	    LastName: user.LastName,
    })

	fmt.Println(resp)

	if err != nil {
		h.Logger.Error("Error Get user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetLifeStyleDataById godoc
// @Security ApiKeyAuth
// @Summary Get lifestyle data by ID
// @Description Retrieves lifestyle data for a specific entry by its ID
// @Tags Lifestyle
// @Accept       json
// @Produce      json
// @Param id path string true "Data ID"
// @Success 200 {object} health.GetLifeStyleDataByIdRes "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/lifestyle/getLifestyleById/{id} [get]
func (h *Handler) GetLifeStyleDataById(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("sadhfuyieudsc", id)

	resp, err := h.Lifestyle.GetLifeStyleDataById(ctx, &health.GetLifeStyleDataByIdReq{Id: id})
	if err != nil {
		h.Logger.Error("Error Get user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}
	fmt.Println(resp)

	ctx.JSON(http.StatusOK, resp)
}


// UpdateLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Update lifestyle data
// @Description Updates a specific lifestyle data entry
// @Tags Lifestyle
// @Accept       json
// @Produce      json
// @Param body body health.UpdateLifeStyleDataReq true "Request body for updating lifestyle data"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 400 {object} models.Error "Invalid request parameters"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/lifestyle/updateLifestyleData [put]
func (h *Handler) UpdateLifeStyleData(ctx *gin.Context) {
	var update health.UpdateLifeStyleDataReq

	if err := ctx.ShouldBindJSON(&update); err != nil {
		h.Logger.Error("Error binding JSON: ", "error", err)
		ctx.JSON(400, models.Error{Message: err.Error()})
		return
	}

	_, err := h.Lifestyle.UpdateLifeStyleData(ctx, &health.UpdateLifeStyleDataReq{Id: update.Id, DataType: update.DataType, DataValue: update.DataValue})
	if err != nil {
		h.Logger.Error("Error Updating user life Style: ", "error", err)
		ctx.JSON(500, models.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Success{Message: "Lifestyle data updated successfully"})
}


// DeleteLifeStyleData godoc
// @Security ApiKeyAuth
// @Summary Delete lifestyle data
// @Description Deletes a specific lifestyle data entry by its ID
// @Tags Lifestyle
// @Accept       json
// @Produce      json
// @Param id path string true "Data ID"
// @Success 200 {object} models.Success "Successful operation"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /api/lifestyle/deleteLifestyleData/{id} [delete]
func (h *Handler) DeleteLifeStyleData(ctx *gin.Context) {
	id := ctx.Param("id")

    _, err := h.Lifestyle.DeleteLifeStyleData(ctx, &health.DeleteLifeStyleDataReq{Id: id})
    if err!= nil {
        h.Logger.Error("Error deleting user life Style: ", "error", err)
        ctx.JSON(500, models.Error{Message: err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, models.Success{Message: "Lifestyle data deleted successfully"})
}