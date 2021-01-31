package db

import (
	"context"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
)

//BorroRel borra una relacion
func BorroRel(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, rel)

	//Si hay un error retorna falso y el error
	if err != nil {
		return false, err
	}

	return true, nil
}
