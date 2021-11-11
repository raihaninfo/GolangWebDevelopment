package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/home.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}
func notFount(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("view/404.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.NotFoundHandler = http.HandlerFunc(notFount)
	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}
