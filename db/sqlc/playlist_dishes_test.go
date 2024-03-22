package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomPlaylistDish(t *testing.T) PlaylistDish {
	playlist := createRandomPlaylist(t)
	dish := createRandomDish(t)

	arg := CreatePlaylist_DishParams{
		OrderID:      util.RandomInt(0, 100),
		PlaylistID:   playlist.ID,
		DishID:       dish.ID,
		DishQuantity: util.RandomInt(1, 5),
	}

	playlistDish, err := testQueries.CreatePlaylist_Dish(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlistDish)

	require.Equal(t, arg.OrderID, playlistDish.OrderID)
	require.Equal(t, arg.PlaylistID, playlistDish.PlaylistID)
	require.Equal(t, arg.DishID, playlistDish.DishID)
	require.Equal(t, arg.DishQuantity, playlistDish.DishQuantity)

	return playlistDish
}

func TestCreatePlaylist_Dish(t *testing.T) {
	createRandomPlaylistDish(t)
}

func TestGetPlaylist_Dish(t *testing.T) {
	playlistDish1 := createRandomPlaylistDish(t)
	playlistDish2, err := testQueries.GetPlaylist_Dish(context.Background(), playlistDish1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, playlistDish2)

	require.Equal(t, playlistDish1.OrderID, playlistDish2.OrderID)
	require.Equal(t, playlistDish1.PlaylistID, playlistDish2.PlaylistID)
	require.Equal(t, playlistDish1.DishID, playlistDish2.DishID)
	require.Equal(t, playlistDish1.DishQuantity, playlistDish2.DishQuantity)
}

func TestUpdatePlaylistDish(t *testing.T) {
	playlistDish1 := createRandomPlaylistDish(t)

	arg := UpdatePlaylist_DishParams{
		ID:           playlistDish1.ID,
		OrderID:      playlistDish1.OrderID,
		PlaylistID:   playlistDish1.PlaylistID,
		DishID:       playlistDish1.DishID,
		DishQuantity: util.RandomInt(0, 5),
		AddedAt:      time.Now(),
	}

	playlistDish2, err := testQueries.UpdatePlaylist_Dish(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, playlistDish2)

	require.Equal(t, arg.ID, playlistDish2.ID)
	require.Equal(t, arg.OrderID, playlistDish2.OrderID)
	require.Equal(t, arg.PlaylistID, playlistDish2.PlaylistID)
	require.Equal(t, arg.DishID, playlistDish2.DishID)
	require.Equal(t, arg.DishQuantity, playlistDish2.DishQuantity)
	require.NotEqual(t, arg.AddedAt, playlistDish2.CreatedAt)
}

func TestDeletePlaylist_Dish(t *testing.T) {
	playlistDish1 := createRandomPlaylistDish(t)
	err := testQueries.DeletePlaylist_Dish(context.Background(), playlistDish1.ID)
	require.NoError(t, err)

	userPlaylist2, err := testQueries.GetPlaylist_Dish(context.Background(), playlistDish1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userPlaylist2)
}

func TestListPlaylist_Dish(t *testing.T) {
	var lastPlaylistDish PlaylistDish
	for i := 0; i < 10; i++ {
		lastPlaylistDish = createRandomPlaylistDish(t)
	}
	arg := ListPlaylist_DishesParams{
		ID:     lastPlaylistDish.ID,
		Limit:  5,
		Offset: 5,
	}

	playlistDishes, err := testQueries.ListPlaylist_Dishes(context.Background(), arg)
	require.NoError(t, err)

	for _, playlistDish := range playlistDishes {
		require.NotEmpty(t, playlistDish)
		require.Equal(t, arg.ID, playlistDish.ID)
	}
}
