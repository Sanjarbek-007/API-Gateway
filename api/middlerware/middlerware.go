package middleware

import (
	tokenn "api-gateway/api/token"
	"fmt"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func Check(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "error": "Authorization header is required",
        })
        return
    }

    parts := strings.Split(authHeader, " ")
	fmt.Println(parts[0])
	fmt.Println(parts[1])
    if len(parts) != 2 || parts[0] != "Bearer" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid authorization header format",
        })
        return
    }

    accessToken := parts[1]
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
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return "unauthorized", http.StatusUnauthorized
    }

    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return "unauthorized", http.StatusUnauthorized
    }

    token := parts[1]
    _, role, err := tokenn.GetUserInfoFromAccessToken(token)
    if err != nil {
        return "error while reading role", 500
    }

    return role, 0
}


func (casb *casbinPermission) CheckPermission(c *gin.Context) (bool, error) {
    act := c.Request.Method
    sub, status := casb.GetRole(c)
    if status != 0 {
        return false, fmt.Errorf("failed to get role: %s", sub)
    }
    obj := c.FullPath()

    ok, err := casb.enforcer.Enforce(sub, obj, act)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "Error": "Internal server error",
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
			c.AbortWithError(500, err)
		}
		if !result {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Forbidden",
			})
		}

		c.Next()
	}
}
