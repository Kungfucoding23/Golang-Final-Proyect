package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var connection = db.MongoCN

// BorrarTweet borra un tweet por su ID
func BorrarTweet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	db := connection.Database("microblogging")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	deleteResult, err := col.DeleteOne(ctx, condicion)

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(deleteResult)
}
