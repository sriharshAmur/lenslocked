package views

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Must(t Template, e error) Template {
	if e != nil {
		panic(e)
	}
	return t
}

func ParseFS(fs embed.FS, patterns ...string) (Template, error) {
	pageWithLayout := append([]string{"layout-page.gohtml"}, patterns...)
	tpl, err := template.ParseFS(fs, pageWithLayout...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template:  %w", err)
	}
	return Template{
		htmpTpl: tpl,
	}, nil
}

type Template struct {
	htmpTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmpTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
