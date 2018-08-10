package views

import "html/template"

// NewView takes in the files needed for the view
// and parses the files. It then returns View
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// View contains a Template which is of type
// *template.Template
type View struct {
	Template *template.Template
}
