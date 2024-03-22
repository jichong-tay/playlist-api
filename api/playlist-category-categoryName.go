package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type playlistCategoryResponse struct {
	CategoryTitle string       `form:"categoryTitle" json:"categoryTitle"`
	PlaylistList  []playlistv1 `form:"list" json:"list"`
}

type listPlaylistByCategory struct {
	Category string `form:"category" binding:"required"`
}

func (server *Server) getPlaylistCategory(ctx *gin.Context) {
	var req listPlaylistByCategory
	var resp playlistCategoryResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var playlistlist playlistv1
	var playlistlists []playlistv1
	var dishes []dish
	var dish dish
	var costCat float64

	categoryplaylistMap := make(map[string][]playlistv1)

	arg := req.Category

	categoryPlaylistsDB, err := server.store.ListPlaylistByCategory(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	for _, PlaylistDB := range categoryPlaylistsDB {

		playlistdishesDB, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, PlaylistDB.ID)
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

			dishes = append(dishes, dish)
			costCat += dishDB.Price * float64(playlistdishDB.DishQuantity)
		}

		playlistlist.ID = PlaylistDB.ID
		playlistlist.Name = PlaylistDB.Name
		playlistlist.Description = PlaylistDB.Description.String
		playlistlist.ImageURL = PlaylistDB.ImageUrl.String
		playlistlist.IsPublic = PlaylistDB.IsPublic
		playlistlist.DeliveryDay = PlaylistDB.DeliveryDay.String
		playlistlist.Category = PlaylistDB.Category.String
		playlistlist.CreatedAt = PlaylistDB.CreatedAt
		playlistlist.UpdatedAt = PlaylistDB.AddedAt
		playlistlist.Dishes = dishes
		playlistlist.Cost = fmt.Sprintf("%.2f", costCat)

		// Check if the category name already exists in the map
		var found bool
		if playlistlists, found = categoryplaylistMap[PlaylistDB.Category.String]; found {
			categoryplaylistMap[PlaylistDB.Category.String] = append(playlistlists, playlistlist)
		} else {
			categoryplaylistMap[PlaylistDB.Category.String] = []playlistv1{playlistlist}
		}

	}

	for categoryTitle, playlistList := range categoryplaylistMap {
		resp.CategoryTitle = categoryTitle
		resp.PlaylistList = playlistList
	}

	ctx.JSON(http.StatusOK, resp)
}
