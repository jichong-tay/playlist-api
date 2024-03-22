package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getPlaylistLatestV2Request struct {
	UserID int64 `form:"user_id" binding:"required,min=0"`
}

func (server *Server) getPlaylistLatestV2(ctx *gin.Context) {
	var req getPlaylistLatestV2Request
	var resp playlistLatestResponsev2

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// build publicPlaylist
	var publicplaylists []playlistv2

	publicPlaylistsDB, err := server.store.ListPlaylistPublicAndCategoryAll(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	publicplaylists, _ = server.maptoModelV2(ctx, publicPlaylistsDB)

	// build userPlaylist
	var userplaylists []playlistv2

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

	userplaylists, _ = server.maptoModelV2(ctx, userPlaylistsDB)

	resp = playlistLatestResponsev2{
		PublicPlaylist: publicplaylists,
		UserPlaylist:   userplaylists,
	}

	ctx.JSON(http.StatusOK, resp)
}
