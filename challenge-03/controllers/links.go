package controllers

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func LinksIndex(w http.ResponseWriter, r *http.Request) {

}
