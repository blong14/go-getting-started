package core

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("tmpl/*.html"))

// RenderTemplate renders an html template file
func RenderTemplate(w http.ResponseWriter, tmpl string, s interface{}) {
	log.Println(templates.DefinedTemplates())
	if err := templates.ExecuteTemplate(w, tmpl+".html", s); err != nil {
		log.Println(err.Error())
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}
