package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("template/home.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tmp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "GET" {
			tmp, err := template.ParseFiles("template/upload.gohtml")
			if err != nil {
				fmt.Println(err.Error())
			}
			tmp.Execute(w, nil)
		}
		r.ParseMultipartForm(10)
		file, fileHeader, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer file.Close()
		fmt.Println(fileHeader.Filename)
		// fmt.Println(fileHeader.Header)
		fmt.Println(fileHeader.Size)

		contenType := fileHeader.Header["Content-Type"][0]
		fmt.Println(contenType)
		var osFile *os.File
		if contenType == "image/jpeg" {
			osFile, err = ioutil.TempFile("images/JPG", "*.jpg")
		} else if contenType == "application/pdf" {
			osFile, err = ioutil.TempFile("images/PDFs", "*.pdf")
		} else if contenType == "image/png" {
			osFile, err = ioutil.TempFile("images/png", "*.png")
		}
		fmt.Println("err", err)
		defer osFile.Close()

		//func ReadAll

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err.Error())
		}
		osFile.Write(fileBytes)

		tmp, err := template.ParseFiles("template/welcome.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		tmp.Execute(w, nil)

	})

	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}
