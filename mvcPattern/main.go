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
	singupView  *views.View
	loginView   *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	chErr(homeView.Render(w, nil))
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	chErr(contactView.Render(w, nil))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	chErr(aboutView.Render(w, nil))
}

func singup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	chErr(singupView.Render(w, nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	chErr(loginView.Render(w, nil))
}

func main() {
	homeView = views.NewView("header", "views/home.gohtml")
	contactView = views.NewView("header", "views/contact.gohtml")
	aboutView = views.NewView("header", "views/about.gohtml")
	singupView = views.NewView("header", "views/singup.gohtml")
	loginView = views.NewView("header", "views/login.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/about", about)
	r.HandleFunc("/singup", singup)
	r.HandleFunc("/login", login)
	http.ListenAndServe(":8080", r)
}

func chErr(err error) {
	if err != nil {
		panic(err)
	}
}
