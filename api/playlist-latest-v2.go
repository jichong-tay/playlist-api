package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getPlaylistLatestV2Request struct {
	UserID int64 `form:"user_id" binding:"required,min=0"`
}

func (server *Server) getPlaylistLatestV2(ctx *gin.Context) {

	var req getPlaylistLatestV2Request
	var resp playlistLatestResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
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
