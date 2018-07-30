package middleware

import (
	"github.com/blong14/goping-web/users"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// InitContextData pulls user from session
func InitContextData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user users.User

		ctx := make(map[string]interface{})

		session := sessions.Default(c)

		if v := session.Get("user"); v != nil {
			user, _ = users.ParseUser(v.(string))
			ctx["user"] = user
		}

		ctx["token"] = csrf.GetToken(c)

		c.Set("context", ctx)

		c.Next()
	}
}
