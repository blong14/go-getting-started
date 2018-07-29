package helpers

import (
	"net/http"
	"net/http/httptest"

	"github.com/blong14/goping-web/controllers"
	"github.com/blong14/goping-web/middleware"
	"github.com/gin-gonic/gin"
	stat "github.com/semihalev/gin-stats"
)

// ExecuteRequest mocks a request to our mux
func ExecuteRequest(req *http.Request, withTemplates bool) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := GetRouter(withTemplates)
	router.ServeHTTP(rr, req)
	return rr
}

// GetRouter creates a router during testing
func GetRouter(withTemplates bool) *gin.Engine {
	router := gin.Default()

	if withTemplates {
		router.LoadHTMLGlob("../templates/*")
		router.Static("/static", "static")
	}

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

	return router
}
