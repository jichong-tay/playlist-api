package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getPlaylistUserRequest struct {
	UserID int64 `form:"user_id" binding:"required,min=0"`
}

func (server *Server) getPlaylistUser(ctx *gin.Context) {
	var req getPlaylistUserRequest
	var resp playlistUserResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

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

	resp = playlistUserResponse{
		UserPlaylist: userplaylists,
	}

	ctx.JSON(http.StatusOK, resp)
}
