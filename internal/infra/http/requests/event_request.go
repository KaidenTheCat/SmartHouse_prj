package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type EventRequest struct {
	Device_uuid uuid.UUID `json:"device_uuid" validate:"required"`
	Action      string    `json:"action" validate:"required"`
}

func (r EventRequest) ToDomainModel() (interface{}, error) {
	return domain.Event{
		Device_uuid: r.Device_uuid,
		Action:      r.Action,
	}, nil
}
