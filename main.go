package main

import (
	"io"
	"log"
	"os"

	"github.com/blong14/goping-web/controllers"
	"github.com/blong14/goping-web/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	stat "github.com/semihalev/gin-stats"
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

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	router.Use(middleware.Sessions())
	router.Use(middleware.Csrf())
	router.Use(middleware.SetCsrf())
	router.Use(middleware.SetUser())
	router.Use(stat.RequestStats())

	router.GET("/", controllers.Index)
	router.GET("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/account/github/callback", controllers.LoginCallback)
	router.GET("/stats", controllers.Stats)

	router.Run(":" + port)
}
