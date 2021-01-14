package db

import (
	"context" //context tiene que ver con lo relacionado con los contextos o entornos de las BD
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConectDB()

// Client crea una nueva instancia de ClientOptions.
// ClientOptions contiene las opciones para configurar la instancia de Cliente
// ApplyURI analiza el URI dado y establece las opciones en consecuencia
var clientOptions = options.Client().ApplyURI("mongodb+srv://ale:ale1402@golangfinalproyect.cl5zb.mongodb.net/microblogging?retryWrites=true&w=majority")

/*ConectDB es la función que me permite conectar a la base de datos*/
func ConectDB() *mongo.Client {
	// Connect crea un nuevo cliente y luego lo inicializa
	// TODO devuelve un contexto vacío que no es nulo
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	// Ping envía un comando ping para verificar que el cliente puede conectarse
	// El parámetro rp se utiliza para determinar qué servidor se selecciona para la operación. Si es nil, se utiliza la preferencia del cliente.
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
	// Si se logra la conexión retorna un 1, sino un 0. Esto es evaluado en la funcion main
	if err != nil {
		return 0
	}
	return 1
}

// information about the URI format https://docs.mongodb.com/manual/reference/connection-string/
