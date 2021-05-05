package views

import (
	"fmt"
	"html/template"
)

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/header.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/footer.gohtml",
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
