package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMode_NewPixKey(testing *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Caio"
	account, err := NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, err := NewPixKey(kind, key, account)

	require.NotEmpty(testing, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(testing, pixKey.Kind, kind)
	require.Equal(testing, pixKey.Key, key)

	kind = "cpf"
	_, err = NewPixKey(kind, key, account)
	require.Nil(testing, err)

	kind = "nome"
	_, err = NewPixKey(kind, key, account)
	require.NotNil(testing, err)
}
