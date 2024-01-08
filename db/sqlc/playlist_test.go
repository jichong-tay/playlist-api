package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jichong-tay/foodpanda-playlist-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomPlaylist(t *testing.T) Playlist {
	arg := CreatePlaylistParams{
		Name:        sql.NullString{String: util.RandomName(), Valid: true},
		Description: sql.NullString{String: util.RandomName(), Valid: true},
		ImageUrl:    sql.NullString{String: util.RandomName(), Valid: true},
		IsPublic:    sql.NullBool{Bool: false, Valid: true},
		DeliveryDay: sql.NullString{String: "01-01-2024", Valid: true},
		Category:    sql.NullString{String: util.RandomName(), Valid: true},
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
		Name:        sql.NullString{String: util.RandomName(), Valid: true},
		Description: sql.NullString{String: util.RandomName(), Valid: true},
		ImageUrl:    sql.NullString{String: util.RandomName(), Valid: true},
		IsPublic:    sql.NullBool{Bool: false, Valid: true},
		DeliveryDay: sql.NullString{String: "01-01-2024", Valid: true},
		Category:    sql.NullString{String: util.RandomName(), Valid: true},
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
	require.Len(t, playlists, 5)

	for _, playlist := range playlists {
		require.NotEmpty(t, playlist)
	}
}
