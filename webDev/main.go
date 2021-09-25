package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type myinfo struct {
	Name       string
	Age        int
	FatherName string
	Pession    string
	Mobile     string
}

var info myinfo

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, info)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/about.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func main() {
	info = myinfo{
		Name:       "Md Abu Raihan",
		Age:        21,
		FatherName: "Robiul Islam",
		Pession:    "Programming",
		Mobile:     "01853566901",
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
