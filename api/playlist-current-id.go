package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO Clean Up
// type getPlaylistCurrentRequestJSON struct {
// 	ID          int64                     `form:"id" json:"id"`
// 	Name        string                    `form:"name" json:"name"`
// 	DeliveryDay string                    `form:"deliveryDay" json:"deliveryDay"`
// 	IsPublic    bool                      `form:"isPublic" json:"isPublic"`
// 	FoodItems   []playlist_dishesMainJSON `form:"foodItems" json:"foodItems"`
// 	Cost        string                    `form:"cost" json:"cost"`
// }

type getPlaylistByIDRequest struct {
	ID int64 `uri:"playlistid" binding:"required,min=1"`
}

func (server *Server) getPlaylistCurrent(ctx *gin.Context) {
	var req getPlaylistByIDRequest

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

	playlistDishes, err := server.store.ListPlaylist_DishesByPlaylistID(ctx, playlist.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	restuarant_foodItems, cost, _ := server.maptoModelFoodItems(ctx, playlistDishes)

	resp := getPlaylistCurrentResponse{
		ID:                   playlist.ID,
		Name:                 playlist.Name,
		DeliveryDay:          playlist.DeliveryDay.String,
		IsPublic:             playlist.IsPublic,
		Restuarant_FoodItems: restuarant_foodItems,
		Cost:                 fmt.Sprintf("%.2f", cost),
	}

	ctx.JSON(http.StatusOK, resp)

}
