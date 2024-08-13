package tokenn

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingkey = "key"
)

func ValidateAccessToken(accessTokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(signingkey), nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}


func ExtractAccessClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
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
        return "", "", errors.New("invalid claims")
    }
    fmt.Println("Claims:", claims)
    userID, userIDOk := claims["user_id"].(string)
    role, roleOk := claims["role"].(string)
	if !userIDOk ||!roleOk {
        return "", "", errors.New("invalid claim types")
    }

    if !userIDOk || !roleOk {
        return "", "", errors.New("invalid claim types")
    }

    return userID, role, nil
}

