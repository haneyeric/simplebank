package db

import (
	"context"
	"testing"
	"time"

	"github.com/haneyeric/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomBalance(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entCreate := createRandomEntry(t, acc)
	entGet, err := testQueries.GetEntry(context.Background(), entCreate.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entGet)

	require.Equal(t, entCreate.ID, entGet.ID)
	require.Equal(t, entCreate.AccountID, entGet.AccountID)
	require.Equal(t, entCreate.Amount, entGet.Amount)
	require.WithinDuration(t, entCreate.CreatedAt, entGet.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	listParams := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), listParams)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, listParams.AccountID, entry.AccountID)
	}
}
