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

func createRandomUserPlaylist(t *testing.T) UserPlaylist {
	user := createRandomUser(t)
	playlist := createRandomPlaylist(t)
	arg := CreateUser_PlaylistParams{
		UserID:       user.ID,
		PlaylistID:   playlist.ID,
		DeliveryDay:  null.NewString("Monday", true),
		DeliveryTime: null.NewTime(time.Now(), true),
		Status:       null.NewString(util.RandomName(), true),
	}

	userPlaylist, err := testQueries.CreateUser_Playlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userPlaylist)

	require.Equal(t, arg.UserID, userPlaylist.UserID)
	require.Equal(t, arg.PlaylistID, userPlaylist.PlaylistID)
	require.Equal(t, arg.DeliveryDay, userPlaylist.DeliveryDay)
	require.Equal(t, arg.DeliveryTime.Time.Format("00:00:00"), userPlaylist.DeliveryTime.Time.Format("00:00:00"))
	require.Equal(t, arg.Status, userPlaylist.Status)

	return userPlaylist
}

func TestCreateUser_Playlist(t *testing.T) {
	createRandomUserPlaylist(t)
}

func TestGetUser_Playlist(t *testing.T) {
	userPlaylist1 := createRandomUserPlaylist(t)
	userPlaylist2, err := testQueries.GetUser_Playlist(context.Background(), userPlaylist1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userPlaylist2)

	require.Equal(t, userPlaylist1.UserID, userPlaylist2.UserID)
	require.Equal(t, userPlaylist1.PlaylistID, userPlaylist2.PlaylistID)
	require.Equal(t, userPlaylist1.DeliveryDay, userPlaylist2.DeliveryDay)
	require.Equal(t, userPlaylist1.DeliveryTime.Time.Format("00:00:00"), userPlaylist2.DeliveryTime.Time.Format("00:00:00"))
	require.Equal(t, userPlaylist1.Status, userPlaylist2.Status)
}

func TestUpdateUserPlaylist(t *testing.T) {
	userPlaylist1 := createRandomUserPlaylist(t)

	arg := UpdateUser_PlaylistParams{
		ID:           userPlaylist1.ID,
		UserID:       userPlaylist1.UserID,
		PlaylistID:   userPlaylist1.PlaylistID,
		DeliveryDay:  null.NewString("Tuesday", true),
		DeliveryTime: null.NewTime(time.Now(), true),
		Status:       null.NewString(util.RandomName(), true),
	}

	userPlaylist2, err := testQueries.UpdateUser_Playlist(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userPlaylist2)

	require.Equal(t, arg.ID, userPlaylist2.ID)
	require.Equal(t, arg.UserID, userPlaylist2.UserID)
	require.Equal(t, arg.PlaylistID, userPlaylist2.PlaylistID)
	require.Equal(t, arg.DeliveryDay, userPlaylist2.DeliveryDay)
	require.Equal(t, arg.DeliveryTime.Time.Format("00:00:00"), userPlaylist2.DeliveryTime.Time.Format("00:00:00"))
	require.Equal(t, arg.Status, userPlaylist2.Status)
}

func TestDeleteUser_Playlist(t *testing.T) {
	userPlaylist1 := createRandomUserPlaylist(t)
	err := testQueries.DeleteUser_Playlist(context.Background(), userPlaylist1.ID)
	require.NoError(t, err)

	userPlaylist2, err := testQueries.GetUser_Playlist(context.Background(), userPlaylist1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userPlaylist2)
}

func TestListUser_Playlist(t *testing.T) {
	var lastUserPlaylist UserPlaylist
	for i := 0; i < 10; i++ {
		lastUserPlaylist = createRandomUserPlaylist(t)
	}
	arg := ListUser_PlaylistsParams{
		UserID: lastUserPlaylist.ID,
		Limit:  5,
		Offset: 5,
	}

	userPlaylists, err := testQueries.ListUser_Playlists(context.Background(), arg)
	require.NoError(t, err)
	for _, userPlaylist := range userPlaylists {
		require.NotEmpty(t, userPlaylist)
		require.Equal(t, arg.UserID, userPlaylist.UserID)
	}
}
