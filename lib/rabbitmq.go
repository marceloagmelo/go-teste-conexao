package lib

import (
	"fmt"
	"log"

	"github.com/marceloagmelo/go-teste-conexao/utils"
	"github.com/marceloagmelo/go-teste-conexao/variaveis"
	"github.com/streadway/amqp"
)

//AutenticarRabbitMQ no rabbitmq
func AutenticarRabbitMQ() (mensagem string) {
	var connectionString = fmt.Sprintf("amqp://%s:%s@%s:%s%s", variaveis.User, variaveis.Password, variaveis.Host, variaveis.Port, variaveis.VHost)
	conn, err := amqp.Dial(connectionString)
	defer conn.Close()
	mensagem = utils.CheckErr(err)

	if mensagem == "" {
		log.Println("Conexão com rabbitmq OK!")
		mensagem = "Conexão com rabbitmq OK!"
	}

	return mensagem
}
