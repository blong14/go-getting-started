package controllers

import (
	"net/http"

	"github.com/blong14/goping-web/core"
)

// Index main application home page
func Index(w http.ResponseWriter, r *http.Request) {
	core.RenderTemplate(w, "index", "")
}

// Login shows login form
func Login(w http.ResponseWriter, r *http.Request) {
	core.RenderTemplate(w, "login", "")
}
