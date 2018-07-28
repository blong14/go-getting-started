package middleware

import (
	"github.com/blong14/goping-web/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetUser pulls user from session
func SetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		session := sessions.Default(c)

		if v := session.Get("user"); v != nil {
			user, _ = models.ParseUser(v.(string))
			c.Set("user", user)
		}

		c.Next()
	}
}
