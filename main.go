package main

import (
	"fmt"
	"lenslock/views"
	"net/http"

	"github.com/gorilla/mux"
)

func home(homeView *views.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		homeView.Render(w, nil)
	}

}

func contact(contactView *views.View) func(w http.
	ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		contactView.Render(w, nil)
	}
}

func login(loginView *views.View) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		loginView.Render(w, nil)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>hello to FAQ page</h2>")
}

func custom404Page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>page not found</h2>")
}

func main() {
	homeView, _ := views.NewView("bootstrap", "views/home.gohtml")
	contactView, _ := views.NewView("bootstrap", "views/contact.gohtml")
	loginView, _ := views.NewView("bootstrap", "views/signup.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home(homeView))
	r.HandleFunc("/contact", contact(contactView))
	r.HandleFunc("/signup", login(loginView))
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(custom404Page)
	http.ListenAndServe(":3000", r)
}
