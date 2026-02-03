package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"index;not null;default:user"`
	CreateAt time.Time
}

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	StartAt     time.Time `gorm:"index;not null"`
	EndAt       time.Time `gorm:"index;not null"`
	IsActive    bool      `gorm:"index;default:true;not null"`

	CreatedBy uint
}

type QRSession struct {
	ID        int       `gorm:"primaryKey"`
	EvenID    string    `gorm:"index;not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiredIn time.Time `gorm:"index;not null"`
	Used      bool      `gorm:"not null;default:false"`
	CreateAt  time.Time
}

type Attandance struct {
	ID       int    `gorm:"primaryKey"`
	EventID  string `gorm:"index;not null"`
	UserID   string `gorm:"index;not null"`
	CreateAt time.Time
}
