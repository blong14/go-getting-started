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
var GithubAuth = &oauth2.Config{
	ClientID:     "53cf24a436e19359232a",                     // @todo replace with OS.env
	ClientSecret: "c071a283179f7b7de054ef9896573c9595829e67", // @todo replace with OS.env
	RedirectURL:  "http://localhost:3000/account/github/callback",
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
