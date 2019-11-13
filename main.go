package main

import (
	"net/http"

	"github.com/vinicius/ViajaNet/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
