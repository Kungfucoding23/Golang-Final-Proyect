package db

import (
	"context"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya esta en la db
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")
	condicion := bson.M{"email": email}
	var resultado models.Usuario

	// FindeOne busca y devuelve un solo resultado. Decode descompondrá el documento representado por este resultado
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
