package jwt

import (
	"errors"
	"fmt"
	"go-gin/pkg/secret"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Create(data string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Second * 60 * 60).Unix(),
	})
	tokenHash, err := token.SignedString([]byte(secret.APP_SECRET))
	if err != nil {
		return "", err
	}
	return tokenHash, nil
}

func Verify(c *gin.Context, token string) (string, error) {
	getToken, _ := jwt.Parse(token, func(getToken *jwt.Token) (interface{}, error) {
		if _, ok := getToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method: %v", getToken.Header["alg"])
		}
		return []byte(secret.APP_SECRET), nil
	})
	claims, ok := getToken.Claims.(jwt.MapClaims)
	if !ok || !getToken.Valid {
		return "", errors.New("Unauthorized")
	}
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return "", errors.New("Unauthorized")
	}
	return claims["data"].(string), nil
}
