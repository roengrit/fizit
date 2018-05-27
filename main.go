package main

import (
	"fizit/app/user"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/layout.html", "template/home.html"))
		tmpl.Execute(w, "")
	})
	r.HandleFunc("/users", user.List)

	http.ListenAndServe(":8080", r)
	// tmpl := template.Must(template.ParseFiles("template/layout.html", "template/home.html"))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.Execute(w, "")
	// })

	//http.ListenAndServe(":8080", nil)
}
