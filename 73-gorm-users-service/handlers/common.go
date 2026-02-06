package handlers

import "github.com/gin-gonic/gin"

func Health(ctx *gin.Context) {
	ctx.String(200, "ok")
}

func Root(ctx *gin.Context) {
	ctx.String(200, "Welcome to users-service")
}

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}
