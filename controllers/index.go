package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index main applicaton home page
func Index(c *gin.Context) {
	ctx := c.GetStringMap("context")
	c.HTML(http.StatusOK, "index.gohtml", ctx)
}
