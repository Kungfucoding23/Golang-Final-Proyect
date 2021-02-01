package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
)

//ConsultaRelacion verifica si existe la relación
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var rel models.Relacion

	rel.UsuarioID = IDUsuario
	rel.UsuarioRelacionID = ID

	//Variable resp donde voy a guardar la respuesta si hay relacion o no
	var resp models.RespuestaConsultaRelacion

	status, err := db.ConsultoRelacion(rel)

	//En lugar de mostrar mensajes de error solo muestra true o false
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	//set heeader
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//Encode para pasar a json el modelo
	json.NewEncoder(w).Encode(resp)
}
