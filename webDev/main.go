package main

import (
	"GolangWebDevelopment/webDev/stract"
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

// type sub struct {
// 	Username string
// 	Userdata string
// }

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

func getform(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/getform.gohtml", "templates/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func thanks(w http.ResponseWriter, r *http.Request) {
	var s stract.Sub
	s.Username = r.FormValue("username")
	s.Userdata = r.FormValue("data")

	tem, err := template.ParseFiles("templates/thanks.gohtml", "templates/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, s)
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
	http.HandleFunc("/getform", getform)
	http.HandleFunc("/thanks", thanks)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
