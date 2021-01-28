package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BorroTweetDB borra un tweet por id
func BorroTweetDB(ID string, UserID string) error {
	// var tweetBorrado bool

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	//El IDUsuario es la variable global, se le guarda el id del usuario en el momento del login, por esa razon no es necesario enviarlo como parametro desde el postman
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)

	//aca evaluo si el deletecount es 0 significa que no borro nada por lo tanto significa que el tweet que se quiere borrar no existe o ya fue borrado
	// if deleteResult.DeletedCount == 0 {
	// 	tweetBorrado = false
	// }

	return err
	// si no hay ningun error retornara nil, habiendo ya borrado el tweet
}
