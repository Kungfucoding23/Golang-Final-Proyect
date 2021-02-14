package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntentoLoginValido(t *testing.T) {
	user, finded := IntentoLogin("ale3@registro.com", "12345seis")
	assert.NotNil(t, user)
	assert.EqualValues(t, true, finded)

	assert.EqualValues(t, "Juan", user.Nombre)
}
