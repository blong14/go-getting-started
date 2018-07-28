package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index main application home page
func Index(c *gin.Context) {
	user, _ := c.Get("user")
	c.HTML(http.StatusOK, "index.gohtml", user)
}
