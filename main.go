package main

import (
	"log"

	"github.com/Kungfucoding23/Golang-Final-Proyect/bd"
	"github.com/Kungfucoding23/Golang-Final-Proyect/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
