package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type queryResult struct {
	Page     int
	Limit    int
	Keyword  string
	OrderBy  string
	SortedBy string
}

func Offset(page int, limit int) int {
	return (page - 1) * limit
}

func QueryStr(c *gin.Context) queryResult {
	getPage := strings.TrimSpace(c.DefaultQuery("page", "1"))
	getLimit := strings.TrimSpace(c.DefaultQuery("limit", "10"))
	getKeyword := strings.TrimSpace(c.DefaultQuery("keyword", ""))
	getOrderBy := strings.TrimSpace(c.DefaultQuery("orderBy", "id"))
	getSortedBy := strings.TrimSpace(c.DefaultQuery("sortedBy", "asc"))

	newPage := Str2Int(getPage)
	if newPage < 1 {
		newPage = 1
	}

	newLimit := Str2Int(getLimit)
	if newLimit < 1 {
		newLimit = 1
	}
	if newLimit > 100 {
		newLimit = 100
	}

	sortedByToLower := strings.ToLower(getSortedBy)
	if sortedByToLower != "asc" && sortedByToLower != "desc" {
		getSortedBy = "asc"
	}

	return queryResult{
		Page:     newPage,
		Limit:    newLimit,
		Keyword:  getKeyword,
		OrderBy:  getOrderBy,
		SortedBy: getSortedBy,
	}
}
