package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)

	// http.HandleFunc("/", handlerFunc)

	http.ListenAndServe(":8080", router)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h2>This is home page</h2>")
// 	} else if r.URL.Path == "/about" {
// 		fmt.Fprint(w, "<h2>This is About page</h2>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h2>This is 404 page</h2>")
// 	}
// }
