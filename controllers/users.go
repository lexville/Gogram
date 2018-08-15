package controllers

import (
	"Gogram/models"
	"Gogram/views"
	"fmt"
	"net/http"
)

// Users contains NewView which is of type *views.View
type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
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
		LoginView: views.NewView(
			"base",
			"users/loginForm",
			"users/login",
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
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}

// LoginForm contains an email of type string
// as well as password of type string
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login is used to veryfy the provided email address and passwors
// and login the user if they are indeed the account holders
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrInvalidEmail:
		fmt.Fprintln(w, "invalid email address")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "invalid password provided")
	case nil:
		fmt.Fprintln(w, user)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
