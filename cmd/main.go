package main

import (
	"github.com/gorilla/handlers"
	"github.com/jbonachera/short"
	"log"
	"net/http"
	"os"
)

func main() {
	site := short.Site{Host: os.Getenv("SITE"), HashMinSize: 5}
	http.HandleFunc("/", site.Redirect)
	http.HandleFunc("/post", site.Post)
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
