package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
)

type playlistLatestResponse struct {
	PublicPlaylist []publicPlaylist `form:"publicPlaylist" json:"publicPlaylist"`
	UserPlaylist   []userPlaylist   `form:"userPlaylist" json:"userPlaylist"`
}
type playlistDish struct {
	ID           int64     `form:"id" json:"id"`
	PlaylistID   int64     `form:"playlistId" json:"playlistId"`
	DishID       int64     `form:"dishId" json:"dishId"`
	DishQuantity int64     `form:"dishQuantity" json:"dishQuantity"`
	CreatedAt    time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `form:"updatedAt" json:"updatedAt"`
}
type playlistList struct {
	ID          int64     `form:"id" json:"id"`
	Name        string    `form:"name" json:"name"`
	Description string    `form:"description" json:"description"`
	ImageURL    string    `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool      `form:"isPublic" json:"isPublic"`
	DeliveryDay string    `form:"deliveryDay" json:"deliveryDay"`
	Category    string    `form:"category" json:"category"`
	CreatedAt   time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `form:"updatedAt" json:"updatedAt"`
	Dishes      []dishes  `form:"dishes" json:"dishes"`
	Cost        string    `form:"cost" json:"cost"`
}
type publicPlaylist struct {
	CategoryTitle string         `form:"categoryTitle" json:"categoryTitle"`
	PlaylistList  []playlistList `form:"list" json:"list"`
}
type userPlaylist struct {
	ID          int64     `form:"id" json:"id"`
	Name        string    `form:"name" json:"name"`
	Description string    `form:"description" json:"description"`
	ImageURL    string    `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool      `form:"isPublic" json:"isPublic"`
	DeliveryDay string    `form:"deliveryDay" json:"deliveryDay"`
	Category    string    `form:"category" json:"category"`
	CreatedAt   time.Time `form:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `form:"updatedAt" json:"updatedAt"`
	Dishes      []dishes  `form:"dishes" json:"dishes"`
	Cost        string    `form:"cost" json:"cost"`
}

type getPlaylistLatestRequest struct {
	UserID         int64 `form:"user_id" binding:"required,min=0"`
	PublicPageID   int64 `form:"public_page_id" binding:"min=1"`
	PublicPageSize int64 `form:"public_page_size" binding:"min=5,max=100"`
	UserPageID     int64 `form:"user_page_id" binding:"min=1"`
	UserPageSize   int64 `form:"user_page_size" binding:"min=5,max=100"`
}

func (server *Server) getPlaylistLatest(ctx *gin.Context) {

	var req getPlaylistLatestRequest
	var resp playlistLatestResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	argPublic := db.ListPlaylistPublicAndCategoryParams{
		Limit:  int32((req.PublicPageSize)),
		Offset: int32((req.PublicPageID - 1) * req.PublicPageSize),
	}

	// build publicPlaylist
	var publicplaylistlist []publicPlaylist
	var publicplaylist publicPlaylist
	var publicplaylistMap = make(map[string][]playlistList)
	var playlistlists []playlistList
	var playlistlist playlistList
	var dishesSlice []dishes
	var dish dishes
	var costPublic float64

	//Loop through the playlist and create a map
	publicPlaylistsDB, err := server.store.ListPlaylistPublicAndCategory(ctx, argPublic)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	for _, publicPlaylistDB := range publicPlaylistsDB {

		playlistdishesDB, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, publicPlaylistDB.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}

		for _, playlistdishDB := range playlistdishesDB {
			dishDB, err := server.store.GetDish(ctx, playlistdishDB.DishID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
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

			dishesSlice = append(dishesSlice, dish)
			costPublic += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		playlistlist.ID = publicPlaylistDB.ID
		playlistlist.Name = publicPlaylistDB.Name
		playlistlist.Description = publicPlaylistDB.Description.String
		playlistlist.ImageURL = publicPlaylistDB.ImageUrl.String
		playlistlist.IsPublic = publicPlaylistDB.IsPublic
		playlistlist.DeliveryDay = publicPlaylistDB.DeliveryDay.String
		playlistlist.Category = publicPlaylistDB.Category.String
		playlistlist.CreatedAt = publicPlaylistDB.CreatedAt
		playlistlist.UpdatedAt = publicPlaylistDB.AddedAt
		playlistlist.Dishes = dishesSlice
		playlistlist.Cost = fmt.Sprintf("%.2f", costPublic)

		// Check if the category name already exists in the map
		var found bool
		if playlistlists, found = publicplaylistMap[publicPlaylistDB.Category.String]; found {
			publicplaylistMap[publicPlaylistDB.Category.String] = append(playlistlists, playlistlist)
		} else {
			publicplaylistMap[publicPlaylistDB.Category.String] = []playlistList{playlistlist}
		}

	}

	for categoryTitle, playlistList := range publicplaylistMap {
		publicplaylist.CategoryTitle = categoryTitle
		publicplaylist.PlaylistList = playlistList

		publicplaylistlist = append(publicplaylistlist, publicplaylist)
	}

	//build userPlaylist
	var userplaylistlist []userPlaylist
	var userplaylist userPlaylist
	var dishesSliceUser []dishes
	var dishUser dishes
	var costUser float64

	argUser := db.ListPlaylistsByUserIDParams{
		UserID: req.UserID,
		Limit:  int32((req.UserPageSize)),
		Offset: int32((req.UserPageID - 1) * req.UserPageSize),
	}

	//Loop through the playlist and create a map
	userPlaylistsDB, err := server.store.ListPlaylistsByUserID(ctx, argUser)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	for _, userPlaylistDB := range userPlaylistsDB {

		playlistdishesDB, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, userPlaylistDB.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}

		for _, playlistdishDB := range playlistdishesDB {
			dishDB, err := server.store.GetDish(ctx, playlistdishDB.DishID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
			}

			dishUser.ID = dishDB.ID
			dishUser.RestaurantID = dishDB.RestaurantID
			dishUser.IsAvailable = dishDB.IsAvailable
			dishUser.Name = dishDB.Name
			dishUser.Description = dishDB.Description.String
			dishUser.Price = dishDB.Price
			dishUser.Cuisine = dishDB.Cuisine.String
			dishUser.ImageURL = dishDB.ImageUrl.String
			dishUser.PlaylistDish = playlistDish{
				ID:           playlistdishDB.ID,
				PlaylistID:   playlistdishDB.PlaylistID,
				DishID:       playlistdishDB.DishID,
				DishQuantity: playlistdishDB.DishQuantity,
				CreatedAt:    playlistdishDB.CreatedAt,
				UpdatedAt:    playlistdishDB.AddedAt,
			}

			dishesSliceUser = append(dishesSliceUser, dishUser)
			costUser += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		userplaylist.ID = userPlaylistDB.ID
		userplaylist.Name = userPlaylistDB.Name
		userplaylist.Description = userPlaylistDB.Description.String
		userplaylist.ImageURL = userPlaylistDB.ImageUrl.String
		userplaylist.IsPublic = userPlaylistDB.IsPublic
		userplaylist.DeliveryDay = userPlaylistDB.DeliveryDay.String
		userplaylist.Category = userPlaylistDB.Category.String
		userplaylist.CreatedAt = userPlaylistDB.CreatedAt
		userplaylist.UpdatedAt = userPlaylistDB.AddedAt
		userplaylist.Dishes = dishesSliceUser
		userplaylist.Cost = fmt.Sprintf("%.2f", costUser)

		userplaylistlist = append(userplaylistlist, userplaylist)

	}

	//build response
	resp = playlistLatestResponse{
		PublicPlaylist: publicplaylistlist,
		UserPlaylist:   userplaylistlist,
	}

	ctx.JSON(http.StatusOK, resp)
}
