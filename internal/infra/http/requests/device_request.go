package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	Serial_number     string  `json:"serial_number" validate:"required"`
	Characteristics   *string `json:"characteristics"`
	Category          string  `json:"category" validate:"required"`
	Units             *string `json:"units"`
	Power_consumption *string `json:"power_consumption"`
}

type UpdateDeviceRequest struct {
	Serial_number     *string `json:"serial_number"`
	Characteristics   *string `json:"characteristics"`
	Category          *string `json:"category"`
	Units             *string `json:"units"`
	Power_consumption *string `json:"power_consumption"`
}

func (r UpdateDeviceRequest) ToDomainModel() (interface{}, error) {
	var (
		serial_number string
		category      string
	)
	if r.Serial_number != nil {
		serial_number = *r.Serial_number
	}

	if r.Category != nil {
		category = *r.Category
	}

	return domain.Device{
		Serial_number:     serial_number,
		Characteristics:   r.Characteristics,
		Category:          category,
		Units:             r.Units,
		Power_consumption: r.Power_consumption,
	}, nil
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
