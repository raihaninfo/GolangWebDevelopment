package main

import (
	"GolangWebDevelopment/mvcPattern/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	aboutView   *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := aboutView.Template.ExecuteTemplate(w, aboutView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("header", "views/home.gohtml")
	contactView = views.NewView("header", "views/contact.gohtml")
	aboutView = views.NewView("header", "views/about.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/about", about)
	http.ListenAndServe(":8080", r)
}
