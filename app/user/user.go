package user

import (
	"html/template"
	"net/http"
)

//List list of user
func List(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/layout.html", "template/user/user-list.html"))
	tmpl.Execute(w, "")
}
