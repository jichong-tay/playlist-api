package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type playlistLatestResponsev2 struct {
	PublicPlaylist []publicplaylistv2 `form:"publicPlaylist" json:"publicPlaylist"`
	UserPlaylist   []userPlaylistv2   `form:"userPlaylist" json:"userPlaylist"`
}

type getPlaylistLatestV2Request struct {
	UserID int64 `form:"user_id" binding:"required,min=0"`
}

type publicplaylistv2 struct {
	ID          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	ImageURL    string `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool   `form:"isPublic" json:"isPublic"`
	DeliveryDay string `form:"deliveryDay" json:"deliveryDay"`
	Category    string `form:"category" json:"category"`
	Cost        string `form:"cost" json:"cost"`
}

type userPlaylistv2 struct {
	ID          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	ImageURL    string `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool   `form:"isPublic" json:"isPublic"`
	DeliveryDay string `form:"deliveryDay" json:"deliveryDay"`
	Category    string `form:"category" json:"category"`
	Cost        string `form:"cost" json:"cost"`
}

func (server *Server) getPlaylistLatestV2(ctx *gin.Context) {

	var req getPlaylistLatestV2Request
	var resp playlistLatestResponsev2

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// build publicPlaylist
	var publicplaylists []publicplaylistv2
	var publicplaylist publicplaylistv2
	// var dishesSlice []dishes // TODO Clean Up
	var dish dishes
	var costPublic float64

	//Loop through the playlist and create a map
	publicPlaylistsDB, err := server.store.ListPlaylistPublicAndCategoryAll(ctx)
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

			// dishesSlice = append(dishesSlice, dish) // TODO Clean Up
			costPublic += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		publicplaylist.ID = publicPlaylistDB.ID
		publicplaylist.Name = publicPlaylistDB.Name
		publicplaylist.ImageURL = publicPlaylistDB.ImageUrl.String
		publicplaylist.IsPublic = publicPlaylistDB.IsPublic
		publicplaylist.DeliveryDay = publicPlaylistDB.DeliveryDay.String
		publicplaylist.Category = publicPlaylistDB.Category.String
		publicplaylist.Cost = fmt.Sprintf("%.2f", costPublic)

		publicplaylists = append(publicplaylists, publicplaylist)
	}

	//build userPlaylist
	var userplaylistlist []userPlaylistv2
	var userplaylist userPlaylistv2
	// var dishesSliceUser []dishes // TODO Clean Up
	var dishUser dishes
	var costUser float64

	argUser := req.UserID

	//Loop through the playlist and create a map
	userPlaylistsDB, err := server.store.ListPlaylistsByUserIDAll(ctx, argUser)
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

			// dishesSliceUser = append(dishesSliceUser, dishUser) // TODO Clean Up
			costUser += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		userplaylist.ID = userPlaylistDB.ID
		userplaylist.Name = userPlaylistDB.Name
		userplaylist.Description = userPlaylistDB.Description.String
		userplaylist.ImageURL = userPlaylistDB.ImageUrl.String
		userplaylist.IsPublic = userPlaylistDB.IsPublic
		userplaylist.DeliveryDay = userPlaylistDB.DeliveryDay.String
		userplaylist.Category = userPlaylistDB.Category.String
		userplaylist.Cost = fmt.Sprintf("%.2f", costUser)

		userplaylistlist = append(userplaylistlist, userplaylist)

	}

	//build response
	resp = playlistLatestResponsev2{
		PublicPlaylist: publicplaylists,
		UserPlaylist:   userplaylistlist,
	}

	ctx.JSON(http.StatusOK, resp)
}
