package variaveis

import (
	"os"
	"time"
)

//DataFormat formato da data
var DataFormat = "02/01/2006 15:04:05"

//DataFormatArquivo formato da data para arquivos
var DataFormatArquivo = "20060102-150405"

//DataHoraAtual a data e hora tual
var DataHoraAtual = time.Now()

//Host nome da máquina
var Host = os.Getenv("HOSTNAME")

//User nome do usuário do banco de dados
var User = os.Getenv("USER")

//Password senha do usuário do banco de dados
var Password = os.Getenv("PASSWORD")

//Database nome do banco de dados
var Database = os.Getenv("DATABASE")

//Port porta
var Port = os.Getenv("PORT")

//VHost do rabbitmq
var VHost = os.Getenv("RABBITMQ_DEFAULT_VHOST")

// AutenticarDatabase mecanismo de autenticação do mongo
var AutenticarDatabase = "N"
