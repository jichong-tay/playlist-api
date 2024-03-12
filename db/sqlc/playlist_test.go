package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func createRandomPlaylist(t *testing.T) Playlist {
	arg := CreatePlaylistParams{
		Name:        util.RandomName(),
		Description: null.NewString(util.RandomName(), true),
		ImageUrl:    null.NewString(util.RandomName(), true),
		IsPublic:    true,
		DeliveryDay: null.NewString("Monday", true),
		Category:    null.NewString(util.RandomName(), true),
	}

	playlist, err := testQueries.CreatePlaylist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlist)

	require.Equal(t, arg.Name, playlist.Name)
	require.Equal(t, arg.Description, playlist.Description)
	require.Equal(t, arg.ImageUrl, playlist.ImageUrl)
	require.Equal(t, arg.IsPublic, playlist.IsPublic)
	require.Equal(t, arg.DeliveryDay, playlist.DeliveryDay)
	require.Equal(t, arg.Category, playlist.Category)

	return playlist
}

func TestCreatePlaylist(t *testing.T) {
	createRandomPlaylist(t)
}

func TestGetPlaylist(t *testing.T) {
	playlist1 := createRandomPlaylist(t)
	playlist2, err := testQueries.GetPlaylist(context.Background(), playlist1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, playlist2)

	require.Equal(t, playlist1.Name, playlist2.Name)
	require.Equal(t, playlist1.Description, playlist2.Description)
	require.Equal(t, playlist1.ImageUrl, playlist2.ImageUrl)
	require.Equal(t, playlist1.IsPublic, playlist2.IsPublic)
	require.Equal(t, playlist1.DeliveryDay, playlist2.DeliveryDay)
	require.Equal(t, playlist1.Category, playlist2.Category)

}

func TestUpdatePlaylist(t *testing.T) {
	playlist1 := createRandomPlaylist(t)

	arg := UpdatePlaylistParams{
		ID:          playlist1.ID,
		Name:        util.RandomName(),
		Description: null.NewString(util.RandomName(), true),
		ImageUrl:    null.NewString(util.RandomName(), true),
		IsPublic:    true,
		DeliveryDay: null.NewString("Monday", true),
		Category:    null.NewString(util.RandomName(), true),
		AddedAt:     time.Now(),
	}

	playlist2, err := testQueries.UpdatePlaylist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlist2)

	require.Equal(t, arg.Name, playlist2.Name)
	require.Equal(t, arg.Description, playlist2.Description)
	require.Equal(t, arg.ImageUrl, playlist2.ImageUrl)
	require.Equal(t, arg.IsPublic, playlist2.IsPublic)
	require.Equal(t, arg.DeliveryDay, playlist2.DeliveryDay)
	require.Equal(t, arg.Category, playlist2.Category)
	require.WithinDuration(t, arg.AddedAt, playlist2.AddedAt, 1*time.Second)

}

func TestDeletePlaylist(t *testing.T) {
	playlist1 := createRandomPlaylist(t)
	err := testQueries.DeletePlaylist(context.Background(), playlist1.ID)
	require.NoError(t, err)

	playlist2, err := testQueries.GetPlaylist(context.Background(), playlist1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, playlist2)
}

func TestListPlaylist(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomPlaylist(t)
	}
	arg := ListPlaylistsParams{
		Limit:  5,
		Offset: 5,
	}

	playlists, err := testQueries.ListPlaylists(context.Background(), arg)
	require.NoError(t, err)
	//require.Len(t, playlists, 5)

	for _, playlist := range playlists {
		require.NotEmpty(t, playlist)
	}
}
