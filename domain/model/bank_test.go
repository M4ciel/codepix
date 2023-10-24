package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewBank(testing *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := NewBank(code, name)

	require.Nil(testing, err)
	require.NotEmpty(testing, uuid.FromStringOrNil(bank.ID))
	require.Equal(testing, bank.Code, code)
	require.Equal(testing, bank.Name, name)

	_, err = NewBank("", "")
	require.NotNil(testing, err)
}
