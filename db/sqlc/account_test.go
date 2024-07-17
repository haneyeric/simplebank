package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/haneyeric/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accCreate := createRandomAccount(t)
	accGet, err := testQueries.GetAccount(context.Background(), accCreate.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accGet)

	require.Equal(t, accCreate.ID, accGet.ID)
	require.Equal(t, accCreate.Owner, accGet.Owner)
	require.Equal(t, accCreate.Balance, accGet.Balance)
	require.Equal(t, accCreate.Currency, accGet.Currency)
	require.WithinDuration(t, accCreate.CreatedAt, accGet.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accCreate := createRandomAccount(t)

	update := UpdateAccountParams{
		ID:      accCreate.ID,
		Balance: util.RandomBalance(),
	}

	accUpdate, err := testQueries.UpdateAccount(context.Background(), update)
	require.NoError(t, err)
	require.NotEmpty(t, accUpdate)

	require.Equal(t, accCreate.ID, accUpdate.ID)
	require.Equal(t, accCreate.Owner, accUpdate.Owner)
	require.Equal(t, update.Balance, accUpdate.Balance)
	require.Equal(t, accCreate.Currency, accUpdate.Currency)
	require.WithinDuration(t, accCreate.CreatedAt, accUpdate.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	accCreate := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), accCreate.ID)
	require.NoError(t, err)

	accDel, err := testQueries.GetAccount(context.Background(), accCreate.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accDel)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	listParams := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accs, err := testQueries.ListAccounts(context.Background(), listParams)
	require.NoError(t, err)
	require.Len(t, accs, 5)

	for _, acc := range accs {
		require.NotEmpty(t, acc)
	}
}
