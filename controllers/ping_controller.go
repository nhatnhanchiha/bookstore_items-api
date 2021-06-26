package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
}

func (p pingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
