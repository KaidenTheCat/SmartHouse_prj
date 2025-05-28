package domain

import (
	"time"

	"github.com/google/uuid"
)

type Measurement struct {
	Id          uint64
	Device_id   uint64
	Device_uuid uuid.UUID
	Room_id     uint64
	Value       string
	CreateDate  time.Time
}
