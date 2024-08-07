package db

import (
	"context"
	"testing"
	"time"

	"github.com/haneyeric/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	userCreate := createRandomUser(t)
	userGet, err := testQueries.GetUser(context.Background(), userCreate.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userGet)

	require.Equal(t, userCreate.Username, userGet.Username)
	require.Equal(t, userCreate.HashedPassword, userGet.HashedPassword)
	require.Equal(t, userCreate.FullName, userGet.FullName)
	require.Equal(t, userCreate.Email, userGet.Email)
	require.WithinDuration(t, userCreate.PasswordChangeAt, userGet.PasswordChangeAt, time.Second)
	require.WithinDuration(t, userCreate.CreatedAt, userGet.CreatedAt, time.Second)
}
