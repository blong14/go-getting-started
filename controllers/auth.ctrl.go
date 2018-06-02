package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GithubAuth config
var GithubAuth *oauth2.Config = &oauth2.Config{
	ClientID:     "", // @todo replace with OS.env
	ClientSecret: "", // @todo replace with OS.env
	RedirectURL:  "http://localhost:3000/account/github/callback",
	Scopes: []string{
		"user:email",
	},
	Endpoint: github.Endpoint,
}

// GitHubInit log in with gh
func GitHubInit(c *gin.Context) {
	url := GithubAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(302, url)
}

// GitHubCallback gh success callback
func GitHubCallback(c *gin.Context) {
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
	log.Println("body", string(body))

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
