package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Kungfucoding23/Golang-Final-Proyect/middlew"
	"github.com/Kungfucoding23/Golang-Final-Proyect/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors" //CORS stands for Cross Origin Resource Sharing
)

// Handlers seteo mi puerto, el handler y pongo a escuchar al servidor
func Handlers() {
	// NewRouter devuelve una nueva instancia de enrutador
	router := mux.NewRouter()
	// por cada EndPoint vamos a tener un renglon que permita manejar la funcion correspondiente
	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	// Getenv recupera el valor de la variable de entorno nombrada por la clave.
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	// AllowAll crea un nuevo controlador de Cors con configuración permisiva que permite todos los orígenes con todos los métodos estándar con cualquier header y credenciales.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
