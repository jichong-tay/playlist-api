package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getPublicPlaylist(ctx *gin.Context) {

	playlistDB, err := server.store.ListPlaylistPublicAndCategoryAll(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	// build publicPlaylist
	publicPlaylist, _ := server.maptoModelV2(ctx, playlistDB)

	resp := playlistResponsev2{
		Playlist: publicPlaylist,
	}

	ctx.JSON(http.StatusOK, resp)
}
