package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5"
	mockdb "github.com/presedo93/wedding/back/db/mock"
	db "github.com/presedo93/wedding/back/db/sqlc"
	"github.com/presedo93/wedding/back/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetGuestAPI(t *testing.T) {
	guest := randomGuest()

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

			jwks := NewMockJWKS("some-user")
			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			// start the server
			server := NewServer(store, jwks)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/guests/%d", tc.id)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			request.Header.Set(authHeader, "Bearer some-token")
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomGuest() db.Guest {
	return db.Guest{
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
