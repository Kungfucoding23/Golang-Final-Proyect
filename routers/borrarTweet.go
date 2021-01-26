package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var connection = db.MongoCN

// BorrarTweet borra un tweet por su ID
func BorrarTweet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")

	db := connection.Database("microblogging")
	col := db.Collection("tweet")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := col.DeleteOne(ctx, filter)

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(deleteResult)
}
