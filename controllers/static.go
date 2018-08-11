package controllers

import "Gogram/views"

// NewStatic returns a pointer to the static struct.
// It contains the templates needed by the static pages
func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("base", "static/home"),
		Contact:  views.NewView("base", "static/contact"),
		NotFound: views.NewView("base", "static/notfound"),
		FAQ:      views.NewView("base", "static/faq"),
	}
}

// Static contains the Home,Contact View which are of type
// *views.View
type Static struct {
	Home     *views.View
	Contact  *views.View
	NotFound *views.View
	FAQ      *views.View
}
