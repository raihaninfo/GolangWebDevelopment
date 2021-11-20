package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginAuth)
	r.HandleFunc("/logout", logout)
	fmt.Println("Listening port :8080")
	http.ListenAndServe(":8080", r)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("temp/404.gohtml", "temp/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tem, err := template.ParseFiles("temp/home.gohtml", "temp/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tem, err := template.ParseFiles("temp/about.gohtml", "temp/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	_, ok := session.Values["username"]
	if ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tem, err := template.ParseFiles("temp/login.gohtml", "temp/header.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	tem.Execute(w, nil)

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func loginAuth(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	var userLenth bool = false
	if username == "raihan" {
		userLenth = true
	}
	var passlenth bool = false
	if password == "password" {
		passlenth = true
	}
	if !userLenth || !passlenth {
		temp, err := template.ParseFiles("temp/login.gohtml", "temp/header.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}
		temp.Execute(w, "Please give me right user name & password")

	} else if userLenth || passlenth {
		session, err := store.Get(r, "login-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
		}
		session.Values["username"] = username
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)

	}

}
