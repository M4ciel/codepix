package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewAccount(testing *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Caio"
	account, err := NewAccount(bank, accountNumber, ownerName)

	require.Nil(testing, err)
	require.NotEmpty(testing, uuid.FromStringOrNil(account.ID))
	require.Equal(testing, account.Number, accountNumber)
	require.Equal(testing, account.Bank.ID, bank.ID)

	_, err = NewAccount(bank, "", ownerName)
	require.NotNil(testing, err)
}
