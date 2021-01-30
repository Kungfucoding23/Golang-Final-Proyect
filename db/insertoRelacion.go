package db

import (
	"context"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
)

//InsertarRelacion guarda la relación
func InsertarRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, rel)
	if err != nil {
		return false, err
	}

	return true, nil
}
