package service

import (
	"errors"
	"presensi-qr-backend/model"
	"presensi-qr-backend/repository"
)

var ErrEmailTaken = errors.New("email already taken")

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) Create(name, email, password, role string) (*model.User, error) {
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, ErrEmailTaken
	}
	if role == "" {
		role = "user"
	}
	if role != "user" && role != "admin" {
		role = "user"
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
