package controllers

import (
	"GolangWebDevelopment/mvcPattern/views"
	"fmt"
	"net/http"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("header", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// GET / singup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "This is signup page")
}
