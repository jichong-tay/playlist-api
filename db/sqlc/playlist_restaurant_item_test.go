package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jichong-tay/foodpanda-playlist-api/util"
	"github.com/stretchr/testify/require"
)

/*
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

*/

func createRandomPlaylist_Restaurant_Item(t *testing.T) PlaylistRestaurantItem {
	arg := CreatePlaylist_Restaurant_ItemParams{
		PlaylistID:             sql.NullInt32{Int32: int32(util.RandomInt(0, 10)), Valid: true},
		RestaurantItemID:       sql.NullInt32{Int32: int32(util.RandomInt(0, 10)), Valid: true},
		RestaurantItemQuantity: sql.NullInt32{Int32: int32(util.RandomInt(0, 10)), Valid: true},
		AddedAt:                sql.NullTime{Time: time.Now(), Valid: true},
	}

	playlistrestaurantitem, err := testQueries.CreatePlaylist_Restaurant_Item(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlistrestaurantitem)

	require.Equal(t, arg.PlaylistID, playlistrestaurantitem.PlaylistID)
	require.Equal(t, arg.PlaylistID, playlistrestaurantitem.PlaylistID)
	require.Equal(t, arg.PlaylistID, playlistrestaurantitem.PlaylistID)

	return playlistrestaurantitem
}

func TestCreatePlaylist_Restaurant_Item(t *testing.T) {
	createRandomPlaylist_Restaurant_Item(t)
}
