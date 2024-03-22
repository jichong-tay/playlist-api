// package db is a package that contains the database access code for the playlist service.
package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jichong-tay/playlist-api/util"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func createRandomDish(t *testing.T) Dish {
	restaurant := createRandomRestaurant(t)
	arg := CreateDishParams{
		RestaurantID: restaurant.ID,
		IsAvailable:  true,
		Name:         util.RandomName(),
		Description:  null.NewString(util.RandomName(), true),
		Price:        float64(util.RandomInt(1, 9999)) / 100,
		Cuisine:      null.NewString(util.RandomName(), true),
		ImageUrl:     null.NewString(util.RandomName(), true),
	}

	dish, err := testQueries.CreateDish(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, dish)

	require.Equal(t, arg.RestaurantID, dish.RestaurantID)
	require.Equal(t, arg.IsAvailable, dish.IsAvailable)
	require.Equal(t, arg.Name, dish.Name)
	require.Equal(t, arg.Description, dish.Description)
	require.Equal(t, arg.Price, dish.Price)
	require.Equal(t, arg.Cuisine, dish.Cuisine)
	require.Equal(t, arg.ImageUrl, dish.ImageUrl)

	return dish
}

func TestCreateDish(t *testing.T) {
	createRandomDish(t)
}

func TestGetDish(t *testing.T) {
	dish1 := createRandomDish(t)
	dish2, err := testQueries.GetDish(context.Background(), dish1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dish2)

	require.Equal(t, dish1.RestaurantID, dish2.RestaurantID)
	require.Equal(t, dish1.IsAvailable, dish2.IsAvailable)
	require.Equal(t, dish1.Name, dish2.Name)
	require.Equal(t, dish1.Description, dish2.Description)
	require.Equal(t, dish1.Price, dish2.Price)
	require.Equal(t, dish1.Cuisine, dish2.Cuisine)
	require.Equal(t, dish1.ImageUrl, dish2.ImageUrl)
}

func TestUpdateDish(t *testing.T) {
	dish1 := createRandomDish(t)

	arg := UpdateDishParams{
		ID:           dish1.ID,
		RestaurantID: dish1.RestaurantID,
		IsAvailable:  false,
		Name:         util.RandomName(),
		Description:  null.NewString(util.RandomName(), true),
		Price:        float64(util.RandomInt(1, 9999)) / 100,
		Cuisine:      null.NewString(util.RandomName(), true),
		ImageUrl:     null.NewString(util.RandomName(), true),
	}

	dish2, err := testQueries.UpdateDish(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, dish2)

	require.Equal(t, arg.ID, dish2.ID)
	require.Equal(t, arg.RestaurantID, dish2.RestaurantID)
	require.Equal(t, arg.IsAvailable, dish2.IsAvailable)
	require.Equal(t, arg.Name, dish2.Name)
	require.Equal(t, arg.Description, dish2.Description)
	require.Equal(t, arg.Price, dish2.Price)
	require.Equal(t, arg.Cuisine, dish2.Cuisine)
	require.Equal(t, arg.ImageUrl, dish2.ImageUrl)
}

func TestDeleteDish(t *testing.T) {
	dish1 := createRandomDish(t)
	err := testQueries.DeleteDish(context.Background(), dish1.ID)
	require.NoError(t, err)

	dish2, err := testQueries.GetDish(context.Background(), dish1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, dish2)
}

func TestListDish(t *testing.T) {
	var lastDish Dish
	for i := 0; i < 10; i++ {
		lastDish = createRandomDish(t)
	}

	arg := ListDishesParams{
		ID:     lastDish.ID,
		Limit:  5,
		Offset: 5,
	}

	dishes, err := testQueries.ListDishes(context.Background(), arg)
	require.NoError(t, err)
	for _, dish := range dishes {
		require.NotEmpty(t, dish)
	}
}
