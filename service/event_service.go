package service

import (
	"errors"
	"presensi-qr-backend/model"
	"presensi-qr-backend/repository"
	"time"
)

var ErrInvalidEventTime = errors.New("invalid event time")

type Service struct {
	eventRepo repository.EventRepository
}

func NewEventServer(repo repository.EventRepository) *Service {
	return &Service{eventRepo: repo}
}

func (s *Service) CreateEvent(title, desc string, startAt, endAt string, createdBy uint) (*model.Event, error) {
	start, err := time.Parse(time.RFC3339, startAt)
	if err != nil {
		return nil, ErrInvalidEventTime
	}
	end, err := time.Parse(time.RFC3339, endAt)
	if err != nil {
		return nil, ErrInvalidEventTime
	}

	e := model.Event{
		Title:       title,
		Description: desc,
		StartAt:     start,
		EndAt:       end,
		IsActive:    true,
		CreatedBy:   createdBy,
	}
	if err := s.eventRepo.Create(&e); err != nil {
		return nil, err
	}

	return &e, nil

}
