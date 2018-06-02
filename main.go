package main

import (
	"log"
	"os"

	"github.com/blong14/goping-web/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	store := cookie.NewStore([]byte("secret"))

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.Use(sessions.Sessions("goping", store))

	router.GET("/", controllers.Index)

	router.GET("/login", controllers.GitHubInit)
	// router.POST("/logout", controllers.GithubAuthLogOut)
	router.GET("/account/github/callback", controllers.GitHubCallback)

	router.Run(":" + port)
}
