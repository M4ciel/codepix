package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewTransaction(testing *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Caio"
	account, _ := NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Giovanna"
	accountDestination, err := NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := NewPixKey(kind, key, accountDestination)

	require.NotEqual(testing, account.ID, accountDestination.ID)

	amount := 3.10
	statusTransaction := "pending"
	descriptionTransaction := "My description"
	transaction, err := NewTransaction(account, amount, pixKey, descriptionTransaction)

	require.Nil(testing, err)
	require.NotNil(testing, uuid.FromStringOrNil(transaction.ID))
	require.Equal(testing, transaction.Amount, amount)
	require.Equal(testing, transaction.Status, statusTransaction)
	require.Equal(testing, transaction.Description, descriptionTransaction)
	require.Empty(testing, transaction.CancelDescription)

	pixKeySameAccount, err := NewPixKey(kind, key, account)

	_, err = NewTransaction(account, amount, pixKeySameAccount, descriptionTransaction)
	require.NotNil(testing, err)

	_, err = NewTransaction(account, 0, pixKey, descriptionTransaction)
	require.NotNil(testing, err)
}

func TestModel_ChangeStatusOfATransaction(testing *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Caio"
	account, _ := NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Giovanna"
	accountDestination, _ := NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "j@j.com"
	pixKey, _ := NewPixKey(kind, key, accountDestination)

	amount := 3.10
	descriptionTransaction := "My description"
	transaction, _ := NewTransaction(account, amount, pixKey, descriptionTransaction)

	transaction.Complete()
	require.Equal(testing, transaction.Status, TransactionCompleted)

	cancelDescription := "Error"
	transaction.Cancel(cancelDescription)
	require.Equal(testing, transaction.Status, TransactionError)
	require.Equal(testing, transaction.CancelDescription, cancelDescription)
}
