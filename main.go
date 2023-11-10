package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html")) //encapsula os templates e devolve 2 retornos (template e msg de erro)

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{
			Nome:       "Camiseta",
			Descricao:  "Preta",
			Preco:      20,
			Quantidade: 3,
		},
		{
			"Tênis",
			"Azul",
			40,
			9,
		},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
