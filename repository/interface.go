package repository

import (
	"presensi-qr-backend/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type EventRepository interface {
	Create(e *model.Event) error
	// FindById(id uint) (*model.Event, error)
	// LastActive() ([]model.Event, error)
}

type QRSessionRepository interface {
	Create(q *model.QRSession) error
	FindByToken(token string) (*model.QRSession, error)
}

type AttandanceRepository interface {
	Create(tx *gorm.DB, a *model.Attandance) error
}
