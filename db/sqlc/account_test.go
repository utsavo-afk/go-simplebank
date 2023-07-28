package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
func TestGetAccount(t *testing.T) {
	acc := createRandomAccount(t)
	args := acc.ID
	account, err := testQueries.GetAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.ID, acc.ID)
	require.Equal(t, account.Balance, acc.Balance)
	require.Equal(t, account.Owner, acc.Owner)
	require.Equal(t, account.Currency, acc.Currency)
	require.Equal(t, account.CreatedAt, acc.CreatedAt)
	require.NotZero(t, account.ID)
}
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)
	args := UpdateAccountParams{
		ID:      acc.ID,
		Balance: gofakeit.Int64(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.ID, args.ID)
	require.Equal(t, account.Balance, args.Balance)
}
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	args := account.ID
	deleteRandomAccount(t, args)
}

// helper func
func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    gofakeit.Name(),
		Balance:  gofakeit.Int64(),
		Currency: gofakeit.CurrencyShort(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
func deleteRandomAccount(t *testing.T, id int64) {
	args := id
	err := testQueries.DeleteAccount(context.Background(), args)
	require.NoError(t, err)
}
