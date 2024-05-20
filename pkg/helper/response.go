package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Paginate struct {
	Page  int `json:"page" example:"0"`
	Limit int `json:"limit" example:"0"`
	Total int `json:"total" example:"0"`
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

func SendDatas(c *gin.Context, msg string, data interface{}, paginate Paginate) {
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
	c.AbortWithStatusJSON(code, gin.H{
		"message": msg,
	})
}

func SendErrors(c *gin.Context, msg string, err interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"errors":  err,
		"message": msg,
	})
}
