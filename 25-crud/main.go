package main

import (
	"crud/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", route.CriarUsuario).Methods(http.MethodPost) // os parametros sao passados implicitamente para a funcao Criarusuario
	log.Fatal(http.ListenAndServe(":3333", r))
}
