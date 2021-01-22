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
	router.HandleFunc("/verperfil", middlew.CheckDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.CheckDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	// Getenv recupera el valor de la variable de entorno nombrada por la clave.
	// Si el valor de la variable de entorno de heroku es vacio se le asigna el puerto 8080
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	// AllowAll crea un nuevo controlador de Cors con configuración permisiva que permite todos los orígenes con todos los métodos estándar con cualquier header y credenciales.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
