package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreatePassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword, err := CreateHashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = ComparePassword(hashedPassword, password)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = ComparePassword(hashedPassword, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := CreateHashedPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword, hashedPassword2)
}
