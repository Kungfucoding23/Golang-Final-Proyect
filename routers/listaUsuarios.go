package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
)

//ListaUsuarios lee la lista de usuarios
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	//capturamos los parametros que vienen de la URL
	//en type indico si quiero los usuarios que yo sigo o los que no sigo
	typeUser := r.URL.Query().Get("type")
	//pagina
	page := r.URL.Query().Get("page")
	//busqueda
	search := r.URL.Query().Get("search")

	//convierto la pagina que viene como string a numero
	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)
	// llamo a la función de bd con el ID del usuario que esta logueado y lo que vino
	result, status := db.LeoTodosUsuarios(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios ", http.StatusBadRequest)
		return
	}

	//si no encontró registros el status va a venir en true y result en vacío
	//viene todo bien

	// establecemos el tipo de Header
	w.Header().Set("Content-type", "application/json")
	// le damos un status created
	w.WriteHeader(http.StatusCreated)
	// le devolvemos la lista de resultados al navegador
	json.NewEncoder(w).Encode(result)
}
