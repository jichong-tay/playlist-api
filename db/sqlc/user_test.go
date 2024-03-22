package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	null "gopkg.in/guregu/null.v4"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomName(),
		Email:        util.RandomName(),
		PasswordHash: hashedPassword,
		Address:      null.NewString(util.RandomName(), true),
		Uuid:         util.RandomString(6),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Address, user.Address)
	require.Equal(t, arg.Uuid, user.Uuid)

	return user
}

func TestCreateUser(t *testing.T) {
	user1 := createRandomUser(t)
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomName(),
		Email:        user1.Email,
		PasswordHash: hashedPassword,
		Address:      null.NewString(util.RandomName(), true),
		Uuid:         util.RandomString(6),
	}

	_, err = testQueries.CreateUser(context.Background(), arg)
	require.Error(t, err)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.Address, user2.Address)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:           user1.ID,
		Username:     util.RandomName(),
		Email:        util.RandomName(),
		PasswordHash: user1.PasswordHash,
		Address:      null.NewString(util.RandomName(), true),
		Uuid:         util.RandomString(6),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.PasswordHash, user2.PasswordHash)
	require.Equal(t, arg.Address, user2.Address)
	require.Equal(t, arg.Uuid, user2.Uuid)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	var lastUser User
	for i := 0; i < 10; i++ {
		lastUser = createRandomUser(t)
	}
	arg := ListUsersParams{
		ID:     lastUser.ID,
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	for _, user := range users {
		require.NotEmpty(t, user)
		require.Equal(t, arg.ID, user.ID)
	}
}
