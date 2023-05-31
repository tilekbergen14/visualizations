package functions

import (
	"groupie-tracker/models"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data models.Data) {
	t, err := template.ParseFiles("views/" + tmpl + ".html")
	if err != nil {
		return
	}
	t.Execute(w, data)
}
