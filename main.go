package main

import (
	"log"
	"net/http"

	"github.com/marceloagmelo/go-teste-conexao/routers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	routers.CarregaRotas()

	log.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
