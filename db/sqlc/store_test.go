package db

import (
	"context"
	"testing"
	"time"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	null "gopkg.in/guregu/null.v4"
)

func createRandomPlaylistTx(t *testing.T) Playlist {
	store := NewStore(testDB)

	user := createRandomUser(t)

	arg := CreatePlaylistTxParams{
		Name:         util.RandomName(),
		Description:  null.NewString(util.RandomName(), true),
		ImageURL:     null.NewString(util.RandomName(), true),
		IsPublic:     true,
		DeliveryDay:  null.NewString("Monday", true),
		Category:     null.NewString(util.RandomName(), true),
		UserID:       user.ID,
		DeliveryTime: null.NewTime(time.Now(), true),
		Status:       null.NewString(util.RandomName(), true),
	}

	playlist, err := store.CreatePlaylistTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlist)

	var userPlaylists []UserPlaylist
	userPlaylists, err = store.ListUser_Playlists(context.Background(), ListUser_PlaylistsParams{
		UserID: user.ID,
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)

	for _, userPlaylist := range userPlaylists {
		require.NotEmpty(t, userPlaylist)
		require.Equal(t, arg.UserID, userPlaylist.UserID)
		require.Equal(t, playlist.ID, userPlaylist.PlaylistID)
	}

	return playlist
}

func TestCreatePlaylistTx(t *testing.T) {
	createRandomPlaylistTx(t)
}

func TestCreatePlaylistTxConcurrency(t *testing.T) {
	store := NewStore(testDB)
	n := 5
	errs := make(chan error)
	playlists := make(chan Playlist)
	users := make(chan User)

	// run n concurrent user playlist transaction
	for i := 0; i < n; i++ {
		go func() {
			user := createRandomUser(t)

			playlist, err := store.CreatePlaylistTx(context.Background(), CreatePlaylistTxParams{
				Name:         util.RandomName(),
				Description:  null.NewString(util.RandomName(), true),
				ImageURL:     null.NewString(util.RandomName(), true),
				IsPublic:     true,
				DeliveryDay:  null.NewString("Monday", true),
				Category:     null.NewString(util.RandomName(), true),
				UserID:       user.ID,
				DeliveryTime: null.NewTime(time.Now(), true),
				Status:       null.NewString(util.RandomName(), true),
			})

			errs <- err
			playlists <- playlist
			users <- user
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		playlist := <-playlists
		require.NotEmpty(t, playlist)

		user := <-users
		require.NotEmpty(t, user)

		var userPlaylists []UserPlaylist
		userPlaylists, err = store.ListUser_Playlists(context.Background(), ListUser_PlaylistsParams{
			UserID: user.ID,
			Limit:  5,
			Offset: 5,
		})
		require.NoError(t, err)

		for _, userPlaylist := range userPlaylists {
			require.NotEmpty(t, userPlaylist)
			require.Equal(t, user.ID, userPlaylist.UserID)
			require.Equal(t, playlist.ID, userPlaylist.PlaylistID)
		}
	}
}
