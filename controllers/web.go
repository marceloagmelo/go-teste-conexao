package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/marceloagmelo/go-teste-conexao/lib"
	"github.com/marceloagmelo/go-teste-conexao/models"
	"github.com/marceloagmelo/go-teste-conexao/variaveis"
)

//Home é a página inicial da aplicação
func Home(c echo.Context) error {
	formDados := models.FormDados{}
	formDados.VHost = "/"

	data := map[string]interface{}{
		"titulo":    "Teste de conexão",
		"formDados": formDados,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

//Conectar os dados no banco de dados
func Conectar(c echo.Context) error {
	tipo := c.FormValue("tipo")
	host := c.FormValue("host")
	database := c.FormValue("database")
	port := c.FormValue("port")
	user := c.FormValue("user")
	password := c.FormValue("password")
	vHost := c.FormValue("vhost")

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

		return c.Render(http.StatusOK, "index.html", data)
	}
	mensagem = "Campos obrigatórios!"

	formDados.Password = ""

	data := map[string]interface{}{
		"titulo":    "Teste de conexão",
		"formDados": formDados,
		"mensagem":  mensagem,
	}

	return c.Render(http.StatusOK, "index.html", data)
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
