package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	pmt, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pmt.Execute(w, nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contect-Type", "text/html")
	fmt.Fprint(w, "This is about page")
}
