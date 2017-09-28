package cache

import (
	"html/template"
	"log"
)

var templates *template.Template

// init all templates
func InitTemplates() {
	templates = template.Must(template.ParseGlob("templates/*"))
}

// render template from cache prior
func RenderTemplate(page string) *template.Template {
	tmpl := templates.Lookup(page)
	if tmpl == nil {
		log.Println("template not in cache:", page)
		tmpl = template.Must(template.ParseFiles("templates/" + page))
	}
	return tmpl
}
