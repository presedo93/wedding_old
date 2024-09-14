package db

import (
	"context"
	"testing"

	"github.com/presedo93/wedding/back/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:       "Fulan",
		Email:      "fulan@email.com",
		Companions: 2,
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Companions, user2.Companions)
}

func TestUpdateUserName(t *testing.T) {
	oldUser := createRandomUser(t)
	newName := util.RandomString(6)

	arg := UpdateUserNameParams{
		ID:   oldUser.ID,
		Name: newName,
	}

	user2, err := testStore.UpdateUserName(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, oldUser.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.ID)

	require.NoError(t, err)
	user2, err := testStore.GetUser(context.Background(), user1.ID)

	require.Error(t, err)
	require.Empty(t, user2)
}
