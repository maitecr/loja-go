package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) //encapsula os templates e devolve 2 retornos (template e msg de erro)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.Execute(w, "index")
}
