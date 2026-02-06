package service

import "presensi-qr-backend/model"

type UserService interface {
	Create(name, email, password, role string) (*model.User, error)
}

type EventService interface {
	CreateEvent(title, desc string, startAt, endAt string, createdBy uint) (*model.Event, error)
}
