package main

import (
	"fmt"
	"presensi-qr-backend/db"
	"presensi-qr-backend/handler"
	"presensi-qr-backend/repository"
	"presensi-qr-backend/router"
	"presensi-qr-backend/service"

	// "presensi-qr-backend/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("file .env dosent exist")
	}
	db := db.Database()

	repoUser := repository.NewGormUser(db)
	repoEvent := repository.NewGormEvent(db)

	srvUser := service.NewUserService(repoUser)
	srvEvent := service.NewEventServer(repoEvent)

	handlerUser := handler.NewUserHandler(srvUser)
	handerEvent := handler.NewHandlerEvent(srvEvent)

	r := router.Route(handlerUser, handerEvent)

	// r
	r.Run(":8000")
}
