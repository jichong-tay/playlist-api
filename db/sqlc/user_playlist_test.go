package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	null "gopkg.in/guregu/null.v4"
)

func createRandomUser_Playlist(t *testing.T) UserPlaylist {

	user := createRandomUser(t)
	playlist := createRandomPlaylist(t)
	arg := CreateUser_PlaylistParams{
		UserID:       user.ID,
		PlaylistID:   playlist.ID,
		DeliveryDay:  null.NewString("Monday", true),
		DeliveryTime: null.NewTime(time.Now(), true),
		Status:       null.NewString(util.RandomName(), true),
	}

	user_playlist, err := testQueries.CreateUser_Playlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user_playlist)

	require.Equal(t, arg.UserID, user_playlist.UserID)
	require.Equal(t, arg.PlaylistID, user_playlist.PlaylistID)
	require.Equal(t, arg.DeliveryDay, user_playlist.DeliveryDay)
	require.Equal(t, arg.DeliveryTime.Time.Format("00:00:00"), user_playlist.DeliveryTime.Time.Format("00:00:00"))
	require.Equal(t, arg.Status, user_playlist.Status)

	return user_playlist
}

func TestCreateUser_Playlist(t *testing.T) {
	createRandomUser_Playlist(t)
}

func TestGetUser_Playlist(t *testing.T) {
	user_playlist1 := createRandomUser_Playlist(t)
	user_playlist2, err := testQueries.GetUser_Playlist(context.Background(), user_playlist1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user_playlist2)

	require.Equal(t, user_playlist1.UserID, user_playlist2.UserID)
	require.Equal(t, user_playlist1.PlaylistID, user_playlist2.PlaylistID)
	require.Equal(t, user_playlist1.DeliveryDay, user_playlist2.DeliveryDay)
	require.Equal(t, user_playlist1.DeliveryTime.Time.Format("00:00:00"), user_playlist2.DeliveryTime.Time.Format("00:00:00"))
	require.Equal(t, user_playlist1.Status, user_playlist2.Status)

}

func TestUpdateUserPlaylist(t *testing.T) {
	user_playlist1 := createRandomUser_Playlist(t)

	arg := UpdateUser_PlaylistParams{
		ID:           user_playlist1.ID,
		UserID:       user_playlist1.UserID,
		PlaylistID:   user_playlist1.PlaylistID,
		DeliveryDay:  null.NewString("Tuesday", true),
		DeliveryTime: null.NewTime(time.Now(), true),
		Status:       null.NewString(util.RandomName(), true),
	}

	user_playlist2, err := testQueries.UpdateUser_Playlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user_playlist2)

	require.Equal(t, arg.ID, user_playlist2.ID)
	require.Equal(t, arg.UserID, user_playlist2.UserID)
	require.Equal(t, arg.PlaylistID, user_playlist2.PlaylistID)
	require.Equal(t, arg.DeliveryDay, user_playlist2.DeliveryDay)
	require.Equal(t, arg.DeliveryTime.Time.Format("00:00:00"), user_playlist2.DeliveryTime.Time.Format("00:00:00"))
	require.Equal(t, arg.Status, user_playlist2.Status)

}

func TestDeleteUser_Playlist(t *testing.T) {
	user_playlist1 := createRandomUser_Playlist(t)
	err := testQueries.DeleteUser_Playlist(context.Background(), user_playlist1.ID)
	require.NoError(t, err)

	user_playlist2, err := testQueries.GetUser_Playlist(context.Background(), user_playlist1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user_playlist2)
}

func TestListUser_Playlist(t *testing.T) {
	var lastUserPlaylist UserPlaylist
	for i := 0; i < 10; i++ {
		lastUserPlaylist = createRandomUser_Playlist(t)
	}
	arg := ListUser_PlaylistsParams{
		UserID: lastUserPlaylist.ID,
		Limit:  5,
		Offset: 5,
	}

	user_playlists, err := testQueries.ListUser_Playlists(context.Background(), arg)
	require.NoError(t, err)
	//require.Len(t, playlists, 5)

	for _, user_playlist := range user_playlists {
		require.NotEmpty(t, user_playlist)
		require.Equal(t, arg.UserID, user_playlist.UserID)
	}
}
