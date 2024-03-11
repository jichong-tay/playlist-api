package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
)

type getPlaylistLatestRequest struct {
	UserID         int64 `form:"user_id" binding:"required,min=0"`
	PublicPageID   int64 `form:"public_page_id" binding:"min=1"`
	PublicPageSize int64 `form:"public_page_size" binding:"min=5,max=100"`
	UserPageID     int64 `form:"user_page_id" binding:"min=1"`
	UserPageSize   int64 `form:"user_page_size" binding:"min=5,max=100"`
}

func (server *Server) getPlaylistLatest(ctx *gin.Context) {

	var req getPlaylistLatestRequest
	var resp playlistLatestResponsev1

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	argPublic := db.ListPlaylistPublicAndCategoryParams{
		Limit:  int32((req.PublicPageSize)),
		Offset: int32((req.PublicPageID - 1) * req.PublicPageSize),
	}

	// build publicPlaylist
	var publicPlaylist []categoryPlaylist

	publicPlaylistDB, err := server.store.ListPlaylistPublicAndCategory(ctx, argPublic)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	publicPlaylist, _ = server.maptoModelCategory(ctx, publicPlaylistDB)

	//build userPlaylist
	var userPlaylist []playlistv1

	argUser := db.ListPlaylistsByUserIDParams{
		UserID: req.UserID,
		Limit:  int32((req.UserPageSize)),
		Offset: int32((req.UserPageID - 1) * req.UserPageSize),
	}

	userPlaylistDB, err := server.store.ListPlaylistsByUserID(ctx, argUser)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	userPlaylist, _ = server.maptoModelV1(ctx, userPlaylistDB)

	//build response
	resp = playlistLatestResponsev1{
		CategoryPlaylist: publicPlaylist,
		UserPlaylist:     userPlaylist,
	}

	ctx.JSON(http.StatusOK, resp)
}
