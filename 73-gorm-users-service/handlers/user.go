package handlers

import (
	"log/slog"
	"time"
	"user-service/database"
	"user-service/models"
	"user-service/security"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UsersDB *database.UserDB
}

func NewUserHandler(userDb *database.UserDB) *UserHandler {
	return &UserHandler{userDb}
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	user := new(models.User)

	err := ctx.Bind(user)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	err = user.Validate()
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	user.Status = "active"
	user.LastUpdated = uint64(time.Now().Unix())
	hash, err := security.HashPassword(user.Password)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	user.Password = hash

	user, err = uh.UsersDB.Create(user)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(400, "Something went wrong")
		ctx.Abort()
		return
	}

	user.Password = "*************"

	ctx.JSON(201, user)

}
