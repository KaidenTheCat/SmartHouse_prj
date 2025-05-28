package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/google/uuid"
)

type EventService interface {
	Save(m domain.Event) (domain.Event, error)
	FindByUuid(uuid uuid.UUID) (domain.Device, error)
}

type eventService struct {
	eventRepo  database.EventRepository
	deviceRepo database.DeviceRepository
}

func NewEventService(er database.EventRepository, dr database.DeviceRepository) eventService {
	return eventService{
		eventRepo:  er,
		deviceRepo: dr,
	}
}

func (s eventService) Save(m domain.Event) (domain.Event, error) {
	event, err := s.eventRepo.Save(m)
	if err != nil {
		log.Printf("eventService.Save(s.eventRepo.Save): %s", err)
		return domain.Event{}, err
	}

	return event, nil
}

func (s eventService) FindByUuid(uuid uuid.UUID) (domain.Device, error) {
	device, err := s.deviceRepo.FindByUuid(uuid)
	if err != nil {
		log.Printf("eventService.FindByUuid(s.deviceRepo.FindByUuid): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}
