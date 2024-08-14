package middleware

import (
	tokenn "api-gateway/api/token"
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func Check(c *gin.Context) {

	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})
		return
	}

	_, err := tokenn.ValidateAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}
	c.Next()
}

func (casb *casbinPermission) GetRole(c *gin.Context) (string, int) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	}
	userId, role, err := tokenn.GetUserInfoFromAccessToken(token)
	if err != nil {
		return "error while reding role", 500
	}
	fmt.Println(userId, role)
	c.Set("user_id", userId)
	return role, 0
}

func (casb *casbinPermission) CheckPermission(c *gin.Context) (bool, error) {
	act := c.Request.Method
	sub, status := casb.GetRole(c)
	if status != 0 {
		return false, fmt.Errorf("error getting role: %v", status)
	}
	obj := c.FullPath()
	log.Println(sub, obj, act)
	ok, err := casb.enforcer.Enforce(sub, obj, act)
	if err != nil {
		fmt.Println("Error enforcing policy", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": fmt.Sprintf("Internal server error: %v", err),
		})
		c.Abort()
		return false, err
	}
	return ok, nil
}


func CheckPermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(c *gin.Context) {
		result, err := casbHandler.CheckPermission(c)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !result {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
			})
			return
		}

		c.Next()
	}
}
