package main

import (
	"log"
	"net/http"
	"os"

	"github.com/blong14/goping-web/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Serve static assets
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// app routes
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
