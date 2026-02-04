package router

import (
	"presensi-qr-backend/handler"

	"github.com/gin-gonic/gin"
)

func Route(userH *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/user", userH.CreateUser)

	return r
}
