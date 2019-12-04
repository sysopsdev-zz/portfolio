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
	c.tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (c Controller) About(w http.ResponseWriter, req *http.Request) {
	c.tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
