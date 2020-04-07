#!/usr/bin/env bash

source setenv.sh

# Criar rede
echo "Criando a rede teste-conexao-net..."
docker network create teste-conexao-net 

# Mysql
echo "Subindo o mysql..."
docker run -d --name mysqldb --network teste-conexao-net  \
-p 3306:3306 \
-e MYSQL_USER=${MYSQL_USER} \
-e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
-e MYSQL_DATABASE=${MYSQL_DATABASE} \
mysql:5.7

# MongoDB
echo "Subindo o mongodb..."
docker run -d --name mongodb --network teste-conexao-net  \
-p 27017:27017 \
-e MONGO_INITDB_ROOT_USERNAME=${MONGO_INITDB_ROOT_USERNAME} \
-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} \
-e MONGO_INITDB_DATABASE=${MONGO_INITDB_DATABASE} \
mongo:4.2

# RabbitMQ
echo "Subindo o rabbitmq..."
docker run -d --name rabbitmq --network teste-conexao-net  \
-p 5672:5672 -p 15672:15672 \
-e RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER} \
-e RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS} \
-e RABBITMQ_ERLANG_COOKIE=${RABBITMQ_ERLANG_COOKIE} \
-e RABBITMQ_DEFAULT_VHOST=${RABBITMQ_DEFAULT_VHOST} \
rabbitmq:3.6.16-management

# Teste de conexao
echo "Subindo o go-teste-conexao..."
docker run -d --name go-teste-conexao --network teste-conexao-net  \
-p 8080:8080 \
masrceloagmelo/go-teste-conexao

# Listando os containers
docker ps