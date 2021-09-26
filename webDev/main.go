package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// type myinfo struct {
// 	Name       string
// 	Age        int
// 	FatherName string
// 	Pession    string
// 	Mobile     string
// }

// var info myinfo

type list []string

var li list

func homeHandler(w http.ResponseWriter, r *http.Request) {

	// ParseFiles
	tem, err := template.ParseFiles("templates/home.gohtml", "templates/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, li)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/about.gohtml", "templates/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func main() {
	// info = myinfo{
	// 	Name:       "Md Abu Raihan",
	// 	Age:        21,
	// 	FatherName: "Robiul Islam",
	// 	Pession:    "Programming",
	// 	Mobile:     "01853566901",
	// }

	li = list{"Mango", "egg", "Apple", "banana", "pineapple", "water"}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}