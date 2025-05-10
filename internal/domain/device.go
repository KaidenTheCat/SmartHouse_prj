package domain

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	Id                uint64
	House_id          uint64
	Room_id           uint64
	Uuid              uuid.UUID
	Serial_number     string
	Characteristics   *string
	Category          string
	Units             *string
	Power_consumption *string
	CreateDate        time.Time
	UpdatedDate       time.Time
	DeleteDate        *time.Time
}
