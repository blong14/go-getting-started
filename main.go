package main

import (
	"io"
	"log"
	"os"

	"github.com/blong14/goping-web/helpers"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func init() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	withTemplates := true

	router := helpers.GetRouter(withTemplates)

	router.Run(":" + port)
}
