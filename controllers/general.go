package controllers

import (
	"html/template"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		return
	}
	c.tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (c Controller) Podcast(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/podcast" {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		return
	}
	c.tpl.ExecuteTemplate(w, "podcast.gohtml", nil)
}

func (c Controller) NotFound(w http.ResponseWriter, req *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		c.tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
		return
	}
}
