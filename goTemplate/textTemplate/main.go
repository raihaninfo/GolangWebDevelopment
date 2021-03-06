package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	emailtext := `hello {{ .FirstName}} {{ .LastName}}
	Thank you for registering at myapp. 
	your login url {{.Url}}
	Thank you so much`

	data := struct {
		FirstName string
		LastName  string
		Url       string
	}{
		"Md Abu",
		"Raihan",
		"facebook.com/raihan.mahmudi.50",
	}
	tpl := template.New("emltmp")
	eml, err := tpl.Parse(emailtext)
	if err != nil {
		fmt.Println(err)
	}
	eml.Execute(os.Stdout, data)
}
