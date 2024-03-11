package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getPublicPlaylist(ctx *gin.Context) {

	playlistsDB, err := server.store.ListPlaylistPublicAndCategoryAll(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	// build publicPlaylist
	publicPlaylists, _ := server.maptoModelV2(ctx, playlistsDB)

	resp := publicPlaylists

	ctx.JSON(http.StatusOK, resp)
}
