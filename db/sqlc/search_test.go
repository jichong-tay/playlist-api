package db

import (
	"context"
	"testing"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func createRandomSearch(t *testing.T) Search {

	user := createRandomUser(t)
	arg := CreateSearchParams{
		UserID:  user.ID,
		Keyword: null.NewString(util.RandomName(), true),
	}

	search, err := testQueries.CreateSearch(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, search)

	require.Equal(t, arg.UserID, search.UserID)
	require.Equal(t, arg.Keyword, search.Keyword)

	return search
}

func TestCreateSearch(t *testing.T) {
	createRandomSearch(t)
}

func TestGetSearch(t *testing.T) {
	search1 := createRandomSearch(t)
	search2, err := testQueries.GetSearch(context.Background(), search1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, search2)

	require.Equal(t, search1.UserID, search2.UserID)
	require.Equal(t, search1.Keyword, search2.Keyword)

}

func TestListSearch(t *testing.T) {
	var lastSearch Search
	for i := 0; i < 10; i++ {
		lastSearch = createRandomSearch(t)
	}
	arg := ListSearchesParams{
		ID:     lastSearch.ID,
		Limit:  5,
		Offset: 0,
	}

	searches, err := testQueries.ListSearches(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, searches)

	for _, search := range searches {
		require.NotEmpty(t, search)
		require.Equal(t, arg.ID, search.ID)
	}

}
