package db

import (
	"context"
	"time"

	"github.com/Kungfucoding23/Golang-Final-Proyect/models"
	"go.mongodb.org/mongo-driver/bson"
)

//LeoTweetsSeguidos lee los tweets de mis seguidores
func LeoTweetsSeguidos(ID string, page int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("relacion")

	skip := (page - 1) * 20

	//condiciones es un slice donde guardo el match que tiene que hacer el ID del usuario que solicita la info con el usuarioid de mi BD, contiene el lugar donde buscara los tweets y que hagan match con el usuario que los busca, es decir que si id este contenido tambien en la estructura del tweet ya que de esta forma sabemos que es un tweet de uno de sus seguidores. Despues los ordenamos por fecha en orden descendente y pasamos la paginación
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	/*
		Ahora con el comando $lookup voy a unir la tabla relacion con la de tweets,
		tiene 4 parámetros necesarios para unir tablas de MongoDB:
		"from": con la tabla queremos unir la tabla relacion,
		"localField": por qué campo local queremos unirlas,
		"foreignField": con que campo externo quiero unirla,
		"as": un alias de como queremos llamar esa tabla, la llamamos igual
	*/

	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	/*
		El comando $unwind hace que no me venga todo como maestro-detalle,
		con ese comando todos los documentos vienen iguales y es más fácil procesar
		la información, aunque venga info repetida
	*/

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	/*
		Ahora creamos el cursor con la función Aggregate del framework nuevo,
		que se ejecuta en la bd MongoDB y me trae el cursor.
		Ese cursor directamente ya tiene todo lo que necesitamos para procesar.
	*/

	cursor, err := col.Aggregate(ctx, condiciones)

	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
