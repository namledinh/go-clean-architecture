package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"http_api/common"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternalServer(err))
				}
			}
		}()
		c.Next()
	}
}