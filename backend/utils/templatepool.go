package utils

import (
	"html/template"
	"log"
)

var (
	HtmlTemplates *template.Template
)

func init()  {
	HtmlTemplates = template.New("views")
	if _, err := HtmlTemplates.ParseGlob("views/template/*.html"); err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}
}