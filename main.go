package main

import (
	"Gogram/controllers"
	"net/http"

	"Gogram/views"

	"github.com/gorilla/mux"
)

var (
	homeView     *views.View
	contactView  *views.View
	faqView      *views.View
	notFoundView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	must(notFoundView.Render(w, nil))
}

func main() {
	homeView = views.NewView(
		"base",
		"views/home.gohtml",
	)
	contactView = views.NewView(
		"base",
		"views/contact.gohtml",
	)
	faqView = views.NewView(
		"base",
		"views/faq.gohtml",
	)
	notFoundView = views.NewView(
		"base",
		"views/notfound.gohtml",
	)
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
