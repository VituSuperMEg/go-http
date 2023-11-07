package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":2000", nil))
}
