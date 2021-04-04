package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	// template parse
	tmp, err := template.ParseFiles("one.gohtml", "two.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	// template exicute
	tmp.Execute(os.Stdout, 100)
}
