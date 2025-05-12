package domain

import (
	"time"

	"github.com/google/uuid"
)

const SENSOR string = "SENSOR"
const ACTUATOR string = "ACTUATOR"

type Device struct {
	Id                uint64
	House_id          uint64
	Room_id           uint64
	Uuid              uuid.UUID
	Serial_number     string
	Characteristics   *string
	Category          string
	Units             *string // for category SENSOR
	Power_consumption *string // for category ACTUATOR
	CreateDate        time.Time
	UpdatedDate       time.Time
	DeleteDate        *time.Time
}
