package service

import "presensi-qr-backend/repository"

type EventService struct {
	eventRepo repository.EventRepository
}

func NewEventRepo(eventRepo repository.EventRepository) *EventService {

}
