package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Layout Dir
var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// NewView Create new view
func NewView(layout string, files ...string) (*View, error) {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return &View{
		Template: t,
		Layout:   layout,
	}, nil
}

// View data
type View struct {
	Template *template.Template
	Layout   string
}

// Render view
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
