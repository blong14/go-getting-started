package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	stat "github.com/semihalev/gin-stats"
)

func init() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func main() {
	router := gin.Default()

	router.Use(stat.RequestStats())

	router.LoadHTMLGlob("tpls/*")

	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/stats", stats)

	router.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
}

func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about.gohtml", nil)
}

func stats(c *gin.Context) {
	c.HTML(http.StatusOK, "stats.gohtml", stat.Report())
}
