package lib

import (
	"context"
	"log"

	"github.com/marceloagmelo/go-teste-conexao/utils"
	"github.com/marceloagmelo/go-teste-conexao/variaveis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//AutenticarMongo no mongo
func AutenticarMongo() (mensagem string) {

	c, mensagem := getClient()
	if mensagem == "" {
		err := c.Ping(context.Background(), readpref.Primary())
		mensagem = utils.CheckErr(err)
	}

	if mensagem == "" {
		log.Println("Conexão com mongo OK!")
		mensagem = "Conexão com mongo OK!"
	}

	return mensagem
}

//getClient recuperar o client
func getClient() (client *mongo.Client, mensagem string) {
	stringConexao := "mongodb://" + variaveis.User + ":" + variaveis.Password + "@" + variaveis.Host + ":" + variaveis.Port + "/" + variaveis.Database

	clientOptions := options.Client().ApplyURI(stringConexao)
	client, err := mongo.NewClient(clientOptions)
	mensagem = utils.CheckErr(err)

	if mensagem == "" {
		err = client.Connect(context.Background())
		mensagem = utils.CheckErr(err)
	}

	return client, mensagem
}
