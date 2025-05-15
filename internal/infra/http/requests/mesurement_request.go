package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type MeasurementRequest struct {
	Device_uuid uuid.UUID `json:"device_uuid" validate:"required"`
	Value       string    `json:"value" validate:"required"`
}

func (r MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		Device_uuid: r.Device_uuid,
		Value:       r.Value,
	}, nil
}
