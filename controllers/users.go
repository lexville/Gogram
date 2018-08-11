package controllers

import (
	"Gogram/views"
	"net/http"
)

// Users contains NewView which is of type *views.View
type Users struct {
	NewView *views.View
}

// NewUsers is used to create a new users controller
// This function will panic if the templates aren't
// parsed correctly and so it should only be used during
// setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView(
			"base",
			"views/users/signupForm.gohtml",
			"views/users/new.gohtml",
		),
	}
}

// New Renders the signup view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}
