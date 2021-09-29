package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// func homeHandler(w) {

// }

// func actionHandler(w http.ResponseWriter, r *http.Request) {
// 	tmp, err := template.ParseFiles("template/action.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	tmp.Execute(w, nil)
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		tmp, err := template.ParseFiles("template/home.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		tmp.Execute(w, nil)
	})

	// http.HandleFunc("/action", actionHandler)

	http.ListenAndServe(":8080", nil)
}
