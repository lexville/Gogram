package controllers

import "Gogram/views"

// NewStatic returns a pointer to the static struct.
// It contains the templates needed by the static pages
func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("base", "views/static/home.gohtml"),
		Contact:  views.NewView("base", "views/static/contact.gohtml"),
		NotFound: views.NewView("base", "views/static/notfound.gohtml"),
		FAQ:      views.NewView("base", "views/static/faq.gohtml"),
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
