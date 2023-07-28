package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}
func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	trn := createRandomTransfer(t, account1, account2)
	transfer, err := testQueries.GetTransfer(context.Background(), trn.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, trn.ID, transfer.ID)
	require.Equal(t, trn.Amount, transfer.Amount)
}
func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	args := ListTransfersParams{
		Limit:         5,
		Offset:        5,
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
	}
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}
	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)
	}
}

// helper funcs
func createRandomTransfer(t *testing.T, accountFrom, accountTo Account) Transfer {
	args := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        gofakeit.Int64(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, accountFrom.ID)
	require.Equal(t, transfer.ToAccountID, accountTo.ID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}
