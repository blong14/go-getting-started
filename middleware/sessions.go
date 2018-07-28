package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Sessions initializes application session
func Sessions() gin.HandlerFunc {
	// @todo: move "secret" to OS.env
	store := cookie.NewStore([]byte("secret"))
	return sessions.Sessions("goping", store)
}
