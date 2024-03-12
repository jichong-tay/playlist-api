package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomPlaylist_Dish(t *testing.T) PlaylistDish {

	playlist := createRandomPlaylist(t)
	dish := createRandomDish(t)

	arg := CreatePlaylist_DishParams{
		OrderID:      util.RandomInt(0, 100),
		PlaylistID:   playlist.ID,
		DishID:       dish.ID,
		DishQuantity: util.RandomInt(1, 5),
	}

	playlist_dish, err := testQueries.CreatePlaylist_Dish(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlist_dish)

	require.Equal(t, arg.OrderID, playlist_dish.OrderID)
	require.Equal(t, arg.PlaylistID, playlist_dish.PlaylistID)
	require.Equal(t, arg.DishID, playlist_dish.DishID)
	require.Equal(t, arg.DishQuantity, playlist_dish.DishQuantity)

	return playlist_dish
}

func TestCreatePlaylist_Dish(t *testing.T) {
	createRandomPlaylist_Dish(t)
}

func TestGetPlaylist_Dish(t *testing.T) {
	playlist_dish1 := createRandomPlaylist_Dish(t)
	playlist_dish2, err := testQueries.GetPlaylist_Dish(context.Background(), playlist_dish1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, playlist_dish2)

	require.Equal(t, playlist_dish1.OrderID, playlist_dish2.OrderID)
	require.Equal(t, playlist_dish1.PlaylistID, playlist_dish2.PlaylistID)
	require.Equal(t, playlist_dish1.DishID, playlist_dish2.DishID)
	require.Equal(t, playlist_dish1.DishQuantity, playlist_dish2.DishQuantity)

}

func TestUpdatePlaylistDish(t *testing.T) {
	playlist_dish1 := createRandomPlaylist_Dish(t)

	arg := UpdatePlaylist_DishParams{
		ID:           playlist_dish1.ID,
		OrderID:      playlist_dish1.OrderID,
		PlaylistID:   playlist_dish1.PlaylistID,
		DishID:       playlist_dish1.DishID,
		DishQuantity: util.RandomInt(0, 5),
		AddedAt:      time.Now(),
	}

	playlist_dish2, err := testQueries.UpdatePlaylist_Dish(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlist_dish2)

	require.Equal(t, arg.ID, playlist_dish2.ID)
	require.Equal(t, arg.OrderID, playlist_dish2.OrderID)
	require.Equal(t, arg.PlaylistID, playlist_dish2.PlaylistID)
	require.Equal(t, arg.DishID, playlist_dish2.DishID)
	require.Equal(t, arg.DishQuantity, playlist_dish2.DishQuantity)
	require.NotEqual(t, arg.AddedAt, playlist_dish2.CreatedAt)

}

func TestDeletePlaylist_Dish(t *testing.T) {
	playlist_dish1 := createRandomPlaylist_Dish(t)
	err := testQueries.DeletePlaylist_Dish(context.Background(), playlist_dish1.ID)
	require.NoError(t, err)

	user_playlist2, err := testQueries.GetPlaylist_Dish(context.Background(), playlist_dish1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user_playlist2)
}

func TestListPlaylist_Dish(t *testing.T) {
	var lastPlaylistDish PlaylistDish
	for i := 0; i < 10; i++ {
		lastPlaylistDish = createRandomPlaylist_Dish(t)
	}
	arg := ListPlaylist_DishesParams{
		ID:     lastPlaylistDish.ID,
		Limit:  5,
		Offset: 5,
	}

	playlist_dishes, err := testQueries.ListPlaylist_Dishes(context.Background(), arg)
	require.NoError(t, err)
	//require.Len(t, playlists, 5)

	for _, playlist_dish := range playlist_dishes {
		require.NotEmpty(t, playlist_dish)
		require.Equal(t, arg.ID, playlist_dish.ID)
	}
}
