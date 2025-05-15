package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uint64
	Device_uuid uuid.UUID
	Room_id     uint64
	Action      string
	CreateDate  time.Time
}
