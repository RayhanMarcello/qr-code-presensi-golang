package db

import (
	"fmt"
	"os"
	"presensi-qr-backend/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		fmt.Println("file .env ga berisi data yg bener")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err connect to db")
	}
	if err := db.AutoMigrate(
		&model.Attandance{},
		&model.Event{},
		&model.QRSession{},
		&model.User{},
	); err != nil {
		panic(err)
	}
	return db
}
