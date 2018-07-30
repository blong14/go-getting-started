package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stat "github.com/semihalev/gin-stats"
)

// Index main applicaton home page
func Index(c *gin.Context) {
	ctx := c.GetStringMap("context")
	c.HTML(http.StatusOK, "index.gohtml", ctx)
}

// Stats data
func Stats(c *gin.Context) {
	c.JSON(http.StatusOK, stat.Report())
}
