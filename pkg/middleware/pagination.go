package middleware

import (
	"github.com/depender/email-sequence-service/constants"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Pagination(perPageLimit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		perPage, err := strconv.ParseInt(c.Query("per_page"), 10, 0)
		if err != nil || perPage > perPageLimit {
			perPage = 20
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil {
			page = 1
		}

		countRequired, err := strconv.ParseBool(c.Query("count_required"))
		if err != nil {
			countRequired = false
		}

		c.Set(constants.Limit, perPage)
		c.Set(constants.Page, page)
		c.Set(constants.CountRequired, countRequired)

		c.Next()
	}
}
