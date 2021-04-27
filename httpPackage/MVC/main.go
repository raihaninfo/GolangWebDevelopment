package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/home.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tmp.Execute(w, nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/about.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tmp.Execute(w, nil)
}
