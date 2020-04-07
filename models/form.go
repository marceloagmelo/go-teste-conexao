package models

import (
	"strconv"

	"github.com/marceloagmelo/go-teste-conexao/lib"
	"github.com/marceloagmelo/go-teste-conexao/variaveis"
)

//FormDados dados doformulário
type FormDados struct {
	Tipo               string `json:"tip"`
	Host               string `json:"host"`
	Database           string `json:"database"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	VHost              string `json:"vhost"`
	AutenticarDatabase string `json:"autenticar-database"`
}

// Conectar interface
type Conectar interface {
	Mysql() string
	Mongo() string
	Rabbitmq() string
}

//Mysql conectar com mysql
func (form FormDados) Mysql() (mensagem string) {
	variaveis.Host = form.Host
	variaveis.User = form.User
	variaveis.Password = form.Password
	variaveis.Database = form.Database
	variaveis.Port = strconv.Itoa(form.Port)

	mensagem = lib.AutenticarMysql()

	return mensagem
}

//Mongo conectar com o mongo
func (form FormDados) Mongo() (mensagem string) {
	variaveis.Host = form.Host
	variaveis.User = form.User
	variaveis.Password = form.Password
	if form.AutenticarDatabase == "S" {
		variaveis.Database = form.Database
	}
	variaveis.Port = strconv.Itoa(form.Port)
	variaveis.AutenticarDatabase = form.AutenticarDatabase

	mensagem = lib.AutenticarMongo()

	return mensagem
}

//Rabbitmq conectar com o rabbitmq
func (form FormDados) Rabbitmq() (mensagem string) {
	variaveis.Host = form.Host
	variaveis.User = form.User
	variaveis.Password = form.Password
	variaveis.VHost = form.VHost
	variaveis.Port = strconv.Itoa(form.Port)

	mensagem = lib.AutenticarRabbitMQ()

	return mensagem
}
