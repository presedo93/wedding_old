// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateGuest(ctx context.Context, arg CreateGuestParams) (Guest, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	DeleteGuest(ctx context.Context, id int64) error
	DeleteProfile(ctx context.Context, id uuid.UUID) error
	DeleteUserGuest(ctx context.Context, profileID uuid.UUID) error
	GetGuest(ctx context.Context, id int64) (Guest, error)
	GetGuests(ctx context.Context, arg GetGuestsParams) ([]Guest, error)
	GetProfile(ctx context.Context, id uuid.UUID) (Profile, error)
	GetProfiles(ctx context.Context, arg GetProfilesParams) ([]Profile, error)
	GetUserGuests(ctx context.Context, profileID uuid.UUID) ([]Guest, error)
	UpdateGuest(ctx context.Context, arg UpdateGuestParams) (Guest, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
}

var _ Querier = (*Queries)(nil)
