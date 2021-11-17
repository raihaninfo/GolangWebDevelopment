package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/", home)
	r.HandleFunc("/login", login)
	// r.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", r)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("temp/404.gohtml", "temp/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("temp/home.gohtml", "temp/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request)  {
	tem, err:= template.ParseFiles("temp/login.gohtml", "temp/header.gohtml")
	if err!=nil{
		panic(err)
	}
	tem.Execute(w, nil)
}
