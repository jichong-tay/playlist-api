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

	userPlaylist, _ := server.store.GetUserPlaylistByPlaylistID(ctx, req.ID)

	deliveryTime := userPlaylist.DeliveryTime.Time.Format("15:04") // TODO:userPlaylist.DeliveryTime.Time.Format("15:04:05")
	if deliveryTime == "00:00" {
		deliveryTime = ""
	}

	resp := currentPlaylist{
		ID:                   playlist.ID,
		Name:                 playlist.Name,
		DeliveryDay:          playlist.DeliveryDay.String,
		DeliveryTime:         deliveryTime,
		IsPublic:             playlist.IsPublic,
		Restuarant_FoodItems: restuarant_foodItems,
		Cost:                 fmt.Sprintf("%.2f", cost),
		Status:               userPlaylist.Status.String,
	}

	ctx.JSON(http.StatusOK, resp)

}
