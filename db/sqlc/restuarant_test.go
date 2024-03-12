package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	null "gopkg.in/guregu/null.v4"
)

func createRandomRestaurant(t *testing.T) Restaurant {
	arg := CreateRestaurantParams{
		Name:        util.RandomName(),
		Description: null.NewString(util.RandomName(), true),
		Location:    null.NewString(util.RandomName(), true),
		Cuisine:     null.NewString(util.RandomName(), true),
		ImageUrl:    null.NewString(util.RandomName(), true),
	}

	restaurant, err := testQueries.CreateRestaurant(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restaurant)

	require.Equal(t, arg.Name, restaurant.Name)
	require.Equal(t, arg.Description, restaurant.Description)
	require.Equal(t, arg.Location, restaurant.Location)
	require.Equal(t, arg.Cuisine, restaurant.Cuisine)
	require.Equal(t, arg.ImageUrl, restaurant.ImageUrl)

	return restaurant
}

func TestCreateRestuarant(t *testing.T) {
	createRandomRestaurant(t)
}

func TestGetResturant(t *testing.T) {
	restaurant1 := createRandomRestaurant(t)
	restaurant2, err := testQueries.GetRestaurant(context.Background(), restaurant1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, restaurant2)

	require.Equal(t, restaurant1.Name, restaurant2.Name)
	require.Equal(t, restaurant1.Description, restaurant2.Description)
	require.Equal(t, restaurant1.Location, restaurant2.Location)
	require.Equal(t, restaurant1.Cuisine, restaurant2.Cuisine)
	require.Equal(t, restaurant1.ImageUrl, restaurant2.ImageUrl)
}

func TestUpdateRestaurant(t *testing.T) {
	Restaurant1 := createRandomRestaurant(t)

	arg := UpdateRestaurantParams{
		ID:          Restaurant1.ID,
		Name:        util.RandomName(),
		Description: null.NewString(util.RandomName(), true),
		Location:    null.NewString(util.RandomName(), true),
		Cuisine:     null.NewString(util.RandomName(), true),
		ImageUrl:    null.NewString(util.RandomName(), true),
	}

	Restaurant2, err := testQueries.UpdateRestaurant(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Restaurant2)

	require.Equal(t, arg.Name, Restaurant2.Name)
	require.Equal(t, arg.Description, Restaurant2.Description)
	require.Equal(t, arg.Location, Restaurant2.Location)
	require.Equal(t, arg.Cuisine, Restaurant2.Cuisine)
	require.Equal(t, arg.ImageUrl, Restaurant2.ImageUrl)

}

func TestDeleteRestuarant(t *testing.T) {
	restaurant1 := createRandomRestaurant(t)
	err := testQueries.DeleteRestaurant(context.Background(), restaurant1.ID)
	require.NoError(t, err)

	restaurant2, err := testQueries.GetRestaurant(context.Background(), restaurant1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, restaurant2)
}

func TestListRestaurants(t *testing.T) {
	var lastRestaurant Restaurant
	for i := 0; i < 10; i++ {
		lastRestaurant = createRandomRestaurant(t)
	}
	arg := ListRestaurantsParams{
		ID:     lastRestaurant.ID,
		Limit:  5,
		Offset: 5,
	}

	restaurants, err := testQueries.ListRestaurants(context.Background(), arg)
	require.NoError(t, err)
	//require.Len(t, restaurants, 5)

	for _, restaurant := range restaurants {
		require.NotEmpty(t, restaurant)
		require.Equal(t, arg.ID, restaurant.ID)
	}
}
