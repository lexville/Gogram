package views

import (
	"net/http"
	"html/template"
	"path/filepath"
)

var (
	layoutDir = "views/layouts/"
	templateExtension = ".gohtml"
)

// View contains a Template which is of type
// *template.Template and the layout whisch is of
// type string
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// NewView takes in the files needed for the view
// and parses the files. It then returns View
func NewView(layout string, files ...string) *View {
	files = append(
		files,
		layoutFiles()...,
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// layoutFiles returns a slice of strings
// representing the layout files that will
// be used in the application
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExtension)
	if err != nil {
		panic(err)
	}
	return files
}
