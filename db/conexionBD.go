package db

import (
	"context" //context tiene que ver con lo relacionado con los contextos o entornos de las BD
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://ale:ale1402@golangfinalproyect.cl5zb.mongodb.net/microblogging?retryWrites=true&w=majority")

/*ConectDB es la función que me permite conectar a la base de datos*/
func ConectDB() *mongo.Client {
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
	log.Println("Conexión exitosa con la BD")
	return client
}

/*CheckConnection es el ping a la BD*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}