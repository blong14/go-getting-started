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

// SetCsrf writes csrf token in the header
func SetCsrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-CSRF-TOKEN", csrf.GetToken(c))
		c.Next()
	}
}
