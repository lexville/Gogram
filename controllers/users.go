package controllers

import (
	"Gogram/models"
	"Gogram/views"
	"fmt"
	"net/http"
)

// Users contains NewView which is of type *views.View
type Users struct {
	NewView *views.View
	us      *models.UserService
}

// SignupForm contains an email of type string
// as well as password of type string
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// NewUsers is used to create a new users controller
// This function will panic if the templates aren't
// parsed correctly and so it should only be used during
// setup
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView(
			"base",
			"users/signupForm",
			"users/new",
		),
		us: us,
	}
}

// New renders the signup view to create a new
// user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a
// user submits it from the signup page. This is used to
// create a new user account
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, form)
}
