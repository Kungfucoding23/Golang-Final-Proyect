package routers

import (
	"net/http"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
)

//BajaRelacion borra la relacion
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es necesario", http.StatusBadRequest)
		return
	}

	var rel models.Relacion
	rel.UsuarioID = IDUsuario
	rel.UsuarioRelacionID = ID

	status, err := db.BorroRel(rel)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la relación: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logró borrar la relación: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
