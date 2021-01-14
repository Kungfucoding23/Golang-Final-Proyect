package main

import (
	"log"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"github.com/Kungfucoding23/Golang-Final-Proyect/handlers"
)

func main() {
	// Si es 0 significa que no se logro la conexion
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	// Si se conecta, llamo al manejador que es donde seteo el puerto y se escucha el servidor
	handlers.Handlers()

}
