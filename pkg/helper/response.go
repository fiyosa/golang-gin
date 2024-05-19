package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPaginate struct {
	Page  int
	Limit int
	Total int
}

func SendSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func SendData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": msg,
	})
}

func SendDatas(c *gin.Context, msg string, data interface{}, paginate IPaginate) {
	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"pagination": paginate,
		"message":    msg,
	})
}

func SendError(c *gin.Context, msg string, statusCode ...int) {
	code := http.StatusBadRequest // Default status code
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.JSON(code, gin.H{
		"message": msg,
	})
}

func SendErrors(c *gin.Context, msg string, err interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors":  err,
		"message": msg,
	})
}
