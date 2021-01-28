package routers

import (
	"net/http"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
)

// BorrarTweet borra un tweet por su ID
func BorrarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err := db.BorroTweetDB(ID, IDUsuario)

	// if !tweetBorrado {
	// 	// fmt.Fprintf(w, "El tweet que desea borrar no existe")
	// 	//la deje comentada xq me daba una advertencia:
	// 	//http: superfluous response.WriteHeader call from github.com/Kungfucoding23/Golang-Final-Proyect/routers.BorrarTweet (borrarTweet.go:31)
	// }

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	//Si llego hasta aca, ya se borro el tweet
	w.WriteHeader(http.StatusOK)

	// json.NewEncoder(w).Encode(deleteResult)
}

//antes de separar el borrarTweet en dos archivos no me daba la advertencia, tengo que revisar porque
