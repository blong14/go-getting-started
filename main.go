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
	router.Use(middleware.InitContextData())
	router.Use(stat.RequestStats())

	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/stats", controllers.Stats)
		authorized.GET("/ping", controllers.Ping)
		authorized.POST("/ping", controllers.DoPing)
	}

	router.GET("/", controllers.Index)
	router.GET("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/account/github/callback", controllers.LoginCallback)

	router.Run(":" + port)
}
