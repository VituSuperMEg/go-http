package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h2>Hora certa : %s</h2>", s)
}
func main() {
	http.HandleFunc("/hora", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":2000", nil))
}
