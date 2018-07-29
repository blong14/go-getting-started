package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// Csrf initializes csrf middleware
func Csrf() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.Abort()
		},
	})
}
