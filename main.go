package main

import (
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	tmpl := template.Must(template.ParseFiles("template/layout.html", "template/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, "")
	})

	http.ListenAndServe(":8080", nil)
}
