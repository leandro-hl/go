package utils

import (
	"html/template"
	"regexp"
)

type pp func() *regexp.Regexp
type tt func() *template.Template

func GetValidPath() pp {
	validPath := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

	return func() *regexp.Regexp {
		return validPath
	}
}

func GetTemplates() tt {
	templates := template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

	return func() *template.Template {
		return templates
	}
}
