package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, templatePath string, info interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, info)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	user := struct {
		Name string
	}{
		Name: "Sriharsh",
	}
	executeTemplate(w, tplPath, user)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath, nil)
}

func wordsHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "words.gohtml")
	executeTemplate(w, tplPath, nil)
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	word := chi.URLParam(r, "word")
	tplPath := filepath.Join("templates", "word.gohtml")
	executeTemplate(w, tplPath, struct {
		Word string
	}{
		Word: word,
	})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/words", wordsHandler)
	r.Get("/words/{word}", wordHandler)
	r.NotFound(errorHandler)
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
