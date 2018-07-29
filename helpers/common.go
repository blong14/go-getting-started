package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"

	"github.com/blong14/goping-web/controllers"
	"github.com/blong14/goping-web/middleware"
	"github.com/gin-gonic/gin"
	stat "github.com/semihalev/gin-stats"
)

// ExecuteRequest mocks a request to our mux
func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := GetRouter("../templates/*")
	router.ServeHTTP(rr, req)
	return rr
}

// GetRouter creates a router during testing
func GetRouter(templatePath string) *gin.Engine {
	router := gin.Default()

	router.SetFuncMap(funcMap())
	router.LoadHTMLGlob(templatePath)
	router.Static("/static", "static")

	router.Use(middleware.Sessions())
	router.Use(middleware.Csrf())
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

func funcMap() (fm template.FuncMap) {
	fm = template.FuncMap{
		"csrfField": csrf,
	}
	return
}

func csrf(token string) (tmpl template.HTML) {
	tmpl = template.HTML(fmt.Sprintf("<input type='hidden' name='_csrf' value='%s'>", token))
	return
}
