package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
)

//LeoTweetsRelacion lee los tweets de todos nuestros seguidores
func LeoTweetsRelacion(w http.ResponseWriter, r *http.Request) {
	//valido que se ingrese una pagina
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	//Convierte la pagina a entero y valida que sea mayor a cero
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar una página mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.LeoTweetsSeguidores(IDUsuario, page)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
