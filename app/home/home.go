package home

import (
	"html/template"
	"net/http"
)

//Home home page
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/layout.html", "template/home.html"))
	tmpl.Execute(w, "")
}
