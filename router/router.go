package router

import (
	"presensi-qr-backend/handler"

	"github.com/gin-gonic/gin"
)

func Route(userH *handler.UserHandler, eventH *handler.HandlerEvent) *gin.Engine {
	r := gin.Default()

	r.POST("/user", userH.CreateUser)
	r.POST("/event", eventH.CreateEventHandler)

	return r
}
