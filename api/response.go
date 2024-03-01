package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
)

type dishes struct {
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

type playlistLatestResponsev2 struct {
	PublicPlaylist []playlistv2 `form:"publicPlaylist" json:"publicPlaylist"`
	UserPlaylist   []playlistv2 `form:"userPlaylist" json:"userPlaylist"`
}

type playlistv2 struct {
	ID          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	ImageURL    string `form:"imageUrl" json:"imageUrl"`
	IsPublic    bool   `form:"isPublic" json:"isPublic"`
	DeliveryDay string `form:"deliveryDay" json:"deliveryDay"`
	Category    string `form:"category" json:"category"`
	Cost        string `form:"cost" json:"cost"`
}

func (server *Server) maptoModelV2(ctx *gin.Context, playlistsDB []db.Playlist) ([]playlistv2, error) {
	var playlists []playlistv2
	var playlist playlistv2
	var dish dishes
	var cost float64

	for _, playlistDB := range playlistsDB {

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

		playlists = append(playlists, playlist)
	}
	return playlists, nil
}
