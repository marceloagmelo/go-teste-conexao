package lib

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/marceloagmelo/go-teste-conexao/utils"
	"github.com/marceloagmelo/go-teste-conexao/variaveis"

	_ "github.com/go-sql-driver/mysql"
)

//AutenticarMysql no mysql
func AutenticarMysql() (mensagem string) {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", variaveis.User, variaveis.Password, variaveis.Host, variaveis.Port, variaveis.Database)
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	mensagem = utils.CheckErr(err)

	if mensagem == "" {
		err = db.Ping()
		mensagem = utils.CheckErr(err)
	}

	if mensagem == "" {
		log.Println("Conexão com mysql OK!")
		mensagem = "Conexão com mysql OK!"
	}

	return mensagem
}
