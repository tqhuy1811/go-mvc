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
		homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	}

}

func contact(contactView *views.View) func(w http.
	ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
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
	r := mux.NewRouter()
	r.HandleFunc("/", home(homeView))
	r.HandleFunc("/contact", contact(contactView))
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(custom404Page)
	http.ListenAndServe(":3000", r)
}
