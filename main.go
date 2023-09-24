package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/sriharshamur/lenslocked/controllers"
	"github.com/sriharshamur/lenslocked/views"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl, func(r *http.Request) interface{} {
		return struct {
			Name string
		}{
			Name: "Sriharsh",
		}
	}))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl, nil))

	tpl, err = views.Parse(filepath.Join("templates", "words.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/words", controllers.StaticHandler(tpl, nil))

	tpl, err = views.Parse(filepath.Join("templates", "word.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/words/{word}", controllers.StaticHandler(tpl, func(r *http.Request) interface{} {
		return struct {
			Word string
		}{
			Word: chi.URLParam(r, "word"),
		}
	}))

	r.NotFound(errorHandler)

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
