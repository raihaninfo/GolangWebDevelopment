package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/home.gohtml", "view/header.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmp, err := template.ParseFiles("view/404.gohtml", "view/header.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/about.gohtml", "view/header.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/login.gohtml", "view/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tmp.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/login", login)
	r.NotFoundHandler = http.HandlerFunc(notFount)
	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}
