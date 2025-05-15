package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type EventService interface {
	Save(m domain.Event) (domain.Event, error)
}

type eventService struct {
	eventRepo database.EventRepository
}

func NewEventService(er database.EventRepository) eventService {
	return eventService{
		eventRepo: er,
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
