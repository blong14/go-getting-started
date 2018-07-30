package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GithubAuth config
var GithubAuth = &oauth2.Config{
	ClientID:     os.Getenv("GH_CLIENT_ID"),
	ClientSecret: os.Getenv("GH_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GH_CLIENT_CALLBACK"),
	Scopes: []string{
		"user:email",
	},
	Endpoint: github.Endpoint,
}

// Login log in with gh
func Login(c *gin.Context) {
	url := GithubAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// LoginCallback success callback
func LoginCallback(c *gin.Context) {
	code := c.Query("code")
	tok, err := GithubAuth.Exchange(oauth2.NoContext, code)

	if err != nil {
		log.Fatal(err)
	}

	session := sessions.Default(c)

	client := GithubAuth.Client(oauth2.NoContext, tok)
	resp, err := client.Get("https://api.github.com/user")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")

	if err != nil {
		log.Println("JSON parse error: ", err)
		return
	}

	session.Set("user", string(prettyJSON.Bytes()))
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

// Logout the user
func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}

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
