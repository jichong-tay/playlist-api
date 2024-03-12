package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
	"gopkg.in/guregu/null.v4"
)

type getUserRequest struct {
	UserID int64 `uri:"userid" binding:"required,min=0"`
}

func (server *Server) getUserPlaylist(ctx *gin.Context) {

	var req getUserRequest
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
