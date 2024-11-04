package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestCreateGuestAPI(t *testing.T) {
	userID := util.RandomID()
	guest := randomGuest(userID)

	body := createGuestBody{
		Name:           guest.Name,
		Phone:          guest.Phone,
		Allergies:      []string{},
		IsVegetarian:   guest.IsVegetarian,
		NeedsTransport: guest.NeedsTransport,
	}

	arg := db.CreateGuestParams{
		ProfileID:      userID,
		Name:           guest.Name,
		Phone:          guest.Phone,
		Allergies:      guest.Allergies,
		IsVegetarian:   guest.IsVegetarian,
		NeedsTransport: guest.NeedsTransport,
	}

	testCases := []struct {
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		name          string
		body          createGuestBody
	}{
		{
			name: "OK",
			body: body,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateGuest(gomock.Any(), gomock.Eq(arg)).Times(1).Return(guest, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGuest(t, recorder.Body, guest)
			},
		},
		{
			name: "InternalError",
			body: body,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateGuest(gomock.Any(), gomock.Eq(arg)).Times(1).Return(guest, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: createGuestBody{},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateGuest(gomock.Any(), gomock.Any()).Times(0)
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

			request, err := http.NewRequest(http.MethodPost, "/api/guests", bytes.NewReader(body))
			require.NoError(t, err)

			request.Header.Set(auth.Header, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetGuestAPI(t *testing.T) {
	userID := util.RandomID()
	guest := randomGuest(userID)

	testCases := []struct {
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		name          string
		id            int64
	}{
		{
			name: "OK",
			id:   guest.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetGuest(gomock.Any(), gomock.Eq(guest.ID)).Times(1).Return(guest, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGuest(t, recorder.Body, guest)
			},
		},
		{
			name: "NotFound",
			id:   guest.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetGuest(gomock.Any(), gomock.Eq(guest.ID)).Times(1).Return(db.Guest{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			id:   guest.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetGuest(gomock.Any(), gomock.Eq(guest.ID)).Times(1).Return(db.Guest{}, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			id:   0,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetGuest(gomock.Any(), gomock.Any()).Times(0)
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

			// start the server
			server := NewServer(store, jwks)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/guests/%d", tc.id)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			request.Header.Set(auth.Header, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdateGuestAPI(t *testing.T) {
	new_name := util.RandomName()

	guest := randomGuest(util.RandomID())
	guest.Name = new_name

	body := updateGuestBody{
		Name: new_name,
	}

	arg := db.UpdateGuestParams{
		ID:   guest.ID,
		Name: pgtype.Text{String: new_name, Valid: true},
	}

	testCases := []struct {
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		name          string
		body          updateGuestBody
		id            int64
	}{
		{
			name: "OK",
			body: body,
			id:   guest.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateGuest(gomock.Any(), gomock.Eq(arg)).Times(1).Return(guest, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGuest(t, recorder.Body, guest)
			},
		},
		{
			name: "InternalError",
			body: body,
			id:   guest.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateGuest(gomock.Any(), gomock.Eq(arg)).Times(1).Return(guest, pgx.ErrTxClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: updateGuestBody{},
			id:   0,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateGuest(gomock.Any(), gomock.Any()).Times(0)
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

			jwks := NewMockJWKS(uuid.New())
			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			server := NewServer(store, jwks)
			recorder := httptest.NewRecorder()

			body, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/api/guests/%d", tc.id)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
			require.NoError(t, err)

			request.Header.Set(auth.Header, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomGuest(profileID uuid.UUID) db.Guest {
	return db.Guest{
		ProfileID:      profileID,
		ID:             int64(util.RandomInt(1, 1000)),
		Name:           util.RandomName(),
		Phone:          util.RandomPhoneNumber(),
		IsVegetarian:   true,
		Allergies:      []string{},
		NeedsTransport: false,
	}
}

func requireBodyMatchGuest(t *testing.T, body *bytes.Buffer, guest db.Guest) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotGuest db.Guest
	err = json.Unmarshal(data, &gotGuest)

	require.NoError(t, err)
	require.Equal(t, guest, gotGuest)
}
