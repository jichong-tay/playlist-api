package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/jichong-tay/playlist-api/db/mock"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gopkg.in/guregu/null.v4"
)

func TestGetPlaylistAPI(t *testing.T) {
	playlist := randomPlaylist()

	testCases := []struct {
		name          string
		playlistID    int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			playlistID: playlist.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPlaylist(gomock.Any(), gomock.Eq(playlist.ID)).
					Times(1).
					Return(playlist, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPlaylist(t, recorder.Body, playlist)
			},
		},

		{
			name:       "NotFound",
			playlistID: playlist.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPlaylist(gomock.Any(), gomock.Eq(playlist.ID)).
					Times(1).
					Return(db.Playlist{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},

		{
			name:       "InternalError",
			playlistID: playlist.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPlaylist(gomock.Any(), gomock.Eq(playlist.ID)).
					Times(1).
					Return(db.Playlist{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},

		{
			name:       "InvalidID",
			playlistID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPlaylist(gomock.Any(), gomock.Any()).
					Times(0)
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

			store := mockdb.NewMockStore(ctrl)
			//build stubs
			tc.buildStubs(store)

			//start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/playlists/%d", tc.playlistID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)

		})

	}
}

func randomPlaylist() db.Playlist {
	return db.Playlist{
		ID:          util.RandomInt(1, 1000),
		Name:        util.RandomName(),
		Description: null.NewString(util.RandomName(), true),
		ImageUrl:    null.NewString(util.RandomName(), true),
		IsPublic:    true,
		DeliveryDay: null.NewString(util.RandomName(), true),
		Category:    null.NewString(util.RandomName(), true),
	}
}

func requireBodyMatchPlaylist(t *testing.T, body *bytes.Buffer, playlist db.Playlist) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotPlaylist db.Playlist
	err = json.Unmarshal(data, &gotPlaylist)
	require.NoError(t, err)
	require.Equal(t, playlist, gotPlaylist)

}
