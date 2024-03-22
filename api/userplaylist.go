package api

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goombaio/namegenerator"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"github.com/jichong-tay/playlist-api/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/guregu/null.v4"
)

type getUserURI struct {
	UserID int64 `uri:"userid" binding:"required,min=0"`
}

func (server *Server) getUserPlaylist(ctx *gin.Context) {
	var req getUserURI
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
			UserID:       req.UserID,
			PlaylistID:   req.PlaylistID,
			Status:       null.NewString("Cancelled", true),
			DeliveryDay:  null.NewString("", true),
			DeliveryTime: null.NewTime(time.Time{}, false),
		}
	}

	if ctx.Request.Method == "POST" {
		arg = db.UpdateStatusForUser_PlaylistParams{
			UserID:       req.UserID,
			PlaylistID:   req.PlaylistID,
			Status:       null.NewString("Pending", true),
			DeliveryDay:  null.NewString("", true),
			DeliveryTime: null.NewTime(time.Time{}, false),
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
	var reqUser getUserURI
	var reqPlaylist currentPlaylist

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// random name generator for playlist
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	// format playlist name
	caser := cases.Title(language.English)
	playlistName := caser.String(fmt.Sprint(name))
	playlistNameDesc := caser.String(fmt.Sprint("Playlist created from ", reqPlaylist.Name))

	// convert string to time
	const timeFormat = "15:04"
	deliveryTime, _ := time.Parse(timeFormat, reqPlaylist.DeliveryTime)

	// random image url
	imageURL, _ := util.RandomImageURL(400, 400)

	arg := db.PlaylistDishTxParams{
		Name:         playlistName,
		Description:  null.NewString(playlistNameDesc, true),
		ImageURL:     null.NewString(imageURL, true),
		IsPublic:     false,
		DeliveryDay:  null.NewString(reqPlaylist.DeliveryDay, true),
		Category:     null.NewString("", true),
		UserID:       reqUser.UserID,
		PlaylistID:   reqPlaylist.ID,
		DeliveryTime: null.NewTime(deliveryTime, true),
		Status:       null.NewString("Pending", true),
		DishItems:    server.maptoModelFoodItemV2(reqPlaylist.RestaurantFoodItems),
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
		Value: fmt.Sprintf("%d", userPlaylistsDB.ID),
	})

	server.getPlaylistCurrent(ctx)
}

func (server *Server) updateUserPlaylist(ctx *gin.Context) {
	var reqUser getUserURI
	var reqPlaylist currentPlaylist

	if err := ctx.ShouldBindUri(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqPlaylist); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	const timeFormat = "15:04" // convert string to time
	deliveryTime, _ := time.Parse(timeFormat, reqPlaylist.DeliveryTime)

	playlist, _ := server.store.GetPlaylist(ctx, reqPlaylist.ID) // get the playlist from the database

	arg := db.PlaylistDishTxParams{
		Name:         reqPlaylist.Name,
		Description:  playlist.Description,
		ImageURL:     playlist.ImageUrl,
		IsPublic:     false,
		DeliveryDay:  null.NewString(reqPlaylist.DeliveryDay, true),
		Category:     playlist.Category,
		UserID:       reqUser.UserID,
		PlaylistID:   reqPlaylist.ID,
		DeliveryTime: null.NewTime(deliveryTime, true),
		Status:       null.NewString("Pending", true),
		DishItems:    server.maptoModelFoodItemV2(reqPlaylist.RestaurantFoodItems),
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

type getPlaylistRandomURI struct {
	Cuisine string  `uri:"cuisine" binding:"required"`
	Num     int     `uri:"num" binding:"required"`
	Budget  float64 `uri:"budget" binding:"required"`
}

func (server *Server) getPlaylistRandom(ctx *gin.Context) {
	var req getPlaylistRandomURI
	// var playlist currentPlaylist
	// var foodItems []foodItem
	empty := make([]string, 0)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	searchArg := sql.NullString{String: strings.ToLower(req.Cuisine), Valid: true}

	dishesDB, err := server.store.ListDishesByCuisine(ctx, searchArg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	selectedDishesDB, err := randomSelectDishesv2(dishesDB, req.Num, req.Budget/(float64(req.Num)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, empty) // to return empty array
		return
	}
	foodItems, cost, err := server.maptoModelDishes(ctx, selectedDishesDB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := currentPlaylist{
		Name: "Random Playlist",
		// IsPublic:             false,
		RestaurantFoodItems: foodItems,
		Cost:                cost,
	}

	ctx.JSON(http.StatusOK, resp)
}

// // randomSelectDishesv1 selects dishes randomly from the list of dishes and will always return
// func randomSelectDishesv1(dishes []db.Dish, count int, price float64) ([]db.Dish, error) {

// 	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	selectedDishes := make([]db.Dish, count)
// 	selectedIndices := make(map[int]bool)

// 	tryCount := 10 //try 10 times to get the dish within the price range
// 	j := 0
// 	if count > len(dishes) { //check if there are enough dishes else return error
// 		return selectedDishes, fmt.Errorf("not enough dishes for selection")
// 	}
// 	for i := 0; i < count; {
// 		randomIndex := randGenerator.Intn(len(dishes))
// 		if !selectedIndices[randomIndex] {
// 			j++
// 			selectedDishes[i] = dishes[randomIndex]
// 			if (selectedDishes[i].Price >= price-5 && selectedDishes[i].Price <= price+5) || j > tryCount {
// 				selectedIndices[randomIndex] = true
// 				i++
// 			}
// 		}
// 	}

// 	return selectedDishes, error(nil)
// }

// randomSelectDishesv2 selects dishes randomly from the list of dishes and will NOT always return
func randomSelectDishesv2(dishes []db.Dish, count int, price float64) ([]db.Dish, error) {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	selectedDishes := make([]db.Dish, count)
	selectedIndices := make(map[int]bool)
	tryCount := 50 // try x times to get the dish within the price range
	j := 0
	if count > len(dishes) { // check if there are enough dishes else return error
		return selectedDishes, fmt.Errorf("not enough dishes for selection")
	}
	for i := 0; i < count; {
		randomIndex := randGenerator.Intn(len(dishes))
		if !selectedIndices[randomIndex] {
			if j > tryCount {
				return selectedDishes, fmt.Errorf("no dishes found")
			}
			j++
			selectedDishes[i] = dishes[randomIndex]
			if selectedDishes[i].Price >= price*0.95 && selectedDishes[i].Price <= price*1.1 {
				selectedIndices[randomIndex] = true
				i++
			}
		}
	}

	return selectedDishes, error(nil)
}
