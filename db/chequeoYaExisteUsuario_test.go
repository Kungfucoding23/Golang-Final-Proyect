package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChequeoUsuarioExiste(t *testing.T) {
	user, exist, id := ChequeoYaExisteUsuario("ale3@registro.com")

	assert.NotNil(t, user)

	assert.EqualValues(t, true, exist)

	assert.EqualValues(t, "600de14b13cff4873ceed9a8", id)
}

func TestChequeoUsuarioNoExiste(t *testing.T) {
	user, exist, id := ChequeoYaExisteUsuario("ale365@registro.com")

	assert.NotNil(t, user)

	assert.EqualValues(t, false, exist)

	assert.EqualValues(t, "000000000000000000000000", id)
}
