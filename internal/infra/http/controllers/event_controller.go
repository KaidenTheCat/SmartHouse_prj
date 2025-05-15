package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
)

type EventController struct {
	eventService app.EventService
}

func NewEventController(ev app.EventService) EventController {
	return EventController{
		eventService: ev,
	}
}

func (c EventController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		room := r.Context().Value(RoomKey).(domain.Room)
		event.Room_id = room.Id

		event, err = c.eventService.Save(event)
		if err != nil {
			log.Printf("EventController.Save(c.eventService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
