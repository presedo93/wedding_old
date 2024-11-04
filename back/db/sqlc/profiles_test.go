package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/presedo93/wedding/back/util"
	"github.com/stretchr/testify/require"
)

func createRandomProfile(t *testing.T) Profile {
	arg := CreateProfileParams{
		ID:               util.RandomID(),
		Name:             util.RandomName(),
		Email:            util.RandomEmail(),
		Phone:            util.RandomPhoneNumber(),
		PictureUrl:       pgtype.Text{String: "", Valid: true},
		CompletedProfile: false,
		AddedGuests:      false,
		AddedSongs:       false,
		AddedPictures:    false,
	}

	profile, err := testStore.CreateProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, profile)

	return profile
}

func TestCreateProfile(t *testing.T) {
	createRandomProfile(t)
}

func TestGetProfile(t *testing.T) {
	profile1 := createRandomProfile(t)
	profile2, err := testStore.GetProfile(context.Background(), profile1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.Name, profile2.Name)
	require.Equal(t, profile1.Phone, profile2.Phone)
}

func TestUpdateProfile(t *testing.T) {
	profile1 := createRandomProfile(t)
	newName := util.RandomName()

	arg := UpdateProfileParams{
		ID:   profile1.ID,
		Name: pgtype.Text{String: newName, Valid: true},
	}

	profile2, err := testStore.UpdateProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, profile2)

	require.Equal(t, profile1.ID, profile2.ID)
	require.Equal(t, arg.Name.String, profile2.Name)
}

func TestDeleteProfile(t *testing.T) {
	profile1 := createRandomProfile(t)
	err := testStore.DeleteProfile(context.Background(), profile1.ID)

	require.NoError(t, err)
	profile2, err := testStore.GetProfile(context.Background(), profile1.ID)

	require.Error(t, err)
	require.Empty(t, profile2)
}
