package middleware

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Sessions initializes application session
func Sessions() gin.HandlerFunc {
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	return sessions.Sessions("goping", store)
}
