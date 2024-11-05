package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/presedo93/wedding/back/auth"
	mockdb "github.com/presedo93/wedding/back/db/mock"
	db "github.com/presedo93/wedding/back/db/sqlc"
	"github.com/presedo93/wedding/back/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateProfileAPI(t *testing.T) {
	userID := util.RandomID()
	profile := randomProfile(userID)

	body := createProfileBody{
		ID:            userID,
		Name:          profile.Name,
		Phone:         profile.Phone,
		Email:         profile.Email,
		PictureUrl:    profile.PictureUrl.String,
		AddedGuests:   profile.AddedGuests,
		AddedSongs:    profile.AddedSongs,
		AddedPictures: profile.AddedPictures,
	}

	arg := db.CreateProfileParams{
		ID:            userID,
		Name:          profile.Name,
		Phone:         profile.Phone,
		Email:         profile.Email,
		PictureUrl:    profile.PictureUrl,
		AddedGuests:   profile.AddedGuests,
		AddedSongs:    profile.AddedSongs,
		AddedPictures: profile.AddedPictures,
	}

	testCases := []struct {
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		name          string
		body          createProfileBody
	}{
		{
			name: "OK",
			body: body,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Eq(arg)).Times(1).Return(profile, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProfile(t, recorder.Body, profile)
			},
		},
		{
			name: "InternalError",
			body: body,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Eq(arg)).Times(1).Return(profile, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: createProfileBody{},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			jwks := NewMockJWKS(userID)
			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			server := NewServer(store, jwks)
			recorder := httptest.NewRecorder()

			body, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/api/profiles", bytes.NewReader(body))
			require.NoError(t, err)

			request.Header.Set(auth.Header, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetProfileAPI(t *testing.T) {
	userID := util.RandomID()
	profile := randomProfile(userID)

	testCases := []struct {
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		name          string
		id            string
	}{
		{
			name: "OK",
			id:   profile.ID.String(),
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(profile.ID)).Times(1).Return(profile, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProfile(t, recorder.Body, profile)
			},
		},
		{
			name: "NotFound",
			id:   profile.ID.String(),
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(profile.ID)).Times(1).Return(db.Profile{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			id:   profile.ID.String(),
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(profile.ID)).Times(1).Return(db.Profile{}, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			jwks := NewMockJWKS(userID)
			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			// start the server
			server := NewServer(store, jwks)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/api/profiles", nil)
			require.NoError(t, err)

			request.Header.Set(auth.Header, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateProfileAPI(t *testing.T) {
	new_name := util.RandomName()

	profile := randomProfile(util.RandomID())
	profile.Name = new_name

	body := updateProfileBody{
		Name: new_name,
	}

	arg := db.UpdateProfileParams{
		ID:   profile.ID,
		Name: pgtype.Text{String: new_name, Valid: true},
	}

	testCases := []struct {
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		name          string
		id            string
		body          updateProfileBody
	}{
		{
			name: "OK",
			body: body,
			id:   profile.ID.String(),
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateProfile(gomock.Any(), gomock.Eq(arg)).Times(1).Return(profile, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProfile(t, recorder.Body, profile)
			},
		},
		{
			name: "InternalError",
			body: body,
			id:   profile.ID.String(),
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateProfile(gomock.Any(), gomock.Eq(arg)).Times(1).Return(profile, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: updateProfileBody{Email: "no_email"},
			id:   "",
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			jwks := NewMockJWKS(profile.ID)
			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			server := NewServer(store, jwks)
			recorder := httptest.NewRecorder()

			body, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPut, "/api/profiles", bytes.NewReader(body))
			require.NoError(t, err)

			request.Header.Set(auth.Header, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomProfile(profileID uuid.UUID) db.Profile {
	return db.Profile{
		ID:               profileID,
		Name:             util.RandomName(),
		Phone:            util.RandomPhoneNumber(),
		Email:            util.RandomEmail(),
		PictureUrl:       pgtype.Text{String: util.RandomUrl(), Valid: true},
		CompletedProfile: true,
		AddedGuests:      false,
		AddedSongs:       false,
		AddedPictures:    false,
	}
}

func requireBodyMatchProfile(t *testing.T, body *bytes.Buffer, profile db.Profile) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProfile db.Profile
	err = json.Unmarshal(data, &gotProfile)

	require.NoError(t, err)
	require.Equal(t, profile, gotProfile)
}
