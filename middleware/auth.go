package middleware

import (
	"net/http"

	"github.com/blong14/goping-web/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRequired handles route authorization
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := controllers.ContextData(c)

		_, ok := ctx["user"]
		if ok {
			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
