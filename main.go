package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sriharshamur/lenslocked/controllers"
	"github.com/sriharshamur/lenslocked/templates"
	"github.com/sriharshamur/lenslocked/views"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl, func(r *http.Request) interface{} {
		return struct {
			Name string
		}{
			Name: "Sriharsh",
		}
	}))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl, func(r *http.Request) interface{} { return nil }))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "words.gohtml"))
	r.Get("/words", controllers.StaticHandler(tpl, func(r *http.Request) interface{} { return nil }))

	tpl = views.Must(views.ParseFS(templates.FS, "word.gohtml"))
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
