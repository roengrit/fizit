package main

import (
	"fizit/app/home"
	"fizit/app/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	r := mux.NewRouter()
	r.HandleFunc("/", home.Home)
	r.HandleFunc("/users", user.List)
	http.ListenAndServe(":8080", r)
}
