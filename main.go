package main

import (
	"curso_golang_alura/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	_ = http.ListenAndServe(":8000", nil)
}
