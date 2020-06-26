package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoC maneja la conexion creada por la funcion ConnectDB
var MongoC = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://m001-student:m001-mongodb-basics@cluster0-f0gab.gcp.mongodb.net/golmonrea?retryWrites=true&w=majority")

// ConnectDB se encarga de conectar la aplicacion a la base de datos
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

//CheckConnection verifica la conexion a la BD
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
