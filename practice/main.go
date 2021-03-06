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

func reg(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/reg.gohtml", "view/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tmp.Execute(w, nil)
}

func reqauth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	var namelenth bool = false
	if 5 <= len(username) && len(username) <= 50 {
		namelenth = true
	}
	passlenth := false
	if 8 <= len(password) && len(password) <= 40 {
		passlenth = true
	}

	if !namelenth || !passlenth {
		tmp, err := template.ParseFiles("view/reg.gohtml", "view/header.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		tmp.Execute(w, "Give me right user name and password")
		return
	} else if namelenth || passlenth {
		tmp, err := template.ParseFiles("view/regauth.gohtml", "view/header.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		tmp.Execute(w, nil)
		return
	}
	fmt.Println(username, password)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/reg", reg)
	r.HandleFunc("/regauth", reqauth)
	r.NotFoundHandler = http.HandlerFunc(notFount)
	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}
