package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	Serial_number     string  `json:"serial_number" validate:"required"`
	Characteristics   *string `json:"characteristics"`
	Category          string  `json:"category" validate:"required"`
	Units             *string `json:"units"`
	Power_consumption *string `json:"power_consumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		Serial_number:     r.Serial_number,
		Characteristics:   r.Characteristics,
		Category:          r.Category,
		Units:             r.Units,
		Power_consumption: r.Power_consumption,
	}, nil
}
