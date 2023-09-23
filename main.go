package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Home Page!!!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=uts-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>Please contact us on <a href=\"mailto:test@test.com\">test@test.com</a></p>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page Not Found", http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "<p>Page Not Found</p>")
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "<p>Page Not Found</p>")
	}
}

func main() {
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/", homeHandler)
	// http.Handle("/", http.HandlerFunc(homeHandler))
	// http.HandleFunc("/contact", contactHandler)
	// http.Handle("/contact", http.HandlerFunc(contactHandler))
	// var router http.HandlerFunc = pathHandler
	// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	var router Router
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", router)
	http.ListenAndServe(":3000", nil)
}
