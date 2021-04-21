package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h2>This is home page</h2>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h2>This is About page</h2>")
	} else if r.URL.Path == "/con" {
		fmt.Fprint(w, "<h2>Contact page</h2>")
	} else {
		fmt.Fprint(w, "<h2>This is 404 page</h2>")
	}

}
