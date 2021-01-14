package db

import "golang.org/x/crypto/bcrypt"

// EncriptarPassword encripta la pass recibida
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //El algoritmo de encriptacion hara (2 elevado al costo) pasados por el texto
	// GenerateFromPassword returns the bcrypt hash of the password at the given cost
	// recibe pass convirtiendolo a byte y el costo
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	// devuelve bytes convitiendo el resultado a string
	return string(bytes), err
}
