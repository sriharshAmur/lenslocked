package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Home Page!!!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=uts-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>Please contact us on <a href=\"mailto:test@test.com\">test@test.com</a></p>")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func wordsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Type any word after the /words to display that on the page. </h1>")
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	word := chi.URLParam(r, "word")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	response := fmt.Sprintf("<h1>The word you typed is %s</h1>", word)
	fmt.Fprint(w, response)
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
