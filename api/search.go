package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"gopkg.in/guregu/null.v4"
)

type searchRequest struct {
	UserID        int64  `uri:"userid" binding:"required,min=0"`
	SearchKeyword string `uri:"search"`
}

func (server *Server) searchDishes(ctx *gin.Context) {

	var req searchRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	searchArg := db.CreateSearchParams{
		UserID:  req.UserID,
		Keyword: null.NewString(strings.ToLower(req.SearchKeyword), true),
	}

	_, err := server.store.CreateSearch(ctx, searchArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	searchDishArg := sql.NullString{
		String: req.SearchKeyword,
		Valid:  true}

	searchDishesDB, err := server.store.SearchDishes(ctx, searchDishArg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp := server.maptoModelSearchDish(searchDishesDB)

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) getRecentSearch(ctx *gin.Context) {

	var req searchRequest
	var recentSearchDB []db.Search

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListSearchesByUserIDParams{
		UserID: req.UserID,
		Limit:  5,
		Offset: 0,
	}

	recentSearchDB, err := server.store.ListSearchesByUserID(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	searchResult := server.maptoModelSearch(recentSearchDB)

	resp := searchResult

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) searchDishesDelete(ctx *gin.Context) {

	var req searchRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.DeleteSearchByKeywordParams{
		UserID:  req.UserID,
		Keyword: null.NewString(strings.ToLower(req.SearchKeyword), true),
	}

	searchResult, err := server.store.DeleteSearchByKeyword(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	if len(searchResult) == 0 {
		ctx.JSON(http.StatusNotFound, errResponse(sql.ErrNoRows))
		return
	}

	ctx.JSON(http.StatusOK, searchResult)
}
