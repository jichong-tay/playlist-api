package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type createPlaylistRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	IsPublic    bool   `json:"is_public"`
	DeliveryDay string `json:"delivery_day"`
	Category    string `json:"category"`
}

func (server *Server) createPlaylist(ctx *gin.Context) {
	var req createPlaylistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreatePlaylistParams{
		Name:        req.Name,
		Description: null.NewString(req.Description, true),
		ImageUrl:    null.NewString(req.ImageUrl, true),
		IsPublic:    req.IsPublic,
		DeliveryDay: null.NewString(req.DeliveryDay, true),
		Category:    null.NewString(req.Category, true),
	}
	playlist, err := server.store.CreatePlaylist(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

type getPlaylistRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPlaylist(ctx *gin.Context) {
	var req getPlaylistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	playlist, err := server.store.GetPlaylist(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}

type listPlaylistRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listPlaylist(ctx *gin.Context) {
	var req listPlaylistRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListPlaylistsParams{
		Limit:  int32((req.PageSize)),
		Offset: int32((req.PageID - 1) * req.PageSize),
	}

	playlists, err := server.store.ListPlaylists(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playlists)
}
