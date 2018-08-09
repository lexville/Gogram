package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")
// 	} else if r.URL.Path == "/contact" {
// 		fmt.Fprint(w, "To get in touch, pleese send an email to <a href=\"mailto:support@gophergram.com\">support@gophergram.com</a>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h2>Sorry Page Not Found!!! :-(</h2>")
// 	}
// }

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "home!\n")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "contact!\n")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Faq!\n")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Oops!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	router.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", router)
}
