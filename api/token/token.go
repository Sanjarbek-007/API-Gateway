package token

import (
	"api-gateway/config"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func ExtractClaims(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Load().ACCESS_TOKEN), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func TokenValid(tokenString string) bool {
	_, err := ExtractClaims(tokenString)
	return err == nil
}

func GetUserId(tokenStr string) string {
	claims, err := ExtractClaims(tokenStr)
	if err == nil {
		return claims.ID
	}
	return ""
}

func GetRole(tokenStr string) string {
	claims, err := ExtractClaims(tokenStr)
	if err == nil {
		return claims.Role
	}
	return ""
}

