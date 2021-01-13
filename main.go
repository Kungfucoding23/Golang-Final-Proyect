package main

import (
	"log"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"github.com/Kungfucoding23/Golang-Final-Proyect/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()

}
