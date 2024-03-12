package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/playlist-api/db/sqlc"
	"github.com/jichong-tay/playlist-api/util"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Address  string `json:"address"`
}

type createUserResponse struct {
	ID       int64       `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Address  null.String `json:"address"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Address:      null.NewString(req.Address, true),
	}
	user, err := server.store.CreateUser(ctx, arg)
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

	rsp := createUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
	}

	ctx.JSON(http.StatusOK, rsp)
}
