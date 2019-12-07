package main

import (
	"github.com/gorilla/handlers"
	"html/template"
	"net/http"
	"os"
	"projects/admindev2/controllers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/about", c.About)
	http.HandleFunc("/resume", c.Resume)
	http.HandleFunc("/podcast", c.Podcast)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
