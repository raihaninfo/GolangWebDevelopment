package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// func homeHandler(w) {

// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("template/home.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tmp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.Method)
		tmp, err := template.ParseFiles("template/upload.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		tmp.Execute(w, nil)
	})

	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}
