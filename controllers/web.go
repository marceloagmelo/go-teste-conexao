package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/marceloagmelo/go-teste-conexao/models"
)

var view = template.Must(template.ParseGlob("views/*.html"))

//Home é a página inicial da aplicação
func Home(w http.ResponseWriter, r *http.Request) {
	formDados := models.FormDados{}
	formDados.VHost = "/"
	formDados.AutenticarDatabase = "N"

	data := map[string]interface{}{
		"titulo":    "Teste de conexão",
		"formDados": formDados,
		"mensagem":  "",
	}

	view.ExecuteTemplate(w, "Index", data)
}

//Conectar os dados no banco de dados
func Conectar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var interf models.Conectar

		tipo := r.FormValue("tipo")
		host := r.FormValue("host")
		database := r.FormValue("database")
		port := r.FormValue("port")
		user := r.FormValue("user")
		password := r.FormValue("password")
		vHost := r.FormValue("vhost")
		autenticarDatabase := r.FormValue("autenticarDatabase")

		var formDados models.FormDados
		formDados.Tipo = tipo
		formDados.Host = host
		formDados.Database = database
		formDados.Port, _ = strconv.Atoi(port)
		formDados.User = user
		formDados.Password = password
		formDados.VHost = vHost
		formDados.AutenticarDatabase = autenticarDatabase

		interf = formDados

		mensagem := ""

		if tipo != "" {
			switch tipo {
			case "mysql":
				mensagem = interf.Mysql()
			case "mongo":
				mensagem = interf.Mongo()
			case "rabbitmq":
				mensagem = interf.Rabbitmq()
			default:
				log.Println("Tipo não encontrado!")
			}

			formDados.Password = ""

			data := map[string]interface{}{
				"titulo":    "Teste de conexão",
				"formDados": formDados,
				"mensagem":  mensagem,
			}

			view.ExecuteTemplate(w, "Index", data)
		}
		mensagem = "Campos obrigatórios!"

		http.Redirect(w, r, "/", 301)
	}
}
