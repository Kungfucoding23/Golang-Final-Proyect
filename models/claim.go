package models

/*
En esta estructura procesamos el JWT que viene del Login, lo vamos a desencriptar de alguna manera dentro del modelo que es un JSON. Ese modelo lo llamamos Claim
JWT es un estandar abierto que define un metodo compacto y autocontenido para encapsular y compartir aserciones (claims) sobre una entidad (subject) de manera segura entre distintas partes mediante el uso de objetos JSON.
*/
import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim es la estructura usada para procesar el JWT */
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
