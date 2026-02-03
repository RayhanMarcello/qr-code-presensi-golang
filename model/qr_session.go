package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"size:255;uniqueIndex;not null"`
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

	CreatedBy uint `gorm:"index;not null"`
	Createor  User `gorm:"foreignKey:CreatedBy"`

	CreatedAt time.Time
}

type QRSession struct {
	ID        uint      `gorm:"primaryKey"`
	EventID   uint      `gorm:"index;not null"`
	Event     Event     `gorm:"foreignKey:EventID"`
	Token     string    `gorm:"type:char(64);uniqueIndex;not null"`
	ExpiredIn time.Time `gorm:"index;not null"`

	CreatedBy uint `gorm:"index;not null"`
	Creator   User `gorm:"foreignKey:CreatedBy"`

	CreateAt time.Time
}

type Attandance struct {
	ID      int    `gorm:"primaryKey"`
	EventID string `gorm:"index;not null"`
	Event   Event  `gorm:"foreignKey:EventID"`

	UserID string `gorm:"index;not null"`
	User   User   `gorm:"foreignKey:UserID"`

	QRSessionID uint      `gorm:"index;not null"`
	QRSession   QRSession `gorm:"foreignKey:QRSessionID"`

	CreateAt time.Time
}
