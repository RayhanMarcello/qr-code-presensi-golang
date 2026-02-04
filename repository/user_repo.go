package repository

import (
	"fmt"
	"presensi-qr-backend/model"

	"gorm.io/gorm"
)

type GormUser struct {
	db *gorm.DB
}

func NewGormUser(db *gorm.DB) *GormUser {
	return &GormUser{db: db}
}

func (r *GormUser) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *GormUser) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println(err.Error())
	}
	return &user, nil
}
