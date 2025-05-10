package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type DeviceDto struct {
	Id                uint64     `json:"id"`
	House_id          uint64     `json:"house_id"`
	Room_id           uint64     `json:"room_id"`
	Uuid              uuid.UUID  `json:"uuid"`
	Serial_number     string     `json:"serial_number,omitempty"`
	Characteristics   *string    `json:"characteristics"`
	Category          string     `json:"category"`
	Units             *string    `json:"units,omitempty"`
	Power_consumption *string    `json:"power_consumption,omitempty"`
	CreateDate        time.Time  `json:"createDate"`
	UpdatedDate       time.Time  `json:"updateDate"`
	DeleteDate        *time.Time `json:"deleteDate,omitempty"`
}

func (d DeviceDto) DomainToDto(h domain.Device) DeviceDto {
	return DeviceDto{
		Id:                h.Id,
		House_id:          h.House_id,
		Room_id:           h.Room_id,
		Uuid:              h.Uuid,
		Serial_number:     h.Serial_number,
		Characteristics:   h.Characteristics,
		Category:          h.Category,
		Units:             h.Units,
		Power_consumption: h.Power_consumption,
		CreateDate:        h.CreateDate,
		UpdatedDate:       h.UpdatedDate,
		DeleteDate:        h.DeleteDate,
	}
}
