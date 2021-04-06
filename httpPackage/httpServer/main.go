package main

import "net/http"

func main() {
	// route
	http.HandleFunc("/", home)
	// http server
	http.ListenAndServe(":8080", nil)
}

// route function
func home(w http.ResponseWriter, r *http.Request) {
	//
}
