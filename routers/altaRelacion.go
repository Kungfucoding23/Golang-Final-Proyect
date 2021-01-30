package routers

import (
	"net/http"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
)

//AltaRelacion hace la insercion del registro de la relacion entre usuarios
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es necesario", http.StatusBadRequest)
		return
	}

	//asigno los valores a la relacion
	var rel models.Relacion
	rel.UsuarioID = IDUsuario
	rel.UsuarioRelacionID = ID

	status, err := db.InsertarRelacion(rel)
	if err != nil {
		http.Error(w, "Hubo un error al insertar la relación: "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se logró insertar la relación: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
