package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
)

type dish struct {
	ID           int64        `form:"id" json:"id"`
	RestaurantID int64        `form:"restaurantId" json:"restaurantId"`
	IsAvailable  bool         `form:"isAvailable" json:"isAvailable"`
	Name         string       `form:"name" json:"name"`
	Description  string       `form:"description" json:"description"`
	Price        float64      `form:"price" json:"price"`
	Cuisine      string       `form:"cuisine" json:"cuisine"`
	ImageURL     string       `form:"imageUrl" json:"imageUrl"`
	CreatedAt    time.Time    `form:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time    `form:"updatedAt" json:"updatedAt"`
	PlaylistDish playlistDish `form:"PlaylistDish" json:"PlaylistDish"`
}

type playlistDish struct {
	ID           int64     `form:"id" json:"id"`
	PlaylistID   int64     `form:"playlistId" json:"playlistId"`
	DishID       int64     `form:"dishId" json:"dishId"`
	DishQuantity int64     `form:"dishQuantity" json:"dishQuantity"`
	CreatedAt    time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `form:"updatedAt" json:"updatedAt"`
}

type categoryPlaylist struct {
	CategoryTitle string       `form:"categoryTitle" json:"categoryTitle"`
	Playlist      []playlistv1 `form:"list" json:"list"`
}
type playlistv1 struct {
	ID          int64     `form:"id" json:"id"`
	Name        string    `form:"name" json:"name"`
	Description string    `form:"description" json:"description"`
	ImageURL    string    `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool      `form:"isPublic" json:"isPublic"`
	DeliveryDay string    `form:"deliveryDay" json:"deliveryDay"`
	Category    string    `form:"category" json:"category"`
	CreatedAt   time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `form:"updatedAt" json:"updatedAt"`
	Dishes      []dish    `form:"dishes" json:"dishes"`
	Cost        string    `form:"cost" json:"cost"`
}

type playlistv2 struct {
	ID          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	ImageURL    string `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool   `form:"isPublic" json:"isPublic"`
	DeliveryDay string `form:"deliveryDay" json:"deliveryDay"`
	Category    string `form:"category" json:"category"`
	Cost        string `form:"cost" json:"cost"`
	Status      string `form:"status" json:"status"`
}

type playlistLatestResponsev1 struct {
	CategoryPlaylist []categoryPlaylist `form:"publicPlaylist" json:"publicPlaylist"`
	UserPlaylist     []playlistv1       `form:"userPlaylist" json:"userPlaylist"`
}

type playlistLatestResponsev2 struct {
	PublicPlaylist []playlistv2 `form:"publicPlaylist" json:"publicPlaylist"`
	UserPlaylist   []playlistv2 `form:"userPlaylist" json:"userPlaylist"`
}
type playlistUserResponse struct {
	UserPlaylist []playlistv2 `form:"userPlaylist" json:"userPlaylist"`
}

type playlistResponsev1 struct {
	Playlist []playlistv1 `form:"list" json:"list"`
}

type playlistResponsev2 struct {
	Playlist []playlistv2 `form:"list" json:"list"`
}

func (server *Server) maptoModelV1(ctx *gin.Context, playlistsDB []db.Playlist) ([]playlistv1, error) {
	var playlists []playlistv1
	var dishes []dish

	for _, PlaylistDB := range playlistsDB {
		var playlist playlistv1
		var dish dish
		var cost float64

		playlistdishesDB, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, PlaylistDB.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return nil, err
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return nil, err
		}

		for _, playlistdishDB := range playlistdishesDB {
			dishDB, err := server.store.GetDish(ctx, playlistdishDB.DishID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errResponse(err))
					return nil, err
				}
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return nil, err
			}

			dish.ID = dishDB.ID
			dish.RestaurantID = dishDB.RestaurantID
			dish.IsAvailable = dishDB.IsAvailable
			dish.Name = dishDB.Name
			dish.Description = dishDB.Description.String
			dish.Price = dishDB.Price
			dish.Cuisine = dishDB.Cuisine.String
			dish.ImageURL = dishDB.ImageUrl.String
			dish.PlaylistDish = playlistDish{
				ID:           playlistdishDB.ID,
				PlaylistID:   playlistdishDB.PlaylistID,
				DishID:       playlistdishDB.DishID,
				DishQuantity: playlistdishDB.DishQuantity,
				CreatedAt:    playlistdishDB.CreatedAt,
				UpdatedAt:    playlistdishDB.AddedAt,
			}

			dishes = append(dishes, dish)
			cost += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		playlist.ID = PlaylistDB.ID
		playlist.Name = PlaylistDB.Name
		playlist.Description = PlaylistDB.Description.String
		playlist.ImageURL = PlaylistDB.ImageUrl.String
		playlist.IsPublic = PlaylistDB.IsPublic
		playlist.DeliveryDay = PlaylistDB.DeliveryDay.String
		playlist.Category = PlaylistDB.Category.String
		playlist.CreatedAt = PlaylistDB.CreatedAt
		playlist.UpdatedAt = PlaylistDB.AddedAt
		playlist.Dishes = dishes
		playlist.Cost = fmt.Sprintf("%.2f", cost)

		playlists = append(playlists, playlist)

	}

	return playlists, nil
}

func (server *Server) maptoModelV2(ctx *gin.Context, playlistsDB []db.Playlist) ([]playlistv2, error) {
	var playlists []playlistv2

	for _, playlistDB := range playlistsDB {
		var playlist playlistv2
		var dish dish
		var cost float64

		playlistdishesDB, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, playlistDB.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return nil, err
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return nil, err
		}

		for _, playlistdishDB := range playlistdishesDB {
			dishDB, err := server.store.GetDish(ctx, playlistdishDB.DishID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errResponse(err))
					return nil, err
				}
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return nil, err
			}

			dish.ID = dishDB.ID
			dish.RestaurantID = dishDB.RestaurantID
			dish.IsAvailable = dishDB.IsAvailable
			dish.Name = dishDB.Name
			dish.Price = dishDB.Price
			dish.Cuisine = dishDB.Cuisine.String
			dish.ImageURL = dishDB.ImageUrl.String
			dish.PlaylistDish = playlistDish{
				ID:           playlistdishDB.ID,
				PlaylistID:   playlistdishDB.PlaylistID,
				DishID:       playlistdishDB.DishID,
				DishQuantity: playlistdishDB.DishQuantity,
				CreatedAt:    playlistdishDB.CreatedAt,
				UpdatedAt:    playlistdishDB.AddedAt,
			}

			cost += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		playlist.ID = playlistDB.ID
		playlist.Name = playlistDB.Name
		playlist.ImageURL = playlistDB.ImageUrl.String
		playlist.IsPublic = playlistDB.IsPublic
		playlist.DeliveryDay = playlistDB.DeliveryDay.String
		playlist.Category = playlistDB.Category.String
		playlist.Cost = fmt.Sprintf("%.2f", cost)

		status, _ := server.store.ListStatusByPlaylistID(ctx, playlistDB.ID)
		playlist.Status = status.String

		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (server *Server) maptoModelCategory(ctx *gin.Context, playlistsDB []db.Playlist) ([]categoryPlaylist, error) {

	var categoryPlaylists []categoryPlaylist
	var categoryplaylist categoryPlaylist
	var playlistMap = make(map[string][]playlistv1)

	var playlists []playlistv1
	var dishes []dish

	for _, PlaylistDB := range playlistsDB {
		var playlist playlistv1
		var dish dish
		var cost float64

		playlistdishesDB, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, PlaylistDB.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return nil, err
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return nil, err
		}

		for _, playlistdishDB := range playlistdishesDB {
			dishDB, err := server.store.GetDish(ctx, playlistdishDB.DishID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errResponse(err))
					return nil, err
				}
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return nil, err
			}

			dish.ID = dishDB.ID
			dish.RestaurantID = dishDB.RestaurantID
			dish.IsAvailable = dishDB.IsAvailable
			dish.Name = dishDB.Name
			dish.Description = dishDB.Description.String
			dish.Price = dishDB.Price
			dish.Cuisine = dishDB.Cuisine.String
			dish.ImageURL = dishDB.ImageUrl.String
			dish.PlaylistDish = playlistDish{
				ID:           playlistdishDB.ID,
				PlaylistID:   playlistdishDB.PlaylistID,
				DishID:       playlistdishDB.DishID,
				DishQuantity: playlistdishDB.DishQuantity,
				CreatedAt:    playlistdishDB.CreatedAt,
				UpdatedAt:    playlistdishDB.AddedAt,
			}

			dishes = append(dishes, dish)
			cost += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		playlist.ID = PlaylistDB.ID
		playlist.Name = PlaylistDB.Name
		playlist.Description = PlaylistDB.Description.String
		playlist.ImageURL = PlaylistDB.ImageUrl.String
		playlist.IsPublic = PlaylistDB.IsPublic
		playlist.DeliveryDay = PlaylistDB.DeliveryDay.String
		playlist.Category = PlaylistDB.Category.String
		playlist.CreatedAt = PlaylistDB.CreatedAt
		playlist.UpdatedAt = PlaylistDB.AddedAt
		playlist.Dishes = dishes
		playlist.Cost = fmt.Sprintf("%.2f", cost)

		// Check if the category name already exists in the map
		var found bool

		if playlists, found = playlistMap[PlaylistDB.Category.String]; found {
			playlistMap[PlaylistDB.Category.String] = append(playlists, playlist)
		} else {
			playlistMap[PlaylistDB.Category.String] = []playlistv1{playlist}
		}
	}

	for categoryTitle, playlist := range playlistMap {
		categoryplaylist.CategoryTitle = categoryTitle
		categoryplaylist.Playlist = playlist

		categoryPlaylists = append(categoryPlaylists, categoryplaylist)
	}

	return categoryPlaylists, nil
}

type currentPlaylist struct {
	ID                   int64                 `form:"id" json:"id"`
	Name                 string                `form:"name" json:"name"`
	DeliveryDay          string                `form:"deliveryDay" json:"deliveryDay"`
	DeliveryTime         string                `form:"deliveryTime" json:"deliveryTime"`
	IsPublic             bool                  `form:"isPublic" json:"isPublic"`
	Restuarant_FoodItems []restaurant_foodItem `form:"foodItems" json:"foodItems"`
	Cost                 string                `form:"cost" json:"cost"`
}

type restaurant_foodItem struct {
	RestaurantName string     `form:"restaurantName" json:"restaurantName"`
	FoodItems      []foodItem `form:"foodItems" json:"foodItems"`
}

type foodItem struct {
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Quantity    int64   `form:"quantity" json:"quantity"`
	Price       float64 `form:"price" json:"price"`
	ImageURL    string  `form:"imageUrl" json:"imageUrl"`
	DishID      int64   `form:"dishId" json:"dishId"`
}

func (server *Server) maptoModelFoodItems(ctx *gin.Context, playlistDishesDB []db.PlaylistDish) ([]restaurant_foodItem, float64, error) {

	var restaurant_fooditem restaurant_foodItem
	var foodItems []foodItem
	var restaurant_foodItems []restaurant_foodItem
	var playlistDishesMap = make(map[string][]foodItem)
	var cost float64

	// Loop through the playlist_dishes and create the map
	for _, playlistDishDB := range playlistDishesDB {
		var fooditem foodItem

		restaurantName, err := server.store.ListRestaurantNameByDishID(ctx, playlistDishDB.DishID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return nil, 0, err
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return nil, 0, err
		}

		// Map database model to JSON response
		dish, err := server.store.GetDish(ctx, playlistDishDB.DishID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return nil, 0, err
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return nil, 0, err
		}

		fooditem.Name = dish.Name
		fooditem.Description = dish.Description.String
		fooditem.Quantity = playlistDishDB.DishQuantity
		fooditem.Price = dish.Price
		fooditem.ImageURL = dish.ImageUrl.String
		fooditem.DishID = dish.ID
		cost += dish.Price * float64(playlistDishDB.DishQuantity)

		// Check if the restaurant name already exists in the map
		var found bool
		if foodItems, found = playlistDishesMap[restaurantName]; found {
			playlistDishesMap[restaurantName] = append(foodItems, fooditem)
		} else {
			playlistDishesMap[restaurantName] = []foodItem{fooditem}
		}

	}

	for restaurantName, dishes := range playlistDishesMap {

		restaurant_fooditem.RestaurantName = restaurantName
		restaurant_fooditem.FoodItems = dishes

		restaurant_foodItems = append(restaurant_foodItems, restaurant_fooditem)
	}

	return restaurant_foodItems, cost, nil
}
