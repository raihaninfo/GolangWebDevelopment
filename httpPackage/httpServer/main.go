package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// route
	http.HandleFunc("/", home)
	// http server
	http.ListenAndServe(":8080", nil)
}

// route function
func home(w http.ResponseWriter, r *http.Request) {
	//
	ptmp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)

}
