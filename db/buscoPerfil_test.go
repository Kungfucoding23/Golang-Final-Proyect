package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscoPerfilRegistroEncontrado(t *testing.T) {
	//Le paso un ID correcto de la base de datos
	perfil, err := BuscoPerfil("600de14b13cff4873ceed9a8")
	//Se espera que el perfil no venga vacio
	assert.NotNil(t, perfil)
	//Y que no haya error
	assert.Nil(t, err)
	assert.EqualValues(t, "600de14b13cff4873ceed9a8", perfil.ID.Hex())
}
