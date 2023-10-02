package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sriharshamur/lenslocked/controllers"
	"github.com/sriharshamur/lenslocked/models"
	"github.com/sriharshamur/lenslocked/templates"
	"github.com/sriharshamur/lenslocked/views"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.Home(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	pgConfig := models.DefaultPostgresConfig()
	db, err := models.Open(pgConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}

	userC := controllers.Users{
		UserService: &userService,
	}
	userC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml"))
	userC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml"))
	r.Get("/signup", userC.New)
	r.Get("/signin", userC.SignIn)
	r.Post("/users", userC.Create)

	tpl = views.Must(views.ParseFS(templates.FS, "words.gohtml"))
	r.Get("/words", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "word.gohtml"))
	r.Get("/words/{word}", controllers.Word(tpl))

	r.NotFound(errorHandler)

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
