package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index main application home page
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}
