package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/create", create)
	r.HandleFunc("/delete", delete)
	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}

func create(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	name := "Raihan"
	session.Values["name"] = name
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl, err := template.ParseFiles("view/create.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, nil)
}

func delete(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	tpl, err := template.ParseFiles("view/delete.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tpl.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/home.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Execute(w, nil)
}
