package domain

import "time"

type Room struct {
	Id          uint64
	HouseId     uint64
	Name        string
	Description *string
	CreateDate  time.Time
	UpdatedDate time.Time
	DeleteDate  *time.Time
}
