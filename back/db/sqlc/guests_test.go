package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/presedo93/wedding/back/util"
	"github.com/stretchr/testify/require"
)

func createRandomGuest(t *testing.T) Guest {
	profile := createRandomProfile(t)

	arg := CreateGuestParams{
		ProfileID:      profile.ID,
		Name:           "Fulan",
		Phone:          "+34 666 666 666",
		IsVegetarian:   false,
		Allergies:      []string{"gluten", "lactose"},
		NeedsTransport: false,
	}

	guest, err := testStore.CreateGuest(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, guest)

	return guest
}

func TestCreateGuest(t *testing.T) {
	createRandomGuest(t)
}

func TestGetGuest(t *testing.T) {
	guest1 := createRandomGuest(t)
	guest2, err := testStore.GetGuest(context.Background(), guest1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, guest2)

	require.Equal(t, guest1.Name, guest2.Name)
	require.Equal(t, guest1.Phone, guest2.Phone)
}

func TestUpdateGuest(t *testing.T) {
	guest1 := createRandomGuest(t)
	newName := util.RandomName()

	arg := UpdateGuestParams{
		ID:   guest1.ID,
		Name: pgtype.Text{String: newName, Valid: true},
	}

	guest2, err := testStore.UpdateGuest(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, guest2)

	require.Equal(t, guest1.ID, guest2.ID)
	require.Equal(t, arg.Name.String, guest2.Name)
}

func TestDeleteGuest(t *testing.T) {
	guest1 := createRandomGuest(t)
	err := testStore.DeleteGuest(context.Background(), guest1.ID)

	require.NoError(t, err)
	guest2, err := testStore.GetGuest(context.Background(), guest1.ID)

	require.Error(t, err)
	require.Empty(t, guest2)
}
