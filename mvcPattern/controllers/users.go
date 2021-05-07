package controllers

import (
	"GolangWebDevelopment/mvcPattern/views"
	"net/http"
)


func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("header", "views/users/new.gohtml")
	}
}

type Users struct {
}

func (u *Users) New(w http.ResponseWriter, r *http.Request){

}