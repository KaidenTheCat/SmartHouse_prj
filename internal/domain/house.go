package domain

import "time"

type House struct {
	Id          uint64
	UserId      uint64
	Name        string
	Description *string
	City        string
	Address     string
	Lat         float64
	Lon         float64
	CreateDate  time.Time
	UpdatedDate time.Time
	DeleteDate  *time.Time
}
