package main

import (
	"loja/routes"
	"net/http"
)

//go mod init
//go mod tidy
//go build

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
