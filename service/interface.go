package service

import "presensi-qr-backend/model"

type UserService interface {
	Create(name, email, password, role string) (*model.User, error)
}
