package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type playlistCategoryResponseAll struct {
	PlaylistList []playlistList `form:"list" json:"list"`
}

func (server *Server) getPublicPlaylist(ctx *gin.Context) {
	var resp playlistCategoryResponseAll

	var playlistlist playlistList
	var playlistlists []playlistList
	var dishesSlice []dishes
	var dish dishes
	var costCat float64

	categoryPlaylistsDB, err := server.store.ListPlaylistPublicAndCategoryAll(ctx)
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

			dishesSlice = append(dishesSlice, dish)
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
		playlistlist.Dishes = dishesSlice
		playlistlist.Cost = fmt.Sprintf("%.2f", costCat)

		playlistlists = append(playlistlists, playlistlist)

	}

	resp.PlaylistList = playlistlists

	ctx.JSON(http.StatusOK, resp)
}
