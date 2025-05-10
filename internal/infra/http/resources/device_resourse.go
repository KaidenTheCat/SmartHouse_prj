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

type DeviceFindDto struct {
	Id                uint64    `json:"id"`
	House_id          uint64    `json:"house_id"`
	Room_id           uint64    `json:"room_id"`
	Uuid              uuid.UUID `json:"uuid"`
	Serial_number     string    `json:"serial_number,omitempty"`
	Characteristics   *string   `json:"characteristics"`
	Category          string    `json:"category"`
	Units             *string   `json:"units,omitempty"`
	Power_consumption *string   `json:"power_consumption,omitempty"`
}

func (d DeviceDto) DomainToDto(dd domain.Device) DeviceDto {
	return DeviceDto{
		Id:                dd.Id,
		House_id:          dd.House_id,
		Room_id:           dd.Room_id,
		Uuid:              dd.Uuid,
		Serial_number:     dd.Serial_number,
		Characteristics:   dd.Characteristics,
		Category:          dd.Category,
		Units:             dd.Units,
		Power_consumption: dd.Power_consumption,
		CreateDate:        dd.CreateDate,
		UpdatedDate:       dd.UpdatedDate,
		DeleteDate:        dd.DeleteDate,
	}
}

func (d DeviceFindDto) DomainToFindDto(dd domain.Device) DeviceFindDto {
	return DeviceFindDto{
		Id:                dd.Id,
		House_id:          dd.House_id,
		Room_id:           dd.Room_id,
		Uuid:              dd.Uuid,
		Serial_number:     dd.Serial_number,
		Characteristics:   dd.Characteristics,
		Category:          dd.Category,
		Units:             dd.Units,
		Power_consumption: dd.Power_consumption,
	}
}
