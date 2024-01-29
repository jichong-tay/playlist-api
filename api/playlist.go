package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type createPlaylistRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	IsPublic    bool   `json:"is_public"`
	DeliveryDay string `json:"delivery_day"`
	Category    string `json:"category"`
}

func (server *Server) createPlaylist(ctx *gin.Context) {
	var req createPlaylistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreatePlaylistParams{
		Name:        req.Name,
		Description: null.NewString(req.Description, true),
		ImageUrl:    null.NewString(req.ImageUrl, true),
		IsPublic:    req.IsPublic,
		DeliveryDay: null.NewString(req.DeliveryDay, true),
		Category:    null.NewString(req.Category, true),
	}
	playlist, err := server.store.CreatePlaylist(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

type getPlaylistRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPlaylist(ctx *gin.Context) {
	var req getPlaylistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	playlist, err := server.store.GetPlaylist(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}

type listPlaylistRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPlaylist(ctx *gin.Context) {
	var req listPlaylistRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListPlaylistsParams{
		Limit:  int32((req.PageSize)),
		Offset: int32((req.PageID - 1) * req.PageSize),
	}

	playlists, err := server.store.ListPlaylists(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlists)
}

type playlist_dishesJSON struct {
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Quantity    int64   `form:"quantity" json:"quantity"`
	Price       float64 `form:"price" json:"price"`
	ImageURL    string  `form:"imageUrl" json:"imageUrl"`
	DishID      int64   `form:"dishId" json:"dishId"`
}

type playlist_dishesMainJSON struct {
	RestaurantName string                `form:"restaurantName" json:"restaurantName"`
	FoodItems      []playlist_dishesJSON `form:"foodItems" json:"foodItems"`
}

type getPlaylistCurrentRequestJSON struct {
	ID          int64                     `form:"id" json:"id"`
	Name        string                    `form:"name" json:"name"`
	DeliveryDay string                    `form:"deliveryDay" json:"deliveryDay"`
	IsPublic    bool                      `form:"isPublic" json:"isPublic"`
	FoodItems   []playlist_dishesMainJSON `form:"foodItems" json:"foodItems"`
	Cost        string                    `form:"cost" json:"cost"`
}

type getPlaylistCurrentResponseJSON struct {
	ID          int64                     `form:"id" json:"id"`
	Name        string                    `form:"name" json:"name"`
	DeliveryDay string                    `form:"deliveryDay" json:"deliveryDay"`
	IsPublic    bool                      `form:"isPublic" json:"isPublic"`
	FoodItems   []playlist_dishesMainJSON `form:"foodItems" json:"foodItems"`
	Cost        string                    `form:"cost" json:"cost"`
}

type getPlaylistByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPlaylistCurrent(ctx *gin.Context) {
	var req getPlaylistByIDRequest
	var resp getPlaylistCurrentResponseJSON

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	playlist, err := server.store.GetPlaylist(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	fmt.Println(playlist)

	playlistDishes, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, playlist.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// Initialize response structs
	var playlistDishesJSON playlist_dishesJSON
	var playlistDishesMainJSON playlist_dishesMainJSON
	var foodItems []playlist_dishesJSON
	var foodItemsMain []playlist_dishesMainJSON
	var playlistDishesMap = make(map[string][]playlist_dishesJSON)
	var cost float64

	// Loop through the playlist_dishes and create the map
	for _, pd := range playlistDishes {
		restaurantName, err := server.store.ListRestaurantNameByDishID(ctx, pd.DishID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}

		// Map database model to JSON response
		dish, err := server.store.GetDish(ctx, pd.DishID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}

		playlistDishesJSON.Name = dish.Name
		playlistDishesJSON.Description = dish.Description.String
		playlistDishesJSON.Quantity = pd.DishQuantity
		playlistDishesJSON.Price = dish.Price
		playlistDishesJSON.ImageURL = dish.ImageUrl.String
		playlistDishesJSON.DishID = dish.ID
		cost += dish.Price * float64(pd.DishQuantity)

		// Check if the restaurant name already exists in the map
		var found bool
		if foodItems, found = playlistDishesMap[restaurantName]; found {
			playlistDishesMap[restaurantName] = append(foodItems, playlistDishesJSON)
		} else {
			playlistDishesMap[restaurantName] = []playlist_dishesJSON{playlistDishesJSON}
		}

	}

	for restaurantName, dishes := range playlistDishesMap {

		playlistDishesMainJSON.RestaurantName = restaurantName
		playlistDishesMainJSON.FoodItems = dishes

		foodItemsMain = append(foodItemsMain, playlistDishesMainJSON)
	}

	resp = getPlaylistCurrentResponseJSON{
		ID:          playlist.ID,
		Name:        playlist.Name,
		DeliveryDay: playlist.DeliveryDay.String,
		IsPublic:    playlist.IsPublic,
		FoodItems:   foodItemsMain,
		Cost:        fmt.Sprintf("%.2f", cost),
	}

	ctx.JSON(http.StatusOK, resp)

}
