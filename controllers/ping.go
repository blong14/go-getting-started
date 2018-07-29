package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping shows ping page
func Ping(c *gin.Context) {
	ctx := c.GetStringMap("context")
	c.HTML(http.StatusOK, "ping.gohtml", ctx)
}

// DoPing pings the url
func DoPing(c *gin.Context) {
	url := c.PostForm("url")

	fmt.Println(url)

	c.Redirect(http.StatusFound, "/ping")
}
