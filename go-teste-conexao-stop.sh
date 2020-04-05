#!/usr/bin/env bash

# Mysql
echo "Finalizando o mysql..."
docker rm -f mysqldb

# MongoDB
echo "Finalizando o mongodb..."
docker rm -f mongodb

# RabbitMQ
echo "Finalizando o rabbitmq..."
docker rm -f rabbitmq

# Teste de conexao
echo "Finalizando o go-teste-conexao..."
docker rm -f go-teste-conexao

# Remover rede
echo "Removendo a rede teste-conexao-net..."
docker network rm teste-conexao-net
