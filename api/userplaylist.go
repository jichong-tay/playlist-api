package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	UserID int64 `uri:"userid" binding:"required,min=0"`
}

func (server *Server) getUserPlaylist(ctx *gin.Context) {

	var req getUserRequest
	var userplaylists []playlistv2

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

	userplaylists, _ = server.maptoModelV2(ctx, userPlaylistsDB)

	resp := playlistResponsev2{
		Playlist: userplaylists,
	}

	ctx.JSON(http.StatusOK, resp)
}
