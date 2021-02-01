package models

//RespuestaConsultaRelacion tiene el valor booleano que se obtiene de consultar la relación
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"` // Como sólo es para devolver en el http no necesito que sea bson, sino solo json.

}
