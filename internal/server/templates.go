package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates *template.Template

func loadTemplates() {
	templates = template.Must(
		template.ParseGlob(filepath.Join("templates", "*.tmpl.html")))
}

func renderTemplate(w http.ResponseWriter, name string, data interface {}) {
	if err := templates.ExecuteTemplate(w, name+".tmpl.html", data); err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
