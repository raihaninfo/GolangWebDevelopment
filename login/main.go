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
	r.HandleFunc("/loginauth", loginAuth)
	// r.HandleFunc("/logout", logout)
	fmt.Println("Listening port :8080")
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

func login(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("temp/login.gohtml", "temp/header.gohtml")
	if err != nil {
		panic(err)
	}
	tem.Execute(w, nil)
}
func loginAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	var userLenth bool = false
	if 5 <= len(username) && 50 >= len(username) {
		userLenth = true
	}
	var passlenth bool = false
	if 5 <= len(password) && 30 >= len(password) {
		passlenth = true
	}
	if !userLenth || !passlenth {
		temp, err := template.ParseFiles("temp/login.gohtml", "temp/header.gohtml")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, "Please give me right user name & password")

	} else if userLenth || passlenth {
		temp, err := template.ParseFiles("temp/loginauth.gohtml", "temp/header.gohtml")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)

	}

}
