package routers

import (
	"net/http"

	"github.com/marceloagmelo/go-teste-conexao/controllers"
)

//CarregaRotas  as rotas
func CarregaRotas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/v1/conectar", controllers.Conectar)
}
