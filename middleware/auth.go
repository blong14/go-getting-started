package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRequired handles route authorization
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.GetStringMap("context")

		_, ok := ctx["user"]
		if ok {
			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
