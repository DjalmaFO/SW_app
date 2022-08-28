package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Retorna a porta do server (configurada no .env)
func GetPort() string {
	return os.Getenv("PORT")
}

// Retorna o nome da tabela (configurado no .env)
func GetTable() string {
	return os.Getenv("TABLE")
}

// Retorna o nome do DataBase (configurado no .env)
func GetDBName() string {
	return os.Getenv("DB_NAME")
}

// Retorna a porta do DB (configurada no .env)
func GetDBPort() string {
	return os.Getenv("DB_PORT")
}

// Retorna o nome do servidor onde o DB está hospedado (configurado no .env)
func GetDBServer() string {
	return os.Getenv("DB_SERVER")
}

// ConnectionDB => Retorna ponteiros de client e context para serem utilizados em transações MongoDB
func ConnectionDB() (*mongo.Client, *context.Context) {
	ctx := context.TODO()
	opts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s", GetDBServer(), GetDBPort()),
	)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err.Error())
	}

	return client, &ctx
}
