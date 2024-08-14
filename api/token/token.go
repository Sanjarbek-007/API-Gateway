package tokenn

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingkey = "key"
)

func ValidateAccessToken(tokenStr string) (bool, error) {
	if tokenStr == "" {
		return false, nil
	}
	
	_, err := ExtractAccessClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccessClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
	})

	if err != nil {
		fmt.Println("ldshgfiuesrhgfvriser", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		fmt.Println("Pasr", err)
		return nil, err
	}

	return &claims, nil
}

func GetUserInfoFromAccessToken(accessTokenString string) (string, string, error) {
	refreshToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(signingkey), nil })
	if err != nil || !refreshToken.Valid {
		return "", "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", err
	}
	userID := claims["user_id"].(string)
	Role := claims["role"].(string)

	return userID, Role, nil
}
