package main

import (
	"admindev/controllers"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tpl)
	index := http.HandlerFunc(c.Index)

	http.Handle("/", loggingFunc(index))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":5000", nil)
}

func loggingFunc(h http.Handler) http.Handler {
	logFile, err := os.OpenFile("admindev.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}
