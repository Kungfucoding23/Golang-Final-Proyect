package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta relacion entre dos usuarios
func ConsultoRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         rel.UsuarioID,
		"usuariorelacionid": rel.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	fmt.Println(resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}

/*
consulta de la relación es un endpoint simple que va a permitir, dándole dos ID, decir si están relacionados. Esta función es de las más importantes de la relación porque va a ser llamada constantemente cuando querramos listar todos los usuarios, indicando el detalle de las relaciones.
*/
