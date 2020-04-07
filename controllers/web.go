package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/marceloagmelo/go-teste-conexao/lib"
	"github.com/marceloagmelo/go-teste-conexao/models"
	"github.com/marceloagmelo/go-teste-conexao/variaveis"
)

var view = template.Must(template.ParseGlob("views/*.html"))

//Home é a página inicial da aplicação
func Home(w http.ResponseWriter, r *http.Request) {
	formDados := models.FormDados{}
	formDados.VHost = "/"

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
		tipo := r.FormValue("tipo")
		host := r.FormValue("host")
		database := r.FormValue("database")
		port := r.FormValue("port")
		user := r.FormValue("user")
		password := r.FormValue("password")
		vHost := r.FormValue("vhost")

		var formDados models.FormDados
		formDados.Tipo = tipo
		formDados.Host = host
		formDados.Database = database
		formDados.Port, _ = strconv.Atoi(port)
		formDados.User = user
		formDados.Password = password
		formDados.VHost = vHost

		mensagem := ""

		if tipo != "" {
			switch tipo {
			case "mysql":
				mensagem = conectarMysql(formDados)
			case "mongo":
				mensagem = conectarMongo(formDados)
			case "rabbitmq":
				mensagem = conectarRabbitMQ(formDados)
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

func conectarMysql(formDados models.FormDados) (mensagem string) {
	variaveis.Host = formDados.Host
	variaveis.User = formDados.User
	variaveis.Password = formDados.Password
	variaveis.Database = formDados.Database
	variaveis.Port = strconv.Itoa(formDados.Port)

	mensagem = lib.AutenticarMysql()

	return mensagem
}

func conectarMongo(formDados models.FormDados) (mensagem string) {
	variaveis.Host = formDados.Host
	variaveis.User = formDados.User
	variaveis.Password = formDados.Password
	variaveis.Database = formDados.Database
	variaveis.Port = strconv.Itoa(formDados.Port)

	mensagem = lib.AutenticarMongo()

	return mensagem
}

func conectarRabbitMQ(formDados models.FormDados) (mensagem string) {
	variaveis.Host = formDados.Host
	variaveis.User = formDados.User
	variaveis.Password = formDados.Password
	variaveis.VHost = formDados.VHost
	variaveis.Port = strconv.Itoa(formDados.Port)

	mensagem = lib.AutenticarRabbitMQ()

	return mensagem
}
