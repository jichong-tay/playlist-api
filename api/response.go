package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

// type playlistResponsev1 struct {
// 	Playlist []playlistv1 `form:"list" json:"list"`
// }

// type playlistResponsev2 struct {
// 	Playlist []playlistv2 `form:"list" json:"list"`
// }

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
	var playlists []playlistv1
	var dishes []dish
	playlistMap := make(map[string][]playlistv1)

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
	ID                  int64                `form:"id" json:"id"`
	Name                string               `form:"name" json:"name"`
	DeliveryDay         string               `form:"deliveryDay" json:"deliveryDay"`
	DeliveryTime        string               `form:"deliveryTime" json:"deliveryTime"`
	IsPublic            bool                 `form:"isPublic" json:"isPublic"`
	RestaurantFoodItems []restaurantFoodItem `form:"foodItems" json:"foodItems"`
	Cost                interface{}          `form:"cost" json:"cost"`
	Status              string               `form:"status" json:"status"`
}

type restaurantFoodItem struct {
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

func (server *Server) maptoModelFoodItems(ctx *gin.Context, playlistDishesDB []db.PlaylistDish) ([]restaurantFoodItem, float64, error) {
	var restaurantFooditem restaurantFoodItem
	var foodItems []foodItem
	var restaurantFoodItems []restaurantFoodItem
	var cost float64
	playlistDishesMap := make(map[string][]foodItem)

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

		restaurantFooditem.RestaurantName = restaurantName
		restaurantFooditem.FoodItems = dishes

		restaurantFoodItems = append(restaurantFoodItems, restaurantFooditem)
	}

	return restaurantFoodItems, cost, nil
}

type searchDish struct {
	DishID         int64   `json:"dishId"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	ImageURL       string  `json:"imageUrl"`
	RestaurantName string  `json:"restaurantName"`
	RestaurantID   int64   `json:"restaurantId"`
}

func (server *Server) maptoModelSearchDish(searchDishesRow []db.SearchDishesRow) []searchDish {
	var searchDishes []searchDish
	var searchDish searchDish

	for _, searchDishRow := range searchDishesRow {

		searchDish.DishID = searchDishRow.DishID
		searchDish.Name = searchDishRow.DishName
		searchDish.Description = searchDishRow.DishDescription.String
		searchDish.Price = searchDishRow.DishPrice
		searchDish.ImageURL = searchDishRow.DishImageurl.String
		searchDish.RestaurantName = searchDishRow.RestaurantName
		searchDish.RestaurantID = searchDishRow.RestaurantID

		searchDishes = append(searchDishes, searchDish)
	}

	return searchDishes
}

type searchResult struct {
	ID       int64  `json:"keyword_id"`
	Keywords string `json:"keyword"`
}

func (server *Server) maptoModelSearch(searches []db.Search) []searchResult {
	var result []searchResult
	caser := cases.Title(language.English)
	for _, s := range searches {
		searchResult := searchResult{
			ID:       s.ID,
			Keywords: caser.String(s.Keyword.String),
		}
		result = append(result, searchResult)
	}
	return result
}

func (server *Server) maptoModelFoodItemV2(restuarantFoodItems []restaurantFoodItem) []db.RestaurantFoodItem {
	var foodItemDB db.FoodItem
	var restuarantFoodItemsDB []db.RestaurantFoodItem
	var restuarantFoodItemDB db.RestaurantFoodItem

	for _, restuarantfooditem := range restuarantFoodItems {
		restuarantFoodItemDB.RestaurantName = restuarantfooditem.RestaurantName

		var foodItemsDB []db.FoodItem
		for _, fooditem := range restuarantfooditem.FoodItems {
			foodItemDB = db.FoodItem{
				Name:        fooditem.Name,
				Description: fooditem.Description,
				Quantity:    fooditem.Quantity,
				Price:       fooditem.Price,
				ImageURL:    fooditem.ImageURL,
				DishID:      fooditem.DishID,
			}
			foodItemsDB = append(foodItemsDB, foodItemDB)
		}

		restuarantFoodItemDB.FoodItems = foodItemsDB

		restuarantFoodItemsDB = append(restuarantFoodItemsDB, restuarantFoodItemDB)
	}

	return restuarantFoodItemsDB
}

func (server *Server) maptoModelDishes(ctx *gin.Context, dishesDB []db.Dish) ([]restaurantFoodItem, float64, error) {
	var restaurantFooditem restaurantFoodItem
	var foodItems []foodItem
	var restaurantFoodItems []restaurantFoodItem
	var cost float64
	playlistDishesMap := make(map[string][]foodItem)

	// Loop through the dishes and create the map
	for _, dishDB := range dishesDB {
		var fooditem foodItem

		restaurantName, err := server.store.ListRestaurantNameByDishID(ctx, dishDB.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return nil, 0, err
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return nil, 0, err
		}

		// Map database model to JSON response
		dish, err := server.store.GetDish(ctx, dishDB.ID)
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
		fooditem.Quantity = 1
		fooditem.Price = dish.Price
		fooditem.ImageURL = dish.ImageUrl.String
		fooditem.DishID = dish.ID
		cost += dish.Price * float64(1)

		// Check if the restaurant name already exists in the map
		var found bool
		if foodItems, found = playlistDishesMap[restaurantName]; found {
			playlistDishesMap[restaurantName] = append(foodItems, fooditem)
		} else {
			playlistDishesMap[restaurantName] = []foodItem{fooditem}
		}

	}

	for restaurantName, dishes := range playlistDishesMap {

		restaurantFooditem.RestaurantName = restaurantName
		restaurantFooditem.FoodItems = dishes

		restaurantFoodItems = append(restaurantFoodItems, restaurantFooditem)
	}

	return restaurantFoodItems, cost, nil
}
