package db

import (
	"context"
	"testing"
	"time"

	"github.com/haneyeric/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accFrom, accTo Account) Transfer {

	arg := CreateTransfersParams{
		FromAccountID: accFrom.ID,
		ToAccountID:   accTo.ID,
		Amount:        util.RandomBalance(),
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	createRandomTransfer(t, acc1, acc2)
}

func TestGetTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	transferCreate := createRandomTransfer(t, acc1, acc2)
	transferGet, err := testQueries.GetTransfer(context.Background(), transferCreate.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferGet)

	require.Equal(t, transferCreate.ID, transferGet.ID)
	require.Equal(t, transferCreate.FromAccountID, transferGet.FromAccountID)
	require.Equal(t, transferCreate.ToAccountID, transferGet.ToAccountID)
	require.Equal(t, transferCreate.Amount, transferGet.Amount)
	require.WithinDuration(t, transferCreate.CreatedAt, transferGet.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, acc1, acc2)
	}

	listParams := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), listParams)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, listParams.ToAccountID, transfer.ToAccountID)
		require.Equal(t, listParams.FromAccountID, transfer.FromAccountID)
		require.True(t, transfer.FromAccountID == acc1.ID || transfer.ToAccountID == acc1.ID)
	}
}
