package middlew

import (
	"net/http"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
)

/*CheckDB es el middleware que me permite conocer el estado de la BD */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
