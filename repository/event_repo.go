package repository

import (
	"fmt"
	"presensi-qr-backend/model"

	"gorm.io/gorm"
)

type GormEvent struct {
	db *gorm.DB
}

func newGormEvent(db *gorm.DB) *GormEvent {
	return &GormEvent{db: db}
}

func (r *GormEvent) Create(e *model.Event) error {
	err := r.db.Create(e).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
