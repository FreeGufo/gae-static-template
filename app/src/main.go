package main

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", notfoundHandler)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	template.Must(template.ParseFiles("templates/notfound.html")).Execute(w, nil)
}
