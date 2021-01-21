package viewmodels

import "html/template"

type EmailMessage struct {
	To      []string
	Header  string
	Title   string
	Message template.HTML
}
