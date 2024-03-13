package api

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"gopkg.in/guregu/null.v4"
)

type getUserUri struct {
	UserID int64 `uri:"userid" binding:"required,min=0"`
}

func (server *Server) getUserPlaylist(ctx *gin.Context) {

	var req getUserUri
	var playlists []playlistv2

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	argUser := req.UserID

	userPlaylistsDB, err := server.store.ListPlaylistsByUserIDAll(ctx, argUser)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	playlists, _ = server.maptoModelV2(ctx, userPlaylistsDB)

	resp := playlists

	ctx.JSON(http.StatusOK, resp)
}

type updateUserPlaylistRequest struct {
	UserID     int64 `uri:"userid" binding:"required,min=0"`
	PlaylistID int64 `uri:"playlistid" binding:"required,min=0"`
}

func (server *Server) updateUserPlaylistStatus(ctx *gin.Context) {

	var req updateUserPlaylistRequest
	var arg db.UpdateStatusForUser_PlaylistParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if ctx.Request.Method == "PUT" {
		arg = db.UpdateStatusForUser_PlaylistParams{
			UserID:     req.UserID,
			PlaylistID: req.PlaylistID,
			Status:     null.NewString("Cancelled", true),
		}
	}

	if ctx.Request.Method == "POST" {
		arg = db.UpdateStatusForUser_PlaylistParams{
			UserID:     req.UserID,
			PlaylistID: req.PlaylistID,
			Status:     null.NewString("Pending", true),
		}
	}

	playlist, err := server.store.UpdateStatusForUser_Playlist(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	if len(playlist) == 0 {
		ctx.JSON(http.StatusNotFound, errResponse(sql.ErrNoRows))
		return
	}

	userPlaylistsDB, err := server.store.ListPlaylistsByUserIDAll(ctx, req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	playlists, _ := server.maptoModelV2(ctx, userPlaylistsDB)

	resp := playlists

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) createUserPlaylist(ctx *gin.Context) {

	var reqUser getUserUri
	var reqPlaylist currentPlaylist

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	//convert string to time
	const timeFormat = "15:04"
	deliveryTime, _ := time.Parse(timeFormat, reqPlaylist.DeliveryTime)

	arg := db.PlaylistDishTxParams{
		Name:         reqPlaylist.Name,
		Description:  null.NewString("", true),
		ImageUrl:     null.NewString("", true),
		IsPublic:     false,
		DeliveryDay:  null.NewString(reqPlaylist.DeliveryDay, true),
		Category:     null.NewString("", true),
		UserID:       reqUser.UserID,
		PlaylistID:   reqPlaylist.ID,
		DeliveryTime: null.NewTime(deliveryTime, true),
		Status:       null.NewString("Pending", true),
		DishItems:    server.maptoModelFoodItemV2(reqPlaylist.Restuarant_FoodItems),
	}

	userPlaylistsDB, err := server.store.CreatePlaylistDishTx(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "playlistid",
		Value: fmt.Sprintf("%d", userPlaylistsDB.ID)})

	server.getPlaylistCurrent(ctx)

}

func (server *Server) updateUserPlaylist(ctx *gin.Context) {

	var reqUser getUserUri
	var reqPlaylist currentPlaylist

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	//convert string to time
	const timeFormat = "15:04"
	deliveryTime, _ := time.Parse(timeFormat, reqPlaylist.DeliveryTime)

	arg := db.PlaylistDishTxParams{
		Name:         reqPlaylist.Name,
		Description:  null.NewString("", true),
		ImageUrl:     null.NewString("", true),
		IsPublic:     false,
		DeliveryDay:  null.NewString(reqPlaylist.DeliveryDay, true),
		Category:     null.NewString("", true),
		UserID:       reqUser.UserID,
		PlaylistID:   reqPlaylist.ID,
		DeliveryTime: null.NewTime(deliveryTime, true),
		Status:       null.NewString("Pending", true),
		DishItems:    server.maptoModelFoodItemV2(reqPlaylist.Restuarant_FoodItems),
	}

	userPlaylistsDB, err := server.store.UpdatePlaylistDishTx(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "playlistid",
		Value: fmt.Sprintf("%d", userPlaylistsDB.ID),
	})

	server.getPlaylistCurrent(ctx)

}

type getPlaylistRandomUri struct {
	Cuisine string  `uri:"cuisine" binding:"required"`
	Num     int     `uri:"num" binding:"required"`
	Budget  float64 `uri:"budget" binding:"required"`
}

func (server *Server) getPlaylistRandom(ctx *gin.Context) {

	var req getPlaylistRandomUri
	// var playlist currentPlaylist
	// var foodItems []foodItem

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	searchArg := sql.NullString{
		String: strings.ToLower(req.Cuisine),
		Valid:  true}

	dishesDB, err := server.store.ListDishesByCuisine(ctx, searchArg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	selectedDishesDB := randomSelectDishes(dishesDB, req.Num, req.Budget/(float64(req.Num)))

	foodItems, cost, err := server.maptoModelDishes(ctx, selectedDishesDB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := currentPlaylist{
		Name: "Random Playlist",
		// IsPublic:             false,
		Restuarant_FoodItems: foodItems,
		Cost:                 cost,
	}

	ctx.JSON(http.StatusOK, resp)
}

func randomSelectDishes(dishes []db.Dish, count int, price float64) []db.Dish {

	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	selectedDishes := make([]db.Dish, count)
	selectedIndices := make(map[int]bool)

	tryCount := 10 //try 10 times to get the dish within the price range
	j := 0
	for i := 0; i < count; {
		randomIndex := randGenerator.Intn(len(dishes))
		if !selectedIndices[randomIndex] {
			j++
			if (selectedDishes[i].Price >= price-5 && selectedDishes[i].Price <= price+5) || j > tryCount {
				selectedDishes[i] = dishes[randomIndex]
				selectedIndices[randomIndex] = true
				i++
			}
		}
	}

	return selectedDishes
}
