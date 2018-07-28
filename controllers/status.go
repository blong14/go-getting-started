package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stat "github.com/semihalev/gin-stats"
)

// Stats data
func Stats(c *gin.Context) {
	c.JSON(http.StatusOK, stat.Report())
}
