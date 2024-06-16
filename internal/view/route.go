package view

import (
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitGin(g *gin.Engine) {
	api := g.Group("/api")
	api.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, map[string]interface{}{"ping": "pong"}) })

	//event
	eventRoute := NewEventRoute()
	api.POST("/event", sdkginext.NewEventHandlerFunc(eventRoute.InitEvent()))
}

func initMiddle(g *gin.Engine) {
	g.Use(gin.Recovery())
}
